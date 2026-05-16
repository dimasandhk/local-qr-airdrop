## 2025-02-21 - Preventing Stored XSS in Inline HTML Generation
**Learning:** When using raw strings for HTML template generation in Go (like `fmt.Sprintf`), user-controlled data such as filenames can easily introduce Stored Cross-Site Scripting (XSS) if not properly sanitized.
**Action:** Always use `html.EscapeString()` from the `html` package when rendering user-provided input into HTML strings in Go.

## 2025-02-21 - Managing Memory Leaks in Global State
**Learning:** In a long-running app, continuously appending to global state variables (like an `uploadedFiles` slice) without a limit will cause a memory leak.
**Action:** When introducing global state variables for recent items, implement a hard cap (e.g., retaining only the last 10 items) to prevent unbounded memory growth.

## 2025-02-21 - FOUC Prevention in Go Inline HTML Templates
**Learning:** When injecting global UI features like Dark Mode into a Go CLI's web views using inline `fmt.Sprintf` templates without a shared templating engine, placing the initialization script at the end of the body causes a Flash of Unstyled Content (FOUC).
**Action:** Always place the theme initialization `<script>` inside the `<head>` tag and ensure CSS and JS snippets are duplicated across all relevant endpoints to maintain a consistent UI state.
