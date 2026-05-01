## 2025-02-21 - Preventing Stored XSS in Inline HTML Generation
**Learning:** When using raw strings for HTML template generation in Go (like `fmt.Sprintf`), user-controlled data such as filenames can easily introduce Stored Cross-Site Scripting (XSS) if not properly sanitized.
**Action:** Always use `html.EscapeString()` from the `html` package when rendering user-provided input into HTML strings in Go.

## 2025-02-21 - Managing Memory Leaks in Global State
**Learning:** In a long-running app, continuously appending to global state variables (like an `uploadedFiles` slice) without a limit will cause a memory leak.
**Action:** When introducing global state variables for recent items, implement a hard cap (e.g., retaining only the last 10 items) to prevent unbounded memory growth.

## 2024-05-01 - Missing Dark Mode capability for late night airdrops
**Learning:** Terminal-based airdrop tools are often used in low-light environments by developers. The bright white default UI for the receive/upload screens can be jarring. Adding a simple, persistent dark mode significantly improves the user experience without requiring backend changes.
**Action:** Always consider the physical environment in which a tool is used. For CLI-spawned web interfaces, providing a dark mode toggle that persists via `localStorage` is a high-value MVP feature.
