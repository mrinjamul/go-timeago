# go-timeago

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/pkgsite.svg)](https://pkg.go.dev/github.com/mrinjamul/go-timeago)
[![GoReportCard](https://goreportcard.com/badge/github.com/mrinjamul/go-timeago)](https://goreportcard.com/report/github.com/mrinjamul/go-timeago)
[![Code style: Standard](https://img.shields.io/badge/code%20style-Standard-blue.svg)]()
[![License: BSD-3 ](https://img.shields.io/badge/License-BSD%203-blue.svg)](https://github.com/mrinjamul/go-timeago/blob/main/LICENSE)

## What is this?

A go module to calculate the time difference between two dates and return the value in words.

No third-party library used.

## How to use this module?

To use this module,

```sh
go get github.com/mrinjamul/go-timeago
```

and then,

```go
import (
    ...
    "github.com/mrinjamul/go-timeago"
)
```

### Examples

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mrinjamul/go-timeago"
)

func main() {
	time := time.Date(2021, 0, 0, 0, 0, 0, 0, time.Local)
	msg, err := timeago.FormatNow(time)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(msg)
}
```

Output:

    4 months ago

## License

This go module is licensed under BSD 3-Clause Copyright © 2021 Injamul Mohammad Mollah <mrinjamul@gmail.com>

## Thanks

This go package is inspired by [mrasif/timeago](https://github.com/mrasif/timeago) written in Java.
