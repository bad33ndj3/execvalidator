# execvalidator

ExecValidator is a library for validating if an executable is available on the system (read: exists in its PATH).

### Install

```shell
$ go install github.com/bad33ndj3/execvalidator@latest
```

### Usage

usage: execvalidator exists [executables to validate]

```shell
# validates that exec is in the PATH
$ execvalidator exists ls
>

# returns an error if executable is not in the PATH
$ execvalidator exists non-existent-executable
> Error: failed to validate executable: exec: "non-existent-executable": executable file not found in $PATH
```  
