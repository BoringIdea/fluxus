# Fluxus Monorepo

Fluxus is an open-source collection of applications and tooling that power the Fluxus ecosystem, including the public website, backend services, smart contracts, and The Graph subgraph. This repository consolidates the previously standalone codebases into a single pnpm-powered monorepo for easier collaboration and consistent tooling.

## Packages

- `website` – Next.js front-end application for the public Fluxus experience.
- `backend` – NestJS API, websocket gateway, and supporting services.
- `contracts` – Hardhat/Foundry smart contracts and deployment scripts.
- `subgraph` – The Graph subgraph definitions, generated types, and tests.

## Getting Started

### Prerequisites

- Node.js 20+
- pnpm 9 (`corepack enable` is recommended)
- Docker (optional, required for local subgraph infrastructure)

### Installation

```bash
pnpm install
```

This command links all workspace dependencies and creates a single lockfile at the repository root.

## Useful Commands

```bash
pnpm dev                # Start the Next.js website locally
pnpm dev:backend        # Start the NestJS backend in watch mode
pnpm --filter contracts test   # Run contract project tests
pnpm --filter subgraph build   # Build the subgraph
```

To run scripts defined inside a specific package, use `pnpm --filter <package> run <script>`.

## Repository Structure

```
.
|── backend/    # NestJS application and related docs
├── contracts/  # Smart contracts, typechain output, hardhat/foundry configs
├── subgraph/   # Graph protocol configuration and mappings
├── website/    # Next.js UI and tooling
├── docs/       # Additional documentation (guides, architecture, ADRs)
└── scripts/    # Cross-project automation and utilities
```

Existing package-level documentation (e.g. API setup, technical design notes) lives alongside each project. Add new cross-cutting docs under `docs/`.

## Contributing

See `CONTRIBUTING.md` for contribution guidelines, code style, and testing expectations. Please note that by contributing you agree to license your work under the BLS license.

## Security

If you discover a security issue, please follow the process described in `SECURITY.md`.

## License

This project is released under the BLS license. See `LICENSE` for details.
