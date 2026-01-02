# Contributing to Fluxus

Thank you for investing time in improving Fluxus! The project thrives on community feedback, code, and documentation contributions.

## Development Workflow

1. Fork the repository and create a feature branch.
2. Run `pnpm install` to bootstrap all workspaces.
3. Make focused, well-tested changes with clear commit messages.
4. Run the relevant checks before opening a pull request:
   - `pnpm -r --if-present run lint`
   - `pnpm -r --if-present run test`
   - Additional package-specific scripts as documented in each workspace.
5. Open a pull request that summarizes the changes and references related issues.

## Coding Standards

- Prefer TypeScript and follow existing project conventions.
- Keep documentation up to date when behavior changes.
- Include tests whenever possible; new features without tests may be delayed.

## Commit Messages & PRs

- Use descriptive, imperative commit messages.
- Keep pull requests focused; large unrelated changes should be split into multiple PRs.
- Provide screenshots or logs for UI or integration changes when appropriate.

## License

By contributing to Fluxus you agree that your contributions are licensed under the BLS license found in `LICENSE`.
