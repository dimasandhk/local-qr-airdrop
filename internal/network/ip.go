package network

import (
	"fmt"
	"net"
)

func GetLocalIP() string {
	ifaces, _ := net.Interfaces()
	var fallbackIP string
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			// Filter: IPv4 only, no loopback, no dummy APIPA
			if ip != nil && ip.To4() != nil && !ip.IsLoopback() && !ip.IsLinkLocalUnicast() {
				fmt.Printf("[%s] -> http://%s:3030\n", i.Name, ip.String())
				if fallbackIP == "" {
					fallbackIP = ip.String()
				}
				if i.Name == "Wi-Fi" || i.Name == "wlan0" || i.Name == "en0" {
					return ip.String()
				}
			}
		}
	}
	return fallbackIP
}
