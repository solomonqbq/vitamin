package network

import (
	"fmt"
	"net"
	"strconv"
	"testing"
)

func TestIP(t *testing.T) {
	fmt.Println(GetLocalIP())
}

func TestScanner(t *testing.T) {
	addrs := make([]*net.TCPAddr, 0)
	addrs = append(addrs, &net.TCPAddr{IP: net.ParseIP("192.168.15.64"), Port: 3301})
	addrs = append(addrs, &net.TCPAddr{IP: net.ParseIP("192.168.15.64"), Port: 3306})
	addrs = append(addrs, &net.TCPAddr{IP: net.ParseIP("192.168.1.1"), Port: 3306})
	result := Scanner(addrs)
	for index, r := range result {
		fmt.Printf("ping %s result is %s\n", addrs[index].IP, strconv.FormatBool(r))
	}
}
