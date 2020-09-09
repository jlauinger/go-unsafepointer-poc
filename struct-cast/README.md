# Architecture dependent types cast POC

This proof of concept shows how architecture-dependent types within a struct type can cause
alignment issues and thus buffer overflow vulnerabilities when structs are casted in-place using
`unsafe`.


## Vulnerability mechanism

`PinkStruct` and `VioletStruct` share the same types except field `B` which has type `int` or
`int64` respectively. Only the `int` type is platform dependent. Therefore, casting instances
of one struct to the other type can work on some platforms but not on others. Notably, casts
work on 64-bit platforms.

In the POC, an instance of `PinkStruct` is cast in-place to type `VioletStruct` using
`unsafe.Pointer`. Then, the fields are printed to standard output. When the fields misalign,
arbitrary memory data is printed.


## Threat model

This vulnerability can be introduced when developers do not carefully test their code on all
target platforms, or a target platform is added to the Go platform later on.


## Execute POC

To run this proof of concept, execute the following commands:

```
GOARCH=amd64 go run main.go
GOARCH=386 go run main.go
```

Expected output:

```
1
42
9000
---
1
38654705664042
702917787932164096
```

