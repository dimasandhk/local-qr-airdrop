package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/atotto/clipboard"
	"github.com/dimasandhk/local-qr-airdrop/internal/network"
	"github.com/dimasandhk/local-qr-airdrop/internal/terminal"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 3030, "Port to run the server on")
	flag.IntVar(&port, "p", 3030, "Port to run the server on (shorthand)")

	flag.Usage = func() {
		fmt.Println("Usage: local-qr-airdrop [options] <path-to-file-or-folder>")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	inputPath := flag.Arg(0)

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

	app.Use(logger.New(logger.Config{
		Format:     "🕒 ${time} | 📱 ${ip} | 📝 ${status} - ${method} ${path}\n",
		TimeFormat: "15:04:05",
	}))

	localIP := network.GetLocalIP(port)

	serverURL := fmt.Sprintf("http://%s:%d", localIP, port)

	fmt.Println("========================================")
	if info.IsDir() {
		fmt.Printf("🎯 Target Directory : %s\n", absPath)
	} else {
		fmt.Printf("🎯 Target File      : %s\n", absPath)
	}
	fmt.Printf("🚀 Accessible via   : %s\n", serverURL)

	// Attempt to copy the URL to the clipboard, silently ignore errors
	err = clipboard.WriteAll(serverURL)
	if err == nil {
		fmt.Println("📋 URL copied to clipboard!")
	}

	fmt.Println("========================================")

	terminal.PrintQRCode(serverURL)

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

	log.Fatal(app.Listen(fmt.Sprintf("[::]:%d", port)))
}
