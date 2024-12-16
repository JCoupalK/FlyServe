<p align="center">
 <img src="https://github.com/JCoupalK/FlyServe/assets/108779415/e2cc0123-bfcc-4191-8a83-6ccf37d8e687"
</p>

# FlyServe
![License](https://img.shields.io/github/license/JCoupalK/FlyServe)
![GitHub issues](https://img.shields.io/github/issues-raw/JCoupalK/FlyServe)
![GitHub go.mod Go version (branch & subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/JCoupalK/FlyServe/main)

FlyServe is a simple web server written in Go, designed to serve files and directories over HTTP. It provides features such as serving specific files, serving directories, basic authentication, and directory listings.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Building from Source](#building-from-source)
- [Usage](#usage)
- [Examples](#examples)
- [Screenshots](#screenshots)
- [Contributing](#contributing)
- [License](#license)

## Features

- Serve specific files
- Serve directories with directory listings
- Basic authentication support
- Customizable port number
- Single binary, no external dependencies

## Installation

1. Download the binary with wget:

    ```shell
    wget https://github.com/JCoupalK/FlyServe/releases/download/1.0/flyserve_linux_amd64_1.0.tar.gz
    ```

2. Unpack it with tar

    ```shell
    tar -xf flyserve_linux_amd64_1.0.tar.gz
    ```

3. Move it to your /usr/local/bin/ (Optional):

    ```shell
    sudo mv flyserve /usr/local/bin/flyserve
    ```

## Building from Source

1. Ensure you have Go installed on your system. You can download Go from [here](https://golang.org/dl/).
2. Clone the repository:

    ```shell
    git clone https://github.com/JCoupalK/FlyServe
    ```

3. Navigate to the cloned directory:

    ```shell
    cd FlyServe
    ```

4. Build the tool:

    ```shell
    go build -o flyserve .
    ```

## Usage

Options:

```txt
-p, --port            Port to serve on (Default is 8080)
-f, --file            Specific file to serve
-d, --directory       Directory to serve files from (Default is current directory)
-u, --username        Username for basic authentication
-pw, --password       Password for basic authentication
-h, --html            Enable auto-resolution of .html files
```

## Examples

```bash
# Serve files from the current directory on port 8080
flyserve

# Serve files from a specific directory on port 9000
flyserve -d /path/to/directory -p 9000

# Serve a specific file on port 8080 with basic authentication
flyserve -f /path/to/file.txt -u username -pw password

# Serve websites by just specifying a directory that contains an index.html
flyserve -d /path/to/website -p 80 -h
```

## Screenshots

<img alt="screenshot1" width="400" src="https://github.com/JCoupalK/FlyServe/assets/108779415/74c10791-62d5-4e68-988a-62dfba2965ff"/>
<br>
<img alt="screenshot2" width="400" src="https://github.com/JCoupalK/FlyServe/assets/108779415/3e93a041-c54e-4c64-96fb-d59ff63e6640"/>
<br>
<img alt="screenshot3" width="400" src="https://github.com/JCoupalK/FlyServe/assets/108779415/da666e86-ae02-4c25-8fe7-cb5f9f8f04c0"/>

## Contributing

Contributions are welcome. If you find a bug or have a feature request, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
