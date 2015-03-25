# portscan
a simple portscanner written in go

# example
```
$ go build portscan.go 
$ ./portscan -host=google.de
open 80
scan finished
```

# arguments
```
$ ./portscan -h
Usage of ./portscan:
  -closed=false: list closed ports (true/false)
  -end=-1: the upper end to scan
  -host="localhost": the host to scan
  -pause="1ms": pause after each port scan (e.g. 5ms)
  -start=80: the lower end to scan
  -timeout="1000ms": timeout (e.g. 50ms or 1s)
  ```
