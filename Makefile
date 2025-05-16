GO_MODULE := github.com/KennedySurianto/pre_tpa_web/

.PHONY:protoc-go
protoc-go:
	protoc --go_opt=module=${GO_MODULE} --go_out=. \
		--go-grpc_opt=module=${GO_MODULE} --go-grpc_out=. \
		./proto/controller/*.proto ./proto/dto/*.proto
