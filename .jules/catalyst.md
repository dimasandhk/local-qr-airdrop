## 2024-04-16 - Download Landing Page for Single Files
**Learning:** Automatically streaming large files directly upon scanning a QR code can result in bad mobile browser UX (inlining large text/JSON files, confusing downloads).
**Action:** Adding a simple intercepting HTML landing page to confirm the file details (name, size) and triggering the actual download provides explicit control and better mobile UX. It also allows triggering the graceful server shutdown cleanly.
