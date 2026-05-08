## 2025-02-21 - Preventing Stored XSS in Inline HTML Generation
**Learning:** When using raw strings for HTML template generation in Go (like `fmt.Sprintf`), user-controlled data such as filenames can easily introduce Stored Cross-Site Scripting (XSS) if not properly sanitized.
**Action:** Always use `html.EscapeString()` from the `html` package when rendering user-provided input into HTML strings in Go.

## 2025-02-21 - Managing Memory Leaks in Global State
**Learning:** In a long-running app, continuously appending to global state variables (like an `uploadedFiles` slice) without a limit will cause a memory leak.
**Action:** When introducing global state variables for recent items, implement a hard cap (e.g., retaining only the last 10 items) to prevent unbounded memory growth.

## 2025-05-08 - Implementing Theme Toggles in Go Single-File Servers
**Learning:** When injecting global UI features (like Dark Mode) into a CLI's web views where HTML templates are generated inline via `fmt.Sprintf` (rather than a shared templating engine), you are forced to duplicate CSS and JS initialization snippets across all endpoints to maintain consistency.
**Action:** When implementing global themes in a system lacking shared layouts, ensure the styling and the FOUC prevention scripts are injected consistently into the `<head>` of every generated HTML page.
