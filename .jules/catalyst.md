## 2025-02-21 - Preventing Stored XSS in Inline HTML Generation
**Learning:** When using raw strings for HTML template generation in Go (like `fmt.Sprintf`), user-controlled data such as filenames can easily introduce Stored Cross-Site Scripting (XSS) if not properly sanitized.
**Action:** Always use `html.EscapeString()` from the `html` package when rendering user-provided input into HTML strings in Go.

## 2025-02-21 - Managing Memory Leaks in Global State
**Learning:** In a long-running app, continuously appending to global state variables (like an `uploadedFiles` slice) without a limit will cause a memory leak.
**Action:** When introducing global state variables for recent items, implement a hard cap (e.g., retaining only the last 10 items) to prevent unbounded memory growth.

## 2025-05-03 - Handling UI Logic without a Templating Engine
**Learning:** When injecting global UI features (like Dark Mode) into a CLI's web views where `fmt.Sprintf` handles HTML templates inline, CSS and JS snippets must be explicitly duplicated across endpoints due to the lack of a shared templating engine.
**Action:** When adding global UI state features (like Dark Mode) in standard Go setups utilizing raw string templates, remember to duplicate the inline CSS/JS components across all relevant routes to maintain consistent styling.
