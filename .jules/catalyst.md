## 2025-02-21 - Preventing Stored XSS in Inline HTML Generation
**Learning:** When using raw strings for HTML template generation in Go (like `fmt.Sprintf`), user-controlled data such as filenames can easily introduce Stored Cross-Site Scripting (XSS) if not properly sanitized.
**Action:** Always use `html.EscapeString()` from the `html` package when rendering user-provided input into HTML strings in Go.

## 2025-02-21 - Managing Memory Leaks in Global State
**Learning:** In a long-running app, continuously appending to global state variables (like an `uploadedFiles` slice) without a limit will cause a memory leak.
**Action:** When introducing global state variables for recent items, implement a hard cap (e.g., retaining only the last 10 items) to prevent unbounded memory growth.
## 2024-05-26 - Implementing Dark Mode Without Templating Engines
**Learning:** The `local-qr-airdrop` application lacks a shared HTML templating engine (like `html/template` blocks) and instead hardcodes entire HTML documents inside inline `fmt.Sprintf` statements within specific Fiber route handlers. This architectural constraint means that global UI features like CSS variables, Dark Mode toggles, and their associated JavaScript initialization scripts cannot be applied to a single layout file. They must be manually duplicated into every individual HTML string literal returned by the application.
**Action:** When adding global frontend features to this specific app, ensure the plan explicitly includes steps to locate and modify every independent `fmt.Sprintf` block serving HTML. Avoid attempting to refactor the app into using an external templating engine (which violates the Catalyst anti-over-engineering rule), and instead duplicate the necessary `<style>` and `<script>` blocks while being exceptionally careful with `%%` escaping in the `Sprintf` formats.
