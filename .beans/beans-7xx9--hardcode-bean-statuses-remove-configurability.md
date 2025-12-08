---
title: Hardcode bean statuses (remove configurability)
status: open
type: task
created_at: 2025-12-08T13:07:39Z
updated_at: 2025-12-08T13:07:39Z
links:
    - parent: beans-58hm
---

## Summary

Remove the configurability of bean statuses from config.yaml and hardcode the following statuses:

- **not-ready**: Not yet ready to be worked on (blocked, needs more info, etc.)
- **ready**: Ready to be worked on
- **in-progress**: Currently being worked on
- **completed**: Finished successfully
- **canceled**: Will not be done

## Rationale

Simplifies the system by removing unnecessary configurability. The hardcoded statuses cover all workflow states.

## Checklist

- [ ] Update `internal/config/config.go` to hardcode statuses instead of reading from config
- [ ] Remove statuses section from config.yaml handling
- [ ] Update `beans init` to not create statuses in config
- [ ] Update any validation logic to use hardcoded statuses
- [ ] Remove any `beans statuses` command if planned
- [ ] Update prompt.md to reflect the hardcoded statuses
- [ ] Update tests
- [ ] Consider migration path for existing beans using old status names (open → ready, done → completed)