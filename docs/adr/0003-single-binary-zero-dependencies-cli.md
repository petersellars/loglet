# 3. Single Binary, Zero Dependencies CLI

Date: 2025-09-02

## Status

Accepted

## Context
Our CLI is used by developers and CI/CD agents across heterogeneous environments (macOS, Windows, Linux; laptops, build containers, ephemeral runners) where network access may be unreliable or restricted and where pre-installing language runtimes or system packages is undesirable. Past friction has come from transient dependency failures, mismatched toolchain versions, and configuration drift between machines.

To reduce this operational burden and guarantee a consistent user experience, we want the CLI to:

* Ship as a single, self-contained executable.
* Use `//go:embed` to package templates directly into the binary.
* Work fully offline, with no runtime dependency installation.
* Simplify distribution and guarantee a consistent environment.

These goals favor a distribution model that minimizes moving parts and external state.

## Decision
Build and distribute the CLI as one statically linked Go binary with zero runtime dependencies, embedding all required templates and default assets via `//go:embed`.

Key aspects:

* **Static build**: Produce statically linked binaries (disable CGO by default) for supported OS/architectures.

  Example build flags:
  * `CGO_ENABLED=0`
  * `-ldflags "-s -w -X main.version=${VERSION} -X main.commit=${COMMIT} -X main.date=${DATE}"`
* **Embedded assets**: Use Go 1.16+ `embed` package to include required templates, schema files, and default configuration.
  ```
  //go:embed templates/*
  var templateFS embed.FS
  ```
  The CLI loads templates from `templateFS` at runtime; no filesystem lookups or network fetches are required.
* **Offline-first behavior**: All operations that rely on templates or default data read exclusively from the embedded assets. Optional online features (e.g., checking for updates) are strictly opt-in and non-blocking.
* **Deterministic releases**: Versioned, reproducible builds published as checksummed artifacts per platform (e.g., `cli_<version>_<os>_<arch>`), with SBOM and signature where applicable.
* **Simple install paths**: Users obtain the binary via direct download or a thin package that only drops the binary onto `PATH` (no post-install hooks, no dependency resolution).

## Consequences

**Positive**

* **Reliability & portability**: Fewer environmental assumptions; works on clean hosts, air-gapped networks, and ephemeral CI runners.
* **Fast startup & execution**: No runtime initialization of external tooling or network calls for templates.
* **Security surface reduction**: No dynamic dependency loading; fewer supply-chain and path injection vectors. Easier to attest, scan, and verify a single artifact.
* **Operational simplicity**: Distribution is one file; support is simpler (users can report the exact version/commit).
* **Determinism**: Embedded templates ensure commands generate consistent output across machines and over time.

**Negative / Trade-offs**

* **Binary size**: Embedding assets increases artifact size (potentially tens of MB). This affects cold downloads and storage.
* **Update cadence**: Any template change requires a new release. There is no out-of-band template hotfix unless we add an override mechanism.
* **Extensibility limits**: No dynamic plugins by default; extending behavior may require rebuilding. (We avoid Go’s runtime plugin system for portability and stability reasons.)
* **Localization/config churn**: Multiple template sets (e.g., per-team) can lead to more frequent releases unless we provide an “external override” flag.
* **CGO-dependent features unavailable**: With `CGO_ENABLED=0`, features that require native libraries (e.g., some cryptographic modules or OS-specific calls) may be constrained. If such a need arises, we’ll gate it with a separate build.

**Mitigations**

* Provide a **user-specified override path** (e.g., `--templates-dir`) that, if present, supersedes embedded assets for power users—while keeping the defaults embedded and offline.
* Implement **semantic versioning** and clear **release notes** to communicate when embedded templates change.
* Offer **delta-aware update helpers** (optional) that download a new binary only when versions differ.
* Include **telemetry toggle** defaulting to off; no network access is required for baseline commands.

## Alternatives Considered
1. **Binary + external templates on disk**
    * *Pros*: Smaller binary; templates can be patched without a new build.
    * *Cons*: Reintroduces file management, path issues, and drift. Breaks offline guarantees. 
    
      **Rejected** due to operational fragility.
2. **Runtime dependency on a package manager (Homebrew, apt, npm, pip, etc.)**
    * *Pros*: Familiar distribution channels; easier updates.
    * *Cons*: Requires bootstrapping a package manager and online access; increases variability across hosts. 
      
      **Rejected** to preserve offline, zero-dep operation. (We may still publish taps/repos that simply install the same single binary.)
3. **Container image distribution**
    * *Pros*: Strong isolation; consistent environment.
    * *Cons*: Requires Docker/OCI runtime, which some environments lack; worse UX for local dev; mounting files can be awkward. 
    
      **Rejected** as the primary vehicle; acceptable as an optional channel.
4. **Dynamic plugin system (Go plugins or exec-based plugin discovery)**
    * *Pros*: Extensibility without rebuilding the core.
    * *Cons*: Go plugins are OS/arch/toolchain fragile; exec-based plugins reintroduce dependencies. 
    
      **Deferred** until a concrete extensibility need emerges.
5. **Polyglot CLI (wrapper around Python/Node)**
    * *Pros*: Rapid iteration, large ecosystem.
    * *Cons*: Requires language runtimes and package resolution; undermines offline and determinism goals. 
    
      **Rejected.**

## Related Links
* Go embed package docs (for `//go:embed`).
* Architectural vision and design principles: [ARCHITECTURE.md](../../ARCHITECTURE.md)
<!-- 
* Build recipe (internal): `build/release.md` — includes `CGO_ENABLED=0`, `-ldflags`, cross-compile targets, checksums/signing, and SBOM generation.
* Security posture (internal): `docs/security/supply-chain.md` — artifact signing, provenance, and scanning.
* Issue tracker epics (internal): “CLI Offline-First Distribution”, “Template Embedding & Overrides”. 
-->
