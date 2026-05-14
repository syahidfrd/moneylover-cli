---
name: moneylover-budgets
description: "Money Lover: Manage budgets (list, add, edit, delete)."
metadata:
  version: 0.1.0
  requires:
    bins:
      - moneylover
    cliHelp: "moneylover budgets --help"
---

# budgets

> **PREREQUISITE:** Read `../moneylover-shared/SKILL.md` for auth, global flags, and output format.

```bash
moneylover budgets <action> [flags]
```

## Commands

### list

List all budgets.

```bash
moneylover budgets list
```

### add

Add a new budget.

```bash
moneylover budgets add --category <ID> --amount <N> --wallet <ID> --start <DATE> --end <DATE> [--repeat]
```

| Flag | Required | Default | Description |
|------|----------|---------|-------------|
| `--category` | ✓ | — | Category ID |
| `--amount` | ✓ | — | Budget amount |
| `--wallet` | ✓ | — | Wallet ID |
| `--start` | ✓ | — | Start date (YYYY-MM-DD) |
| `--end` | ✓ | — | End date (YYYY-MM-DD) |
| `--repeat` | — | false | Repeat monthly |

### edit

Edit a budget.

```bash
moneylover budgets edit --id <ID> --category <ID> --amount <N> --wallet <ID> --start <DATE> --end <DATE> [--repeat]
```

| Flag | Required | Default | Description |
|------|----------|---------|-------------|
| `--id` | ✓ | — | Budget ID |
| `--category` | ✓ | — | Category ID |
| `--amount` | ✓ | — | Budget amount |
| `--wallet` | ✓ | — | Wallet ID |
| `--start` | ✓ | — | Start date (YYYY-MM-DD) |
| `--end` | ✓ | — | End date (YYYY-MM-DD) |
| `--repeat` | — | false | Repeat monthly |

### delete

Delete a budget.

```bash
moneylover budgets delete --id <ID>
```

| Flag | Required | Description |
|------|----------|-------------|
| `--id` | ✓ | Budget ID |

## Examples

```bash
moneylover budgets list
moneylover budgets add --category "253310201C3B408585E3CFDFD68E83A3" --amount 500000 --wallet "57F837F5CC7741728E264465383B5153" --start 2026-05-01 --end 2026-05-31 --repeat
moneylover budgets edit --id "budget123" --category "253310201C3B408585E3CFDFD68E83A3" --amount 600000 --wallet "57F837F5CC7741728E264465383B5153" --start 2026-06-01 --end 2026-06-30
moneylover budgets delete --id "budget123"
```
