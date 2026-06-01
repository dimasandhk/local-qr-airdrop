## 2025-02-21 - Preventing Stored XSS in Inline HTML Generation
**Learning:** When using raw strings for HTML template generation in Go (like `fmt.Sprintf`), user-controlled data such as filenames can easily introduce Stored Cross-Site Scripting (XSS) if not properly sanitized.
**Action:** Always use `html.EscapeString()` from the `html` package when rendering user-provided input into HTML strings in Go.

## 2025-02-21 - Managing Memory Leaks in Global State
**Learning:** In a long-running app, continuously appending to global state variables (like an `uploadedFiles` slice) without a limit will cause a memory leak.
**Action:** When introducing global state variables for recent items, implement a hard cap (e.g., retaining only the last 10 items) to prevent unbounded memory growth.
## 2024-06-01 - Inline HTML Modding for Dark Mode
**Learning:** Adding features like Dark Mode to this app requires modifying inline HTML templates via string replacement in `main.go`. There's no unified template system, so JS/CSS must be duplicated for multiple endpoints (`/` and `/upload`).
**Action:** Use script-based HTML replacement to guarantee precise formatting and always explicitly escape percent signs (`%%`) in Go `fmt.Sprintf` CSS templates.
