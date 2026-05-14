---
name: moneylover-wallets
description: "Money Lover: Manage wallets (list, add, edit, delete)."
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

### edit

Edit an existing wallet.

```bash
moneylover wallets edit --id <ID> --name <NAME> [flags]
```

| Flag | Required | Default | Description |
|------|----------|---------|-------------|
| `--id` | ✓ | — | Wallet ID |
| `--name` | ✓ | — | Wallet name |
| `--icon` | — | icon_80 | Wallet icon |
| `--currency-id` | — | 44 | Currency ID |
| `--account-type` | — | 0 | Account type |

### delete

Delete a wallet.

```bash
moneylover wallets delete --id <ID>
```

| Flag | Required | Description |
|------|----------|-------------|
| `--id` | ✓ | Wallet ID |

## Examples

```bash
moneylover wallets list
moneylover wallets add --name "BCA" --currency-id 44 --icon icon_80
moneylover wallets add --name "Savings" --account-type 5 --balance 1000000
moneylover wallets edit --id "57F837F5CC7741728E264465383B5153" --name "BCA Main"
moneylover wallets delete --id "57F837F5CC7741728E264465383B5153"
```
