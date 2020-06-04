# openfaas-gin-template
Golang templates for OpenFaaS using HTTP frame https://github.com/gin-gonic/gin

this template default enable go module, `GO111MODULE=on`


## usage

```Go
$ faas template pull https://github.com/chennqqi/golang-gin
$ faas new --lang golang-http <fn-name>
```


## advantages
- use gin as http route, we can handle more api in one function
- a simple frame to write golang http functions

## disadvantages
- can't catch http context.Done 