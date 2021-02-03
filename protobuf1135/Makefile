PROTO_OUTPUT := .
PROTO_INPUT := .
PROTO_MAPPINGS := Mmodel.proto=github.com/cristaloleg/public-bugs/protobuf1135/model

install-go-tools:
	cat tools.go | grep _ | grep \".*\" -o | xargs -tI % go install %

proto:
	protoc \
		-I ${PROTO_INPUT} \
		--gogoslick_out=${PROTO_MAPPINGS}:${PROTO_OUTPUT}/model ${PROTO_INPUT}/model.proto

	protoc \
		-I ${PROTO_INPUT} \
		--gogoslick_out=${PROTO_MAPPINGS},plugins=grpc:${PROTO_OUTPUT}/svc ${PROTO_INPUT}/svc.proto

run:
	go run .
