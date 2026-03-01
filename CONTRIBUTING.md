# Contributing to Sarvam AI Go SDK

Thank you for your interest in contributing to this project. We welcome contributions from the community to help improve this SDK.

## How to Contribute

### Reporting Bugs

If you find a bug, please open an issue on GitHub. Include:
- A clear description of the issue.
- Steps to reproduce the problem.
- Expected vs. actual behavior.
- Relevant code snippets or error logs.

### Feature Requests

For new features or enhancements, please open an issue to discuss the proposal before starting implementation.

### Pull Requests

1. Fork the repository.
2. Create a new branch for your feature or fix.
3. Ensure your code follows existing patterns and naming conventions.
4. Add or update tests for your changes.
5. Ensure all tests pass by running `go test ./...`.
6. Submit a pull request with a detailed description of your changes.

## Development Setup

- Go 1.25.5 or later is required.
- The project uses standard Go modules for dependency management.
- Run `go build ./...` to verify compilation.
- Run `go test ./...` to execute the test suite.

## Coding Standards

- Use idiomatic Go practices.
- Ensure all public functions, structs, and constants are documented with doc comments.
- Maintain consistent error handling and wrapping patterns.
- Follow the established functional options pattern for API configuration.

## License

By contributing to this project, you agree that your contributions will be licensed under the project's MIT License.
