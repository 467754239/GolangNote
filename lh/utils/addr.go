package utils

import (
	"net"
)

func LocalAddrpool() ([]string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	var addrpool []string
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				addrpool = append(addrpool, ipnet.IP.String())
			}
		}
	}
	return addrpool, nil
}
