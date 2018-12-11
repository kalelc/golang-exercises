# Golang EX

#### Fetch

###### Build app

`go build main.go`

###### Run app

`./main www.google.com`

###### Response

```
http://www.google.com
200 OK
<!doctype html><html itemscope="" itemtype="http://schema.org/WebPage" lang="es-419"><head>...
```

#### Github/issue

###### Build app

`go build *.go`

###### Run app

`./github repo:golang/go`

###### Response

```
29120 issues:
#29174 eliasnaur 2018-12-11 13:39:50 +0000 UTC x/tools/internal/lsp: handle module paths below root path
#29173 zbindenre 2018-12-11 10:04:02 +0000 UTC Command timeout not triggered for interactive shell command.
#29172  bradfitz 2018-12-10 22:20:21 +0000 UTC x/build/cmd/gopherbot: repeatedly trying to close CLs
#29171    teivah 2018-12-10 19:40:41 +0000 UTC proposal: Go 2: Partial Functions
#29170   fortuna 2018-12-10 17:12:59 +0000 UTC Missing #include in os/user/getgrouplist_darwin.go
#29169   bcmills 2018-12-10 15:21:27 +0000 UTC cmd/go: emit an error in 'go mod vendor' if any replacement is in the vendor dir
...
```
