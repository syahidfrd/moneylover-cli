# moneylover-cli

Unofficial CLI for the [Money Lover](https://moneylover.me) personal finance app.

## Install

### Homebrew (macOS/Linux)

```bash
brew install syahidfrd/tap/moneylover-cli
```

### Go install

```bash
go install github.com/syahidfrd/moneylover-cli@latest
```

### Shell script

```bash
curl -sSL https://raw.githubusercontent.com/syahidfrd/moneylover-cli/main/install.sh | bash
```

### Binary download

Download from [GitHub Releases](https://github.com/syahidfrd/moneylover-cli/releases).

## Setup

```bash
# Step 1: Get OAuth login URL
moneylover auth login
# → {"login_url":"https://oauth.moneylover.me/auth?..."}

# Step 2: Open the URL in browser, login, then save the redirect URL
moneylover auth callback "https://web.moneylover.me/login?access_token=...&refresh_token=...&expire=...&status=true"
# → {"status":"login_success"}
```

> Note: After login, the browser may show a "not found" page. This is expected. Just copy the full URL from the address bar.

## Usage

```bash
# Auth
moneylover auth login
moneylover auth callback "<redirect_url>"
moneylover auth status

# Wallets
moneylover wallets list
moneylover wallets add --name "BCA" --currency-id 44
moneylover wallets edit --id <wallet_id> --name "New Name"
moneylover wallets delete --id <wallet_id>

# Transactions
moneylover transactions list --wallet all --start 2026-05-01 --end 2026-05-31
moneylover transactions add --wallet <wallet_id> --category <category_id> --amount 50000 --note "Lunch" --date 2026-05-14
moneylover transactions edit --id <tx_id> --wallet <wallet_id> --category <category_id> --amount 60000 --note "Dinner" --date 2026-05-14
moneylover transactions delete --id <tx_id>
moneylover transactions debts --wallets <wallet_id_1>,<wallet_id_2>

# Categories
moneylover categories list
moneylover categories add --name "Transport" --type 2 --icon ic_category_transport --wallet <wallet_id>
moneylover categories edit --id <category_id> --name "New Name"
moneylover categories delete --id <category_id>

# Budgets
moneylover budgets list
moneylover budgets add --category <category_id> --amount 500000 --wallet <wallet_id> --start 2026-05-01 --end 2026-05-31 --repeat
moneylover budgets edit --id <budget_id> --category <category_id> --amount 600000 --wallet <wallet_id> --start 2026-05-01 --end 2026-05-31
moneylover budgets delete --id <budget_id>

# Events
moneylover events list

# User
moneylover user info
```

## Configuration

| Option | Env Var | Default | Description |
|--------|---------|---------|-------------|
| `--token-path` | `MONEYLOVER_TOKEN_PATH` | `~/.config/moneylover/token.json` | Path to token file |

## Output

All commands output JSON to stdout. Errors are written to stderr as JSON:

```json
{"error": "not authenticated. Run 'moneylover auth login' first"}
```

Exit code `0` on success, `1` on error.

## Disclaimer

This is an unofficial tool that uses a reverse-engineered API from Money Lover's web app. It is not affiliated with or endorsed by Finsify (the company behind Money Lover). The API may change at any time without notice.

## License

MIT
