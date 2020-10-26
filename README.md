# Go unsafe.Pointer vulnerability POCs

This is a series of proof of concept Go programs showing how the use of `unsafe.Pointer` can lead
to different vulnerabilities. There are seven examples in total.

These examples accompany the blog post series [Exploitation Exercise with unsafe.Pointer in Go](https://dev.to/jlauinger/exploitation-exercise-with-unsafe-pointer-in-go-information-leak-part-1-1kga). The blog series comprises the following parts:

 1. [Information Leak](https://dev.to/jlauinger/exploitation-exercise-with-unsafe-pointer-in-go-information-leak-part-1-1kga)
 2. [Code Flow Redirection](https://dev.to/jlauinger/exploitation-exercise-with-go-unsafe-pointer-code-flow-redirection-part-2-5hgm)
 3. [ROP and Spawning a Shell](https://dev.to/jlauinger/exploitation-exercise-with-go-unsafe-pointer-rop-and-spawning-a-shell-part-3-4mm7)
 4. [SliceHeader Literals in Go create a GC Race and Flawed Escape-Analysis](https://dev.to/jlauinger/sliceheader-literals-in-go-create-a-gc-race-and-flawed-escape-analysis-exploitation-with-unsafe-pointer-on-real-world-code-4mh7)

These blog posts are written as part of my work on my Master's thesis at the [Software Technology Group](https://www.stg.tu-darmstadt.de/stg/homepage.en.jsp) at TU Darmstadt.


## Information leak POC

This proof of concept shows how casting buffers of differing lengths using `unsafe.Pointer` potentially leads
to a buffer overflow, resulting in an information leak in this POC.

A possible threat model to introduce this code pattern is a miscommunication within a software development team.

The exploit code along with instructions to execute it is located in the `information-leak` directory.


## Code flow redirection POC

This proof of concept shows how an incorrect array cast can lead to a code flow redirection
vulnerability, executing a function built into the binary.

A possible threat model is a user-supplied field length in a client/server protocol that gets decoded using unsafe
operations for efficiency.

The exploit code along with instructions to execute it is located in the `code-flow-redirection` directory.


## Code injection POC

This proof of concept shows how an array is cast to a slice without proper length checks, thus
creating a buffer overflow that is used to inject arbitrary code by spawning a shell using
return-oriented programming (ROP).

A possible threat model is a cast of fixed-length data to a slice without correct length checks, or
user-supplied length information in a client/server protocol.

The exploit code along with instructions to execute it is located in the `code-injection` directory.


## Slice cast GC race condition POC

This proof of concept shows how a common, insecure casting pattern for slice types leads to a
garbage collector race condition that causes a use-after-free vulnerability.

The exploit code along with instructions to execute it is located in the `race-slice` directory.


## Escape analysis flaw POC

This proof of concept shows how a common, insecure casting pattern for slice types leads to a
flawed escape analysis that creates a dangling pointer vulnerability.

The exploit code along with instructions to execute it is located in the `escape-analysis` directory.


## Architecture-dependent struct cast POC

This proof of concept shows how architecture-dependent types within a struct type can cause
alignment issues and thus buffer overflow vulnerabilities when structs are casted in-place using
`unsafe`.

The exploit code along with instructions to execute it is located in the `struct-cast` directory.


## go-fuse bug POC

This proof of concepts shows how to exploit a bug that leads to incorrect length information in
a dynamically created slice.

The exploit code along with instructions to execute it is located in the `go-fuse` directory.

