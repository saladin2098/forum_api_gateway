#!/bin/sh

CURRENT_DIR=$1
rm -rf ${CURRENT_DIR}/genproto
mkdir -p ${CURRENT_DIR}/genproto

# Find all directories that contain .proto files, excluding the .git directory
for x in $(find ${CURRENT_DIR}/forum_protos -type d \( ! -path "*/.git*" \)); do
  if ls ${x}/*.proto 1> /dev/null 2>&1; then
    protoc -I=${x} -I/usr/local/include \
      --go_out=${CURRENT_DIR}/genproto --go_opt=paths=source_relative \
      --go-grpc_out=${CURRENT_DIR}/genproto --go-grpc_opt=paths=source_relative ${x}/*.proto
  fi
done
