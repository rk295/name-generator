# Random name generator

[![Go Reference](https://pkg.go.dev/badge/github.com/rk295/name-generator.svg)](https://pkg.go.dev/github.com/rk295/name-generator)

The Yak was shaved. 

Build with `make` . 

It has some options:

```
% name-generator --help
A random name generator

Usage:
  name-generator [flags]

Flags:
  -h, --help               help for name-generator
  -n, --number int         Number of names to generate (default 1)
  -r, --random             Append a random 6 digit number
  -s, --separator string   Separator to use between words (default "-")
  -t, --types strings      Types to include (default [colours,dogs])
```

## Adding more words

Drop a file in [ `lib/data` ](lib/data) , format is one per line, file must end with `.txt` .

## Yak shaver hall of shame

<a href="https://github.com/rk295/name-generator/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=rk295/name-generator" />
</a>

Made with [contrib.rocks](https://contrib.rocks).
