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
			html := fmt.Sprintf(`
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Download %s</title>
				<style>
					body { font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif; display: flex; justify-content: center; align-items: center; height: 100vh; margin: 0; background-color: #f4f4f9; color: #333; }
					.container { text-align: center; background: #fff; padding: 2rem; border-radius: 12px; box-shadow: 0 4px 6px rgba(0,0,0,0.1); max-width: 90%%; }
					h1 { font-size: 1.5rem; margin-bottom: 0.5rem; word-break: break-all; }
					p { color: #666; margin-bottom: 1.5rem; }
					.btn { display: inline-block; background: #007bff; color: #fff; text-decoration: none; padding: 10px 24px; border-radius: 6px; font-weight: bold; font-size: 1.1rem; transition: background 0.2s; cursor: pointer; border: none; }
					.btn:hover { background: #0056b3; }
				</style>
			</head>
			<body>
				<div class="container">
					<h1>%s</h1>
					<p>Size: %s</p>
					<a href="/download" class="btn" onclick="setTimeout(() => fetch('/shutdown'), 1000)">Download File</a>
				</div>
			</body>
			</html>
			`, filepath.Base(absPath), filepath.Base(absPath), formatSize(info.Size()))
			c.Set("Content-Type", "text/html")
			return c.SendString(html)
		})

		app.Get("/download", func(c *fiber.Ctx) error {
			c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filepath.Base(absPath)))
			return c.SendFile(absPath)
		})

		app.Get("/shutdown", func(c *fiber.Ctx) error {
			go func() {
				fmt.Println("\n✅ Download completed. Shutting down server...")
				app.Shutdown()
			}()
			return c.SendStatus(200)
		})
	}

	log.Fatal(app.Listen("[::]:3030"))
}

func formatSize(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}
