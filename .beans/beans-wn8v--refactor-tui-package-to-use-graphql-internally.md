---
title: Refactor TUI package to use GraphQL internally
status: completed
type: task
created_at: 2025-12-09T07:43:17Z
updated_at: 2025-12-09T08:02:43Z
---

Refactor the TUI package to use the GraphQL resolver pattern internally, consistent with how `beans list` already works.

## Checklist

- [x] Add new GraphQL fields (`duplicates`, `duplicatedBy`, `related`, `relatedTo`) to schema.graphqls
- [x] Implement new relationship resolvers in schema.resolvers.go
- [x] Run `mise codegen` to regenerate GraphQL code
- [x] Add resolver to App struct in tui.go
- [x] Refactor listModel to use GraphQL resolver
- [x] Refactor detailModel to use GraphQL resolver for link resolution
- [x] Update collectTagsWithCounts() to use resolver
- [x] Update beansChangedMsg handler to use resolver
- [x] Run tests to verify changes