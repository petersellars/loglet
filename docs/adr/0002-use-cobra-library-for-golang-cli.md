# 2. Use Cobra Library for Golang CLI

Date: 2025-08-24

## Status

Accepted

## Context
We are developing a **command-line interface (CLI)** tool in **Golang** to manage and automate workflows. The CLI needs to support:

* Multiple commands and subcommands
* Consistent flag parsing
* Built-in help and usage messages
* A scalable structure that can grow as new features are added

While Go’s standard library provides `flag` for basic argument parsing, it lacks support for advanced CLI requirements like subcommands, autocompletion, and structured command hierarchies. We needed a mature, well-supported, and widely adopted CLI framework to improve developer productivity and deliver a better user experience.

## Decision
We have decided to use the [Cobra](https://github.com/spf13/cobra) library to implement the CLI.

Rationale:

* Cobra provides **first-class support for commands and subcommands**, making it easier to structure complex CLIs.
* It automatically **generates help and usage messages**, reducing the need for boilerplate code.
* Built-in **flag parsing** integrates cleanly with Go’s `pflag` package.
* Supports **bash/zsh/fish autocompletion** and **man page generation** out of the box.
* Maintained by **spf13**, widely used in the Go ecosystem, and battle-tested in tools like:
  * [Kubernetes kubectl](https://kubernetes.io/docs/reference/kubectl/)
  * [Helm](https://helm.sh/)
  * [Hugo](https://gohugo.io/)

By adopting Cobra, we ensure our CLI is **scalable, user-friendly**, and **maintainable**.

## Consequences

### Positive Impacts

* **Improved developer productivity**: Reduces boilerplate and simplifies command/flag handling.
* **Better user experience**: Automatic help, usage, and completion features improve usability.
* **Easier maintenance and scalability**: Clear separation of commands makes it easy to add or remove features.
* **Community support**: Widely adopted with extensive documentation and examples.

### Negative Impacts / Trade-offs

* **Additional dependency**: Increases the project’s reliance on an external library.
* **Learning curve**: Developers need to become familiar with Cobra’s patterns and APIs.
* **Opinionated structure**: Cobra encourages a particular way of organizing CLI code, which may require refactoring existing code if migrating later.

## Alternatives Considered

1. **Go standard `flag` package**
    * Pros: No external dependency, part of the stdlib.
    * Cons: Lacks support for subcommands, complex flag parsing, and built-in help generation.

2. **urfave/cli**
    * Pros: Mature, widely used, simpler API.
    * Cons: Less active maintenance than Cobra, fewer advanced features like autocompletion and man page generation.

3. **Kingpin**
    * Pros: Declarative syntax, minimal boilerplate.
    * Cons: Project is effectively deprecated, not recommended for new projects.

Cobra was chosen for its **maturity, active maintenance, feature set, and ecosystem adoption**.

## Related Links
* [Cobra GitHub Repository](https://github.com/spf13/cobra)
* [Kubernetes CLI (kubectl) using Cobra](https://github.com/kubernetes/kubectl)
* [Helm CLI built with Cobra](https://helm.sh/docs/helm/)
