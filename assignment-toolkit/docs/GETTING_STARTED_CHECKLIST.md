# ✅ Getting Started Checklist - Do It Yourself!

## 🎯 What You're Going to Learn

By following this checklist, you'll learn to:
- Set up your own assignment template workspace
- Create assignment templates using the CLI
- Understand the YAML format
- Validate and improve template quality
- Use templates as blueprints for LMS content

## 📋 Your Step-by-Step Checklist

### ☐ **Step 1: Choose Your Workspace**

**⚠️ CRITICAL: Git Repository Location**

**❌ DO NOT create workspace inside an existing Git repository!**
```bash
# ❌ WRONG - Don't do this inside existing Git repos
cd /existing-git-repo/some-folder
assignment-toolkit init  # Creates nested Git issues!

# ✅ CORRECT - Create workspace outside existing repos
cd ~  # Home directory (safe)
mkdir assignment-templates
cd assignment-templates
```

**Safe locations for your workspace:**
```bash
# Pick a directory where you want to work
# Examples:
cd ~/Documents/assignment-templates
# OR
mkdir ~/Desktop/my-templates && cd ~/Desktop/my-templates
# OR
cd ~/assignment-work
# OR
cd /home/username/my-templates
```

**💡 Tips**: 
- Choose somewhere you'll remember and can easily access
- Make sure it's NOT inside another Git repository
- Use your home directory or Documents folder for safety

---

### ☐ **Step 2: Initialize Your Workspace**
```bash
# Run the init command (use the full path)
/mnt/LMS-database/repos/English-Foundation-Adventure/assignment-toolkit/assignment-toolkit init
```

**What will happen:**
- The toolkit will ask YOU questions
- Answer them with your information
- It will create files and folders for you

**Questions you'll see:**
- `Author name:` [Type your name]
- `Email:` [Type your email]

---

### ☐ **Step 3: Explore What Was Created**
```bash
# Look at your new workspace
ls -la

# Check your configuration
cat .assignment-config.yaml

# See the directory structure
ls templates/
ls resources/
ls packages/
```

**What to look for:**
- `.assignment-config.yaml` - Your personal settings
- `templates/` - Directory for templates
- `resources/` - Directory for files (images, PDFs, etc.)
- `packages/` - Directory for packaged assignments

---

### ☐ **Step 4: See Available Assignment Types**
```bash
# List all assignment types you can create
/mnt/LMS-database/repos/English-Foundation-Adventure/assignment-toolkit/assignment-toolkit types
```

**What you'll see:**
- All available assignment types
- How they map to your LMS
- Shortcuts and aliases you can use

---

### ☐ **Step 5: Create Your First Template**
```bash
# Start with something simple (multiple choice)
/mnt/LMS-database/repos/English-Foundation-Adventure/assignment-toolkit/assignment-toolkit create multiple-choice
```

**What will happen:**
- Interactive wizard will start
- You'll answer questions about your assignment
- A YAML file will be created

**Questions you might see:**
- `Assignment title:` [Type something like "Sample Quiz"]
- `Description:` [Describe what the assignment is about]
- `Category:` [Subject area like "Geography"]
- `Question:` [Type a sample question]
- `Options:` [Type answer choices]

---

### ☐ **Step 6: Look at What You Created**
```bash
# List your assignments
/mnt/LMS-database/repos/English-Foundation-Adventure/assignment-toolkit/assignment-toolkit list

# Look at the YAML file that was created
ls -la *.yaml
cat sample-quiz.yaml  # Or whatever filename was created
```

**What to understand:**
- The YAML format structure
- How your answers became structured data
- The different sections (metadata, assignment, resources, etc.)

---

### ☐ **Step 7: Validate Your Template**
```bash
# Check the quality of your template
/mnt/LMS-database/repos/English-Foundation-Adventure/assignment-toolkit/assignment-toolkit validate your-file.yaml
```

**What you'll learn:**
- Quality scoring (0-100 points)
- What makes a good assignment template
- Warnings and suggestions for improvement

---

### ☐ **Step 8: Try Different Assignment Types**
```bash
# Create different types to understand the differences
/mnt/LMS-database/repos/English-Foundation-Adventure/assignment-toolkit/assignment-toolkit create essay
/mnt/LMS-database/repos/English-Foundation-Adventure/assignment-toolkit/assignment-toolkit create matching
/mnt/LMS-database/repos/English-Foundation-Adventure/assignment-toolkit/assignment-toolkit create code-submission
```

**What you'll discover:**
- Different types have different question formats
- Some require different information
- How the interactive wizard adapts to each type

---

### ☐ **Step 9: Edit Templates Manually**
```bash
# Try editing a YAML file directly
code sample-quiz.yaml  # Or nano, gedit, etc.

# Make some changes, then validate again
/mnt/LMS-database/repos/English-Foundation-Adventure/assignment-toolkit/assignment-toolkit validate sample-quiz.yaml
```

**What you'll learn:**
- How to modify templates directly
- YAML syntax and structure
- How changes affect validation scores

---

### ☐ **Step 10: Package Your Templates**
```bash
# Create a distributable package
/mnt/LMS-database/repos/English-Foundation-Adventure/assignment-toolkit/assignment-toolkit package sample-quiz.yaml
```

**What this does:**
- Creates a folder with your template
- Includes any resource files
- Adds documentation
- Makes it easy to share with colleagues

---

## 🎓 **Learning Objectives**

After completing this checklist, you should understand:

✅ **How to set up your own workspace**  
✅ **How to create assignment templates interactively**  
✅ **How the YAML format works**  
✅ **How to validate and improve template quality**  
✅ **How different assignment types are structured**  
✅ **How to edit templates manually**  
✅ **How to package templates for sharing**  

## 🚀 **Your Next Steps**

Once you're comfortable with template creation:

1. **Create templates for your subjects** (math, English, science, etc.)
2. **Use templates as blueprints** when creating content in your LMS
3. **Share templates** with other teachers
4. **Set up LMS sync** when you want to upload directly

## 🆘 **If You Get Stuck**

### **Command Help**
```bash
# Get help for any command
assignment-toolkit --help
assignment-toolkit create --help
assignment-toolkit validate --help
```

### **Example Files**
```bash
# Look at working examples
ls /mnt/LMS-database/repos/English-Foundation-Adventure/assignment-toolkit/examples/
cat /mnt/LMS-database/repos/English-Foundation-Adventure/assignment-toolkit/examples/multiple-choice-basic.yaml
```

### **Documentation**
- **Full Guide**: `docs/portable_assignment_development_guide.md`
- **Type Conflicts**: `docs/ASSIGNMENT_TYPE_CONFLICTS_GUIDE.md`
- **Quick Reference**: `docs/portable_assignments_quick_reference.md`

---

## 🎯 **Remember**

- **Learn by doing** - follow each step yourself
- **Experiment** - try different assignment types
- **Validate often** - check your work as you go
- **Start simple** - begin with multiple-choice, then try others
- **Read the YAML** - understanding the format is key

**You've got this! 🚀**
