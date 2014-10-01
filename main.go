package main

import "github.com/visionmedia/docopt"
import "github.com/segmentio/go-log"
import "net/http"
import "net/url"
import "fmt"
import "encoding/json"
import "bytes"
import "io/ioutil"
import "os"

const Version = "0.0.1"

const Api = "http://localhost:3000"

const Usage = `
  Usage:
   registry search <type> <name>
   registry register <type> <name> <url>
   registry unregister <type> <name>
   registry -h | --help
   registry --version

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
	t := args["<type>"].(string)

	if args["search"].(bool) {
		api.Search(name, t)
	} else if args["register"].(bool) {
		api.Register(name, args["<url>"].(string), t)
	} else if args["unregister"].(bool) {
		api.Unregister(name, t)
	}
}

type Client struct {
	http *http.Client
}

func (c *Client) Search(name string, t string) {
	p := url.Values{}
	p.Add("name", name)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/types/%s/packages/search?%s", Api, t, p.Encode()), nil)
	log.Check(err)

	resp, err := c.http.Do(req)
	log.Check(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Check(err)

	var d interface{}
	err = json.Unmarshal(body, &d)
	log.Check(err)

	b, err := json.MarshalIndent(d, "", "  ")
	log.Check(err)

	os.Stdout.Write(b)
}

func (c *Client) Register(name string, u string, t string) {
	p := url.Values{}
	p.Add("name", name)
	p.Add("url", u)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/types/%s/packages?%s", Api, t, p.Encode()), nil)
	log.Check(err)

	_, err = c.http.Do(req)
	log.Check(err)

	fmt.Printf("Registered %s package \"%s\"\n", t, name)
}

func (c *Client) Unregister(name string, t string) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/types/%s/packages/%s", Api, t, name), nil)
	log.Check(err)

	_, err = c.http.Do(req)
	log.Check(err)

	fmt.Printf("Unregistered %s package \"%s\"\n", t, name)
}

func jsonBytes(j map[string]string) *bytes.Buffer {
	b, err := json.Marshal(j)
	log.Check(err)
	return bytes.NewBuffer(b)
}
