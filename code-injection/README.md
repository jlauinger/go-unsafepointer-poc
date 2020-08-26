# Code Injection POC

This proof of concept shows how an array is cast to a slice without proper length checks, thus
creating a buffer overflow that is used to inject arbitrary code by spawning a shell using
return-oriented programming (ROP).


## Vulnerability mechanism

A buffer of length 8 bytes is cast into a `[]byte` slice with improper length configuration using
the `unsafe` package and slice header representation from the `reflect` package. Then, using a
`bufio` reader data is taken from the standard input and written to the slice, causing a buffer
overflow. This is a size-constrained version of a classic `gets()`-based buffer overflow
vulnerability.

The `exploit_rop.py` file runs the program and supplies a carefully crafted input that overwrites
the stack up to the point where the stored return address is located. It gets overwritten with
a series of ROP gadgets, that is small fragments of assembly code that end with a `ret` instruction,
thus chaining the gadgets is as simple as concatenating their addresses.

The exploit payload uses the `mprotect` syscall to mark a memory region as `rwx`, then uses the
`read` syscall to write data to that memory region. It supplies assembly code that spawns a shell
using the `system` syscall to the `read` syscall, and finally jumps to the memory region. The
`mprotect` call prevents DEP, and because Go binaries are statically linked to a rather large binary
that contains lots of ROP gadgets ASLR is not effective.

Further information can be found in the blog post [Exploitation Exercise with Go unsafe.Pointer: ROP and Spawning a Shell (Part 3)](https://dev.to/jlauinger/exploitation-exercise-with-go-unsafe-pointer-rop-and-spawning-a-shell-part-3-4mm7)


## Threat model

This type of buffer overflow can happen when fixed-length data is cast to a slice without proper
checks for matching lengths, using the `unsafe` package for efficiency reasons.


## Execute POC

To run this proof of concept code, execute the following commands:

```
go build main.go
./exploit_rop.py
```

You need the pwntools package for Python 2. You need to compile with the Go compiler version go1.15 linux/amd64
to make sure the ROP gadget addresses align. If you use a different version then the addresses might be different.
The blog post explains how to use Ropper to find the gadgets so you can update the addresses.

Expected output:

```
[+] Starting local process './main': pid 75369
[*] Switching to interactive mode
$ id
uid=1000(johannes) gid=1000(johannes) groups=1000(johannes),54(lock),1001(plugdev)
$
```

