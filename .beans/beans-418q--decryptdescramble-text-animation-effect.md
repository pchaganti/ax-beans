---
# beans-418q
title: Decrypt/descramble text animation effect
status: completed
type: feature
priority: normal
created_at: 2026-03-14T17:40:55Z
updated_at: 2026-03-14T18:27:49Z
---

Add a typewriter-style text animation where the last few characters 'descramble' from random glyphs into the real text. Apply to workspace descriptions, agent activity labels, and similar short text elements.

## Summary of Changes

- Created `frontend/src/lib/actions/decryptText.ts` — a Svelte action that animates text with a typewriter + descramble effect (characters cycle through random glyphs at the leading edge before locking in)
- Applied the effect to: agent activity labels, tool invocation titles, info messages, subagent activity lines, workspace descriptions, and bean mini cards in the sidebar
- Added `immediate` option so historical/existing content renders instantly on mount, and only fresh/new content gets animated
- Fixed animation timing drift by preserving fractional remainder across frames
