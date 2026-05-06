# atlas-mob-vault-index

`atlas-mob-vault-index` explores mobile workflows with a small Go codebase and local fixtures. The technical goal is to create a Go reference implementation for vault workflows, centered on storage recovery, log and snapshot fixtures, and replay consistency checks.

## Why It Exists

The point is to make a small domain rule concrete enough that a reader can change it and immediately see what broke.

## Atlas Mob Vault Index Review Notes

For a quick review, compare `form pressure` with `form pressure` before reading the middle cases.

## Features

- `fixtures/domain_review.csv` adds cases for form pressure and sync drift.
- `metadata/domain-review.json` records the same cases in structured form.
- `config/review-profile.json` captures the read order and the two review questions.
- `examples/atlas-mob-vault-walkthrough.md` walks through the case spread.
- The Go code includes a review path for `form pressure` and `form pressure`.
- `docs/field-notes.md` explains the strongest and weakest cases.

## Architecture Notes

The repository has two validation layers: the original compact policy fixture and the domain review fixture. They are separate so one can change without hiding failures in the other.

The Go implementation avoids hidden state so fixture changes are easy to reason about.

## Usage

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

## Tests

That command is also the regression path. It verifies the domain cases and catches mismatches between the CSV, metadata, and code.

## Limitations And Roadmap

The fixture set is small enough to audit by hand. The next useful expansion is malformed input coverage, not extra surface area.
