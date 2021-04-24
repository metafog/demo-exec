
```
docker run -d planetrio/demo-exec
```

## Docker Build
```
GOOS=linux GOARCH=386 go build ./src-reverse/reverse.go
GOOS=linux GOARCH=386 go build ./src-capitalize/capitalize.go 
docker build -t demo-exec .
```
