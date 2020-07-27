# Struct cast bug

Run with:

```
GOARCH=amd64 go run main.go
GOARCH=386 go run main.go
```

and observe the difference. `go-safer` can help identify pitfalls such as this.

