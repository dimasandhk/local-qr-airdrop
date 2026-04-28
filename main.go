package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/atotto/clipboard"
	"github.com/dimasandhk/local-qr-airdrop/internal/network"
	"github.com/dimasandhk/local-qr-airdrop/internal/terminal"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	uploadedFiles   []string
	uploadedFilesMu sync.Mutex
)

func main() {
	var port int
	var receiveMode bool
	flag.IntVar(&port, "port", 3030, "Port to run the server on")
	flag.IntVar(&port, "p", 3030, "Port to run the server on (shorthand)")
	flag.BoolVar(&receiveMode, "receive", false, "Enable receive mode to upload files to PC")
	flag.BoolVar(&receiveMode, "r", false, "Enable receive mode to upload files to PC (shorthand)")

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

	if receiveMode && !info.IsDir() {
		log.Fatalf("Error: Receive mode requires a directory path, not a file.")
	}

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		BodyLimit:             100 * 1024 * 1024,
	})

	app.Use(logger.New(logger.Config{
		Format:     "🕒 ${time} | 📱 ${ip} | 📝 ${status} - ${method} ${path}\n",
		TimeFormat: "15:04:05",
	}))

	localIP := network.GetLocalIP(port)

	serverURL := fmt.Sprintf("http://%s:%d", localIP, port)

	fmt.Println("========================================")
	if receiveMode {
		fmt.Printf("📥 Receive Mode     : ACTIVE\n")
		fmt.Printf("🎯 Save Directory   : %s\n", absPath)
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

	// Serve either a single file or a whole directory based on user input
	if receiveMode {
		app.Get("/", func(c *fiber.Ctx) error {
			uploadedFilesMu.Lock()
			recentUploadsHTML := ""
			if len(uploadedFiles) > 0 {
				recentUploadsHTML += `<div class="card" style="margin-top: 20px; text-align: left;">`
				recentUploadsHTML += `<h3>🕒 Recent Uploads</h3>`
				recentUploadsHTML += `<ul style="padding-left: 20px;">`
				for i := len(uploadedFiles) - 1; i >= 0; i-- {
					recentUploadsHTML += fmt.Sprintf(`<li>%s</li>`, html.EscapeString(uploadedFiles[i]))
				}
				recentUploadsHTML += `</ul>`
				recentUploadsHTML += `</div>`
			}
			uploadedFilesMu.Unlock()

			html := fmt.Sprintf(`<!DOCTYPE html>
<html data-theme="light">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<style>
		:root {
			--bg-color: #ffffff;
			--text-color: #000000;
			--card-bg: #ffffff;
			--card-border: #ddd;
			--input-bg: #ffffff;
			--input-text: #000000;
		}
		[data-theme="dark"] {
			--bg-color: #121212;
			--text-color: #e0e0e0;
			--card-bg: #1e1e1e;
			--card-border: #333;
			--input-bg: #2d2d2d;
			--input-text: #e0e0e0;
		}
		body { font-family: sans-serif; padding: 20px; text-align: center; max-width: 600px; margin: auto; background-color: var(--bg-color); color: var(--text-color); transition: background-color 0.3s, color 0.3s; }
		.btn { background: #007bff; color: white; border: none; padding: 12px 24px; border-radius: 5px; font-size: 16px; margin-top: 20px; width: 100%%; cursor: pointer; }
		.btn:hover { background: #0056b3; }
		input[type=file] { margin: 20px 0; padding: 10px; border: 1px solid var(--card-border); border-radius: 5px; width: 100%%; box-sizing: border-box; background-color: var(--input-bg); color: var(--input-text); }
		.card { border: 1px solid var(--card-border); padding: 20px; border-radius: 8px; box-shadow: 0 4px 8px rgba(0,0,0,0.1); background-color: var(--card-bg); transition: background-color 0.3s, border-color 0.3s; }
		.theme-btn { position: absolute; top: 10px; right: 10px; padding: 5px 10px; border-radius: 5px; cursor: pointer; border: 1px solid var(--card-border); background: var(--card-bg); color: var(--text-color); }
	</style>
</head>
<body>
	<button onclick="toggleTheme()" class="theme-btn" id="theme-toggle">🌙</button>
	<div class="card">
		<h2>📥 Send File to PC</h2>
		<p>Select a file from your device to send.</p>
		<form action="/upload" method="post" enctype="multipart/form-data">
			<input type="file" name="file" required><br>
			<input type="submit" value="Upload File" class="btn">
		</form>
	</div>
	%s
	<script>
		const htmlEl = document.documentElement;
		const toggleBtn = document.getElementById('theme-toggle');
		const savedTheme = localStorage.getItem('theme') || 'light';
		htmlEl.setAttribute('data-theme', savedTheme);
		toggleBtn.textContent = savedTheme === 'dark' ? '☀️' : '🌙';

		function toggleTheme() {
			const currentTheme = htmlEl.getAttribute('data-theme');
			const newTheme = currentTheme === 'light' ? 'dark' : 'light';
			htmlEl.setAttribute('data-theme', newTheme);
			localStorage.setItem('theme', newTheme);
			toggleBtn.textContent = newTheme === 'dark' ? '☀️' : '🌙';
		}
	</script>
</body>
</html>`, recentUploadsHTML)
			return c.Type("html").SendString(html)
		})

		app.Post("/upload", func(c *fiber.Ctx) error {
			file, err := c.FormFile("file")
			if err != nil {
				return c.Status(400).SendString("Error uploading file")
			}

			savePath := filepath.Join(absPath, filepath.Base(file.Filename))
			err = c.SaveFile(file, savePath)
			if err != nil {
				return c.Status(500).SendString("Error saving file")
			}

			uploadedFilesMu.Lock()
			uploadedFiles = append(uploadedFiles, filepath.Base(file.Filename))
			if len(uploadedFiles) > 10 {
				uploadedFiles = uploadedFiles[len(uploadedFiles)-10:]
			}
			uploadedFilesMu.Unlock()

			successHtml := fmt.Sprintf(`<!DOCTYPE html>
<html data-theme="light">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<style>
		:root {
			--bg-color: #ffffff;
			--text-color: #000000;
			--card-bg: #ffffff;
			--card-border: #ddd;
		}
		[data-theme="dark"] {
			--bg-color: #121212;
			--text-color: #e0e0e0;
			--card-bg: #1e1e1e;
			--card-border: #333;
		}
		body { font-family: sans-serif; padding: 20px; text-align: center; max-width: 600px; margin: auto; background-color: var(--bg-color); color: var(--text-color); transition: background-color 0.3s, color 0.3s; }
		.btn { background: #28a745; color: white; border: none; padding: 12px 24px; border-radius: 5px; font-size: 16px; margin-top: 20px; text-decoration: none; display: inline-block; }
		.btn:hover { background: #218838; }
		.card { border: 1px solid var(--card-border); padding: 20px; border-radius: 8px; box-shadow: 0 4px 8px rgba(0,0,0,0.1); background-color: var(--card-bg); transition: background-color 0.3s, border-color 0.3s; }
		.theme-btn { position: absolute; top: 10px; right: 10px; padding: 5px 10px; border-radius: 5px; cursor: pointer; border: 1px solid var(--card-border); background: var(--card-bg); color: var(--text-color); }
	</style>
</head>
<body>
	<button onclick="toggleTheme()" class="theme-btn" id="theme-toggle">🌙</button>
	<div class="card">
		<h2>✅ Success!</h2>
		<p>Successfully uploaded: <strong>%s</strong></p>
		<a href="/" class="btn">Upload Another File</a>
	</div>
	<script>
		const htmlEl = document.documentElement;
		const toggleBtn = document.getElementById('theme-toggle');
		const savedTheme = localStorage.getItem('theme') || 'light';
		htmlEl.setAttribute('data-theme', savedTheme);
		toggleBtn.textContent = savedTheme === 'dark' ? '☀️' : '🌙';

		function toggleTheme() {
			const currentTheme = htmlEl.getAttribute('data-theme');
			const newTheme = currentTheme === 'light' ? 'dark' : 'light';
			htmlEl.setAttribute('data-theme', newTheme);
			localStorage.setItem('theme', newTheme);
			toggleBtn.textContent = newTheme === 'dark' ? '☀️' : '🌙';
		}
	</script>
</body>
</html>`, html.EscapeString(file.Filename))
			return c.Type("html").SendString(successHtml)
		})
	} else if info.IsDir() {
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
