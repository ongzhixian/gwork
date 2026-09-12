# gwork

Go monorepo workspace

## Working with Modules

Assuming you are in C:\src\gwork, and you want to run your application like so:
`go run .\sample-hello\`

You need to:
1.  Make sure sample-hello has go.mod 
    Done by going to directory and running something like:

```cli
go mod init sample/hello
```

2.  Add to workspace (go.work)

```cli
go work use .\sample-hello\
go work sync
```

Note: You can always run using a command-line `go run .\sample-cli\main.go`
      The downside is that you have remember the entry point file.
      But this presents no great difficulty if the we always stick to `main.go` 
      as the entry point.


## Directories

sample-* are meant to be sample proof-of-concept applications use for learning certain 
aspects of Go.
