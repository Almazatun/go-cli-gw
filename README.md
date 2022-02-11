# 💧 CLI command

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
# Get pass by platform name
$ ./go-cli-gw get -plm=<FULL_PLATFORM_NAME>
# Get pass list by special characters
$ ./go-cli-gw get -pcr=<SPECIAL_CHARACTERS>
# Add pass
$ ./go-cli-gw add -email=bla -password=bla -username=bla -platform=bla -des=bla
```
