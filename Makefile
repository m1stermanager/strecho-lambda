OUTPUT=dist/hello-world
PACKAGE=dist/package.yaml

$(OUTPUT):
	GOOS=linux go build -o $(OUTPUT) src/cmd/lambda-test.go

$(PACKAGE): $(OUTPUT)
	sam package --template-file bin/sam/template.yaml --s3-bucket strecho --output-template-file $(PACKAGE)

deploy: $(PACKAGE)
	sam deploy --template-file $(PACKAGE) --stack-name strecho --capabilities CAPABILITY_IAM

clean:
	rm -rf dist
