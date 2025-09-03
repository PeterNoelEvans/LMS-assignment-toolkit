# 🚀 Portable Assignments - Quick Reference Card

## 📋 Essential Commands

```bash
# Setup YOUR workspace (one-time)
# ⚠️  IMPORTANT: Don't create inside existing Git repos!

# 1. Go to SAFE location (outside any Git repository)
cd ~  # Home directory is always safe
mkdir assignment-templates && cd assignment-templates

# 2. Initialize (toolkit will ask YOU questions)
/path/to/assignment-toolkit/assignment-toolkit init

# 3. Optional: Configure LMS (only if you plan to sync)
assignment-toolkit config set lms-endpoint "https://lms-pne.uk"
```

## 🚨 **Git Repository Warning**

**❌ NEVER do this:**
```bash
cd /existing-git-repo/subfolder
assignment-toolkit init  # Creates nested Git conflicts!
```

**✅ ALWAYS do this:**
```bash
cd ~  # or ~/Documents, ~/Desktop, etc.
mkdir my-templates && cd my-templates
assignment-toolkit init  # Safe!
```

# Create assignments
assignment-toolkit create multiple-choice
assignment-toolkit create matching
assignment-toolkit create code-submission

# Quality check
assignment-toolkit validate my-assignment.yaml

# List assignments
assignment-toolkit list

# Sync with LMS (when online)
assignment-toolkit sync my-assignment.yaml
```

## 🎯 Assignment Types

| Portable Type | Command | LMS Type | Best For |
|---------------|---------|----------|----------|
| `multiple-choice` | `create multiple-choice` | `multiple-choice` | Quizzes, tests |
| `essay` | `create essay` | `writing-long` | Essays, papers |
| `matching` | `create matching` | `matching` | Vocabulary, concepts |
| `code-submission` | `create code-submission` | `code-submission` | Programming |
| `speaking` | `create speaking` | `speaking` | Presentations |
| `listening` | `create listening` | `listening` | Audio comprehension |

**🔄 Shortcuts**: `mcq`→multiple-choice, `tf`→true-false, `code`→code-submission

**📋 See all types**: `assignment-toolkit types`

## 📝 YAML Structure (Quick)

```yaml
metadata:
  id: "unique-id"
  version: "1.0.0"
  author: "Your Name"

assignment:
  title: "Assignment Title"
  description: "What students need to do"
  type: "multiple-choice"
  points: 5
  
  questions:
    question: "Your question?"
    options: ["A", "B", "C", "D"]
    correctAnswer: "B"

resources:
  - title: "Reference"
    local_path: "./resources/file.pdf"
```

## ⚡ Workflow Shortcuts

### Offline Development
```bash
# 1. Create workspace
mkdir ~/travel-assignments && cd ~/travel-assignments
assignment-toolkit init

# 2. Batch create
assignment-toolkit create multiple-choice  # Quiz 1
assignment-toolkit create matching         # Exercise 1
assignment-toolkit create writing          # Essay 1

# 3. Validate all
for f in *.yaml; do assignment-toolkit validate "$f"; done
```

### Online Sync
```bash
# Sync all assignments
for f in *.yaml; do assignment-toolkit sync "$f"; done

# Check what's ready to sync
assignment-toolkit list
```

## 🔧 Common Fixes

### Validation Errors
```bash
# Missing title
assignment:
  title: "Add this!"  # Required

# Missing questions for quiz types  
questions:
  question: "Add question here"
  options: ["A", "B", "C"]
  correctAnswer: "A"
```

### File Issues
```bash
# Check file exists
ls -la my-assignment.yaml

# Fix permissions
chmod 644 my-assignment.yaml

# Validate YAML syntax
assignment-toolkit validate my-assignment.yaml
```

## 📁 Directory Structure

```
my-assignments/
├── .assignment-config.yaml  # Config
├── math-quiz.yaml          # Assignments
├── english-essay.yaml
├── resources/              # Files
│   ├── images/
│   └── documents/
└── templates/              # Reusable templates
```

## 🎯 Quality Checklist

- [ ] Title is descriptive and specific
- [ ] Description explains what to do
- [ ] Questions match the assignment type
- [ ] Points value is appropriate
- [ ] Learning objectives included
- [ ] Resources added if needed
- [ ] Validation score > 90

## 🔄 Sync Status

```bash
# Test connection
assignment-toolkit config test

# Check sync status
assignment-toolkit list

# Force sync if needed
assignment-toolkit sync --force my-assignment.yaml
```

## 📞 Emergency Commands

```bash
# Get help
assignment-toolkit --help
assignment-toolkit create --help

# Debug mode
DEBUG=true assignment-toolkit sync my-assignment.yaml

# Validate examples
assignment-toolkit validate examples/multiple-choice-basic.yaml
```

## 🎨 Templates

```bash
# Copy working example
cp examples/multiple-choice-basic.yaml my-new-quiz.yaml

# Edit and customize
nano my-new-quiz.yaml

# Validate before using
assignment-toolkit validate my-new-quiz.yaml
```

---

**💡 Pro Tips:**
- Work offline, sync when convenient
- Validate before syncing
- Use descriptive filenames
- Keep resources under 10MB
- Version control with git

**🆘 Need Help?**
- Full guide: `docs/portable_assignment_development_guide.md`
- Examples: `assignment-toolkit/examples/`
- Command help: `assignment-toolkit [command] --help`
