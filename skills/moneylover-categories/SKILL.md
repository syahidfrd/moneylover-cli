---
name: moneylover-categories
description: "Money Lover: Manage categories (list, add, edit, delete)."
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

### edit

Edit a category.

```bash
moneylover categories edit --id <ID> --name <NAME> [--icon <ICON>]
```

| Flag | Required | Default | Description |
|------|----------|---------|-------------|
| `--id` | ✓ | — | Category ID |
| `--name` | ✓ | — | Category name |
| `--icon` | — | — | Category icon |

### delete

Delete a category.

```bash
moneylover categories delete --id <ID>
```

| Flag | Required | Description |
|------|----------|-------------|
| `--id` | ✓ | Category ID |

## Examples

```bash
moneylover categories list
moneylover categories add --name "Transport" --type 2 --icon ic_category_transport --wallet "57F837F5CC7741728E264465383B5153"
moneylover categories add --name "Salary" --type 1 --wallet "57F837F5CC7741728E264465383B5153"
moneylover categories edit --id "253310201C3B408585E3CFDFD68E83A3" --name "Education"
moneylover categories delete --id "253310201C3B408585E3CFDFD68E83A3"
```
