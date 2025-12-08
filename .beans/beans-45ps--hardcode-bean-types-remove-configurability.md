---
title: Hardcode bean types (remove configurability)
status: open
type: task
created_at: 2025-12-08T13:07:39Z
updated_at: 2025-12-08T13:07:39Z
links:
    - parent: beans-58hm
---

## Summary

Remove the configurability of bean types from config.yaml and hardcode the following types:

- **milestone**: A target release or checkpoint
- **epic**: A thematic container for related work
- **feature**: A user-facing capability or enhancement
- **task**: A concrete piece of work to complete
- **bug**: Something that is broken and needs fixing

## Rationale

Simplifies the system by removing unnecessary configurability. The hardcoded types cover all common use cases.

## Checklist

- [ ] Update `internal/config/config.go` to hardcode types instead of reading from config
- [ ] Remove types section from config.yaml handling
- [ ] Update `beans init` to not create types in config
- [ ] Update any validation logic to use hardcoded types
- [ ] Remove `beans types` command if it exists (or mark beans-de5h as canceled)
- [ ] Update prompt.md to reflect the hardcoded types
- [ ] Update tests