# the-registry cli

```
Usage:
 registry search <type> <name>
 registry show <type> <name>
 registry register <type> <name> <url>
 registry unregister <type> <name>
 registry -h | --help
 registry --version

Options:
  -h, --help       output help information
  -v, --version    output version
```

## Example

```
$ registry register duo moombahton https://github.com/travisjeffery/moombahton.git
> Registered duo package "moombahton"

$ registry search duo moombahton
[
  {
    "name": "moombahton",
    "type": "duo",
    "url": "https://github.com/travisjeffery/moombahton.git"
  }
]
```
