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

	res, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("Failed to send http reqeust: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode >= 300 {
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
