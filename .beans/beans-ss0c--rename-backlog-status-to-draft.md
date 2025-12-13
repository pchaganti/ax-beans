---
title: Rename 'backlog' status to 'draft'
status: completed
type: task
priority: normal
created_at: 2025-12-13T01:58:26Z
updated_at: 2025-12-13T02:01:05Z
---

Rename the backlog status to draft to better communicate that these beans need refinement before being actionable.

## Changes
- Rename backlog → draft in status definitions
- Update prompt.tmpl to exclude draft status when finding work
- Change default status for new beans from backlog to todo