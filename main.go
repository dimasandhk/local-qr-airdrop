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
	var receiveMode bool
	flag.IntVar(&port, "port", 3030, "Port to run the server on")
	flag.IntVar(&port, "p", 3030, "Port to run the server on (shorthand)")
	flag.BoolVar(&receiveMode, "receive", false, "Enable receive mode to upload files to this device")
	flag.BoolVar(&receiveMode, "r", false, "Enable receive mode (shorthand)")

	flag.Usage = func() {
		fmt.Println("Usage: local-qr-airdrop [options] <path-to-file-or-folder>")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if !receiveMode && flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	inputPath := "."
	if flag.NArg() > 0 {
		inputPath = flag.Arg(0)
	}

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

	if receiveMode && !info.IsDir() {
		log.Fatalf("Error: Receive mode requires a directory path, not a file.")
	}

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		BodyLimit:             100 * 1024 * 1024, // 100MB
	})

	app.Use(logger.New(logger.Config{
		Format:     "🕒 ${time} | 📱 ${ip} | 📝 ${status} - ${method} ${path}\n",
		TimeFormat: "15:04:05",
	}))

	localIP := network.GetLocalIP(port)

	serverURL := fmt.Sprintf("http://%s:%d", localIP, port)

	fmt.Println("========================================")
	if receiveMode {
		fmt.Printf("🎯 Receive Dir      : %s\n", absPath)
	} else if info.IsDir() {
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

	if receiveMode {
		app.Get("/", func(c *fiber.Ctx) error {
			c.Type("html")
			return c.SendString(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>local-qr-airdrop - Upload</title>
    <style>
        body { font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif; display: flex; flex-direction: column; align-items: center; justify-content: center; height: 100vh; margin: 0; background-color: #f5f5f7; color: #1d1d1f; }
        .container { background: white; padding: 2rem; border-radius: 12px; box-shadow: 0 4px 6px rgba(0,0,0,0.1); text-align: center; max-width: 400px; width: 90%; }
        h1 { margin-top: 0; font-size: 1.5rem; }
        input[type="file"] { margin: 1.5rem 0; width: 100%; box-sizing: border-box; }
        button { background-color: #0071e3; color: white; border: none; padding: 10px 20px; border-radius: 6px; font-size: 1rem; cursor: pointer; width: 100%; }
        button:hover { background-color: #0077ed; }
        #status { margin-top: 1rem; font-size: 0.9rem; color: #666; }
    </style>
</head>
<body>
    <div class="container">
        <h1>📤 Upload File</h1>
        <form id="uploadForm" enctype="multipart/form-data">
            <input type="file" id="fileInput" name="file" required>
            <button type="submit">Upload</button>
        </form>
        <div id="status"></div>
    </div>
    <script>
        document.getElementById('uploadForm').addEventListener('submit', async (e) => {
            e.preventDefault();
            const status = document.getElementById('status');
            const fileInput = document.getElementById('fileInput');
            const file = fileInput.files[0];

            if (!file) return;

            status.textContent = 'Uploading...';
            status.style.color = '#0071e3';

            const formData = new FormData();
            formData.append('file', file);

            try {
                const response = await fetch('/upload', {
                    method: 'POST',
                    body: formData
                });

                if (response.ok) {
                    status.textContent = '✅ Upload successful!';
                    status.style.color = 'green';
                    fileInput.value = '';
                } else {
                    status.textContent = '❌ Upload failed: ' + response.statusText;
                    status.style.color = 'red';
                }
            } catch (error) {
                status.textContent = '❌ Error: ' + error.message;
                status.style.color = 'red';
            }
        });
    </script>
</body>
</html>
`)
		})

		app.Post("/upload", func(c *fiber.Ctx) error {
			file, err := c.FormFile("file")
			if err != nil {
				return c.Status(fiber.StatusBadRequest).SendString("File upload failed")
			}

			// Prevent directory traversal attacks
			filename := filepath.Base(file.Filename)
			savePath := filepath.Join(absPath, filename)

			if err := c.SaveFile(file, savePath); err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Could not save file")
			}

			fmt.Printf("📥 Received File : %s\n", filename)
			return c.SendString("File uploaded successfully")
		})
	} else {
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
	}

	log.Fatal(app.Listen(fmt.Sprintf("[::]:%d", port)))
}
