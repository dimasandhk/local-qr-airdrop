## 2025-04-16 - Headless Clipboard Integration Constraint
**Learning:** CLI clipboard integration on Linux relies on external display server tools (`xclip`, `xsel`, Wayland). In headless environments or typical containerized sandboxes, these tools are often missing or fail to execute.
**Action:** Always wrap clipboard interactions in CLI tools with silent error handling to ensure they degrade gracefully and don't crash the application in environments without a window manager.
## 2025-04-20 - Bidirectional File Transfer
**Learning:** For a local file-sharing utility, the ability to send files is only half the user journey. The "missing link" was bidirectional transfer—allowing users to seamlessly pull files from a device to the host without requiring complex FTP or cloud setups.
**Action:** When evaluating simple server utilities, always consider if the inverse operation (e.g., upload vs download) can be implemented with minimal overhead to double the utility of the application.
## 2025-05-18 - Client-Side State Persistence MVP
**Learning:** MVP features that introduce new states (like Dark Mode) can feel broken if that state resets across page navigations. In a local/serverless context, simple `localStorage` is the missing link to make the feature feel polished without adding backend complexity.
**Action:** Always persist client-side UI preferences (like Dark Mode) using `localStorage` to ensure consistency across page reloads or navigations, even for simple MVP implementations.
