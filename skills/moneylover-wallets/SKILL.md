---
name: moneylover-wallets
description: "Money Lover: Manage wallets (list, add)."
metadata:
  version: 0.1.0
  requires:
    bins:
      - moneylover
    cliHelp: "moneylover wallets --help"
---

# wallets

> **PREREQUISITE:** Read `../moneylover-shared/SKILL.md` for auth, global flags, and output format.

```bash
moneylover wallets <action> [flags]
```

## Commands

### list

List all wallets.

```bash
moneylover wallets list
```

### add

Add a new wallet.

```bash
moneylover wallets add --name <NAME> [flags]
```

| Flag | Required | Default | Description |
|------|----------|---------|-------------|
| `--name` | ✓ | — | Wallet name |
| `--currency-id` | — | 44 (IDR) | Currency ID |
| `--icon` | — | icon_80 | Wallet icon |
| `--account-type` | — | 0 | Account type (0=normal, 2=linked, 4=credit card, 5=saving) |
| `--balance` | — | — | Initial balance amount |

## Examples

```bash
moneylover wallets list
moneylover wallets add --name "BCA" --currency-id 44 --icon icon_80
moneylover wallets add --name "Savings" --account-type 5 --balance 1000000
```
