# 📁 Setting Up Separate Assignment Templates Repository

## 🎯 Why Separate Repository?

Perfect for your travel-based workflow:
- ✅ Smaller, faster to sync
- ✅ No production code on travel devices
- ✅ Easy to share with colleagues
- ✅ Platform-specific builds included
- ✅ Independent version control

## 🚀 Setup Instructions

### **Step 1: Create Repository Structure**
```bash
# Create your templates repository
mkdir ~/assignment-templates
cd ~/assignment-templates

# Initialize Git
git init
git branch -M main
```

### **Step 2: Copy Toolkit and Builds**
```bash
# Copy the toolkit with all platform builds
cp -r /mnt/LMS-database/repos/English-Foundation-Adventure/assignment-toolkit/* .

# Verify you have everything
ls -la
ls builds/
```

### **Step 3: Create Repository Files**
```bash
# Create README for your repository
cat > README.md << 'EOF'
# My Assignment Templates

Personal assignment template repository for LMS content development.

## Quick Start

1. **Setup on new computer:**
   ```bash
   ./setup-platform.sh  # Linux/Mac
   setup-platform.bat   # Windows
   ```

2. **Create templates:**
   ```bash
   ./assignment-toolkit init
   ./assignment-toolkit create essay
   ```

3. **Validate quality:**
   ```bash
   ./assignment-toolkit validate my-template.yaml
   ```

## Platform Support

- ✅ Ubuntu/Debian Linux
- ✅ Windows 11
- ✅ macOS (Intel)

## Directory Structure

- `builds/` - Platform-specific executables
- `examples/` - Sample templates
- `docs/` - Documentation (copied from main LMS)
- Your `.yaml` files - Your assignment templates

## Usage

This repository is for **template development only**:
- Design assignment structures offline
- Validate quality and format
- Use as blueprints for LMS content creation
- Share templates with colleagues

EOF

# Create .gitignore
cat > .gitignore << 'EOF'
# Workspace files (each user creates their own)
.assignment-config.yaml

# Temporary files
*.tmp
*.log
.DS_Store
Thumbs.db

# IDE files
.vscode/
.idea/
*.swp
*.swo

# OS files
desktop.ini
EOF
```

### **Step 4: Initial Commit**
```bash
# Add files to Git
git add .
git commit -m "Initial assignment templates repository

- Cross-platform toolkit included
- Platform builds for Ubuntu, Debian, Windows 11
- Example templates and documentation
- Setup scripts for easy deployment"

# Create GitHub repository (optional)
# Go to GitHub, create new repository called "assignment-templates"
# Then:
git remote add origin https://github.com/yourusername/assignment-templates.git
git push -u origin main
```

### **Step 5: Test on Current Platform**
```bash
# Test the setup
./setup-platform.sh

# Create a workspace to test
mkdir test-workspace
cd test-workspace
../assignment-toolkit init

# Create a sample template
../assignment-toolkit create essay
../assignment-toolkit validate *.yaml
```

## 🧳 **Using on Travel Computers**

### **Setup on New Computer:**
```bash
# Clone your repository
git clone https://github.com/yourusername/assignment-templates.git
cd assignment-templates

# Run platform setup
./setup-platform.sh      # Linux
setup-platform.bat       # Windows

# Create your workspace
mkdir my-work
cd my-work
../assignment-toolkit init

# Start creating templates
../assignment-toolkit create multiple-choice
```

### **Syncing Work:**
```bash
# Before leaving computer
git add .
git commit -m "Created geography templates on Ubuntu laptop"
git push

# On next computer
git pull
# Continue working...
```

## 📁 **Repository Structure**
```
assignment-templates/           # Your new repository
├── README.md                  # Repository documentation
├── .gitignore                 # Git ignore rules
├── setup-platform.sh          # Linux/Mac setup
├── setup-platform.bat         # Windows setup
├── builds/                    # Platform-specific executables
│   ├── linux/assignment-toolkit
│   ├── windows/assignment-toolkit.exe
│   └── mac/assignment-toolkit-intel
├── examples/                  # Sample templates
│   ├── multiple-choice-basic.yaml
│   ├── matching-geography.yaml
│   └── code-submission-python.yaml
├── docs/                      # Documentation (copied)
│   ├── GETTING_STARTED_CHECKLIST.md
│   ├── CROSS_PLATFORM_COMPATIBILITY_GUIDE.md
│   └── portable_assignments_quick_reference.md
├── main.go                    # Source code (for rebuilding if needed)
├── types.go
├── commands.go
├── sync.go
└── go.mod                     # Go dependencies

# Your workspaces (not in Git)
work-laptop/                   # Created by you with 'init'
travel-desktop/                # Created by you with 'init'
```

## 🎯 **Benefits for Your Workflow:**

✅ **Travel-Optimized**: Small repo, fast clone/sync  
✅ **Self-Contained**: Everything needed in one place  
✅ **Cross-Platform**: Works on all your computers  
✅ **Secure**: No production LMS code  
✅ **Shareable**: Easy to collaborate with colleagues  
✅ **Professional**: Proper version control and documentation  

## 🚀 **Next Steps:**

1. **Follow the setup instructions above**
2. **Test on your current computer**  
3. **Create GitHub repository**
4. **Test cloning and setup on another computer**
5. **Start creating your assignment templates!**

This gives you a professional, portable, secure way to develop assignment templates across all your devices! 🌍✨
