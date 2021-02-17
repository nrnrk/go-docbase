package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/nrnrk/go-docbase"
)

func main() {
	// Create client
	client := docbase.NewClient(os.Getenv(`DOCBASE_TEAM`), os.Getenv(`DOCBASE_TOKEN`))

	// Create a new post!
	createdPost, err := client.CreatePost(
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
		log.Fatal(err)
	}
	fmt.Printf("Succeeded to Create Post:\n%#v\n", *createdPost)

	// Get the existing post
	gotPost, err := client.GetPost(
		context.Background(),
		createdPost.ID,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Succeeded to Get Post:\n%#v\n", *gotPost)

	// Update the existing post
	updatedPost, err := client.UpdatePost(
		context.Background(),
		&docbase.UpdatePostRequest{
			ID:     gotPost.ID,
			Body:   `Simple Body`,
			Scope:  docbase.ScopeEveryone,
			Tags:   nil,
			Groups: nil,
			Notice: false,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Succeeded to Update Post:\n%#v\n", *updatedPost)

	// List tags
	tags, err := client.ListTags(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(`Succeeded to List Tags!`)
	for _, tag := range tags {
		fmt.Printf("%#v\n", *tag)
	}
}
