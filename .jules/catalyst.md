## 2025-02-21 - Preventing Stored XSS in Inline HTML Generation
**Learning:** When using raw strings for HTML template generation in Go (like `fmt.Sprintf`), user-controlled data such as filenames can easily introduce Stored Cross-Site Scripting (XSS) if not properly sanitized.
**Action:** Always use `html.EscapeString()` from the `html` package when rendering user-provided input into HTML strings in Go.

## 2025-02-21 - Managing Memory Leaks in Global State
**Learning:** In a long-running app, continuously appending to global state variables (like an `uploadedFiles` slice) without a limit will cause a memory leak.
**Action:** When introducing global state variables for recent items, implement a hard cap (e.g., retaining only the last 10 items) to prevent unbounded memory growth.
## 2025-02-21 - Ensuring Required Playwright Binaries are Installed for Verification
**Learning:** When using Playwright for the first time in a fresh environment or container, the Python package installation (`pip install playwright`) is not sufficient. The browser binaries (like Chromium) must also be explicitly downloaded, otherwise verification scripts will crash with 'Executable doesn't exist' errors.
**Action:** When using Playwright for frontend verification, ensure the required browser binaries are downloaded by executing `python3 -m playwright install` in the bash session before running your verification script.
