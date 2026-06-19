## 2025-02-21 - Preventing Stored XSS in Inline HTML Generation
**Learning:** When using raw strings for HTML template generation in Go (like `fmt.Sprintf`), user-controlled data such as filenames can easily introduce Stored Cross-Site Scripting (XSS) if not properly sanitized.
**Action:** Always use `html.EscapeString()` from the `html` package when rendering user-provided input into HTML strings in Go.

## 2025-02-21 - Managing Memory Leaks in Global State
**Learning:** In a long-running app, continuously appending to global state variables (like an `uploadedFiles` slice) without a limit will cause a memory leak.
**Action:** When introducing global state variables for recent items, implement a hard cap (e.g., retaining only the last 10 items) to prevent unbounded memory growth.

## 2025-02-21 - Avoiding Flash of Unstyled Content (FOUC)
**Learning:** When implementing Dark Mode using inline HTML templates, placing the initialization script at the end of the `<body>` can cause a Flash of Unstyled Content (FOUC) while the page loads.
**Action:** Place the theme initialization `<script>` tag inside the `<head>` and manipulate `document.documentElement` instead of `document.body` to ensure the theme is applied before rendering.

## 2025-02-21 - Escaping Percent Signs in fmt.Sprintf
**Learning:** When using `fmt.Sprintf` to generate HTML strings containing inline CSS, percent signs used for dimensions (like `width: 100%;`) will cause "unknown format verb" errors during compilation.
**Action:** Always explicitly escape percent signs as `%%` (e.g., `width: 100%%;`) when dealing with `fmt.Sprintf` format strings.
