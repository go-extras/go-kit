# go-kit

[![Build Status](https://github.com/go-extras/go-kit/actions/workflows/go-test.yml/badge.svg)](https://github.com/go-extras/go-kit/actions/workflows/go-test.yml)

This is a Go module that provides additional functionality to the standard library.
It contains various packages and utilities that can be used to simplify Go programming.

## Features
- Package `must` offers a convenient approach for transforming a two-value function
  into a single-value function by throwing a panic if an error is returned as the second value
  in the original function.
- Package `contextualjson` provides a JSON marshaler that allows specifying a context
  and custom handlers for the serialization of struct fields.
- Package `logger` provides interfaces for logging with various levels of verbosity and functionality.
- Package `pubsub` provides provides a simple publish-subscribe messaging system.

## Installation
To use this module in your Go project, simply run the following command:

```bash
go get github.com/go-extras/go-kit
```

## Usage
To use the packages provided by this module, simply import them in your Go code.

For more details on how to use each package, please refer to the individual package documentation.

## Contribution
Contributions are welcome! Please feel free to submit any issues or pull requests.

## License
This module is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Author

[Denis Voytyuk](https://github.com/denisvmedia)
