# Information Leak POC

This proof of concept shows how casting buffers of differing lengths using `unsafe.Pointer` can lead to
a buffer overflow, resulting in an information leak vulnerability in this case.


## Vulnerability mechanism

Using the arbitrary type casting privilege of `unsafe.Pointer`, a `[8]byte` array value is converted into
a `[25]byte` value. When read, this leaks additional data from the memory.


## Threat model

A possible threat model is a large software company where different teams work at a client and a server
application, respectively, using a shared communication protocol to exchange data. When the protocol was
agreed upon, the server team printed the architecture diagram and hung it on the office wall. Later,
the teams agreed upon a change of the protocol, but the diagram on the wall was not updated.

When an engineer implements the protocol, they look at the diagram and read the incorrect length for a
field. In code review, all engineers from the team also look at the diagram and verify the false
information. Thus, the type of vulnerability shown here gets merged into a production software.

Further information can be found in the blog post [Exploitation Exercise with unsafe.Pointer in Go: Information Leak (Part 1)](https://dev.to/jlauinger/exploitation-exercise-with-unsafe-pointer-in-go-information-leak-part-1-1kga)


## Execute POC

To run this proof of concept, execute the following command:

```
go run main.go
```

Expected output:

```
AAAAAAAAl33t-h4xx0r-w1ns!
```

