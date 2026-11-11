default:
    @just --list

# release dry run: all platforms + the cask, into dist/, nothing published.
# real releases: push a v* tag (see .goreleaser.yaml)
snapshot:
    HOMEBREW_TAP_TOKEN=unused goreleaser release --snapshot --clean

test:
    GOWORK=off go test ./...

# build from source into ~/.local/bin
install:
    mkdir -p ~/.local/bin
    GOWORK=off go build -o ~/.local/bin/proxima .
