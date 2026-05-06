# atlas-mob-vault-index

`atlas-mob-vault-index` is a Go project for Mobile workflows. It turns create a Go reference implementation for vault workflows, centered on storage recovery, log and snapshot fixtures, and replay consistency checks into a small local model with readable fixtures and a direct verification command.

## Reading Atlas Mob Vault Index

Start with the README, then open `metadata/project.json` to check the constants behind the examples. After that, `fixtures/cases.csv` shows the compact path and `examples/extended_cases.csv` gives a wider look at the same rule.

## Purpose

The goal is to capture the core behavior in code and make the surrounding assumptions obvious. A reader should be able to run the verifier, open the fixtures, and understand why each decision was made.

## What It Does

- Models local state with deterministic scoring and explicit review decisions.
- Uses fixture data to keep sync pressure changes visible in code review.
- Includes extended examples for form constraints, including `surge` and `degraded`.
- Documents offline paths tradeoffs in `docs/operations.md`.
- Runs locally with a single verification command and no external credentials.

## Design Sketch

The project is organized around a compact model rather than a large framework. Inputs are scored, classified, and checked against golden fixtures. The constants live in code and are mirrored in metadata so documentation drift is easy to catch. The Go layout uses small packages and table-oriented tests so the behavior stays easy to follow.

## Files Worth Reading

- `policy`: Go package with the core model
- `cmd`: small command entry point
- `fixtures`: compact golden scenarios
- `examples`: expanded scenario set
- `metadata`: project constants and verification metadata
- `docs`: operations and extension notes
- `scripts`: local verification and audit commands
- `go.mod`: Go module metadata

## Setup

Use a normal shell with Go available on `PATH`. The verifier is written as a PowerShell script because the portfolio was assembled on Windows.

## Usage

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

This runs the language-level build or test path against the compact fixture set.

## Verification

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/audit.ps1
```

The audit command checks repository structure and README constraints before it delegates to the verifier.

## Fixture Notes

The extended cases are not random smoke tests. `degraded` keeps pressure on the review path, while `surge` shows the model when capacity and weight are strong enough to clear the threshold.

## Limits

The examples cover useful edges, not every edge. A larger version would add malformed-input tests, richer reports, and deeper domain parsers.

## Next Directions

- Add a short report command that prints the score breakdown for a single scenario.
- Add malformed input fixtures so the failure path is as visible as the happy path.
- Split the scoring constants into a typed configuration object and validate it before use.
- Add one more mobile workflows fixture that focuses on a malformed or borderline input.
