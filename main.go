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
		BodyLimit:             100 * 1024 * 1024, // 100 MB body limit for uploads
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
	fmt.Printf("📤 Upload route     : %s/upload\n", serverURL)

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

	app.Get("/upload", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		return c.SendString(`
<!DOCTYPE html>
<html>
<head>
	<title>Upload File</title>
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<style>
		body { font-family: sans-serif; padding: 2rem; max-width: 600px; margin: auto; }
		.btn { padding: 0.5rem 1rem; background: #007bff; color: white; border: none; cursor: pointer; border-radius: 4px;}
	</style>
</head>
<body>
	<h2>Upload a File</h2>
	<form action="/upload" method="POST" enctype="multipart/form-data">
		<input type="file" name="file" required style="margin-bottom: 1rem; display: block;" />
		<button type="submit" class="btn">Upload</button>
	</form>
</body>
</html>
		`)
	})

	app.Post("/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("Upload failed: %v", err))
		}

		// Sanitize the filename to prevent path traversal
		filename := filepath.Base(file.Filename)

		var targetDir string
		if info.IsDir() {
			targetDir = absPath
		} else {
			targetDir = filepath.Dir(absPath)
		}

		destPath := filepath.Join(targetDir, filename)

		// Save the file
		if err := c.SaveFile(file, destPath); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Could not save file: %v", err))
		}

		return c.SendString(fmt.Sprintf("File %s uploaded successfully!", filename))
	})

	log.Fatal(app.Listen(fmt.Sprintf("[::]:%d", port)))
}
