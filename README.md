# Argus CLI
Collection of various tools bundled in a CLI.

# Environment setup
https://github.com/spf13/cobra

```sh
go install github.com/spf13/cobra-cli@latest
```

## Cobra generator
https://github.com/spf13/cobra-cli/blob/main/README.md

## Setup project in current directory

```sh
cobra-cli init . argus
```

## New command

```sh
cobra-cli add github
```

## Add a nested command

```sh
cobra-cli add dependabot -p 'github'
```
