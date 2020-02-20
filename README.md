# GoatCore
[![Go Report Card](https://goreportcard.com/badge/github.com/goatcms/goatcore)](https://goreportcard.com/report/github.com/goatcms/goatcore)
[![GoDoc](https://godoc.org/github.com/goatcms/goatcore?status.svg)](https://godoc.org/github.com/goatcms/goatcore)

GoatCore is a set od main goat packages. It includes interfaces, basic libraries and startup types.

## Improve golang compilation speed
```
go install -a github.com/mattn/go-sqlite3
```
This command install the package into your $GOPATH.
Right now, you likely have an older version installed under $GOPATH/pkg/ and therefore Go is recompiling it for every build.

## Improve test execution time
You can set *AsyncTestReapeat* at workers\\main.go to lower value like 200 or 500. Tests will less-restrictic for simultaneously processing errors. A good site is faster tests execution.

## Run test
To run test define system environment

 * GOATCORE_TEST_SMTP_FROM_ADDRESS - sender email
 * GOATCORE_TEST_SMTP_TO_ADDRESS - reciver email
 * GOATCORE_TEST_SMTP_SERVER - SMTP server URL (like smtp.gmail.com:465)
 * GOATCORE_TEST_SMTP_USERNAME - SMTP server username
 * GOATCORE_TEST_SMTP_PASSWORD - SMTP server password
 * GOATCORE_TEST_SMTP_IDENTITY - SMTP server identity
 * GOATCORE_TEST_DOCKER - Run docker tests if defined and equals to "YES". The tests require docker daemon started.

On linux/unix/mac add envs by
```
export GOATCORE_TEST_SMTP_FROM_ADDRESS="test@email.com"
export GOATCORE_TEST_SMTP_TO_ADDRESS="test@email.com"
export GOATCORE_TEST_SMTP_SERVER="smtp.gmail.com:465"
export GOATCORE_TEST_SMTP_USERNAME="YOUR_TEST_USER"
export GOATCORE_TEST_SMTP_PASSWORD="YOUR_TEST_USER_PASSWORD"
export GOATCORE_TEST_SMTP_IDENTITY=""
export GOATCORE_TEST_DOCKER="YES"
go test ./...
```

## About
* GoatCore - Szybki Start. https://youtu.be/tqVxzNiJT5g (PL) [Slides](https://docs.google.com/presentation/d/1S0UoP-Js6r7FJxglgSql9kLkRjwFCm7D-ossH4Lz4UA/edit#slide=id.p)
* https://www.youtube.com/watch?v=DUKq5WMz4Y8 (PL) [Slides](https://docs.google.com/presentation/d/1i6e8XM8zZ5FsxIAEqxYjYziafBZt7N-ADYtKY5ENsVc/edit#slide=id.p)
* P.I.W.O (Poznańska Impreza Wolnego Oprogramownaia) - 2017
[Nagranie (PL)](https://www.youtube.com/watch?v=r5etsT7r5No) [Slides](https://docs.google.com/presentation/d/1i6e8XM8zZ5FsxIAEqxYjYziafBZt7N-ADYtKY5ENsVc/edit#slide=id.p)
