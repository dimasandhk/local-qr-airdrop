## 2025-02-14 - Inline Template State Management
**Learning:** When injecting global UI features (like Dark Mode) into this CLI's web views using inline `fmt.Sprintf` templates, CSS and JS snippets must be duplicated across endpoints due to the lack of a shared templating engine.
**Action:** When adding simple UI features, ensure you explicitly replace the targeted strings in all relevant route handlers (e.g., both `/` and `/upload` for a dark mode toggle) within `main.go`.
