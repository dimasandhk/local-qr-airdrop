## 2025-02-21 - Preventing Stored XSS in Inline HTML Generation
**Learning:** When using raw strings for HTML template generation in Go (like `fmt.Sprintf`), user-controlled data such as filenames can easily introduce Stored Cross-Site Scripting (XSS) if not properly sanitized.
**Action:** Always use `html.EscapeString()` from the `html` package when rendering user-provided input into HTML strings in Go.

## 2025-02-21 - Managing Memory Leaks in Global State
**Learning:** In a long-running app, continuously appending to global state variables (like an `uploadedFiles` slice) without a limit will cause a memory leak.
**Action:** When introducing global state variables for recent items, implement a hard cap (e.g., retaining only the last 10 items) to prevent unbounded memory growth.
## 2025-02-21 - FOUC Prevention in Inline Go Templates
**Learning:** When adding a dark mode toggle to inline Go HTML templates using vanilla JS, a Flash of Unstyled Content (FOUC) will occur if the initializing script manipulates `document.body` because it hasn't parsed yet.
**Action:** Always place the initializing `<script>` inside the `<head>` tag and manipulate `document.documentElement` (`:root`) to immediately set the `data-theme` attribute before the DOM renders.

## 2025-02-21 - Duplicating Code in Lack of Shared Templates
**Learning:** This CLI app lacks a shared templating engine, relying entirely on string formatting (`fmt.Sprintf`) for raw HTML responses inside route handlers.
**Action:** When injecting global UI features (like Dark Mode CSS and JS), snippets must be explicitly duplicated across all endpoint templates.
