## 2025-02-21 - Preventing Stored XSS in Inline HTML Generation
**Learning:** When using raw strings for HTML template generation in Go (like `fmt.Sprintf`), user-controlled data such as filenames can easily introduce Stored Cross-Site Scripting (XSS) if not properly sanitized.
**Action:** Always use `html.EscapeString()` from the `html` package when rendering user-provided input into HTML strings in Go.

## 2025-02-21 - Managing Memory Leaks in Global State
**Learning:** In a long-running app, continuously appending to global state variables (like an `uploadedFiles` slice) without a limit will cause a memory leak.
**Action:** When introducing global state variables for recent items, implement a hard cap (e.g., retaining only the last 10 items) to prevent unbounded memory growth.

## 2025-02-14 - Dark Mode Theme Toggle Implementation
**Learning:** Injecting global UI features like dark mode into inline HTML templates in Go requires replicating CSS and JS scripts across multiple endpoint responses due to the lack of a shared templating engine. The inline `<script>` tags for reading local storage and setting initial classes must be placed in the `<head>` instead of `<body>` to prevent a Flash of Unstyled Content (FOUC).
**Action:** When adding global UI features to simple CLI tools that return raw HTML strings, ensure the logic is duplicated across all relevant endpoints or consider extracting a shared header/footer template string to avoid maintenance issues.
