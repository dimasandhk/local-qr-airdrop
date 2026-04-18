## 2025-04-16 - Headless Clipboard Integration Constraint
**Learning:** CLI clipboard integration on Linux relies on external display server tools (`xclip`, `xsel`, Wayland). In headless environments or typical containerized sandboxes, these tools are often missing or fail to execute.
**Action:** Always wrap clipboard interactions in CLI tools with silent error handling to ensure they degrade gracefully and don't crash the application in environments without a window manager.
## 2024-04-18 - Avoid committing build artifacts
**Learning:** Build artifacts like compiled binaries (`main`) and log files (`output.log`) can easily slip into the repository when executing local tests or build commands during development.
**Action:** Always clean up temporary files and execute `git status` to verify the staging area before submitting a commit.
