## 2025-04-16 - Headless Clipboard Integration Constraint
**Learning:** CLI clipboard integration on Linux relies on external display server tools (`xclip`, `xsel`, Wayland). In headless environments or typical containerized sandboxes, these tools are often missing or fail to execute.
**Action:** Always wrap clipboard interactions in CLI tools with silent error handling to ensure they degrade gracefully and don't crash the application in environments without a window manager.
