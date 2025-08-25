# Architecture & Design Principles

This document explains the **architectural vision** and **design principles** behind this project.  
It serves as a reference for contributors when making changes, ensuring we maintain consistency, stability, and long-term maintainability.

---

## Overview

This project provides a **scaffolding and automation tool** that prioritizes:

- **Ease of use** → a single binary, zero dependencies, works offline.
- **Robustness** → deterministic behavior, secure defaults, predictable results.
- **Extensibility** → pluggable CLI-based extensions without compromising stability.

Our architecture follows a **hexagonal (ports and adapters) style**,separating **core business rules** from **infrastructure concerns**.
This allows us to evolve the tool without breaking users or forcing them to adopt internal implementation details.

---

## Goals & Principles

### **1. Single Binary, Zero Dependencies**
> Ship as a single, self-contained executable.

- Uses `//go:embed` to package templates directly into the binary.
- Works fully offline, with no runtime dependency installation.
- Simplifies distribution and guarantees a consistent environment.

---

### **2. Hexagonal Architecture**
> Separate the **domain use-cases** from **adapters** and infrastructure.

- The **core domain logic** knows nothing about I/O, CLI parsing, or template storage.
- **Adapters** (CLI, file system, network, etc.) handle integrations.
- Enables testing the domain in isolation.
- Keeps the codebase flexible and maintainable.

---

### **3. Deterministic & Idempotent Scaffolding**
> Running the same command twice should yield the same result.

- Uses checksums to ensure generated outputs are verifiable.
- Supports `--dry-run` and `--force` flags for safe experimentation.
- Guarantees repeatable scaffolding across environments.

---

### **4. Secure by Default**
> Security isn’t optional — it’s built-in.

- Templates are **signed** and **version-locked** to prevent tampering.
- Deployment configs follow **least-privilege** principles.
- Security-conscious defaults avoid accidental vulnerabilities.

---

### **5. Evolution-Friendly**
> Designed to change safely over time.

- Templates use a **versioned schema** to ensure backward compatibility.
- **Migration hooks** help upgrade templates without breaking existing setups.
- **Fitness functions** and tests preserve quality as the system evolves.

---

### **6. Pluggable, but Stable**
> Extensible, without compromising reliability.

- Extensions are provided via **well-defined CLI contracts**.
- Avoids Go’s native plugin mechanism to ensure **portability** and **stability**.
- Allows contributors to build powerful add-ons without coupling to internal APIs.

---

## Contributing with These Principles in Mind

When adding new features or changing existing behavior:

- **Start from the goals above** — does the change align with our principles?
- **Prefer extension points** over hardcoding assumptions.
- **Preserve determinism** — avoid introducing hidden side effects.
- **Keep security in mind** — secure by default, never optional.

Before making major changes, please read the [CONTRIBUTING.md](./CONTRIBUTING.md) guide.

---

## Architecture Decision Records (ADRs)

Significant architectural decisions are documented as ADRs under [`/docs/adr`](./docs/adr).

These records explain **why** we made certain choices and provide historical context for future contributors

---

## Future Architecture Notes

- The **hexagonal structure** will be enforced over time with tooling and documentation.

---

## Related Documents

- [Architecture Decision Records](./docs/adr/README.md) - document significant choices
- [CONTRIBUTING.md](./CONTRIBUTING.md) — how to propose changes
- [README.md](./README.md) — quick start guide for users
