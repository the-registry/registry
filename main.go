package main

import "github.com/visionmedia/docopt"
import "github.com/segmentio/go-log"
import "fmt"

const Version = "0.0.1"

const Usage = `
  Usage:
   versions search <name>
   versions register <name>
   versions -h | --help
   versions --version

  Options:
    -h, --help       output help information
    -v, --version    output version
`

func main() {
	args, err := docopt.Parse(Usage, nil, true, Version, false)
	log.Check(err)

	name := args["<name>"].(string)

	fmt.Println(name)

	if args["search"].(bool) {

	} else if args["register"].(bool) {

	}
}
