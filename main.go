package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: local-qr-airdrop <path-to-file-or-folder>")
		os.Exit(1)
	}

	inputPath := os.Args[1]

	// filepath.Abs automatically handles relative vs absolute paths
	absPath, err := filepath.Abs(inputPath)
	if err != nil {
		log.Fatalf("Error resolving path: %v", err)
	}

	info, err := os.Stat(absPath)
	if os.IsNotExist(err) {
		log.Fatalf("Error: The path '%s' does not exist.", absPath)
	} else if err != nil {
		log.Fatalf("Error accessing path: %v", err)
	}

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	localIP := getLocalIPs()
	fmt.Println("========================================")
	if info.IsDir() {
		fmt.Printf("🎯 Target Directory : %s\n", absPath)
	} else {
		fmt.Printf("🎯 Target File      : %s\n", absPath)
	}
	fmt.Printf("🚀 Accessible via   : http://%s:3030\n", localIP)
	fmt.Println("========================================")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Serving: " + absPath)
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


