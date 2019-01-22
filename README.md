# lambda-go-sample
Serverless Application System. (AWS Lambda (Go) + API Gateway + Dynamo DB)

## Prerequisites
[AWS CLI](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/setup-awscli.html)

[SAM CLI](https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/sam-cli-requirements.html)

## Setup
```
cp .envrc-example .envrc
```

## Start Server
```
$ cd sam-app/
$ sam local start-api
```

## Build & Deploy
```
$ cd sam-app/

$ make build
$ make package
$ make deploy
```
