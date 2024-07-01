# Temp Mail Detector Golang API

This repo contains the required code for you to make an API request to the [Temp Mail Detector](https://tempmaildetector.com) service in the Go programming language.

Temporary email addresses can cause issues for services which provide a freemium model or which offer a trial. While we understand that temporary emails are great at preserving privacy, there is a need to control where and when they can be used.

Below you will find an example implementation and json response of this library:

### Example response
```json
{
  "email": "soma+email@host.com",
  "score": 100,
  "meta": {
    "block_list": true,
    "domain_age": 1,
    "website_resolves": false,
    "random_characters": true,
    "accepts_all_addresses": false,
    "uses_plus": false
  }
}
```

### Example usage
```Go
package main

import (
	"fmt"
	"log"

	"github.com/TempMailDetector/TempMailDetector-Go"
)

func main() {
	apiKey := "YOUR_API_KEY"
	client := tempmaildetector.NewClient(apiKey)

	email := "temp+mail@host.com"
	response, err := client.CheckEmail(email)
	if err != nil {
		log.Fatalf("Error checking email: %v", err)
	}

	fmt.Printf("Email: %s\nScore: %d\nMeta: %+v\n", response.Email, response.Score, response.Meta)
}
```
