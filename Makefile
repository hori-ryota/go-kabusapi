.PHONY: generate
generate:
	cd ./internal/codegen && go mod download && go run . -dstDir ../../kabusapi
