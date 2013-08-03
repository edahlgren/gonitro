INSTALL
----------
Requires >= go1.1.*

(if you don't have nitro installed)

```git clone git@github.com:bumptech/nitro.git```
* (install libsodium https://github.com/jedisct1/libsodium/releases)
* (install redo https://github.com/apenwarr/redo)
* (install libev dev libraries)
* ```redo```
* ```sudo redo install```

```
go get github.com/edahlgren/gonitro
go install github.com/edahlgren/gonitro
```

RUN EXAMPLE
----------
```
go install github.com/edahlgren/gonitro/example
$GOPATH/bin/example
```

RUN TESTS
----------
```
go install github.com/edahlgren/gonitro/test
$GOPATH/bin/test
```
