# inetaddr


### Goal
`inetaddr` is a tool that allows get external IP address  simply, faster and securely:

- allows get IP address from many DNS servers
- allows get IP address from many HTTP servers

### Install
```shell
go get github.com/basti0nz/inetaddr
```

### Example

```go

package main

import (
	"fmt"
	"github.com/basti0nz/inetaddr"
)

const (
	OnlyDNS = iota
	OnlyHTTP
	BothMethod
)

func main() {
	ip, err := inetaddr.IpAddrString(OnlyDNS)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip)

	ipIP, err := inetaddr.IpAddrIP(OnlyHTTP)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ipIP.String())

	ip, err = inetaddr.IpAddrString(BothMethod)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip)
}
```

## TODO:

- use proxy

## License

MIT

### Issues
Please open *issues* here: [New Issue](https://github.com/basti0nz/inetaddr/issues)

### Suggestions and improvements are welcome.

-Valentyn Nastenko(c 2021) https://github.com/versus