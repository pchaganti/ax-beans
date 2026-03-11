---
# beans-mvh3
title: 'Fix worktree agent merge: set receive.denyCurrentBranch before pushing to main'
status: completed
type: bug
priority: normal
created_at: 2026-03-11T21:28:32Z
updated_at: 2026-03-11T21:28:36Z
---

Worktree agents that try to push their branch to main via `git push . HEAD:main` will fail if the main repo has `receive.denyCurrentBranch` set to its default value (refuse). The agent instructions need to include a step to set `receive.denyCurrentBranch updateInstead` before attempting the push.

## Summary of Changes

Added a git config step to the worktree agent merge instructions in `internal/graph/agent_helpers.go` that sets `receive.denyCurrentBranch updateInstead` on the main repo before the agent pushes to main.
