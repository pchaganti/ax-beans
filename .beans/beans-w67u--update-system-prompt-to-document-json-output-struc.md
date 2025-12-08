---
title: Update system prompt to document JSON output structure
status: scrapped
type: task
created_at: 2025-12-08T14:48:53Z
updated_at: 2025-12-08T14:51:21Z
---

The beans CLI JSON output wraps results in an object with `success`, `beans`, and `count` fields. The system prompt should clarify that agents need to use `.beans[]` (not `.[] `) when parsing with jq.

## Problem

Running:
```bash
beans list --json | jq -r '.[] | ...'
```

Fails with:
```
jq: error: Cannot index boolean with string "title"
```

Because `.[] ` iterates over all top-level values including the boolean `success` field.

## Solution

Update the system prompt (prompt.md) to show the correct jq pattern:
```bash
beans list --json | jq -r '.beans[] | ...'
```

## Checklist

- [ ] Update prompt.md to document the JSON structure
- [ ] Add example jq commands showing correct usage