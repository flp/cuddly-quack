# cuddly-quack
Purpose: satisfy the drafting fix.

## Get Started

To get started, first install the [`go`](http://golang.org) programming language.

Once installed, you can use the `go` tool to get the project:

```
$ go get github.com/flp/cuddly-quack
```

At this point, the app should be available in your `$GOPATH/src/github.com/flp/cuddly-quack` directory.  You can then `cd` into that directory and build the app using `make`:

```
$ cd $GOPATH/src/github.com/flp/cuddly-quack
$ make
```

At this point, you should have a `cuddly-quack` binary in the root directory of the app.  You can run the server by executing the binary:

```
$ ./bin/cuddly-quack
```

The server should be running on port `:8000`.