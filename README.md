# GoatCore
[![CircleCI](https://circleci.com/gh/goatcms/goatcore.svg?style=svg)](https://circleci.com/gh/goatcms/goatcore)
[![Go Report Card](https://goreportcard.com/badge/github.com/goatcms/goatcore)](https://goreportcard.com/report/github.com/goatcms/goatcore)
[![GoDoc](https://godoc.org/github.com/goatcms/goatcore?status.svg)](https://godoc.org/github.com/goatcms/goatcore)
[![codecov.io](https://codecov.io/github/goatcms/goatcore/coverage.svg?branch=master)](https://codecov.io/github/goatcms/goatcore?branch=master)

GoatCore is a set od main goat packages. It includes interfaces, basic libraries and startup types.

## Improve golang compilation speed
```
go install -a github.com/mattn/go-sqlite3
```
This command install the package into your $GOPATH.
Right now, you likely have an older version installed under $GOPATH/pkg/ and therefore Go is recompiling it for every build.

## Improve test execution time
You can set *AsyncTestReapeat* at workers\\main.go to lower value like 200 or 500. Tests will less-restrictic for simultaneously processing errors. A good site is faster tests execution.

## About
* GoatCore - Szybki Start. https://youtu.be/tqVxzNiJT5g (PL) [Slides](https://docs.google.com/presentation/d/1S0UoP-Js6r7FJxglgSql9kLkRjwFCm7D-ossH4Lz4UA/edit#slide=id.p)
* https://www.youtube.com/watch?v=DUKq5WMz4Y8 (PL) [Slides](https://docs.google.com/presentation/d/1i6e8XM8zZ5FsxIAEqxYjYziafBZt7N-ADYtKY5ENsVc/edit#slide=id.p)
* P.I.W.O (Poznańska Impreza Wolnego Oprogramownaia) - 2017
[Nagranie (PL)](https://www.youtube.com/watch?v=r5etsT7r5No) [Slides](https://docs.google.com/presentation/d/1i6e8XM8zZ5FsxIAEqxYjYziafBZt7N-ADYtKY5ENsVc/edit#slide=id.p)
