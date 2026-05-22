## 2025-02-21 - Preventing Stored XSS in Inline HTML Generation
**Learning:** When using raw strings for HTML template generation in Go (like `fmt.Sprintf`), user-controlled data such as filenames can easily introduce Stored Cross-Site Scripting (XSS) if not properly sanitized.
**Action:** Always use `html.EscapeString()` from the `html` package when rendering user-provided input into HTML strings in Go.

## 2025-02-21 - Managing Memory Leaks in Global State
**Learning:** In a long-running app, continuously appending to global state variables (like an `uploadedFiles` slice) without a limit will cause a memory leak.
**Action:** When introducing global state variables for recent items, implement a hard cap (e.g., retaining only the last 10 items) to prevent unbounded memory growth.
## 2024-05-22 - Injecting Global UI Features into Decentralized Go Templates
**Learning:** Adding a global feature like a theme toggle to a Go application that uses inline `fmt.Sprintf` HTML templates without a shared templating engine requires careful duplication of both CSS and initialization logic across every distinct endpoint string to maintain state and avoid FOUC.
**Action:** When adding global UI functionality (like themes or navbars) in such an architecture, explicitly document and implement the logic across all relevant string variables simultaneously, and rely on `localStorage` in the `<head>` to persist state across page loads instead of server-side sessions.
