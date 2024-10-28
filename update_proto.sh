#!/bin/bash

# Directory where the proto files will be saved locally
LOCAL_PROTO_DIR="external/proto"

# Temporary directory for the proto repository
TMP_PROTO_REPO="/tmp/faruqi-protos"

# Check if the proto repo already exists locally
if [ -d "$TMP_PROTO_REPO" ]; then
    # If it exists, update the repository
    echo "Updating existing proto repository..."
    git -C $TMP_PROTO_REPO pull
else
    # Clone the repository if it does not exist
    echo "Cloning proto repository..."
    git clone --depth 1 https://github.com/faruqii/faruqi-protos.git $TMP_PROTO_REPO
fi

# Remove old proto files from local project directory
rm -rf $LOCAL_PROTO_DIR

# Create local proto directory if it doesn't exist
mkdir -p $LOCAL_PROTO_DIR

# Copy updated proto files to the local project directory
cp -r $TMP_PROTO_REPO/proto/products/*.proto $LOCAL_PROTO_DIR
cp -r $TMP_PROTO_REPO/proto/users/*.proto $LOCAL_PROTO_DIR

echo "Proto files updated in $LOCAL_PROTO_DIR"
