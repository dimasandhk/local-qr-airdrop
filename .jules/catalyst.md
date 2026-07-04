## 2025-02-21 - Preventing Stored XSS in Inline HTML Generation
**Learning:** When using raw strings for HTML template generation in Go (like `fmt.Sprintf`), user-controlled data such as filenames can easily introduce Stored Cross-Site Scripting (XSS) if not properly sanitized.
**Action:** Always use `html.EscapeString()` from the `html` package when rendering user-provided input into HTML strings in Go.

## 2025-02-21 - Managing Memory Leaks in Global State
**Learning:** In a long-running app, continuously appending to global state variables (like an `uploadedFiles` slice) without a limit will cause a memory leak.
**Action:** When introducing global state variables for recent items, implement a hard cap (e.g., retaining only the last 10 items) to prevent unbounded memory growth.
## 2025-02-21 - Avoiding FOUC in Static Go Apps
**Learning:** When injecting theme toggles into simple Go web applications without a templating engine or build step, placing the initialization logic at the end of the `<body>` can cause a Flash of Unstyled Content (FOUC).
**Action:** Always inject theme initialization `<script>` blocks inside the `<head>` tag and manipulate `document.documentElement` to apply the theme instantly before the DOM body parses.
