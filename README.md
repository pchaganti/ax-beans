# Beans

An agentic-first issue tracker. Store and manage issues as markdown files in your project's `.beans/` directory.

## Installation

### Homebrew

```bash
brew install hmans/beans/beans
```

## Usage

```bash
beans init          # Initialize a .beans/ directory
beans list          # List all beans
beans show <id>     # Show a bean's contents
beans create "Title" # Create a new bean
beans status <id> <status>  # Change status (open, in-progress, done)
beans archive       # Delete all done beans
```

All commands support `--json` for machine-readable output.

## Contributing

This project currently does not accept contributions -- it's just way too early for that!
But if you do have suggestions or feedback, please feel free to open an issue.
