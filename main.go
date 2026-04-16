package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/dimasandhk/local-qr-airdrop/internal/network"
	"github.com/dimasandhk/local-qr-airdrop/internal/terminal"
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

	localIP := network.GetLocalIP()
	
	fmt.Println("========================================")
	if info.IsDir() {
		fmt.Printf("🎯 Target Directory : %s\n", absPath)
	} else {
		fmt.Printf("🎯 Target File      : %s\n", absPath)
	}
	fmt.Printf("🚀 Accessible via   : http://%s:3030\n", localIP)
	fmt.Println("========================================")

	terminal.PrintQRCode("http://" + localIP + ":3030")

	// Serve either a single file or a whole directory based on user input
	if info.IsDir() {
		// Serve all files inside the directory
		app.Static("/", absPath, fiber.Static{
			Browse: true, // Enables a built-in file browser UI
		})
	} else {
		// Serve the single file directly on the root path
		app.Get("/", func(c *fiber.Ctx) error {
			return c.SendFile(absPath)
		})
	}

	log.Fatal(app.Listen("[::]:3030"))
}
