package websock

import (
	"net"
)

var ip string

func GetIP() string {
	netInterfaceaddr, _ := net.InterfaceAddrs()

	for _, addr := range netInterfaceaddr {
		netIP, ok := addr.(*net.IPNet)

		if ok && !netIP.IP.IsLoopback() && netIP.IP.To4() != nil {
			ip = netIP.IP.String()
			//fmt.Println(ip)
		}
	}
	return ip
}
