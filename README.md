# godd 

A simplified implementation of the [`dd`](https://en.wikipedia.org/wiki/Dd_(Unix)) utility written in Go

## Supported command line arguments

### `if`
#### usage -if=FILE.

This argument, if supplied, makes the godd read from FILE instead of stdin.

### `of`
#### usage -of=FILE.

This argument, if supplied, makes the godd write to FILE instead of stdout.

### `bs`
#### usage -bs=BYTES.

This argument, if supplied, makes the godd read and write up to BYTES bytes at a time (the default is 1024).

### `count`
#### usage -count=N.

This argument, if supplied, makes the godd copy only N input blocks.