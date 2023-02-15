build-HelloWorldFunction:
	GOARCH=arm64 GOOS=linux go build -tags lambda.norpc -o ./bootstrap
	cp bootstrap $(ARTIFACTS_DIR)/bootstrap