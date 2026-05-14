---
name: moneylover-transactions
description: "Money Lover: Manage transactions (list, add, debts)."
metadata:
  version: 0.1.0
  requires:
    bins:
      - moneylover
    cliHelp: "moneylover transactions --help"
---

# transactions

> **PREREQUISITE:** Read `../moneylover-shared/SKILL.md` for auth, global flags, and output format.

```bash
moneylover transactions <action> [flags]
```

## Commands

### list

List transactions by wallet and date range.

```bash
moneylover transactions list --start <DATE> --end <DATE> [--wallet <ID>]
```

| Flag | Required | Default | Description |
|------|----------|---------|-------------|
| `--wallet` | — | all | Wallet ID or "all" |
| `--start` | ✓ | — | Start date (YYYY-MM-DD) |
| `--end` | ✓ | — | End date (YYYY-MM-DD) |

### add

Add a new transaction.

```bash
moneylover transactions add --wallet <ID> --category <ID> --amount <N> --date <DATE> [--note <TEXT>]
```

| Flag | Required | Default | Description |
|------|----------|---------|-------------|
| `--wallet` | ✓ | — | Wallet ID |
| `--category` | ✓ | — | Category ID |
| `--amount` | ✓ | — | Transaction amount |
| `--date` | ✓ | — | Display date (YYYY-MM-DD) |
| `--note` | — | "" | Transaction note/description |

### debts

List debt/loan transactions.

```bash
moneylover transactions debts --wallets <IDs>
```

| Flag | Required | Description |
|------|----------|-------------|
| `--wallets` | ✓ | Comma-separated wallet IDs |

## Examples

```bash
moneylover transactions list --wallet all --start 2026-05-01 --end 2026-05-31
moneylover transactions list --wallet "57F837F5CC7741728E264465383B5153" --start 2026-05-01 --end 2026-05-14
moneylover transactions add --wallet "57F837F5CC7741728E264465383B5153" --category "A2AC40F7745746D1843B11E6F7640B50" --amount 50000 --note "Lunch" --date 2026-05-14
moneylover transactions debts --wallets "57F837F5CC7741728E264465383B5153,065EFEA4415444CE9B289C1328A6EEB3"
```
