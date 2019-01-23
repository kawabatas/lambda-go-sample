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

### example
```
$ curl http://127.0.0.1:3000/users

$ curl -X POST http://127.0.0.1:3000/users -d '{"name":"hoge","age":20}'

$ curl -X DELETE http://127.0.0.1:3000/users/hoge/20
```

echo framework
```
$ curl http://127.0.0.1:3000/echo_users

$ curl -X POST http://127.0.0.1:3000/echo_users -d 'name=hoge' -d 'age=20'

$ curl -X DELETE http://127.0.0.1:3000/echo_users/hoge/20
```

## Build & Deploy
```
$ cd sam-app/

$ make build
$ make package
$ make deploy
```
