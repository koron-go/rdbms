# koron-go/rdbms

[![GoDoc](https://godoc.org/github.com/koron-go/rdbms?status.svg)](https://godoc.org/github.com/koron-go/rdbms)
[![CircleCI](https://img.shields.io/circleci/project/github/koron-go/rdbms/master.svg)](https://circleci.com/gh/koron-go/rdbms/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/koron-go/rdbms)](https://goreportcard.com/report/github.com/koron-go/rdbms)

Discern kind of RDBMS from the driver (`*sql.DB`).

```go

import (
    "database/sql"
    "github.com/koron-go/rdbms"
)

db := sql.Open("foobar://...")
k, err := rdbms.Detect(db)
// TODO: work with k (and err)
```

## Support more RDBMS

### More drivers

Add a `Driver` instance (pair of path prefix and kind of RDBMS) to
`knownDrivers` in driver.go file.

### More kinds of RDBMS

Add a `RDBMS` instance as const in rdbms.go, then add drivers.

## Misc

*   [Drivers](https://github.com/golang/go/wiki/SQLDrivers)
