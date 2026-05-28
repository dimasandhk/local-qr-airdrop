## 2025-02-21 - Preventing Stored XSS in Inline HTML Generation
**Learning:** When using raw strings for HTML template generation in Go (like `fmt.Sprintf`), user-controlled data such as filenames can easily introduce Stored Cross-Site Scripting (XSS) if not properly sanitized.
**Action:** Always use `html.EscapeString()` from the `html` package when rendering user-provided input into HTML strings in Go.

## 2025-02-21 - Managing Memory Leaks in Global State
**Learning:** In a long-running app, continuously appending to global state variables (like an `uploadedFiles` slice) without a limit will cause a memory leak.
**Action:** When introducing global state variables for recent items, implement a hard cap (e.g., retaining only the last 10 items) to prevent unbounded memory growth.
## 2025-02-21 - Adding global UI features via inline HTML
**Learning:** When adding global UI features (like Dark Mode) to an application that serves inline HTML templates directly from handlers (without a shared templating engine), necessary CSS styles and JavaScript initialization logic must be duplicated across all endpoints that return HTML to maintain consistency.
**Action:** When working with inline HTML templates in Go, explicitly check for and duplicate required global assets (like theme scripts or base CSS variables) across all relevant template strings to ensure a uniform user experience.
