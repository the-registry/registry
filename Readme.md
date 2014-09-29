# the-registry cli

```
Usage:
 registry search <name> <type>
 registry register <name> <type> <url>
 registry unregister <name> <type>
 registry -h | --help
 registry --version

Options:
  -h, --help       output help information
  -v, --version    output version
```

## Example

```
$ registry register timecop rubygem https://github.com/travisjeffery/timecop.git
> 201 Created

$ register search timecop
> 200 OK
[
  {
    "name": "timecop",
    "type": "rubygem",
    "url": "https://github.com/travisjeffery/timecop.git"
  }
]
```
