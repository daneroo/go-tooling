# Go Tooling in Action Notes
Great video on golang tooling by [Francesc Campoy](https://campoy.cat/) from the golang team.

## Summary
installation:
```bash
brew install go-delve/delve/delve
go get -u gopkg.in/alecthomas/gometalinter.v1
gometalinter.v1 --install
```

## References
[Youtube Video on Tooling](https://www.youtube.com/watch?v=uBjoTxosSys&index=6&list=WL)

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