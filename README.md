
# HCURL

hcurl is a command-line tool primarily designed for testing HTTP and HTTPS endpoints. It was created using the Cobra CLI package for Go, offering a user-friendly interface.

## Installation

1. Visit the [GitHub Releases](https://github.com/vishnukumarkvs/hcurl/releases) page to download the latest version of hcurl for your specific operating system.

2. Extract the downloaded archive.

3. Move the extracted binary to a directory included in your system's PATH environment variable.

## Usage

### Getting Help

To learn about subcommands and flags of hcurl, use the `-h or --help` option:
```
hcurl -h

hcurl <subcommand> -h
```
### Subcommands

`get`: Send an HTTP GET request.

`post`: Send an HTTP POST request.

### Global Flags

`-h, --help`: Show help for the command or subcommand.

`-n, --count`: Number of times to repeat the request.

`-H, --header`: Set a custom header for the request (can be used multiple times).

### Subcommand-specific Flags

`post: -b, --body`: Provide the request body (supports inline JSON or path to a JSON file prefixed with @).

## Examples

**Send a GET request**

`hcurl get https://dummyjson.com/products/1`

**Send a GET request 3 times**

`hcurl get https://dummyjson.com/products/1 -n 3`

**Send a GET request with a custom header**

`hcurl get https://dummyjson.com/products/1 -H 'x-api-key:xxxxx'`

**Send a POST request with an inline JSON body**

`hcurl post https://dummyjson.com/products/add -b '{"Title": "BMW"}'`

**Send a POST request with a JSON file as the body**

`hcurl post https://dummyjson.com/products/add -b @path/to/json/file`

## Contributing

Contributions are welcome! Please follow the standard contributing guidelines when submitting issues or pull requests.

## License
This project is licensed under the Apache License.
