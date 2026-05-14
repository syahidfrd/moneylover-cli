---
name: moneylover-categories
description: "Money Lover: Manage categories (list, add)."
metadata:
  version: 0.1.0
  requires:
    bins:
      - moneylover
    cliHelp: "moneylover categories --help"
---

# categories

> **PREREQUISITE:** Read `../moneylover-shared/SKILL.md` for auth, global flags, and output format.

```bash
moneylover categories <action> [flags]
```

## Commands

### list

List all categories.

```bash
moneylover categories list
```

### add

Add a new category.

```bash
moneylover categories add --name <NAME> --wallet <ID> [flags]
```

| Flag | Required | Default | Description |
|------|----------|---------|-------------|
| `--name` | ✓ | — | Category name |
| `--wallet` | ✓ | — | Wallet ID |
| `--type` | — | 2 | Category type (1=income, 2=expense) |
| `--icon` | — | ic_category_other | Category icon |

## Examples

```bash
moneylover categories list
moneylover categories add --name "Transport" --type 2 --icon ic_category_transport --wallet "57F837F5CC7741728E264465383B5153"
moneylover categories add --name "Salary" --type 1 --wallet "57F837F5CC7741728E264465383B5153"
```
