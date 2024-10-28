#!/bin/bash

# Directory where the proto files will be saved locally
LOCAL_PROTO_DIR="external/proto"

# Remove the old proto files
rm -rf $LOCAL_PROTO_DIR

# Clone the proto repository 
git clone --depth 1 https://github.com/faruqii/faruqi-protos.git /tmp/faruqi-protos

# Copy the proto files to the local project directory
mkdir -p $LOCAL_PROTO_DIR
cp -r /tmp/faruqi-protos/proto/products/*.proto $LOCAL_PROTO_DIR
cp -r /tmp/faruqi-protos/proto/users/*.proto $LOCAL_PROTO_DIR

# Clean up the temporary clone
rm -rf /tmp/faruqi-protos

echo "Proto files updated in $LOCAL_PROTO_DIR"
