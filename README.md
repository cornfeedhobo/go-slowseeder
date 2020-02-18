# go-slowseeder [![Go Report Card](https://goreportcard.com/badge/github.com/cornfeedhobo/go-slowseeder)](https://goreportcard.com/report/github.com/cornfeedhobo/go-slowseeder)

Package slowseeder implements a drop-in replacement for a rand source
intended for cryptographic key generation.

It has been designed to be simple to comprehend. Generation is
deterministic from a seed, uses multiple layered hashing functions,
and is parameterized to easily extend the time spent during each
iteration, making brute force and pre-computation more difficult.

## Contributing

This project could use some tests. PRs are welcome.

## Limitations

The way golang consumes entropy while performing cryptographic functions
is subject to change at any time. Normally this wouldn't matter, but because
this creates a deterministic stream, the final product will change as well.

## Is it any good

[Yes](http://news.ycombinator.com/item?id=3067434)
