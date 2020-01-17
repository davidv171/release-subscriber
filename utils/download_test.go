package utils

import (
	"fmt"
	"os"
	"testing"
)

func TestDownload(t *testing.T) {
	type args struct {
		url *string
		dst string
	}
	var url = "http://www.golang-book.com/public/pdf/gobook.pdf"
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Basic test from their github", args{
			url: &url,
			dst: "book.pdf",
		}, 0},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Download(tt.args.url, tt.args.dst); got != tt.want {
				t.Errorf("Download() = %v, want %v", got, tt.want)

				err := os.Remove("/go/src/github.com/davidv171/release-subscriber/utils/gobook.pdf")
				if err != nil {
					fmt.Errorf("Error during deletion in test stages : %s", err)
				}

			}
		})
	}
}
