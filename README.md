
# Rename CLI

`rename` is a simple Golang command-line tool that reads a file and replaces specific strings based on a list of replacements defined in a YAML configuration file.

## Features
- Replace multiple strings in a file based on predefined pairs.
- Configuration file stored in `~/.rename/conf.yaml`.

## Installation

1. Clone the repository and navigate into the directory:
   ```bash
   git clone https://github.com/scotty-c/rename.git
   cd rename
   ```

2. Install the required dependencies:
   ```bash
   go get -u github.com/spf13/cobra
   go get -u github.com/spf13/viper
   ```

3. Build the project:
   ```bash
   go build -o rename
   ```

4. Create the configuration file:
   ```bash
   mkdir -p ~/.rename
   touch ~/.rename/conf.yaml
   ```

## Usage

1. Add your replacements to the configuration file (`~/.rename/conf.yaml`):

   Example:
   ```yaml
   replacements:
     dockerhub.io: dockerhub.com
     quay.io: quay.x
     example.org: example.com
   ```

2. Run the `rename` command:

   ```bash
   ./rename <file-to-process>
   ```

   This will process the specified file, replacing the strings as per the configuration.

## Configuration

The configuration file is located at `~/.rename/conf.yaml`. It contains key-value pairs under the `replacements` section, where the key is the string to be replaced, and the value is the new string.

Example configuration:

```yaml
replacements:
  old-string: new-string
  example.com: example.org
```

## License

This project is licensed under the MIT License.

