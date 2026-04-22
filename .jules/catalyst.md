## 2025-04-16 - Headless Clipboard Integration Constraint
**Learning:** CLI clipboard integration on Linux relies on external display server tools (`xclip`, `xsel`, Wayland). In headless environments or typical containerized sandboxes, these tools are often missing or fail to execute.
**Action:** Always wrap clipboard interactions in CLI tools with silent error handling to ensure they degrade gracefully and don't crash the application in environments without a window manager.

## 2024-05-20 - Adding Receive Mode to a Unidirectional Airdrop Tool
**Learning:** For a single-device server meant to be used across multiple local devices without explicit client installs, turning the tool from unidirectional (send-only) to bidirectional (send and receive via a flag) drastically increases utility. Generating a simple web UI form from within the CLI to be served to the remote device creates a cross-platform, zero-install upload mechanism.
**Action:** When building CLI apps that serve web endpoints to local devices, consider offering endpoints that accept data back from those clients, turning the CLI host into a temporary bidirectional hub.
