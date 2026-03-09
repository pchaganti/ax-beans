---
# beans-424n
title: Add drag-and-drop reordering to backlog view
status: completed
type: feature
priority: normal
created_at: 2026-03-09T17:08:28Z
updated_at: 2026-03-09T17:35:58Z
order: zzz
---

Make beans in the backlog view draggable for manual reordering, reusing the same fractional indexing and ordering mechanism from the board view.

## Summary of Changes

- Extracted shared drag-drop ordering logic from BoardView into `dragOrder.ts` (ensureOrdered, computeOrder, applyDrop)
- Created `backlogDrag.svelte.ts` reactive state module for managing drag state across the recursive BeanItem tree
- Added drag-and-drop to BeanItem.svelte with drop indicators and opacity feedback
- Wired up top-level backlog list in +page.svelte as a drop zone
- Refactored BoardView.svelte to use shared `applyDrop` instead of inline logic
- Added `dragBean` helper to backlog page object and e2e drag-and-drop test

- Added three-zone hover detection (top 25% = reorder above, middle 50% = reparent, bottom 25% = reorder below)
- Added `applyReparent` to `dragOrder.ts` with cycle detection
- Ring highlight visual feedback on reparent target
- E2e test for drag-to-reparent
