package docbase

import (
	"errors"
	"time"
)

type CreatePostRequest struct {
	Title  string  `json:"title"`
	Body   string  `json:"body"`
	Draft  bool    `json:"draft"`
	Scope  Scope   `json:"scope"`
	Tags   []Tag   `json:"tags,omitempty"`
	Groups []Group `json:"groups,omitempty"`
	Notice bool    `json:"notice"`
}

func (r *CreatePostRequest) Validate() error {
	if r == nil {
		return errors.New(`CreatePostRequest must be set`)
	}
	if r.Title != `` {
		return errors.New(`Title must not be empty`)
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

// FIXME: implement
type Comment struct{}

// FIXME: implement
type Attachment struct{}
