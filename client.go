package docbase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type (
	Client interface {
		CreatePost(ctx context.Context, req *CreatePostRequest) (*Post, error)
		UpdatePost(ctx context.Context, req *UpdatePostRequest) (*Post, error)
		GetPost(ctx context.Context, id uint) (*Post, error)
	}

	client struct {
		team       string
		token      string
		apiVersion string
		httpClient httpClient
	}
)

func NewClient(team, token string) Client {
	return &client{
		team:       team,
		token:      token,
		apiVersion: `2`,
		httpClient: http.DefaultClient,
	}
}

func (c *client) CreatePost(ctx context.Context, req *CreatePostRequest) (*Post, error) {
	httpReq, err := c.createPostRequest(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("Failed to create http reqeust: %w", err)
	}

	post, err := c.executeAPI(httpReq)
	if err != nil {
		return nil, fmt.Errorf("DocBase API error: %w", err)
	}

	return post, nil
}

func (c *client) UpdatePost(ctx context.Context, req *UpdatePostRequest) (*Post, error) {
	httpReq, err := c.createUpdateRequest(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("Failed to create http reqeust: %w", err)
	}

	post, err := c.executeAPI(httpReq)
	if err != nil {
		return nil, fmt.Errorf("DocBase API error: %w", err)
	}

	return post, nil
}

func (c *client) GetPost(ctx context.Context, id uint) (*Post, error) {
	httpReq, err := c.createGetRequest(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("Failed to create http reqeust: %w", err)
	}

	post, err := c.executeAPI(httpReq)
	if err != nil {
		return nil, fmt.Errorf("DocBase API error: %w", err)
	}

	return post, nil
}

func (c *client) executeAPI(httpReq *http.Request) (*Post, error) {
	res, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("Failed to send http reqeust: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("Got error from docbase server: %s", res.Status)
		}

		return nil, fmt.Errorf("Got error from docbase server: %s, body: %s", res.Status, body)
	}

	var post *Post
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read body: %w", err)
	}
	err = json.Unmarshal(body, &post)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal body to json: %w", err)
	}

	return post, nil
}

func (c *client) createPostRequest(ctx context.Context, req *CreatePostRequest) (*http.Request, error) {
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal request to json: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(
		ctx,
		`POST`,
		fmt.Sprintf("https://api.docbase.io/teams/%s/posts", c.team),
		bytes.NewReader(jsonReq),
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %w", err)
	}

	httpReq.Header.Add(`X-DocBaseToken`, c.token)
	httpReq.Header.Add(`Content-Type`, `application/json`)
	return httpReq, nil
}

func (c *client) createUpdateRequest(ctx context.Context, req *UpdatePostRequest) (*http.Request, error) {
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal request to json: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(
		ctx,
		`PATCH`,
		fmt.Sprintf("https://api.docbase.io/teams/%s/posts/%d", c.team, req.ID),
		bytes.NewReader(jsonReq),
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %w", err)
	}

	httpReq.Header.Add(`X-DocBaseToken`, c.token)
	httpReq.Header.Add(`Content-Type`, `application/json`)
	return httpReq, nil
}

func (c *client) createGetRequest(ctx context.Context, id uint) (*http.Request, error) {
	httpReq, err := http.NewRequestWithContext(
		ctx,
		`GET`,
		fmt.Sprintf("https://api.docbase.io/teams/%s/posts/%d", c.team, id),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %w", err)
	}

	httpReq.Header.Add(`X-DocBaseToken`, c.token)
	return httpReq, nil
}
