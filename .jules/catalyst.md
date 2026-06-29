## 2025-02-21 - Preventing Stored XSS in Inline HTML Generation
**Learning:** When using raw strings for HTML template generation in Go (like `fmt.Sprintf`), user-controlled data such as filenames can easily introduce Stored Cross-Site Scripting (XSS) if not properly sanitized.
**Action:** Always use `html.EscapeString()` from the `html` package when rendering user-provided input into HTML strings in Go.

## 2025-02-21 - Managing Memory Leaks in Global State
**Learning:** In a long-running app, continuously appending to global state variables (like an `uploadedFiles` slice) without a limit will cause a memory leak.
**Action:** When introducing global state variables for recent items, implement a hard cap (e.g., retaining only the last 10 items) to prevent unbounded memory growth.
## 2026-06-29 - Dark Mode in Inline Go Templates
**Learning:** When injecting global UI features like Dark Mode into this CLI's web views using inline fmt.Sprintf templates, CSS and JS snippets must be duplicated across endpoints due to the lack of a shared templating engine. Also, explicit %% escaping is required for CSS rules spanning 100%%.
**Action:** Duplicate snippets for theme toggling securely across inline templates and use document.documentElement to prevent FOUC.
