package network

import (
	"fmt"
	"net"
	"strings"
)

func GetLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println(err.Error())
		return GetFirstLocalIP()
	}
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0]
}

func GetFirstLocalIP() string {
	addrs, _ := GetAllAddrs()
	if addrs != nil && len(addrs) != 0 {
		return addrs[0]
	} else {
		return "127.0.0.1"
	}
}

func GetAllAddrs() (addrs []string, err error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	addrs = make([]string, 0)
	for _, inter := range interfaces {
		all_addrs, _ := inter.Addrs()
		for _, addr := range all_addrs {
			str := strings.Split(addr.String(), "/")
			if str[0] != "127.0.0.1" && str[0] != "::1" {
				addrs = append(addrs, str[0])
			}
		}
	}
	return
}

func GetLocalMac() (macs []string, err error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	macs = make([]string, 0)
	for _, inter := range interfaces {
		addrs, _ := inter.Addrs()
		for _, addr := range addrs {
			fmt.Println(addr)
		}

		mac := inter.HardwareAddr.String()
		if mac != "" {
			macs = append(macs, mac)
		}
	}
	return
}
