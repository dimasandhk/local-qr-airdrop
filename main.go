package main

import (
	"log"
	"net"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	localIP := getLocalIPs()
	fmt.Println("accessible via http://" + localIP + ":3030")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! from " + localIP)
	})

	log.Fatal(app.Listen("[::]:3030"))
}

func getLocalIPs() string {
	ifaces, _ := net.Interfaces()
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
				if i.Name == "Wi-Fi" {
					return ip.String()
				}
			}
		}
	}
	return ""
}


