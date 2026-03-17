---
# beans-pw99
title: Replace model chooser with effort/thinking mode control
status: completed
type: feature
priority: normal
created_at: 2026-03-17T07:20:45Z
updated_at: 2026-03-17T07:28:19Z
---

Remove the model switcher from the agent composer UI and replace it with effort level controls (low, medium, high, max). When effort changes on a running session, kill and respawn the process. Disable buttons while the agent is busy.

## Summary of Changes

Replaced the model chooser (Sonnet/Opus/Haiku) with thinking effort controls (Low/Med/High/Max) across the full stack:

**Backend:**
- Renamed `Session.Model` → `Session.Effort` in agent types
- Replaced `SetModel()` with `SetEffort()` in manager — now kills running processes on change (like SetPlanMode/SetActMode)
- Changed `buildClaudeArgs` to emit `--effort <level>` instead of `--model <name>`
- Updated GraphQL schema: replaced `model` field and `setAgentModel` mutation with `effort` and `setAgentEffort`
- Removed the `model` parameter from `sendAgentMessage` mutation
- Updated resolver, agent helpers, and tests

**Frontend:**
- Replaced model selector buttons with effort level buttons (Low/Med/High/Max)
- High is highlighted by default (when no effort is set)
- Buttons are disabled while agent is running
- Updated store, operations, and component wiring
- Removed model-related GraphQL operations
