# Ctrail-Parser
A go based CLI app to parse through CloudTrail entries using [JMESPath](http://jmespath.org/) syntax

```
go get -v github.com/austevick/ctrail-parser/
```

## Installing
Download either the windows, linux or mac binary for your system. Extract the zip file and put the binary somewhere on your system.

You can see precompiled binaries for Linux, Mac and Windows from the [releases page](https://github.com/austevick/ctrail-parser/releases/)
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
    - Optionally add the bin directory to your PATH by running:
    ```
        export PATH=$PATH:$GOPATH/bin
    ```
    - The project uses [Glide](https://github.com/Masterminds/glide) for package management.
        - If you want to use glide as well, do the following:
            - Download and install glide. You can use one of the latest [releases](https://github.com/Masterminds/glide/releases) for your OS. Just download and extract the archive and put the binary somewhere on your path.
            - Download ctrail-parser by running `go get github.com/austevick/ctrail-parser/`
            - Download dependencies by running `glide update`
    - If you don't want to use glide, you can download ctrail-parser as well as it's dependencies by running:
    ```
        go get -v github.com/austevick/ctrail-parser/
    ```
    - Compile ctrail-parser by running:
    ```
        GOOS=darwin GOARCH=amd64 go build -o $GOPATH/bin github.com/austevick/ctrail-parser/
    ```    
    Use either darwin, linux, or windows for the value of `GOOS` depending on your OS. Also, you can set `GOARCH` to 386 if your OS does not support a 64bit architecture.
    - Test if ctrail-parser compiled correctly by running `ctrail-parser -h`


## Usage
```
ctrail-parser -h
Usage of ./ctrail-parser:
  -logs string
    	Directory for logs (default "/home/user/")
  -query string
    	JMESPath query expression. Defaults to all records (default "Records[*]")
```

By default, ctrail-parser looks for logs in the current directory. You can use the -logs flag to change where ctrail-parser looks for logs.

It will first attempt to look for any files with a .json.gz extension. If it finds one or more files, it will extract the contents of the file (into memory) append all the json into a single slice and apply the JMESPath query onto the slice. If it does not find any .json.gz files, it will look through any directories and try to find any .json.gz files.

For the actual query, the syntax looks something like this:
```
-query="Records[?eventSource=='iam.amazonaws.com']|[?contains(eventName,'List') == \`false\`]|[?contains(eventName,'Get') == \`false\`]|[?requestParameters.roleName=='aws-elasticbeanstalk-service-role']"
```
Due to the way the JSON for CloudTrail entries is structured, the query HAS TO begin with Records[].
