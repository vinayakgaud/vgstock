# VGStock — Decisions

Log of architecture and tooling decisions as they're made. Each entry: date,
decision, reason. Append, never rewrite history.

## 2026-09-22 — Foundations

| # | Decision | Choice | Why |
|---|----------|--------|-----|
| 1 | Repository layout | Monorepo, language-first (`backend/`, `frontend/`, `ml/`) | One clone/history; simple mental map; splits into service-first with a rename later if needed |
| 2 | Backend language | Go | Concurrency for streaming ticks, static binary, small footprint |
| 3 | Frontend framework | Next.js + TypeScript | React ecosystem, charting/components (Lightweight Charts, shadcn/ui), SSR |
| 4 | Package manager | pnpm (Node runtime) | Fast, strict dependency isolation, industry standard for Next.js on Vercel |
| 5 | License | Proprietary (all rights reserved) | Personal/commercial product; no rights granted without written agreement |
| 6 | Deployment | Per-service, not one chunk | Backend (Go) → Railway; Frontend (Next.js) → Vercel; ML sidecar later → own deploy/scaling |
| 7 | Scaling model | Services scale independently | Each service is its own deployable even though code lives in one repo; CPU-heavy ML isolated |
| 8 | Polyglot boundary | Language-first folders, per-service deployables | Repo organized by language; each Go service ships as its own image/config |
| 9 | Git hooks | lefthook only (root `.lefthook.yml`), husky avoided | Single hook manager; husky would fight lefthook over `.git/hooks` |
| 10 | CI | `.github/workflows/ci.yml` — Go vet/test/build + Next lint/build on push/PR | Same checks as local `task verify`, run in the cloud |
| 11 | ESLint | Pinned to `eslint@9.39.5` (9.x line), not v10 | `eslint-plugin-react@7.37.5` incompatible with ESLint 10 (`getFilename` crash); 9.x is the only working major for this plugin stack |
| 12 | Go entrypoint layout | `backend/cmd/vgstock/` (binary named after product) | `cmd/api` felt vague/"just an API"; product entrypoint carries the brand; future tools get their own `cmd/<name>` |
| 13 | Go bootstrap Taskfile commands | Adopted from prior project, Windows-safe: fmt/fmt:check/test/test:race/build/run/tidy | `gofmt -l` alone exits 0 (decorative); real check must fail on unformatted. Local shell is cmd, so checks wrap PowerShell explicitly |
| 14 | HTTP router | Go 1.22+ stdlib `http.ServeMux`, zero web deps | Health route + few endpoints don't justify a dep; chi (or others) sits on net/http so migration later is an afternoon, not a rewrite. Revisit when route/middleware composition hurts |