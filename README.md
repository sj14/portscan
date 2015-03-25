# portscan
a simple portscanner written in go

# examples
```
$ ./portscan -host=google.de
open 80
scan finished
```

```
$ ./portscan.go -start=1656 -end=1740 -timeout=250ms
open 1656
open 1660
open 1657
scan finished in 109.437457ms
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
