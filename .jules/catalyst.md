## 2025-02-21 - Preventing Stored XSS in Inline HTML Generation
**Learning:** When using raw strings for HTML template generation in Go (like `fmt.Sprintf`), user-controlled data such as filenames can easily introduce Stored Cross-Site Scripting (XSS) if not properly sanitized.
**Action:** Always use `html.EscapeString()` from the `html` package when rendering user-provided input into HTML strings in Go.

## 2025-02-21 - Managing Memory Leaks in Global State
**Learning:** In a long-running app, continuously appending to global state variables (like an `uploadedFiles` slice) without a limit will cause a memory leak.
**Action:** When introducing global state variables for recent items, implement a hard cap (e.g., retaining only the last 10 items) to prevent unbounded memory growth.

## 2024-05-24 - Avoiding FOUC in Single-File Web Servers
**Learning:** When injecting global UI features like Dark Mode into web views served directly from Go string literals, the initialization script must be placed inside the `<head>` tag rather than at the end of the `<body>` to prevent a Flash of Unstyled Content (FOUC).
**Action:** Always inject theme initialization logic in the `<head>` when working with static HTML templates.
