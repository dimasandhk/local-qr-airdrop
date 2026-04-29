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
<html>
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<style>
		:root {
			--bg-color: #ffffff;
			--text-color: #333333;
			--card-bg: #ffffff;
			--border-color: #ddd;
			--btn-bg: #007bff;
			--btn-hover: #0056b3;
		}
		[data-theme="dark"] {
			--bg-color: #1a1a1a;
			--text-color: #f0f0f0;
			--card-bg: #2a2a2a;
			--border-color: #444;
			--btn-bg: #0056b3;
			--btn-hover: #007bff;
		}
		body { font-family: sans-serif; padding: 20px; text-align: center; max-width: 600px; margin: auto; background-color: var(--bg-color); color: var(--text-color); transition: background-color 0.3s, color 0.3s; }
		.btn { background: var(--btn-bg); color: white; border: none; padding: 12px 24px; border-radius: 5px; font-size: 16px; margin-top: 20px; width: 100%%; cursor: pointer; transition: background 0.3s; }
		.btn:hover { background: var(--btn-hover); }
		input[type=file] { margin: 20px 0; padding: 10px; border: 1px solid var(--border-color); border-radius: 5px; width: 100%%; box-sizing: border-box; background-color: var(--card-bg); color: var(--text-color); }
		.card { border: 1px solid var(--border-color); padding: 20px; border-radius: 8px; box-shadow: 0 4px 8px rgba(0,0,0,0.1); background-color: var(--card-bg); transition: background-color 0.3s, border-color 0.3s; }
		.theme-toggle { position: absolute; top: 20px; right: 20px; background: none; border: none; font-size: 24px; cursor: pointer; padding: 0; }
	</style>
	<script>
		// Apply theme early to prevent flash
		const savedTheme = localStorage.getItem("theme") || (window.matchMedia("(prefers-color-scheme: dark)").matches ? "dark" : "light");
		document.documentElement.setAttribute("data-theme", savedTheme);
		function toggleTheme() {
			const currentTheme = document.documentElement.getAttribute("data-theme");
			const newTheme = currentTheme === "dark" ? "light" : "dark";
			document.documentElement.setAttribute("data-theme", newTheme);
			localStorage.setItem("theme", newTheme);
		}
	</script>
</head>
<body>
	<button class="theme-toggle" onclick="toggleTheme()" aria-label="Toggle Dark Mode">🌓</button>
	<div class="card">
		<h2>📥 Send File to PC</h2>
		<p>Select a file from your device to send.</p>
		<form action="/upload" method="post" enctype="multipart/form-data">
			<input type="file" name="file" required><br>
			<input type="submit" value="Upload File" class="btn">
		</form>
	</div>
	%s
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
<html>
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<style>
		:root {
			--bg-color: #ffffff;
			--text-color: #333333;
			--btn-bg: #28a745;
			--btn-hover: #218838;
		}
		[data-theme="dark"] {
			--bg-color: #1a1a1a;
			--text-color: #f0f0f0;
			--btn-bg: #218838;
			--btn-hover: #28a745;
		}
		body { font-family: sans-serif; padding: 20px; text-align: center; max-width: 600px; margin: auto; background-color: var(--bg-color); color: var(--text-color); transition: background-color 0.3s, color 0.3s; }
		.btn { background: var(--btn-bg); color: white; border: none; padding: 12px 24px; border-radius: 5px; font-size: 16px; margin-top: 20px; text-decoration: none; display: inline-block; transition: background 0.3s; }
		.btn:hover { background: var(--btn-hover); }
		.theme-toggle { position: absolute; top: 20px; right: 20px; background: none; border: none; font-size: 24px; cursor: pointer; padding: 0; }
	</style>
	<script>
		// Apply theme early to prevent flash
		const savedTheme = localStorage.getItem("theme") || (window.matchMedia("(prefers-color-scheme: dark)").matches ? "dark" : "light");
		document.documentElement.setAttribute("data-theme", savedTheme);
		function toggleTheme() {
			const currentTheme = document.documentElement.getAttribute("data-theme");
			const newTheme = currentTheme === "dark" ? "light" : "dark";
			document.documentElement.setAttribute("data-theme", newTheme);
			localStorage.setItem("theme", newTheme);
		}
	</script>
</head>
<body>
	<button class="theme-toggle" onclick="toggleTheme()" aria-label="Toggle Dark Mode">🌓</button>
	<h2>✅ Success!</h2>
	<p>Successfully uploaded: <strong>%s</strong></p>
	<a href="/" class="btn">Upload Another File</a>
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
