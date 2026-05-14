---
name: moneylover-shared
description: "moneylover CLI: Shared patterns for authentication, global flags, and output formatting."
metadata:
  version: 0.1.0
  requires:
    bins:
      - moneylover
---

# moneylover — Shared Reference

## Installation

The `moneylover` binary must be on `$PATH`. Install via:

```bash
# Homebrew
brew install syahidfrd/tap/moneylover-cli

# Go install
go install github.com/syahidfrd/moneylover-cli@latest
```

## Authentication

```bash
# Step 1: Get OAuth login URL
moneylover auth login
# → {"login_url":"https://oauth.moneylover.me/auth?..."}

# Step 2: User opens URL in browser, logs in, gets redirected
# Step 3: Save token from redirect URL
moneylover auth callback "<redirect_url>"
# → {"status":"login_success"}

# Check auth status
moneylover auth status
# → {"user_id":"...","client":"...","expires_at":"...","valid":true}
```

## Global Flags

| Flag | Env Var | Default | Description |
|------|---------|---------|-------------|
| `--token-path` | `MONEYLOVER_TOKEN_PATH` | `~/.config/moneylover/token.json` | Path to token file |

## Output Format

All commands output JSON to stdout.

Success:
```json
{
  "data": { ... }
}
```

Error (stderr):
```json
{"error": "not authenticated. Run 'moneylover auth login' first"}
```

Exit codes: `0` = success, `1` = error.

## CLI Syntax

```bash
moneylover <resource> <action> [flags]
```

## Important Notes

- All dates use `YYYY-MM-DD` format
- Wallet ID can be a specific ID or `"all"` for transactions list
- Category type: `1` = Income, `2` = Expense
- Wallet account type: `0` = Normal, `2` = Linked, `4` = Credit card, `5` = Saving
