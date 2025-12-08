---
title: Change JSON output to return array directly
status: completed
type: task
created_at: 2025-12-08T14:50:06Z
updated_at: 2025-12-08T14:51:21Z
---

Change the JSON output format from wrapped object to plain array.

## Current (wrapped)
```json
{
  "success": true,
  "beans": [...],
  "count": 3
}
```

## New (direct array)
```json
[...]
```

This allows the intuitive jq pattern:
```bash
beans list --json | jq -r '.[] | ...'
```

## Checklist
- [ ] Update list command JSON output
- [ ] Update any other commands that use the same pattern
- [ ] Update tests
- [ ] Revert prompt.md changes (no longer needed)