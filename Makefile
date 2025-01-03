.PHONE: all
all: gen-go generate

.PHONY: gen-go
gen-go:
	cd protobuf && protoc \
		--go_out=./pbdeviceregistry --go_opt=paths=source_relative \
		--go-grpc_out=./pbdeviceregistry --go-grpc_opt=paths=source_relative \
		pbdeviceregistry.proto

	cd protobuf && protoc \
		--go_out=./pbdriver --go_opt=paths=source_relative \
		--go-grpc_out=./pbdriver --go-grpc_opt=paths=source_relative \
		pbdriver.proto

	cd protobuf && protoc \
		--go_out=./pbtaskmaster --go_opt=paths=source_relative \
		--go-grpc_out=./pbtaskmaster --go-grpc_opt=paths=source_relative \
		pbtaskmaster.proto

	cd protobuf && protoc \
		--go_out=./pbdataproxy --go_opt=paths=source_relative \
		--go-grpc_out=./pbdataproxy --go-grpc_opt=paths=source_relative \
		pbdataproxy.proto

	cd protobuf && protoc \
		--go_out=./pbdriveroperator --go_opt=paths=source_relative \
		--go-grpc_out=./pbdriveroperator --go-grpc_opt=paths=source_relative \
		pbdriveroperator.proto

.PHONY: generate
generate:
	oapi-codegen --config ./oapi-codegen.yaml -o ./openapi/rfc/rfc7807/rfc7807.go -package rfc7807 \
		./openapi/rfc/rfc7807/rfc7807.yaml

	oapi-codegen --config ./oapi-codegen.yaml -o ./openapi/openhes/driver/driverdata/driverdata.go -package driverdata \
		./openapi/openhes/driver/driverdata/driverdata.yaml

	oapi-codegen --config ./oapi-codegen.yaml -o ./openapi/openhes/job/job.go -package job \
		-import-mapping '"../driver/driverdata/driverdata.yaml":"github.com/cybroslabs/hes-2-apis/openapi/openhes/driver/driverdata","../driver/driver.yaml":"github.com/cybroslabs/hes-2-apis/openapi/openhes/driver","../device/device.yaml":"github.com/cybroslabs/hes-2-apis/openapi/openhes/device"' \
		./openapi/openhes/job/job.yaml

	oapi-codegen --config ./oapi-codegen.yaml -o ./openapi/openhes/driver/driver.go -package driver \
		-import-mapping '"../../openhes/driver/driverdata/driverdata.yaml":"github.com/cybroslabs/hes-2-apis/openapi/openhes/driver/driverdata"' \
		./openapi/openhes/driver/driver.yaml

	oapi-codegen --config ./oapi-codegen.yaml -o ./openapi/openhes/device/device.go -package device \
		-import-mapping '"../driver/driver.yaml":"github.com/cybroslabs/hes-2-apis/openapi/openhes/driver"' \
		./openapi/openhes/device/device.yaml

	find ./openapi -type f -name '*.go' | xargs sed -i '' -e "s|union\( *\)json\.RawMessage|Union\1json.RawMessage|g; s|t\.union|t.Union|g; s|action\.union|action.Union|g"
