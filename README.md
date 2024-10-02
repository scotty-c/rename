# Rename CLI

`rename` is a Golang-based command-line application that provides two key functionalities:
1. **Search**: Searches for specific strings in files or directories based on a configuration file.
2. **Change**: Replaces specific strings in files or directories with alternative strings, as specified in the configuration file.

The configuration for search-and-replace operations is stored in a YAML file located at `~/.rename/conf.yaml`.

## Features

- Recursively searches through directories for specific strings.
- Replaces found strings with their replacements.
- Supports custom string replacement configurations via YAML.
- Two main commands: `search` and `change`.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/scotty-c/rename.git
   ```

2. Navigate to the project directory:

   ```bash
   cd rename
   ```

3. Build the CLI:

   ```bash
   make build
   ```

4. (Optional) Move the binary to your system's PATH:

   ```bash
   sudo mv rename /usr/local/bin/
   ```

## Configuration

The configuration file is located at `~/.rename/conf.yaml` and defines the strings to search for and their replacements. The file should follow this format:

```yaml
replacements:
  "dockerhub.io": "dockerhub.com"
  "quay.io": "quay.x"
```

Each entry under `replacements` is a key-value pair where the key is the string to search for and the value is the replacement string.

## Usage

The CLI supports two commands:

### 1. `rename search <file or directory>`

Searches for specific strings in the provided file or directory based on the `conf.yaml` configuration.

#### Example:

```bash
rename search /path/to/file.txt
rename search /path/to/directory
```

### 2. `rename change <file or directory>`

Replaces specific strings in the provided file or directory with their corresponding replacements as defined in the `conf.yaml` file.

#### Example:

```bash
rename change /path/to/file.txt
rename change /path/to/directory
```

## Running Tests

You can run tests for the project using the following command:

```bash
make test
```

The tests will dynamically create a temporary configuration file for testing purposes and will clean it up afterward.

## Development

### Project Structure

The project follows the best practices for Go CLI applications by separating the core logic into the `internal` package. The key files are:

- **cmd/root.go**: Defines the main `rename` commands (`search` and `change`).
- **internal/config.go**: Handles loading the configuration from the `~/.rename/conf.yaml` file.
- **internal/processor.go**: Handles file searching and string replacements.
- **internal/config_test.go**: Unit tests for configuration loading.
- **internal/processor_test.go**: Unit tests for file searching and processing.

### Contributing

1. Fork the repository.
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request.

## License

This project is licensed under the MIT License.