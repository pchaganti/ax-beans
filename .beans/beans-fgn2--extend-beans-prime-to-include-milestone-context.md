---
# beans-fgn2
title: Extend beans prime to include milestone context
status: todo
type: feature
created_at: 2025-12-14T11:37:59Z
updated_at: 2025-12-14T11:37:59Z
---

## Summary

The `beans prime` command generates a prompt that helps agents understand the beans system. It should be extended to inject additional project-specific knowledge, specifically:

1. **List of milestones** - Show all milestones defined in the project
2. **Current milestone(s)** - Highlight which milestone(s) are currently in-progress
3. **Milestone progress** - Optionally show completion status (e.g., X of Y children completed)

## Motivation

When an agent starts working on a project, understanding the current milestone context helps them:
- Prioritize work that aligns with current goals
- Understand the broader project roadmap
- Make better decisions about which beans to pick up

## Checklist

- [ ] Query for all milestones in the project
- [ ] Identify in-progress milestones
- [ ] Format milestone information in a clear, readable way for the prompt
- [ ] Include milestone progress (completed/total children) if available
- [ ] Add the milestone section to the prime output
- [ ] Test with projects that have no milestones (graceful handling)
- [ ] Test with projects that have multiple in-progress milestones