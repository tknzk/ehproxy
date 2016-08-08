# ehproxy

Extend Request Header Proxy powered by goproxy

# Usage

```
exproxy -port=:8888 -host=127.0.0.1:3000 --extend-headers=x_extend_proxy:exproxy,x_extend_header:exetend_value
```

```
Usage of ehproxy:
  -extend-headers string
    default extend header, 'GoProxy:go-proxy (default "GoProxy:go-proxy")
  -host string
    default target host, '127.0.0.1:3000' (default "127.0.0.1:3000")
  -port string
    default server port, ':8888', ':8080', ':80'... (default ":8888")
  -verbose
    default verbose, 'false'
```

# Instrallation

```
go get -u github.com/tknzk/ehproxy
```

# License

link:./LICENSE[The MIT License].


# Author

tknzk <tkkm.knzk@gmail.com>
