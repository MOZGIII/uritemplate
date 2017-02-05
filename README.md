# urltemplate

A simple command line utility that performs RFC 6570 URL Template expansion and does some other related things.

Supports level 4 of RFC 6570.

## Installation

```shell
$ go get github.com/MOZGIII/uritemplate
```

## Usage

When run without arguments it displays detailed commands description.

```shell
$ uritemplate
usage: uritemplate [<flags>] <command> [<args> ...]

A command-line RFC6570 (URI Template) expander.

Flags:
  --help     Show context-sensitive help (also try --help-long and --help-man).
  --newline  Print newline at the end of output (use --no-newline to ommit
             newlines).

Commands:
  help [<command>...]
    Show help.

  expand [<flags>] <template>
    Print the expanded URI template.

  varnames [<flags>] <template>
    Print variable names found in the template.

  regexp <template>
    Print a regexp that matches the template.
```

### URL Template expansion

For help check out `uritemplate expand --help`.

For more info on how exansion works, see [section 3.2 of RFC 6570](https://tools.ietf.org/html/rfc6570#section-3.2).

#### Simple variables

```
$ uritemplate expand "https://google.com/{?q}" --var q=test
https://google.com/?q=test

$ uritemplate expand "http://google.com/{?q}" --var "q=My Query"
http://google.com/?q=My%20Query
```

You can pass multiple `--var` options to expand more variables.

#### JSON variables

You can fully utilize level 4 expansion with JSON-encoded arguments.

```
$ uritemplate expand "http://google.com/{?q}" --json '{ "q": ["My Query", "My Other Query"] }'
http://google.com/?q=My%20Query,My%20Other%20Query

$ uritemplate expand "http://google.com/{?q*}" --json '{ "q": ["My Query", "My Other Query"] }'
http://google.com/?q=My%20Query&q=My%20Other%20Query

$ uritemplate expand "http://google.com/{?q*}" --json '{ "q": { "a": "My Query", "b": "My Other Query" } }'
http://google.com/?a=My%20Query&b=My%20Other%20Query

$ uritemplate expand "http://google.com/{?q}" --json '{ "q": { "a": "My Query", "b": "My Other Query" } }'
http://google.com/?q=a,My%20Query,b,My%20Other%20Query
```

Multiple `--json` options are allowed.

## Development

Feel free to send pull requests.
Make sure code is formatted via `go fmt` and builds correctly.
