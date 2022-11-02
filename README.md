# Random name generator

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
  -t, --types strings      Types to include (default [colours,dogs,metals,trees,gladiators,greek])
```

## Adding more words

Drop a file in `data` , format is one per line, file must end with `.txt` , run make to rebuild the binary with the new data file.
