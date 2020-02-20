# GoatCore
[![Go Report Card](https://goreportcard.com/badge/github.com/goatcms/goatcore)](https://goreportcard.com/report/github.com/goatcms/goatcore)
[![GoDoc](https://godoc.org/github.com/goatcms/goatcore?status.svg)](https://godoc.org/github.com/goatcms/goatcore)

GoatCore is the main project contains dump functions, primitives and base/independent libraries. It includes interfaces, packages, dependency injection system, base application architecture and external modules like terminal, pipelines commands etc.

## Contents
* [Quick start guide](docs/quick_start.md)
* [Architecture](docs/app/architecture.md)

## Structure
* app/* - contains application architecture
* dependency/* - contains dependency injection libraries
* filesystem/* - contains filesystem interfaces and accessor for virtual and operating system filesystem
* goathtml/* - contains loader and provider for **HTML template** (it separate layout views, helpers and layout)
* goatmail/* - SMTP (Simple Mail Transfer Protocol) client
* goattext/* - contains loader and provider for **plain text template** (it separate layout views, helpers and layout)
* i18n/* - provide internationalization libraries
* messages/* - contains forms messages wrapper
* repositories/* - contains libraries for VCS (Version Control System). The git is partly covered now.
* testbase/* - contains test helpers
* workers/* - base library to support your concurrency programming

## Tests

### Concurrency tests accuracy
Improve test execution time by set **AsyncTestReapeat** (at workers/main.go) to lower value (like 200 or 500). Tests will less-restrictive but faster. Increase AsyncTestReapeat to improve test accuracy.

### Run tests
Define the system environment to tests runs.

 * GOATCORE_TEST_SMTP_FROM_ADDRESS - sender email
 * GOATCORE_TEST_SMTP_TO_ADDRESS - reciver email
 * GOATCORE_TEST_SMTP_SERVER - SMTP server URL (like smtp.gmail.com:465)
 * GOATCORE_TEST_SMTP_USERNAME - SMTP server username
 * GOATCORE_TEST_SMTP_PASSWORD - SMTP server password
 * GOATCORE_TEST_SMTP_IDENTITY - SMTP server identity
 * GOATCORE_TEST_DOCKER - Run docker tests if defined and equals to "YES". The tests require docker daemon started.

On Linux/Unix/macOS add environment variables and run go tests by terminal.
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

## Tips&Tricks

### Build library once
To build library "once" you can run:
```
go install -a github.com/goatcms/goatcore
```
This command install the package into your $GOPATH.
It can be source of side effects. Your library version won't be update automatically.
The pre-installed version is installed under $GOPATH/pkg/

You can use the trick to pre-build heavy libraries like go-sqlite3.
```
go install -a github.com/mattn/go-sqlite3
```

## External sources

### In English
In progress

### In Polish
* GoatCore - Szybki Start. https://youtu.be/tqVxzNiJT5g (PL) [Slides](https://docs.google.com/presentation/d/1S0UoP-Js6r7FJxglgSql9kLkRjwFCm7D-ossH4Lz4UA/edit#slide=id.p)
* https://www.youtube.com/watch?v=DUKq5WMz4Y8 (PL) [Slides](https://docs.google.com/presentation/d/1i6e8XM8zZ5FsxIAEqxYjYziafBZt7N-ADYtKY5ENsVc/edit#slide=id.p)
* P.I.W.O (Poznańska Impreza Wolnego Oprogramownaia) - 2017
[Nagranie (PL)](https://www.youtube.com/watch?v=r5etsT7r5No) [Slides](https://docs.google.com/presentation/d/1i6e8XM8zZ5FsxIAEqxYjYziafBZt7N-ADYtKY5ENsVc/edit#slide=id.p)
