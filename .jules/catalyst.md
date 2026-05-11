## 2025-02-21 - Preventing Stored XSS in Inline HTML Generation
**Learning:** When using raw strings for HTML template generation in Go (like `fmt.Sprintf`), user-controlled data such as filenames can easily introduce Stored Cross-Site Scripting (XSS) if not properly sanitized.
**Action:** Always use `html.EscapeString()` from the `html` package when rendering user-provided input into HTML strings in Go.

## 2025-02-21 - Managing Memory Leaks in Global State
**Learning:** In a long-running app, continuously appending to global state variables (like an `uploadedFiles` slice) without a limit will cause a memory leak.
**Action:** When introducing global state variables for recent items, implement a hard cap (e.g., retaining only the last 10 items) to prevent unbounded memory growth.
## 2025-02-23 - Inline Template Duplication constraint
**Learning:** Due to the lack of a shared template engine in this minimal CLI web server, global UI features like theme toggles must have their CSS and JS logic duplicated verbatim across multiple isolated `fmt.Sprintf` HTML string definitions.
**Action:** When implementing app-wide styling or scripts, explicitly audit all HTML response strings to ensure the feature is applied uniformly to avoid broken experiences during navigation.
