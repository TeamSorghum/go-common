## About

[![Go Reference](https://pkg.go.dev/badge/github.com/teamsorghum/go-common.svg)](https://pkg.go.dev/github.com/teamsorghum/go-common)

This repository contains the implementations of some commonly used libraries.

## Development

Requirements:

1. Install go version specified in `go.mod` and `.tool-versions`. You can use [asdf](https://github.com/asdf-vm/asdf) to manage multiple versions.
2. Install [golangci-lint](https://golangci-lint.run/welcome/install/).
3. Install [mockgen](https://github.com/uber-go/mock). If you are using asdf to manage versions, execute `asdf reshim golang` after installing mockgen.
4. Install [wire](https://github.com/google/wire). If you are using asdf to manage versions, execute `asdf reshim golang` after installing wire.
5. Install container environment. You can choose either [docker](https://www.docker.com/) or [podman](https://podman.io/).
6. Execute the following commands to install git hooks:
  - Pre-commit hook will run checks and tests before committing: `ln -s ../../githooks/pre-commit .git/hooks/`
  - Commit-msg hook will check where commit message conforms to [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/): `ln -s ../../githooks/commit-msg .git/hooks/`

Test:

1. `cd deployments && cp .env.example .env` and edit `.env`, then `source .env`.
2. Launch containers `cd deployments && docker compose up -d`.
3. Now you can run tests via `go test` command.
