# Slice Cast GC Race POC

This proof of concept shows how a common, insecure casting pattern for slice types leads to a
garbage collector (GC) race condition that causes a use-after-free vulnerability.


## Vulnerability mechanism

One Goroutine (thread) continuously allocates and frees heap memory, thus increasing the pressure
on the garbage collector. It will run when the heap size doubles since the last run, therefore the
GC will run very often in this POC.

The main Goroutine concurrently and repeatedly reads a line from the standard input, remembers the
first character in a separate variable, converts the string to a `[]byte` slice, then reads another
line into a string. If the first byte in the resulting `[]byte` slice is different from the stored
first byte in the string, the POC prints a `win` message.

The conversion between `string` and `[]byte` slice is done using an insecure in-place casting pattern:
`unsafe.Pointer` is used to convert the string to its internal representation of length and address
of its underlying data, which is available through the `reflect` package. Then, a slice header object
is created from scratch, the data is copied over, and the new header is cast to a `[]byte` slice
using `unsafe.Pointer` again.

The problem is that the garbage collector does not treat the address of the underlying data array of
the newly created slice header as a reference type because it is of type `uintptr` (non-reference)
and the header is created from scratch instead of being derived from a real slice value. Thus, if
the garbage collector runs exactly between creating the slice header and casting the new header to
an actual slice value, then the underlying array is freed and we have a use-after-free vulnerability.
This is a race condition against the GC.

In this POC, the freed memory is reused by the second line that is read from standard input, but in
a real-world scenario this could be arbitrary data including fields of a different, concurrent
Goroutine, and following code could also write the data at the invalid slice.

Further information can be found in the blog post [SliceHeader Literals in Go create a GC Race and Flawed Escape-Analysis. Exploitation with unsafe.Pointer on Real-World Code](https://dev.to/jlauinger/sliceheader-literals-in-go-create-a-gc-race-and-flawed-escape-analysis-exploitation-with-unsafe-pointer-on-real-world-code-4mh7)


## Threat model

This code pattern is taken from real-world code. It is very common and used by very large projects.
Use our linter tool [go-safer](https://github.com/jlauinger/go-safer) to identify usages of this
pattern.

Repeated execution is possible e.g. if the insecure casting pattern is included in a code path that
gets executed upon generating a server response for a request, because then an attacker can flood
the server with requests. If the program crashes, it would usually be restarted by a daemon supervisor,
so that case does not mitigate the risk.


## Execute POC

To run this proof of concept code, execute the following command:

```
go build main.go
./exploit.py | ./main
```

Expected output:

```
win! after 676 iterations
```

