# No Hassle Unique Identifiers in Go (golang)

[![Go Report Card](https://goreportcard.com/badge/gitlab.com/rwxrob/uniq)](https://goreportcard.com/report/gitlab.com/rwxrob/uniq) [![Coverage](https://gocover.io/_badge/gitlab.com/rwxrob/uniq)](https://gocover.io/gitlab.com/rwxrob/uniq) [![GoDoc](https://godoc.org/gitlab.com/rwxrob/uniq?status.svg)](https://godoc.org/gitlab.com/rwxrob/uniq)

`go get gitlab.com/rwxrob/uniq`

Package `uniq` is a utility package that provides common random unique identifiers in UUID, Base32, and n*2 random hexadecimal characters.

    6c671957-2f39-4ce5-9f0e-e8d5ec53bfde (16 bytes, 36 chars, hex-)
    H6M0STKP0MTSU0493GERQDCSJ5BMF3VO     (20 bytes, 32 chars, base32)
    5b ...                               (n bytes, n*2 chars, hex)

When a simple random identifier is all that is needed `Base32()` provides a better alternative to `UUID()`. It takes less space (32 characters), is safe for use with all filesystems including case insensitive ones, and provides additional randomness increased from 2^128 (uuid) to 2^160 (base32).

This package includes the following convenience commands as well for use when integrating with shell scripts:

* `uuid`
* `uid32`
* `randhex [COUNT]`
