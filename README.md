# 📱 local QR Airdrop

**A blazing-fast CLI tool to share local files instantly from your terminal to your phone using QR codes. No cables, no Bluetooth, no cloud.**

[![Made with Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev/)
[![Powered by Fiber](https://img.shields.io/badge/Powered%20by-Fiber-blue.svg)](https://gofiber.io/)

## 💡 What is it?
Ever needed to quickly send a video, log file, or image from your computer to your phone, but didn't want to upload it to Google Drive or deal with AirDrop compatibility? 

**`qrop`** turns your machine into a temporary, single-file local web server. It binds to your local Wi-Fi IP address, generates an ASCII QR code directly in your terminal, and streams the file to any device that scans it. Once the download is complete, the server automatically shuts itself down.

## ✨ Features
* **Zero Client Install:** The receiving device just needs a camera and a web browser.
* **Fully Local & Private:** Your files never touch the internet or a cloud server.
* **Blazing Fast:** Built on top of [Fiber](https://gofiber.io/) (using `fasthttp`), utilizing zero-copy techniques for massive file streaming.
* **Fire-and-Forget:** Automatically handles graceful shutdown the second your file finishes downloading.
* **Cross-Platform:** Compiles to a single standalone binary for Windows, Mac, or Linux.

## 🚀 How to use it

Run the command and pass the file you want to share:

```bash
# Run with a text file
go run main.go ./secrets.txt

# Or a massive video file
go run main.go ./my-vacation-video.mp4
```

1. A QR code will instantly appear in your terminal.
2. Point your phone's camera at the screen.
3. Tap the link to download the file directly to your device.
4. `qrop` prints a success message and closes. That's it!

## 🛠️ Under the Hood (Tech Stack)
This project was built to explore Go's powerful concurrency and networking capabilities:
* **[Golang](https://go.dev/):** The core language.
* **[Fiber](https://gofiber.io/):** An Express-inspired, highly performant web framework for Go. Used to handle the HTTP streaming and graceful shutdowns.
* **[qrterminal](https://github.com/mdp/qrterminal):** Generates half-block ASCII QR codes for terminal rendering.
* **Goroutines & Channels:** Used to run the HTTP server asynchronously and listen for completion signals to trigger the shutdown.

---
*Built with ❤️ to learn Go.*

***

### 💡 Pro-Tip for your GitHub Repo:
Go to your GitHub repo settings, and in the "About" section on the right side, paste this as your short description:
> *A blazing-fast CLI tool to instantly share files over local Wi-Fi using terminal-generated QR codes. Built with Go & Fiber.* 

Add tags like: `golang`, `cli`, `fiber`, `file-sharing`, `qr-code`, `p2p`.