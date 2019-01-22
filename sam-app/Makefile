.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./hello-world/hello-world

build:
	GOOS=linux GOARCH=amd64 go build -o hello-world/hello-world ./hello-world
	GOOS=linux GOARCH=amd64 go build -o user-fetcher/user-fetcher ./user-fetcher

package:
	aws cloudformation package --template-file template.yaml --output-template-file output-template.yaml --s3-bucket lambda-go-sample-111

deploy:
	aws cloudformation deploy --template-file output-template.yaml --stack-name lambda-go-sample-111 --capabilities CAPABILITY_IAM
