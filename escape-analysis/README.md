# Escape Analysis Flaw POC

This proof of concept shows how a common, insecure casting pattern for slice types leads to a
flawed escape analysis that creates a dangling pointer vulnerability.


## Vulnerability mechanism

In the `main` function, a string is obtained from a function and then printed to the standard output.
The `GetString` function declares a constant byte slice and then uses a common, insecure casting
pattern to cast it into a string value in-place. The resulting string is then printed to standard
output and returned to the caller.

Because the insecure slice cast creates a string header from scratch instead of by deriving it from
a real string, Go escape analysis fails to see a connection between the `[]byte` parameter to
`BytesToString` and its `string` return value. Therefore, when the `[]byte` slice is created in
`GetString`, Go escape analysis infers that it does not escape because it does not in `BytesToString`
and is never used afterwards. Thus, it is placed on the stack.

The resulting string value has the same underlying data array because it was created by an in-place
cast. It's data is therefore located on the stack of `BytesToString`. The print in that function
succeeds because the stack exists at that point, but when `BytesToString` returns the resulting
string, its stack is destroyed and therefore a dangling pointer gets returned. The print in the
`main` function reads arbitrary memory where the stack once was.

Further information can be found in the blog post [SliceHeader Literals in Go create a GC Race and Flawed Escape-Analysis. Exploitation with unsafe.Pointer on Real-World Code](https://dev.to/jlauinger/sliceheader-literals-in-go-create-a-gc-race-and-flawed-escape-analysis-exploitation-with-unsafe-pointer-on-real-world-code-4mh7)


## Threat model

This code pattern is taken from real-world code. It is very common and used by very large projects.
Use our linter tool [go-safer](https://github.com/jlauinger/go-safer) to identify usages of this
pattern.


## Execute POC

To run this proof of concept code, execute the following command:

```
go run main.go
```

Expected output:

```
GetString:abcdefgh
main:
```

