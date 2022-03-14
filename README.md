# Golang Static Build



## The basics

Run the go build command to compile the code into an executable file:

```golang
go build -o app main.go
```

From the command line in the project directory, execute the app binary which was just created:

```golang
./app
```

## Cross compiling

With Go you can build the same application for running in different operating systems. This can be achieved by simply providing the required environment variables.

```golang
$ GOOS=darwin GOARCH=amd64 go build -o app-amd64-darwin *.go
```

Depending on the Go code, there might be missing some dependencies, if `CGO` is required, then additional packages musted be installed before cross compiling.

## Static builds

Sometimes different versions of libc, net libraries, missing required system files would cause some issues, if the static build is not handled properly.

* for this example `CGO` is not required, so disable it by setting the environment variable `CGO_ENABLED` to `0`
* to force a rebuild of packages that are already up-to-date, use the go build flag `-a`
* make sure to use the built-in packages and not the system by tagging with `netgo`
* the `ldflags -w` disables **DWARF debugging information** making the file be smaller

```golang
CGO_ENABLED=0 go build \
    -a \
    -tags netgo \
    -ldflags '-w' \
    -o app main.go
```

```golang
go build -a -tags netgo,timetzdata -ldflags '-w -extldflags "-static"' -o app *.go
```