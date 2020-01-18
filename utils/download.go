package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type ResponseError struct {
	Code int
	Url  string
}

func (e ResponseError) Error() string {
	return "Non-200 return code from HTTP while downloading: " + e.Url + " " +  strconv.Itoa(e.Code)
}

/* Download downloads a file into destination by *url */
func Download(url string, dst string) error {

	fmt.Println("Downloading ", url, " to ", dst)
	resp, err := http.Get(url)
	if err != nil {
		return errors.New("Error while downloading")
	}

	if resp.StatusCode != 200 {
		return ResponseError{resp.StatusCode, url}
	}

	out, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("Couldn't create file: %v", err)
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)

	if err != nil {
		return fmt.Errorf("Could not copy file: %v", err)
	}
	return nil
}
