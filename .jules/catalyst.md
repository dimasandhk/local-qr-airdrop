## 2025-02-13 - Inline Themes
**Learning:** Adding a global feature like dark mode into a Go CLI app that serves raw HTML strings using `fmt.Sprintf` requires duplicating the CSS and JS toggle logic across every distinct HTML string block, as there is no centralized template engine to inherit from.
**Action:** Always check the scope of the templates in the Go file and replicate the initialization logic (`<script>` in `<head>`), CSS variables, and toggle UI in all relevant `fmt.Sprintf` calls to ensure a consistent experience across endpoints.
