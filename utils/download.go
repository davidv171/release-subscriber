package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/* Download downloads a file into destination by *url */
func Download(url *string, dst string) int {

	fmt.Println("Downloading ", *url, " to ", dst)
	resp, err := http.Get(*url)
	if err != nil {
		fmt.Println("Error while downloading ", err)
		return 1
	}

	if resp.StatusCode != 200 {
		fmt.Println("Non-zero return code from HTTP while downloading ", url)
		return 1
	}
	out, err := os.Create(dst)
	if err != nil {
		fmt.Println("Couldn't create file")
		return 1
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)

	if err != nil {
		fmt.Println("Could not copy file...")
		return 1
	}
	return 0
}
