package main

import "github.com/visionmedia/docopt"
import "github.com/segmentio/go-log"
import "net/http"
import "fmt"
import "encoding/json"
import "bytes"

const Version = "0.0.1"

const Api = "http://localhost:3000"

const Usage = `
  Usage:
   versions search <name> <type>
   versions register <name> <type> <url>
   versions -h | --help
   versions --version

  Options:
    -h, --help       output help information
    -v, --version    output version
`

func main() {
	args, err := docopt.Parse(Usage, nil, true, Version, false)
	log.Check(err)

	api := Client{
		http: &http.Client{},
	}

	name := args["<name>"].(string)

	if args["search"].(bool) {
		// search api
	} else if args["register"].(bool) {
		api.Register(name, args["<url>"].(string), args["<type>"].(string))
	}
}

type Client struct {
	http *http.Client
}

func (c *Client) Search(name string, t string) {
	j := map[string]string{
		"name": name,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s/packages", Api, t), jsonBytes(j))
	log.Check(err)

	resp, err := c.http.Do(req)
	log.Check(err)

	fmt.Println(resp.Status)
	fmt.Println(resp.Body)
}

func (c *Client) Register(name string, url string, t string) {
	j := map[string]string{
		"name": name,
		"url":  url,
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s/packages", Api, t), jsonBytes(j))
	log.Check(err)

	resp, err := c.http.Do(req)
	log.Check(err)

	fmt.Println(resp.Status)
}

func jsonBytes(j map[string]string) *bytes.Buffer {
	b, err := json.Marshal(j)
	log.Check(err)
	return bytes.NewBuffer(b)
}
