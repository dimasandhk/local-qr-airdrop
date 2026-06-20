## 2025-02-21 - Preventing Stored XSS in Inline HTML Generation
**Learning:** When using raw strings for HTML template generation in Go (like `fmt.Sprintf`), user-controlled data such as filenames can easily introduce Stored Cross-Site Scripting (XSS) if not properly sanitized.
**Action:** Always use `html.EscapeString()` from the `html` package when rendering user-provided input into HTML strings in Go.

## 2025-02-21 - Managing Memory Leaks in Global State
**Learning:** In a long-running app, continuously appending to global state variables (like an `uploadedFiles` slice) without a limit will cause a memory leak.
**Action:** When introducing global state variables for recent items, implement a hard cap (e.g., retaining only the last 10 items) to prevent unbounded memory growth.

## 2025-02-12 - Dark Mode Implementation in Single-File Templates
**Learning:** Injecting feature-rich UI elements (like Dark Mode) into this CLI's web views requires duplicating CSS and JS snippets across endpoints due to the lack of a shared templating engine. Inline initialization `<script>` tags must be placed inside `<head>` to manipulate `document.documentElement` to prevent a Flash of Unstyled Content (FOUC).
**Action:** When adding global UI features to such architectures, plan for duplicated snippet injection and handle state safely within the inline constraints.
