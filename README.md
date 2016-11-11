# Go Tooling in Action Notes
[Great video on golang tooling](https://www.youtube.com/watch?v=uBjoTxosSys&index=6&list=WL)
 by [Francesc Campoy](https://campoy.cat/) from the golang team.

[Found this](https://www.youtube.com/watch?v=6gdVhHMxNTo) for Vendoring: use [`glide`](https://glide.sh/)

## Summary
installation:
```bash
# Debug
brew install go-delve/delve/delve

# Linting
go get -u gopkg.in/alecthomas/gometalinter.v1
gometalinter.v1 --install

# Load testing (with pprof instrumentation)
go get github.com/tsliwowicz/go-wrk
# NOT: error in video: go get github.com/adjust/go-wrk

# profiling
go get github.com/uber/go-torch
# extra tools for torch: these need to be in the path though...
cd $GOPATH/src/github.com/uber/go-torch
git clone https://github.com/brendangregg/FlameGraph.git
```

## References
- [Francesc Campoy on Tooling](https://www.youtube.com/watch?v=uBjoTxosSys&index=6&list=WL)
- [Vendoring Deconstructed](https://www.youtube.com/watch?v=6gdVhHMxNTo)

### Basic commands

### Build
```bash
go build
GOOS=linux go build
GOOS=windows go build

go install
```

### Go Get
```bash
go get github.com/golang/example/hello
ls $GOPATH/src/github.com/golang/example/hello/
$GOPATH/bin/hello
```

### List
```bash
go list -f '{{ .Name }}: {{ .Doc }}'
go list -f '{{ .Imports }}'
```

### Doc
```bash
go doc
go doc fmt
# run as local server
godoc -http=:6060
```

### Vet
```bash
go vet
```

### Debugging with `delve`
[Install for OSX](https://github.com/derekparker/delve/blob/master/Documentation/installation/osx/install.md)
Use brew, simpler because of certs thing.
```
brew install go-delve/delve/delve
```

### Testing
integrated in VS-Code (with coverage)

## Perfomance
Using go test Benchmarks
```
go test -bench .
```

Using go-wrk:
```
go-wrk -d 50 http://localhost:8080/daniel@golang.org
```

Using pprof:

Visit: http://localhost:8080/debug/pprof/
```
# while running go-rk in other console...
go tool pprof --seconds=5 http://localhost:8080/debug/pprof/profile
(pprof) top
...
(pprof) web
```

Using go-torch:

Thses need special script (above) to make svg flame graphs..
```
# while running go-wrk in other console...
cd $GOPATH/src/github.com/uber/go-torch
go build
./go-torch -t 5
```
