# 04-env-var-go

## Create a new project with source code

```bash
mkdir 04-env-var-go
cd 04-env-var-go
go mod init env-var-go
touch main.go
```

Use the source code of the third exercise and use an environment variable to configure the application (to set the name of the text file).

```bash
wasmedge --dir . --env FILE_NAME="hello.txt" main.wasm
```


> Ref: https://gobyexample.com/environment-variables
