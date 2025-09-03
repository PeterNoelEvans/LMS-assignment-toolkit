#!/bin/bash

# Cross-Platform Assignment Toolkit Setup Script
echo "🌍 Assignment Toolkit Cross-Platform Setup"
echo "=========================================="

# Detect platform
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    PLATFORM="linux"
    EXECUTABLE="assignment-toolkit"
    echo "📋 Detected: Linux (Ubuntu/Debian compatible)"
elif [[ "$OSTYPE" == "darwin"* ]]; then
    PLATFORM="mac"
    EXECUTABLE="assignment-toolkit-intel"
    echo "📋 Detected: macOS"
elif [[ "$OSTYPE" == "msys" ]] || [[ "$OSTYPE" == "win32" ]]; then
    PLATFORM="windows"
    EXECUTABLE="assignment-toolkit.exe"
    echo "📋 Detected: Windows"
else
    echo "❌ Unsupported platform: $OSTYPE"
    echo "💡 Supported platforms: Linux, macOS, Windows"
    exit 1
fi

# Check if build exists
BUILD_PATH="builds/$PLATFORM/$EXECUTABLE"
if [[ ! -f "$BUILD_PATH" ]]; then
    echo "❌ Build not found: $BUILD_PATH"
    echo "🔧 Building for $PLATFORM..."
    
    case $PLATFORM in
        "linux")
            go build -o "$BUILD_PATH"
            ;;
        "windows")
            GOOS=windows GOARCH=amd64 go build -o "$BUILD_PATH"
            ;;
        "mac")
            GOOS=darwin GOARCH=amd64 go build -o "$BUILD_PATH"
            ;;
    esac
    
    if [[ $? -ne 0 ]]; then
        echo "❌ Build failed"
        exit 1
    fi
    echo "✅ Build completed: $BUILD_PATH"
fi

# Copy to current directory
cp "$BUILD_PATH" "./assignment-toolkit"
chmod +x "./assignment-toolkit" 2>/dev/null  # Linux/Mac only

echo "✅ Assignment Toolkit ready for $PLATFORM!"
echo ""
echo "🚀 Quick Test:"
echo "  ./assignment-toolkit --help"
echo ""
echo "🎯 Next Steps:"
echo "  1. Create a workspace: mkdir ~/my-templates && cd ~/my-templates"
echo "  2. Initialize: ~/path/to/assignment-toolkit init"
echo "  3. Create template: ~/path/to/assignment-toolkit create essay"
echo ""
echo "📚 Documentation:"
echo "  docs/GETTING_STARTED_CHECKLIST.md"
echo "  docs/CROSS_PLATFORM_COMPATIBILITY_GUIDE.md"
