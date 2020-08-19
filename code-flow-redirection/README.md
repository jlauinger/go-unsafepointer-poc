# Code Flow Redirection POC

This proof of concept shows how an incorrect array cast can lead to a code flow redirection
vulnerability, executing a function built into the binary.


## Vulnerability mechanism

Using `unsafe.Pointer`, a fixed-length array is converted to a wrong length. Afterwards, data
is written to the array, resulting in a buffer overflow. The data overwrites the stored return
address on the stack with the address of the `win` function. When executed, the `win` function
prints a message before the program crashes due to stack corruption.


## Threat model

This type of vulnerability can be introduced if `unsafe` operations are used for efficiency when
decoding a client/server protocol, and the length of a specific field is supplied by an attacker,
e.g. because it is encoded in the protocol.

Further information can be found in the blog post [Exploitation Exercise with Go unsafe.Pointer: Code Flow Redirection (Part 2)](https://dev.to/jlauinger/exploitation-exercise-with-go-unsafe-pointer-code-flow-redirection-part-2-5hgm)


## Execute POC

To run this proof of concept code, execute the following command:

```
go run main.go
```

Expected output:

```
win!
unexpected fault address 0x0
fatal error: fault
[signal SIGSEGV: segmentation violation code=0x80 addr=0x0 pc=0x49ba5e]

goroutine 1 [running]:
...
```

