package cmd

import (
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"testing"
)

func TestGetNewestRelease(t *testing.T) {
	type args struct {
		owner  string
		repo   string
		client *github.Client
		ctx    context.Context
	}
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, nil)

	client := github.NewClient(tc)

	tests := []struct {
		name string
		args args
		want int
	}{{"Basic test for github", args{
		owner:  "google",
		repo:   "go-github",
		client: client,
		ctx:    ctx,
	}, 200},
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNewestRelease(tt.args.owner, tt.args.repo); got != tt.want {
				t.Errorf("GetNewestRelease() = %v, want %v", got, tt.want)
			}
		})
	}
}
