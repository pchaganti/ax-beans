---
title: Refactor beans prompt to use text/template
status: completed
type: task
priority: normal
created_at: 2025-12-09T09:02:36Z
updated_at: 2025-12-09T09:03:37Z
---

Convert cmd/prompt.go from using strings.Builder to Go's text/template package for generating the dynamic sections of the prompt output.