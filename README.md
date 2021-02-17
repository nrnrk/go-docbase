# go-docbase

Golang client for [DocBase](https://docbase.io/)

## Features

- Create / Update / Get DocBase posts
- List DocBase tags

## Usage

1. Generate DocBase token here (https://[team].docbase.io/settings/tokens) ([for more details](https://help.docbase.io/posts/45703#%E3%82%A2%E3%82%AF%E3%82%BB%E3%82%B9%E3%83%88%E3%83%BC%E3%82%AF%E3%83%B3%E4%BD%9C%E6%88%90%E6%96%B9%E6%B3%95))

2. Create client with DocBase team & DocBase token

### Example

```go
import "github.com/nrnrk/go-docbase"

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
        Tags:   []string{`golang`, `sample`},
        Groups: nil,
        Notice: false,
    },
)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Succeeded to Create! Post: %#v\n", *createdPost)
```

## Release Notes

### 0.1.0

- Support basic post & tag API
