package client

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func DoRequest(client *http.Client, url string) ([]byte, error) {
	fmt.Println("Making request to:", url)
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// read response body regardless of status code
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Response body:", string(body))
		return nil, errors.New("unexpected status code: " + resp.Status)
	}

	return body, nil
}
