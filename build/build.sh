#!/bin/zsh

# Name of your program
PROGRAM_NAME="proxima"
# List of OS/architecture combinations to build
PLATFORMS=("windows/amd64" "linux/amd64" "linux/arm64" "darwin/amd64" "darwin/arm64")

DIST_DIR="../dist"

# Cleanup previous builds
echo "Cleaning up previous builds..."
rm -f "$DIST_DIR"/"$PROGRAM_NAME"-*

# Build loop
for PLATFORM in "${PLATFORMS[@]}"; do
    # Split the platform string into OS and ARCH
    GOOS=${PLATFORM%/*}
    GOARCH=${PLATFORM#*/}
    
    # Set the output file name. The extension for Windows is .exe.
    OUTPUT_NAME="$PROGRAM_NAME-$GOOS-$GOARCH"
    if [ "$GOOS" = "windows" ]; then
        OUTPUT_NAME="$OUTPUT_NAME.exe"
    fi

    # Build the binary for the platform
    echo "Building for $GOOS $GOARCH..."
    GOOS=$GOOS GOARCH=$GOARCH go build -o "$DIST_DIR"/"$OUTPUT_NAME" $PROGRAM_NAME.go

    # Check if build was successful
    if [ $? -ne 0 ]; then
        echo "An error has occurred! Aborting the script execution..."
        exit 1
    fi
done

echo "Builds completed successfully."

