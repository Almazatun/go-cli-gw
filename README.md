# CLI command

## Installation
```bash
# Clone the repository
git clone https://github.com/Almazatun/go-cli-gw.git
# Enter into the directory
cd go-cli-gw/
# Install the dependencies
go mod download
```

### Starting the application

```bash
# Build
$ go build
```

### CLI command

```bash
# All
$ ./go-cli-gw get -all
# Get by platform name
$ ./go-cli-gw get -plm test
# Add pass
$ ./go-cli-gw add -des "bla" -email "bla" -password "bla" -platform "bla" -username "bla"
```
