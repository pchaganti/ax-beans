---
title: Fix roadmap to show all planned work
status: completed
type: bug
priority: normal
created_at: 2025-12-09T09:23:25Z
updated_at: 2025-12-09T09:25:35Z
---

The roadmap command currently only shows items under milestones or epics. It should show ALL planned work:

- Sorted by milestone (with a section at the end for beans not assigned to a milestone)
- Within each milestone section, grouped by epic (with a section for beans not assigned to an epic)

This means orphan items (features, bugs, tasks not under any milestone) should appear in an 'Unscheduled' section at the bottom.