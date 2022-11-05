package main

import (
	_ "embed"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type xferConfig struct {
	ServerEndpoint string
}

var Version = "development"

func main() {

	switch len(os.Args) {
	case 1:
		// help
		fmt.Printf("xfer (version %s) - help:\n\nSimply pass in the filename you wish to upload!\n", Version)
	case 2:
		// upload file
		link, token, err := upload(os.Args[0])
		if err != nil {
			log.Fatalf("Failed to upload file: %s", err)
		}
		fmt.Printf("Link: %s\nDelete token: %s\n", link, token)
	default:
		log.Fatalf("Only one parameter expected, I.E. the file name to upload")
	}

}

// upload
// returns link, token and error
func upload(f string) (string, string, error) {

	c, err := loadConfig()
	if err != nil {
		return "", "", err
	}

	ep := c.ServerEndpoint + "/" + rndName()

	data, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	req, err := http.NewRequest("PUT", ep, data)
	if err != nil {
		return "", "", err
	}
	req.Header.Set("Content-Type", "text/plain")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()

	// TODO check response codes ???

	hdr := res.Header.Get("x-url-delete")
	tmp := strings.Split(hdr, "/")
	if len(tmp) > 1 {
		token := tmp[len(tmp)-1]
		link := hdr[0 : len(hdr)-len(token)-1]
		return link, token, nil
	} else {
		return "", "", errors.New("invalid or missing x-url-delete response from server")
	}

}
