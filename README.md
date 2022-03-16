# Golang Static Build

## Cross compiling

With Go you can build the same application for running in different operating systems. This can be achieved by simply providing the required environment variables.

```golang
$ GOOS=darwin GOARCH=amd64 go build -o app-amd64-darwin *.go
```

## CGO

Cgo enables the creation of Go packages that call C code.

The cgo tool is enabled by default for native builds on systems where it is expected to work. It is disabled by default when cross-compiling. You can control this by setting the `CGO_ENABLED` environment variable when running the go tool: set it to 1 to enable the use of cgo, and to 0 to disable it. The go tool will set the build constraint "cgo" if cgo is enabled. The special import "C" implies the "cgo" build constraint, as though the file also said "// +build cgo". Therefore, if cgo is disabled, files that import "C" will not be built by the go tool.

When cross-compiling, you must specify a C cross-compiler for cgo to use. You can do this by setting the generic `CC_FOR_TARGET` or the more specific `CC_FOR_${GOOS}_${GOARCH}` (for example, `CC_FOR_linux_arm`) environment variable when building the toolchain using make.bash, or you can set the `CC` environment variable any time you run the go tool.


## Go Packages

### Net

The DNS resolver in the net package has almost always used cgo to access the system interface. A change in Go `1.5` means that on most Unix systems DNS resolution will no longer require cgo, which simplifies execution on those platforms.

The decision of how to run the resolver applies at **run time**, not build time. The netgo build tag that has been used to enforce the use of the Go resolver is no longer necessary, although it still works.

```go
GODEBUG=netdns=1 prints a one-time strategy decision. (cgo or go DNS lookups)
GODEBUG=netdns=2 prints the per-lookup strategy as a function of the hostname.

The new "netcgo" build tag forces cgo DNS lookups.

GODEBUG=netdns=go (or existing build tag "netgo") forces Go DNS resolution.
GODEBUG=netdns=cgo (or new build tag "netcgo") forces libc DNS resolution.

Options can be combined with e.g. GODEBUG=netdns=go+1 or GODEBUG=netdns=2+cgo.
```

### Time

Package tzdata provides an embedded copy of the timezone database. If this package is imported anywhere in the program, then if the time package cannot find tzdata files on the system, it will use this embedded information.

Importing this package will increase the size of a program by about 450 KB.

This package should normally be imported by a program's main package, not by a library. Libraries normally shouldn't decide whether to include the timezone database in a program.

This package will be automatically imported if you build with -tags timetzdata.

### User

For most Unix systems, this package has two internal implementations of resolving user and group ids to names, and listing supplementary group IDs. One is written in pure Go and parses `/etc/passwd` and `/etc/group`. The other is cgo-based and relies on the standard C library (libc) routines such as `getpwuid_r`, `getgrnam_r`, and `getgrouplist`.

When `cgo` is available, and the required routines are implemented in `libc` for a particular platform, cgo-based (libc-backed) code is used. This can be overridden by using osusergo build tag, which enforces the pure Go implementation.

## Static builds

Sometimes different versions of libc, net libraries, missing required system files would cause some issues, if the static build is not handled properly.

* **`-a`** forces a rebuild of packages that are already up-to-date
* **`ldflags -w`** disables **DWARF debugging information** making the file be smaller

```golang
CGO_ENABLED=0 go build \
    -a \
    -ldflags '-w' \
    -o app main.go
```

```golang
go build -a -tags timetzdata -ldflags '-w -extldflags "-static"' -o app *.go
```
