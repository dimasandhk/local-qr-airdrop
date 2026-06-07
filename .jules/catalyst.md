## 2025-02-21 - Preventing Stored XSS in Inline HTML Generation
**Learning:** When using raw strings for HTML template generation in Go (like `fmt.Sprintf`), user-controlled data such as filenames can easily introduce Stored Cross-Site Scripting (XSS) if not properly sanitized.
**Action:** Always use `html.EscapeString()` from the `html` package when rendering user-provided input into HTML strings in Go.

## 2025-02-21 - Managing Memory Leaks in Global State
**Learning:** In a long-running app, continuously appending to global state variables (like an `uploadedFiles` slice) without a limit will cause a memory leak.
**Action:** When introducing global state variables for recent items, implement a hard cap (e.g., retaining only the last 10 items) to prevent unbounded memory growth.
## 2025-02-21 - FOUC Prevention in Inline Theme Toggles
**Learning:** When injecting theme toggle scripts into raw string literals (like Go's `fmt.Sprintf`), the script needs to run as early as possible to prevent a Flash of Unstyled Content (FOUC). However, at the time the `<head>` is parsed, `document.body` is not yet available, leading to JavaScript errors if the theme is applied there.
**Action:** When implementing Dark Mode or other theme toggles using inline HTML templates, place the initialization `<script>` inside the `<head>` tag and manipulate `document.documentElement` instead of `document.body`.
