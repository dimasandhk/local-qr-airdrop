## 2025-02-21 - Preventing Stored XSS in Inline HTML Generation
**Learning:** When using raw strings for HTML template generation in Go (like `fmt.Sprintf`), user-controlled data such as filenames can easily introduce Stored Cross-Site Scripting (XSS) if not properly sanitized.
**Action:** Always use `html.EscapeString()` from the `html` package when rendering user-provided input into HTML strings in Go.

## 2025-02-21 - Managing Memory Leaks in Global State
**Learning:** In a long-running app, continuously appending to global state variables (like an `uploadedFiles` slice) without a limit will cause a memory leak.
**Action:** When introducing global state variables for recent items, implement a hard cap (e.g., retaining only the last 10 items) to prevent unbounded memory growth.
## 2024-05-30 - Dark Mode Implementation in CLI Web Views
**Learning:** When injecting global UI features like Dark Mode into this CLI application, because it uses inline `fmt.Sprintf` string literals for HTML responses rather than a shared templating engine, any shared CSS variables or Javascript toggle logic must be explicitly duplicated across each individual string literal definition.
**Action:** When implementing UI/UX enhancements (like themes or generic utilities), ensure they are added to all specific handler responses (like the upload form and the success page) to maintain a consistent experience.
