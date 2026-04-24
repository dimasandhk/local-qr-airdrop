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
			html := `<!DOCTYPE html>
<html>
<head>
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<style>
		body { font-family: sans-serif; padding: 20px; text-align: center; max-width: 600px; margin: auto; transition: background 0.3s, color 0.3s; }
		body.dark-mode { background: #121212; color: #ffffff; }
		.btn { background: #007bff; color: white; border: none; padding: 12px 24px; border-radius: 5px; font-size: 16px; margin-top: 20px; width: 100%; cursor: pointer; }
		.btn:hover { background: #0056b3; }
		input[type=file] { margin: 20px 0; padding: 10px; border: 1px solid #ccc; border-radius: 5px; width: 100%; box-sizing: border-box; transition: background 0.3s, color 0.3s; }
		.card { border: 1px solid #ddd; padding: 20px; border-radius: 8px; box-shadow: 0 4px 8px rgba(0,0,0,0.1); transition: background 0.3s, border-color 0.3s; }
		.dark-mode .card { background: #1e1e1e; border-color: #333; }
		.dark-mode input[type=file] { color: #fff; background: #333; border-color: #555; }
	</style>
	<script>
		function toggleDarkMode() {
			document.body.classList.toggle('dark-mode');
			localStorage.setItem('darkMode', document.body.classList.contains('dark-mode'));
		}
		window.onload = function() {
			if (localStorage.getItem('darkMode') === 'true') {
				document.body.classList.add('dark-mode');
			}
		}
	</script>
</head>
<body>
	<div class="card">
		<h2>📥 Send File to PC</h2>
		<p>Select a file from your device to send.</p>
		<form action="/upload" method="post" enctype="multipart/form-data">
			<input type="file" name="file" required><br>
			<input type="submit" value="Upload File" class="btn">
		</form>
		<button type="button" onclick="toggleDarkMode()" class="btn" style="background:#555; margin-top:10px;">Toggle Dark Mode</button>
	</div>
</body>
</html>`
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

			successHtml := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<style>
		body { font-family: sans-serif; padding: 20px; text-align: center; max-width: 600px; margin: auto; transition: background 0.3s, color 0.3s; }
		body.dark-mode { background: #121212; color: #ffffff; }
		.btn { background: #28a745; color: white; border: none; padding: 12px 24px; border-radius: 5px; font-size: 16px; margin-top: 20px; text-decoration: none; display: inline-block; }
		.btn:hover { background: #218838; }
	</style>
	<script>
		function toggleDarkMode() {
			document.body.classList.toggle('dark-mode');
			localStorage.setItem('darkMode', document.body.classList.contains('dark-mode'));
		}
		window.onload = function() {
			if (localStorage.getItem('darkMode') === 'true') {
				document.body.classList.add('dark-mode');
			}
		}
	</script>
</head>
<body>
	<h2>✅ Success!</h2>
	<p>Successfully uploaded: <strong>%s</strong></p>
	<a href="/" class="btn">Upload Another File</a>
	<button type="button" onclick="toggleDarkMode()" class="btn" style="background:#555; margin-top:10px; width:100%%; box-sizing: border-box;">Toggle Dark Mode</button>
</body>
</html>`, file.Filename)
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
