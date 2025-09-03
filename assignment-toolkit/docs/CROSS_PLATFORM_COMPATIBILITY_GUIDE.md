# 🌍 Cross-Platform Compatibility Guide

## 🎯 Overview

The Assignment Toolkit is designed to work seamlessly across Ubuntu, Debian, and Windows 11. However, there are some platform-specific considerations when working across different computers while traveling.

## 🖥️ Platform Support Matrix

| Feature | Ubuntu | Debian | Windows 11 | Notes |
|---------|--------|--------|------------|-------|
| Go Toolkit | ✅ | ✅ | ✅ | Needs separate builds |
| YAML Files | ✅ | ✅ | ✅ | Universal format |
| File Paths | ⚠️ | ⚠️ | ⚠️ | Path separators differ |
| Text Editors | ✅ | ✅ | ✅ | All platforms supported |
| Git Integration | ✅ | ✅ | ✅ | Works everywhere |

## 🔧 Platform-Specific Setup

### **Ubuntu/Debian (Linux)**
```bash
# Build toolkit (if needed)
cd assignment-toolkit
go build -o assignment-toolkit

# Make executable
chmod +x assignment-toolkit

# Optional: Install globally
sudo cp assignment-toolkit /usr/local/bin/

# Usage
./assignment-toolkit init
assignment-toolkit init  # If installed globally
```

### **Windows 11**
```cmd
# Build toolkit (if needed)
cd assignment-toolkit
go build -o assignment-toolkit.exe

# Usage (Command Prompt)
assignment-toolkit.exe init

# Usage (PowerShell)
.\assignment-toolkit.exe init

# Usage (Git Bash - recommended)
./assignment-toolkit.exe init
```

## ⚠️ **Key Differences to Watch**

### **1. File Paths** 🚨 IMPORTANT

**Problem**: Different path separators
- **Linux**: `/home/user/assignments/resources/image.png`
- **Windows**: `C:\Users\User\assignments\resources\image.png`

**Solution**: Use relative paths in YAML files
```yaml
# ✅ GOOD - Works on all platforms
resources:
  - local_path: "./resources/world-map.png"
  - local_path: "./documents/study-guide.pdf"

# ❌ BAD - Platform-specific
resources:
  - local_path: "/home/user/assignments/resources/world-map.png"  # Linux only
  - local_path: "C:\Users\User\assignments\resources\world-map.png"  # Windows only
```

### **2. Executable Names**

**Linux (Ubuntu/Debian)**:
```bash
./assignment-toolkit init
```

**Windows**:
```cmd
assignment-toolkit.exe init
```

**Solution**: Create platform-specific scripts

### **3. Text Editors**

| Platform | Recommended Editors |
|----------|-------------------|
| Ubuntu | `code`, `nano`, `gedit`, `vim` |
| Debian | `nano`, `vim`, `mousepad`, `code` |
| Windows 11 | `code`, `notepad++`, `notepad` |

## 🚀 **Cross-Platform Solutions**

### **1. Build for All Platforms**

Create builds for each platform you use:

```bash
# On Linux (build for all platforms)
cd assignment-toolkit

# Linux build
go build -o assignment-toolkit-linux

# Windows build  
GOOS=windows GOARCH=amd64 go build -o assignment-toolkit-windows.exe

# Create platform-specific directories
mkdir -p builds/linux builds/windows
cp assignment-toolkit-linux builds/linux/assignment-toolkit
cp assignment-toolkit-windows.exe builds/windows/assignment-toolkit.exe
```

### **2. Portable Workspace Setup**

Create a workspace that works everywhere:

```bash
# Create portable workspace structure
mkdir portable-assignments
cd portable-assignments

# Create cross-platform config
cat > .assignment-config.yaml << 'EOF'
author: "Your Name"
email: "your.email@school.edu"
license: "CC-BY-SA-4.0"
language: "en"

# No absolute paths - keep portable
lms_endpoint: ""
api_key: ""

defaults:
  points: "1"
  auto_grade: "true"
  published: "false"
  quarter: "Q1"

# Use relative paths only
templates:
  multiple-choice: "./templates/multiple-choice.yaml"
  essay: "./templates/essay.yaml"
EOF

# Create directory structure
mkdir -p templates resources packages examples
```

### **3. Cross-Platform Scripts**

**Linux/Mac Script** (`run-toolkit.sh`):
```bash
#!/bin/bash
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
"$SCRIPT_DIR/builds/linux/assignment-toolkit" "$@"
```

**Windows Script** (`run-toolkit.bat`):
```cmd
@echo off
set SCRIPT_DIR=%~dp0
"%SCRIPT_DIR%builds\windows\assignment-toolkit.exe" %*
```

**Cross-Platform Script** (`run-toolkit.py`):
```python
#!/usr/bin/env python3
import os
import sys
import platform
import subprocess

script_dir = os.path.dirname(os.path.abspath(__file__))

if platform.system() == "Windows":
    executable = os.path.join(script_dir, "builds", "windows", "assignment-toolkit.exe")
else:
    executable = os.path.join(script_dir, "builds", "linux", "assignment-toolkit")

subprocess.run([executable] + sys.argv[1:])
```

### **4. Sync with Cloud Storage**

Keep your workspace synchronized across devices:

```bash
# Option 1: Git repository
git init
git add .
git commit -m "Initial assignment templates"
git remote add origin https://github.com/yourusername/assignment-templates.git
git push -u origin main

# Option 2: Cloud sync folder
# Put workspace in Dropbox/OneDrive/Google Drive folder
mkdir ~/Dropbox/assignment-templates
cd ~/Dropbox/assignment-templates
```

## 📁 **Recommended Directory Structure**

```
assignment-templates/  (synced folder)
├── .assignment-config.yaml
├── builds/
│   ├── linux/
│   │   └── assignment-toolkit
│   └── windows/
│       └── assignment-toolkit.exe
├── scripts/
│   ├── run-toolkit.sh      # Linux/Mac
│   ├── run-toolkit.bat     # Windows
│   └── run-toolkit.py      # Cross-platform Python
├── templates/
├── resources/
├── examples/
└── README.md
```

## 🔄 **Platform Migration Workflow**

### **Moving from Linux to Windows**
```bash
# On Linux - prepare for Windows
git add .
git commit -m "Linux work session complete"
git push

# On Windows - continue work
git pull
# Use assignment-toolkit.exe instead of assignment-toolkit
```

### **Moving from Windows to Linux**
```cmd
# On Windows - prepare for Linux
git add .
git commit -m "Windows work session complete"  
git push

# On Linux - continue work
git pull
# Use ./assignment-toolkit instead of assignment-toolkit.exe
```

## ⚠️ **Platform-Specific Issues & Solutions**

### **Windows 11 Specific**

**Issue**: PowerShell execution policy
```powershell
# Fix execution policy if needed
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
```

**Issue**: Path length limits
```cmd
# Use shorter paths on Windows
cd C:\assignments  # Instead of C:\Users\VeryLongUsername\Documents\...
```

**Issue**: Antivirus blocking executable
- Add assignment-toolkit.exe to antivirus exceptions
- Use Windows Defender exclusions

### **Ubuntu/Debian Specific**

**Issue**: Go not installed (if rebuilding needed)
```bash
# Install Go on Ubuntu/Debian
sudo apt update
sudo apt install golang-go

# Verify installation
go version
```

**Issue**: Permission denied
```bash
# Fix permissions
chmod +x assignment-toolkit
```

### **Universal Solutions**

**File Encoding**: Always use UTF-8
```yaml
# Add to YAML files if needed
# -*- coding: utf-8 -*-
```

**Line Endings**: Git handles this automatically
```bash
# Configure Git for cross-platform line endings
git config --global core.autocrlf true   # Windows
git config --global core.autocrlf input  # Linux/Mac
```

## 🎯 **Best Practices for Multi-Platform Work**

### **1. Use Relative Paths Always**
```yaml
# ✅ GOOD - Works everywhere
resources:
  - local_path: "./resources/image.png"
  - local_path: "../shared/document.pdf"

# ❌ BAD - Platform-specific
resources:
  - local_path: "/home/user/assignments/image.png"
  - local_path: "C:\Users\User\assignments\image.png"
```

### **2. Consistent Naming**
```bash
# Use lowercase, hyphen-separated names
my-essay-template.yaml
geography-quiz.yaml
python-coding-exercise.yaml

# Avoid spaces and special characters
# my essay template.yaml  ❌
# my_essay_template.yaml  ✅ (acceptable)
```

### **3. Cloud Sync Strategy**
```bash
# Keep workspace in synced folder
~/Dropbox/assignment-templates/
~/OneDrive/assignment-templates/
~/Google Drive/assignment-templates/
```

### **4. Version Control**
```bash
# Use Git for proper cross-platform sync
git init
git add .
git commit -m "Template updates"
git push origin main
```

## 🧳 **Travel Setup Checklist**

### **Before Traveling**
- [ ] Commit all work to Git
- [ ] Ensure workspace is in synced folder
- [ ] Test toolkit on destination platform
- [ ] Download offline documentation
- [ ] Backup important templates

### **While Traveling**
- [ ] Use relative paths only
- [ ] Work in synced directory
- [ ] Validate work regularly
- [ ] Commit changes frequently

### **After Traveling**
- [ ] Sync with main repository
- [ ] Test templates on original platform
- [ ] Merge any conflicts
- [ ] Update documentation if needed

## 🔧 **Platform-Specific Installation**

### **Quick Setup Script**

Create this script to set up on any new machine:

```bash
#!/bin/bash
# setup-assignment-toolkit.sh

echo "🚀 Setting up Assignment Toolkit..."

# Detect platform
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    PLATFORM="linux"
    EXECUTABLE="assignment-toolkit"
elif [[ "$OSTYPE" == "msys" ]] || [[ "$OSTYPE" == "win32" ]]; then
    PLATFORM="windows"
    EXECUTABLE="assignment-toolkit.exe"
else
    echo "❌ Unsupported platform: $OSTYPE"
    exit 1
fi

echo "📋 Detected platform: $PLATFORM"

# Copy appropriate executable
cp "builds/$PLATFORM/$EXECUTABLE" ./assignment-toolkit

# Make executable (Linux only)
if [[ "$PLATFORM" == "linux" ]]; then
    chmod +x assignment-toolkit
fi

echo "✅ Assignment Toolkit ready!"
echo "💡 Run: ./assignment-toolkit --help"
```

## 🎉 **Summary: You're Covered!**

✅ **Works on all your platforms** (Ubuntu, Debian, Windows 11)  
✅ **YAML files are universal** (work everywhere)  
✅ **Git sync keeps everything in sync**  
✅ **Relative paths prevent platform issues**  
✅ **Documentation covers all scenarios**  
✅ **No platform-specific dependencies**  

**The system is designed for exactly your multi-platform, travel-based workflow!** 🌍✈️💻
