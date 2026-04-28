## 2025-02-21 - Preventing Stored XSS in Inline HTML Generation
**Learning:** When using raw strings for HTML template generation in Go (like `fmt.Sprintf`), user-controlled data such as filenames can easily introduce Stored Cross-Site Scripting (XSS) if not properly sanitized.
**Action:** Always use `html.EscapeString()` from the `html` package when rendering user-provided input into HTML strings in Go.

## 2025-02-21 - Managing Memory Leaks in Global State
**Learning:** In a long-running app, continuously appending to global state variables (like an `uploadedFiles` slice) without a limit will cause a memory leak.
**Action:** When introducing global state variables for recent items, implement a hard cap (e.g., retaining only the last 10 items) to prevent unbounded memory growth.

## 2025-02-21 - Maintainability of Embedded HTML Web Views
**Learning:** Integrating UI features like Dark Mode into single-file web views generated from CLI tools (using `fmt.Sprintf`) necessitates repeating CSS and JS snippets across multiple endpoint responses (e.g., index and upload success pages) due to the lack of a shared layout/template system, highlighting a design trade-off between deployment simplicity and codebase maintainability in MVP web architectures.
**Action:** For features requiring extensive styling or global JS logic, consider whether the trade-off of maintaining duplicate strings is worth the UI improvement, or if a lightweight templating engine (like `html/template`) should be introduced.
