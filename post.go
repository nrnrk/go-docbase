package docbase

import (
	"errors"
	"time"
)

type CreatePostRequest struct {
	Title  string   `json:"title"`
	Body   string   `json:"body"`
	Draft  bool     `json:"draft"`
	Scope  Scope    `json:"scope"`
	Tags   []string `json:"tags,omitempty"`
	Groups []uint   `json:"groups,omitempty"`
	Notice bool     `json:"notice"`
}

func (r *CreatePostRequest) Validate() error {
	if r == nil {
		return errors.New(`CreatePostRequest must be set`)
	}
	if r.Title != `` {
		return errors.New(`Title must not be empty`)
	}
	if r.Scope == ScopeGroup && len(r.Groups) == 0 {
		return errors.New(`Group must be specified when the scope is group`)
	}
	return nil
}

type UpdatePostRequest struct {
	ID     uint
	Title  string   `json:"title,omitempty"`
	Body   string   `json:"body,omitempty"`
	Draft  bool     `json:"draft,omitempty"`
	Scope  Scope    `json:"scope,omitempty"`
	Tags   []string `json:"tags,omitempty"`
	Groups []uint   `json:"groups,omitempty"`
	Notice bool     `json:"notice,omitempty"`
}

func (r *UpdatePostRequest) Validate() error {
	if r == nil {
		return errors.New(`UpdatePostRequest must be set`)
	}
	if r.ID == 0 {
		return errors.New(`ID must be specified`)
	}
	if r.Scope == ScopeGroup && len(r.Groups) == 0 {
		return errors.New(`Group must be specified when the scope is group`)
	}
	if r.Scope != ScopeGroup && len(r.Groups) > 0 {
		return errors.New(`Group cannot be set when the scope is not group`)
	}
	return nil
}

type Post struct {
	ID            uint         `json:"id"`
	Title         string       `json:"title"`
	Body          string       `json:"body"`
	Draft         bool         `json:"draft"`
	Archived      bool         `json:"archived"`
	URL           string       `json:"url"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
	Scope         Scope        `json:"scope"`
	Tags          []Tag        `json:"tags"`
	User          User         `json:"user"`
	StartsCount   uint         `json:"stars_count"`
	GoodJobsCount uint         `json:"good_jobs_count"`
	SharingURL    *time.Time   `json:"sharing_url"`
	Comments      []Comment    `json:"comments"`
	Groups        []Group      `json:"groups"`
	Attachments   []Attachment `json:"attachments"`
}

type Scope string

const (
	ScopeEveryone Scope = `everyone`
	ScopeGroup    Scope = `group`
	ScopePrivate  Scope = `private`
)

type Tag struct {
	Name string `json:"name"`
}

type Comment struct {
	ID        string    `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"user"`
}

type Attachment struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Size      uint      `json:"size"`
	URL       string    `json:"url"`
	Markdown  string    `json:"markdown"`
	CreatedAt time.Time `json:"created_at"`
}
