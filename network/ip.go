package network

import (
	"fmt"
	"github.com/xeniumd-china/vitamin/concurrent"
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

//检测IP端口
func Scanner(addrs []*net.TCPAddr) (result []bool) {
	if addrs == nil {
		return nil
	}
	if len(addrs) == 0 {
		return make([]bool, 0)
	}
	result = make([]bool, len(addrs))
	cdl := concurrent.NewCountDownLatch(len(addrs))
	for index, addr := range addrs {
		go func(index int, addr *net.TCPAddr) {
			defer cdl.CountDown()

			conn, err := net.DialTCP("tcp", nil, addr)
			if err != nil {
				result[index] = false
			} else {
				result[index] = true
			}
			if conn != nil {
				conn.Close()
			}
		}(int(index), addr)
	}
	cdl.Await()
	return result
}
