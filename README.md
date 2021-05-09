[![github](https://github.com/almeida-raphael/arpc_code_generator/workflows/Unit%20Tests/badge.svg)](https://github.com/almeida-raphael/arpc_code_generator)
[![codecov](https://codecov.io/gh/almeida-raphael/arpc_code_generator/branch/master/graph/badge.svg)](https://codecov.io/gh/almeida-raphael/arpc_code_generator)
# aRPC Code Generator
A code generator for aRPC protocol, this generator reads `*.arpc.go` files and creates all code needed to 
use run your code remotely over aRPC.

If you want more details on the protocol [click here](https://github.com/almeida-raphael/arpc).

### Installation
Run the code bellow to get the latest version of the code generator: 
```sh
go install github.com/almeida-raphael/arpc_code_generator@latest
```

### Usage
```sh
arpc_code_generator -input-path $PATH_TO_ARPC_DEFINITIONS_DIR -packages-root-path $PATH_TO_PACKAGE_CREATION_DIR
```

### Examples
[Here](https://github.com/almeida-raphael/arpc_examples) you can find example files to run `arpc_gen` command and test the aRPC protocol.

### Help
run `arpc_code_generator -help` to get generator args help

### WIP
This project is currently a work in progress and should be not used in production environments.

### Authors
* [Raphael C. Almeida](https://github.com/almeida-raphael)
* [Vitor Vasconcellos](https://github.com/HeavenVolkoff)
* [Ericson Soares](https://github.com/fogodev)
