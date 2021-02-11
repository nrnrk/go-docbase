package main

import (
	"context"
	"fmt"
	"os"

	"github.com/nrnrk/go-docbase"
)

func main() {
	client := docbase.NewClient(os.Getenv(`DOCBASE_TEAM`), os.Getenv(`DOCBASE_TOKEN`))
	post, err := client.CreatePost(
		context.Background(),
		&docbase.CreatePostRequest{
			Title:  `From go-docbase`,
			Body:   `Sample Body`,
			Draft:  false,
			Scope:  docbase.ScopePrivate,
			Tags:   nil,
			Groups: nil,
			Notice: false,
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Succeeded! Post: %#v\n", *post)
}
