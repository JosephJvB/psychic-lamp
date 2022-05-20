This seems to work
You have to open each go module as a separate vscode project
ie: go.mod must be at root of vscode project

/my-lambda-project
  - template.yaml
  - samconfig.toml
  - /function1
    - handler.go
      - func main() {}
      - `go mod init my-lambda-project/function1`
      - `go mod edit my-lambda-project/shared=../shared`
      - `import "my-lambda-project/shared/models"`
      - `import "my-lambda-project/shared/clients"`
  - /function2
    - handler.go
      - func main() {}
      - `go mod init my-lambda-project/function2`
      - `go mod edit my-lambda-project/shared=../shared`
      - `import "my-lambda-project/shared/models"`
      - `import "my-lambda-project/shared/clients"`
  - /shared
    - `go mod init my-lambda-project/shared`
    - /models
      - `package models`
      - requests.go
      - responses.go
    - /clients
      - `package clients`
      - ddb.go
      - s3.go