# ctrail-parser
A go based CLI app to parse through CloudTrail entries using JMESPath syntax


## Installing
Download either the windows, linux or mac binary for your system. Binaries are all compiled as 64bit. If you need a 32 bit, you need to build the binary from source and specify the environment variable GOARCH=386

## Building
Assuming you already have go installed and on your path, you need to do the following:
- setup a workspace.
    - Create the following directory structure (The root folder's name doesn't matter but the three contained folders MUST be called src, pkg and bin):
    ```
        |-- parser/
          |-- src/
          |-- pkg/
          |-- bin/
    ```
    - Set the `GOPATH` environment variable by cd'ing into the workspace and running:
    ```
        export GOPATH=`pwd`
    ```
    - Optionally amend your PATH variable to also include `$GOPATH/bin` so the compiled binary is on your PATH
    - Download parser's source and put the contents in `$GOPATH/src/github.com/austevick/parser/`
    - Download [Glide](https://github.com/Masterminds/glide) which parser uses as a package manager
        - You can skip using Glide if you handle the vendoring yourself. Parser's only dependency is [go-jmespath](github.com/jmespath/go-jmespath) which is used to filter the json records from CloudTrail
        - If you handle vendoring yourself (and are using go 1.6), setup a vendor/ folder in the root of parser's source code (where the main.go file is) and download the go-jmespath package into (you will have to make this directory layout) `$GOPATH/src/github.com/austevick/parser/vendor/jmespath/go-jmespath/`
    - Assuming the compiled glide binary is on your PATH, cd to `$GOPATH/src/github.com/austevick/parser` and run `glide init`
    - Compile the parser binary by running:
    ```
        GOOS=darwin GOARCH=amd64 go build github.com/austevick/parser/
    ```    
    Replace the value of `GOOS` with either darwin, linux or windows depending on your OS. If your OS does not support a 64bit architecture, replace `GOARCH` with 386

# Usage
```
parser -h
Usage of parser:
  -logs string
    	Directory for logs (default "/home/user/")
  -query string
    	JMESPath query expression. Defaults to all records (default "Records[*]")
```

parser.exe -query="Records[?eventSource=='iam.amazonaws.com']|[?contains(eventName,'List') == \`false\`]|[?contains(eventName,'Get') == \`false\`]|[?requestParameters.roleName=='aws-elasticbeanstalk-service-role']"
