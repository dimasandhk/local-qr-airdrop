## 2025-04-16 - Headless Clipboard Integration Constraint
**Learning:** CLI clipboard integration on Linux relies on external display server tools (`xclip`, `xsel`, Wayland). In headless environments or typical containerized sandboxes, these tools are often missing or fail to execute.
**Action:** Always wrap clipboard interactions in CLI tools with silent error handling to ensure they degrade gracefully and don't crash the application in environments without a window manager.
## 2024-05-24 - Increasing Default Fiber Body Limit for Uploads
**Learning:** Fiber framework defaults to a 4MB body limit, which blocks standard file-sharing utilities when receiving even moderately sized files via multipart forms.
**Action:** Always increase `fiber.Config.BodyLimit` (e.g., to 100MB) when building endpoints meant to accept user file uploads in Go Fiber apps.
