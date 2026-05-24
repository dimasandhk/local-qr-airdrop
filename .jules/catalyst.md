## 2025-02-21 - Preventing Stored XSS in Inline HTML Generation
**Learning:** When using raw strings for HTML template generation in Go (like `fmt.Sprintf`), user-controlled data such as filenames can easily introduce Stored Cross-Site Scripting (XSS) if not properly sanitized.
**Action:** Always use `html.EscapeString()` from the `html` package when rendering user-provided input into HTML strings in Go.

## 2025-02-21 - Managing Memory Leaks in Global State
**Learning:** In a long-running app, continuously appending to global state variables (like an `uploadedFiles` slice) without a limit will cause a memory leak.
**Action:** When introducing global state variables for recent items, implement a hard cap (e.g., retaining only the last 10 items) to prevent unbounded memory growth.

## 2025-02-21 - Avoiding Shared Templating Engine Pitfalls
**Learning:** In CLI applications that embed UI via inline `fmt.Sprintf` HTML templates rather than a shared templating engine or external assets folder, injecting global features (like a Dark Mode toggle) requires duplicating the raw CSS and initialization scripts across all relevant endpoints to ensure visual consistency without over-engineering an asset pipeline.
**Action:** When adding global UI features to simple Go CLI web views, explicitly implement and verify the feature across all inline HTML endpoints to maintain consistency, avoiding the temptation to over-engineer a central template registry for small MVPs.
