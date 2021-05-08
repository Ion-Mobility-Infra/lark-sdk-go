# LarkSuite golang SDK

Open source LarkSuite golang SDK

## Generate go struct

```bash
# source directory of json files: sample
# destination directory of generated golang structs: model
go run cmd/generate.go
```

This command can be run again whenever we update our `sample` directory with new json files.

## Running Examples

```bash
# load environment variables used to call LarkSuite APIs
export `cat .env`

# Run our example
go run example/main.go
```
