# Generate
Generate go file from FlatBuffers schema
```
$ flatc --go --gen-object-api schema/fbs/user.fbs
```

Generate testdata
```
$ go run cmd/generate.go 
```

# Run
```bash
$ go run main.go -b=false
Alloc = 180 KB
Alloc = 209 KB
Alloc = 211 KB
```

```bash
$ go run main.go -b=true
Alloc = 181 KB
Alloc = 4431 KB
Alloc = 4433 KB
```