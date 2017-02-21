# GoatCore
GoatCore is a set od main goat packages. It includes interfaces, basic libraries and startup types.

## Improve golang compilation speed
```
go install -a github.com/mattn/go-sqlite3
```
This command install the package into your $GOPATH.
Right now, you likely have an older version installed under $GOPATH/pkg/ and therefore Go is recompiling it for every build.

## Improve test execution time
You can set *AsyncTestReapeat* at workers\\main.go to lower value like 200/500. Simultaneously tests will less-restrictic and faster.
