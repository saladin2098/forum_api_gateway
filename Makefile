CURRENT_DIR=$(shell pwd)

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}
swag-gen:
	~/go/bin/swag init -g ./api/router.go -o docs force 1	