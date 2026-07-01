## 2025-02-21 - Preventing Stored XSS in Inline HTML Generation
**Learning:** When using raw strings for HTML template generation in Go (like `fmt.Sprintf`), user-controlled data such as filenames can easily introduce Stored Cross-Site Scripting (XSS) if not properly sanitized.
**Action:** Always use `html.EscapeString()` from the `html` package when rendering user-provided input into HTML strings in Go.

## 2025-02-21 - Managing Memory Leaks in Global State
**Learning:** In a long-running app, continuously appending to global state variables (like an `uploadedFiles` slice) without a limit will cause a memory leak.
**Action:** When introducing global state variables for recent items, implement a hard cap (e.g., retaining only the last 10 items) to prevent unbounded memory growth.
## 2025-02-21 - Avoiding Tool Hallucination in Execution Plans
**Learning:** The plan reviewer explicitly rejects plans that use fake or non-existent tools like `submit`. Standard actions (like "Create a PR") must be described plainly without implying they are a tool call.
**Action:** When creating execution plans, verify that every tool explicitly named in the plan matches a confirmed, available tool in the documentation (like `write_file`, `run_in_bash_session`, etc.). Use plain language for abstract concepts like submitting a PR.
