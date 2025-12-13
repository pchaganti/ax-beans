---
title: Multi-select mass editing in TUI
status: completed
type: feature
created_at: 2025-12-13T02:27:04Z
updated_at: 2025-12-13T02:27:04Z
---

Add ability to mark multiple beans using space key and apply status/type/priority changes to all selected beans at once.

## Implementation
- Added selection state (map[string]bool) to list model
- Space key toggles selection on current bean
- Green * marker shows before selected beans
- s/t/P keys now work with multiple selected beans
- Selection count shown in footer
- Esc clears selection before clearing filter
- Selection persists when entering detail view