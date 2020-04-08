package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/shurcooL/githubv4"

	"golang.org/x/oauth2"
)

var (
	apiToken = flag.String("apitoken", "", "use api token")
)

func main() {
	flag.Parse()

	if *apiToken == "" {
		fmt.Println("required apitoken")
		os.Exit(1)
	}

	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: *apiToken},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	// 投げるクエリ
	var query struct {
		Viewer struct {
			Login     githubv4.String
			CreatedAt githubv4.DateTime
		}
	}

	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("    Login:", query.Viewer.Login)
	fmt.Println("CreatedAt:", query.Viewer.CreatedAt)
}
