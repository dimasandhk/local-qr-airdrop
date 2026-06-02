## 2025-02-21 - Preventing Stored XSS in Inline HTML Generation
**Learning:** When using raw strings for HTML template generation in Go (like `fmt.Sprintf`), user-controlled data such as filenames can easily introduce Stored Cross-Site Scripting (XSS) if not properly sanitized.
**Action:** Always use `html.EscapeString()` from the `html` package when rendering user-provided input into HTML strings in Go.

## 2025-02-21 - Managing Memory Leaks in Global State
**Learning:** In a long-running app, continuously appending to global state variables (like an `uploadedFiles` slice) without a limit will cause a memory leak.
**Action:** When introducing global state variables for recent items, implement a hard cap (e.g., retaining only the last 10 items) to prevent unbounded memory growth.

## 2025-02-21 - Adding UI Features to Inline Go Templates
**Learning:** When injecting global UI features (like Dark Mode) into a CLI's web views that use inline `fmt.Sprintf` templates rather than a dedicated templating engine, CSS and JS snippets must be carefully duplicated across all relevant endpoints (e.g., `GET /` and `POST /upload`).
**Action:** When adding global UI functionality to such an architecture, ensure the necessary scripts and styles are applied to every HTML string returned by the web server to maintain consistency and prevent broken state across navigation.
