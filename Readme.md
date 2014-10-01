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
$ registry register rubygem timecop https://github.com/travisjeffery/timecop.git
> 201 Created

$ registry search rubygem timecop
> 200 OK
[
  {
    "name": "timecop",
    "type": "rubygem",
    "url": "https://github.com/travisjeffery/timecop.git"
  }
]
```
