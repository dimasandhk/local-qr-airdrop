## 2025-02-21 - Preventing Stored XSS in Inline HTML Generation
**Learning:** When using raw strings for HTML template generation in Go (like `fmt.Sprintf`), user-controlled data such as filenames can easily introduce Stored Cross-Site Scripting (XSS) if not properly sanitized.
**Action:** Always use `html.EscapeString()` from the `html` package when rendering user-provided input into HTML strings in Go.

## 2025-02-21 - Managing Memory Leaks in Global State
**Learning:** In a long-running app, continuously appending to global state variables (like an `uploadedFiles` slice) without a limit will cause a memory leak.
**Action:** When introducing global state variables for recent items, implement a hard cap (e.g., retaining only the last 10 items) to prevent unbounded memory growth.

## 2025-02-21 - Avoiding Shared Port Conflicts in UI Tests
**Learning:** During repeated frontend verification testing of web servers (like `go run main.go`), lingering background processes can cause "address already in use" errors when attempting to restart the server on the same port, leading to test failures.
**Action:** Always proactively terminate any existing processes on the target port (e.g., `kill -9 $(lsof -t -i :3030) 2>/dev/null || true`) before starting the server for UI testing.
