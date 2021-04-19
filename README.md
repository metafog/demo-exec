# Planetr Demo: demo-exec
Run a 'docker exec' command to execute command line applications

## Docker usage on local machine 
```
docker run -d planetrio/demo-exec
```

## Deploy and run the same docker image on Planetr decentralized cloud

```
$ planetr dcu-run planetrio/demo-exec g.micro
INSTANCE ID STATUS TYPE IMAGE NAME PORTS CREATED AT
c1tafjn2hraq06kud0bg Pending g.micro planetrio/demo-exec 2021-04-17T14:48:06.568676+05:30
```

Please wait for the status to become 'Running'

```
$ planetr dcu-ps
INSTANCE ID STATUS TYPE IMAGE NAME PORTS CREATED AT
c1tafjn2hraq06kud0bg Running g.micro planetrio/demo-exec 2021-04-17T14:48:06.568676+05:30
```

```planetrio/demo-exec``` docker image has two commands ```reverse``` and ```capitalize``` in ```/home``` folder

```
$ planetr dcu-exec c1tatan2hraq6c4g3rlg '/home/reverse "TOM AND JERRY"'
YRREJ DNA MOT
```

Delete the DCU.

```
$ planetr dcu-rm c1tatan2hraq6c4g3rlg
INSTANCE ID STATUS TYPE IMAGE NAME PORTS CREATED AT
c1tatan2hraq6c4g3rlg Deleting g.micro planetrio/demo-exec 2021-04-17T14:48:06.568676+05:30
```

Check status again.

```
$ planetr dcu-ps
INSTANCE ID STATUS TYPE IMAGE NAME PORTS CREATED AT
```

## Docker Build (If in case)
```
GOOS=linux GOARCH=386 go build ./src-reverse/reverse.go
GOOS=linux GOARCH=386 go build ./src-capitalize/capitalize.go 
docker build -t demo-exec .
```
