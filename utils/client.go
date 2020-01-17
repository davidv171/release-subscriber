package utils

import (
	"github.com/google/go-github/github"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

func GetClient() (context.Context, *github.Client) {
	ctx := context.Background()
	//Our endpoints dont need authentication
	tc := oauth2.NewClient(ctx, nil)
	client := github.NewClient(tc)

	return ctx, client
}
