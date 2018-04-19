## getopt(3) for go

I get what `flag` is doing, and it's pretty, but it's also a lot more
typing than I want in simple cases.

The `getopt(3)` interface is clunky in some ways related to C's
design; it relies on being called repeatedly, and sentinel values,
because allocation is expensive and there's no multiple returns.

But I really like the expressiveness of the option string, for the
common case with single-letter options, and I like the ability to
combine short options.

### Invocation

`opts, remaining, err := getopt.GetOpt(args, optstring)'

Takes a slice of strings and an option string, returns a map of provided
options, remaining arguments as a slice of strings, and an error if
an error occurred.

```
type Option {
    Value   string
    Int     int64
    Float   float64
}

type Options map[string]*Option
```

### Option Strings

An option string is a list of allowed options. Letters are
allowed options. A letter followed by punctuation indicates
an option which can take an additional parameter as a value:

	:	Any string
	#	Integer number
	.	Floating-point number

Flags which take a value consume additional arguments
starting immediately after the argument containing the flag
itself.  Thus, given `getopt.GetOpt(args, "a:b:c"):

    "-ab" "foo" "bar" "-c" "baz"

yields:

    a: foo
    b: bar
    c: baz

The special option `--` indicates the end of option parsing,
as does any non-option encountered.

## Future Directions

I'm considering cool features like allowing additional
parameters to `GetOpt` which are used to specify where to
store values, but in the past I've found that this was
not especially *useful* to me even though I always think
it sounds neat, so I'm not doing it yet.
