# Review Journal

This journal records the domain cases that matter before widening the public API.

The local checks classify each case as `ship`, `watch`, or `hold`. That gives the project a small review vocabulary that matches its mobile workflows focus without claiming live deployment or external usage.

## Cases

- `baseline`: `form pressure`, score 155, lane `ship`
- `stress`: `sync drift`, score 205, lane `ship`
- `edge`: `local state`, score 186, lane `ship`
- `recovery`: `conflict cost`, score 200, lane `ship`
- `stale`: `form pressure`, score 212, lane `ship`

## Note

The useful failure mode here is a wrong decision on a named case, not a vague style disagreement.
