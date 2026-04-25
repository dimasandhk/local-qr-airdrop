## 2025-04-16 - Headless Clipboard Integration Constraint
**Learning:** CLI clipboard integration on Linux relies on external display server tools (`xclip`, `xsel`, Wayland). In headless environments or typical containerized sandboxes, these tools are often missing or fail to execute.
**Action:** Always wrap clipboard interactions in CLI tools with silent error handling to ensure they degrade gracefully and don't crash the application in environments without a window manager.
## 2025-04-20 - Bidirectional File Transfer
**Learning:** For a local file-sharing utility, the ability to send files is only half the user journey. The "missing link" was bidirectional transfer—allowing users to seamlessly pull files from a device to the host without requiring complex FTP or cloud setups.
**Action:** When evaluating simple server utilities, always consider if the inverse operation (e.g., upload vs download) can be implemented with minimal overhead to double the utility of the application.
## 2023-10-27 - Backend configs are not features
**Learning:** For the Catalyst persona, purely backend configuration tweaks (like increasing file upload limits) do not qualify as "new features". The task explicitly demands tangible UI/UX additions or user-facing product capabilities.
**Action:** Always prioritize visible, user-interacting elements (like Dark Mode, Search bars, or Export buttons) when selecting the daily build.
