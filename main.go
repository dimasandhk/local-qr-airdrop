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
)

func main() {
	port := flag.Int("port", 3030, "Port to run the local server on")

	flag.Usage = func() {
		fmt.Println("🚀 local-qr-airdrop - Share files instantly to your phone via QR code")
		fmt.Println("\nUsage:")
		fmt.Println("  local-qr-airdrop [options] <path-to-file-or-folder>")
		fmt.Println("\nOptions:")
		flag.PrintDefaults()
		fmt.Println("\nExamples:")
		fmt.Println("  local-qr-airdrop ./secrets.txt        # Share a single file")
		fmt.Println("  local-qr-airdrop ./my-folder          # Share a whole folder")
		fmt.Println("  local-qr-airdrop -port 8080 ./file.go # Run on port 8080")
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

	localIP := network.GetLocalIP(*port)

	serverURL := fmt.Sprintf("http://%s:%d", localIP, *port)

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

	log.Fatal(app.Listen(fmt.Sprintf("[::]:%d", *port)))
}
