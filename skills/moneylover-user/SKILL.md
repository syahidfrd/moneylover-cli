---
name: moneylover-user
description: "Money Lover: Get user information."
metadata:
  version: 0.1.0
  requires:
    bins:
      - moneylover
    cliHelp: "moneylover user --help"
---

# user

> **PREREQUISITE:** Read `../moneylover-shared/SKILL.md` for auth, global flags, and output format.

```bash
moneylover user <action>
```

## Commands

### info

Get current user information.

```bash
moneylover user info
```

Returns user profile including email, subscription status, device info, and client settings.

## Examples

```bash
moneylover user info
```
