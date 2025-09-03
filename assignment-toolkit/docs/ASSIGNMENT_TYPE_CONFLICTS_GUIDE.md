# 🔄 Assignment Type Conflicts & Resolution Guide

## 🚨 The Problem

Your LMS already supports these assignment types:
```
multiple-choice, true-false, matching, drag-and-drop, writing, writing-long, 
speaking, assignment, listening, line-match, phoneme-build, image-upload, code-submission
```

The portable assignment toolkit needed to work with these existing types while also providing:
- More intuitive names (e.g., "essay" instead of "writing-long")
- Conflict-free development
- Backward compatibility
- Clear mapping between portable and LMS formats

## ✅ The Solution: Smart Type Mapping

The toolkit now includes an intelligent type mapping system that:

1. **Resolves Conflicts**: Maps portable types to correct LMS types
2. **Provides Aliases**: Offers shortcuts and intuitive names
3. **Validates Input**: Catches typos and suggests corrections
4. **Maintains Compatibility**: Works seamlessly with existing LMS

## 📋 Type Mapping Reference

### 🎯 Direct Mappings (No Conflicts)
| Portable Type | LMS Type | LMS Subtype | Notes |
|---------------|----------|-------------|-------|
| `multiple-choice` | `multiple-choice` | - | Direct match |
| `true-false` | `true-false` | - | Direct match |
| `matching` | `matching` | - | Direct match |
| `speaking` | `speaking` | - | Direct match |
| `listening` | `listening` | - | Direct match |
| `code-submission` | `code-submission` | - | Direct match |
| `image-upload` | `image-upload` | - | Direct match |

### 🔄 Smart Mappings (Conflict Resolution)
| Portable Type | LMS Type | LMS Subtype | Reason |
|---------------|----------|-------------|---------|
| `essay` | `writing-long` | - | More intuitive name |
| `writing-short` | `writing` | - | Clearer distinction |
| `quiz` | `multiple-choice` | - | Common terminology |
| `presentation` | `speaking` | - | Specific use case |
| `comprehension` | `listening` | - | Educational context |
| `programming` | `code-submission` | - | Alternative name |

### 🎯 Drag-and-Drop Subtypes
| Portable Type | LMS Type | LMS Subtype | Use Case |
|---------------|----------|-------------|----------|
| `drag-drop-ordering` | `drag-and-drop` | `ordering` | Sequence tasks |
| `drag-drop-categorization` | `drag-and-drop` | `categorization` | Grouping tasks |
| `drag-drop-fill-blank` | `drag-and-drop` | `fill-blank` | Fill in blanks |
| `drag-drop-labeling` | `drag-and-drop` | `labeling` | Label diagrams |
| `drag-drop-image-caption` | `drag-and-drop` | `image-caption` | Caption images |

### 🎓 LMS-Specific Types
| Portable Type | LMS Type | Notes |
|---------------|----------|-------|
| `line-match` | `line-match` | Your LMS-specific type |
| `phoneme-build` | `phoneme-build` | Your LMS-specific type |
| `generic-assignment` | `assignment` | Fallback type |

### 🔄 Aliases (Shortcuts)
| Alias | Maps To | Example |
|-------|---------|---------|
| `mcq`, `mc` | `multiple-choice` | `create mcq` |
| `tf`, `t/f` | `true-false` | `create tf` |
| `match` | `matching` | `create match` |
| `code` | `code-submission` | `create code` |
| `dnd` | `drag-drop-ordering` | `create dnd` |
| `oral` | `speaking` | `create oral` |
| `audio` | `listening` | `create audio` |

## 🛠 How It Works

### 1. **Type Resolution Process**
```bash
# User input: "essay"
# 1. Check direct mappings: not found
# 2. Check aliases: not found  
# 3. Resolve to: writing-long
# 4. Result: Creates "essay" assignment, syncs as "writing-long"
```

### 2. **Error Handling**
```bash
# User input: "multiplechoice" (typo)
❌ Unknown assignment type: multiplechoice
📝 Did you mean one of these?
  • multiple-choice - Multiple choice questions
  • quiz - Quiz assignment (mapped to multiple-choice)

💡 Use 'assignment-toolkit types' to see all available types
```

### 3. **Smart Suggestions**
```bash
# User input: "drag"
❌ Unknown assignment type: drag
📝 Did you mean one of these?
  • drag-drop-ordering - Drag and drop ordering
  • drag-drop-categorization - Drag and drop categorization
  • drag-drop-fill-blank - Drag and drop fill in blanks
```

## 🎯 Usage Examples

### ✅ **Conflict-Free Creation**
```bash
# Create using intuitive names
assignment-toolkit create essay
# → Creates "essay" assignment
# → Will sync to LMS as "writing-long"

assignment-toolkit create quiz  
# → Creates "quiz" assignment
# → Will sync to LMS as "multiple-choice"

assignment-toolkit create presentation
# → Creates "presentation" assignment  
# → Will sync to LMS as "speaking"
```

### ✅ **Using Aliases**
```bash
# Quick shortcuts
assignment-toolkit create mcq    # → multiple-choice
assignment-toolkit create tf     # → true-false
assignment-toolkit create code   # → code-submission
assignment-toolkit create dnd    # → drag-drop-ordering
```

### ✅ **LMS-Specific Types**
```bash
# Use your specialized types
assignment-toolkit create line-match     # → line-match
assignment-toolkit create phoneme-build  # → phoneme-build
```

### ✅ **Drag-and-Drop Variants**
```bash
# Specific drag-and-drop types
assignment-toolkit create drag-drop-ordering        # → drag-and-drop (ordering)
assignment-toolkit create drag-drop-categorization  # → drag-and-drop (categorization)
assignment-toolkit create drag-drop-fill-blank      # → drag-and-drop (fill-blank)
```

## 🔍 Discovery Commands

### **List All Types**
```bash
assignment-toolkit types
```
Shows all available types organized by category with LMS mappings.

### **Get Help for Unknown Types**
```bash
assignment-toolkit create unknown-type
# → Shows suggestions and help
```

### **Interactive Selection**
```bash
assignment-toolkit create
# → Shows menu of all types with descriptions
```

## ⚠️ Important Notes

### **1. YAML File Format**
Your assignment files still use the portable type names:
```yaml
assignment:
  title: "My Essay Assignment"
  type: "essay"  # ← Portable type name
  # ... rest of assignment
```

### **2. LMS Synchronization**
During sync, the toolkit automatically converts:
```bash
assignment-toolkit sync my-essay.yaml
# Converts: essay → writing-long
# Sends to LMS as: type: "writing-long"
```

### **3. Backward Compatibility**
Existing assignments continue to work:
```yaml
assignment:
  type: "multiple-choice"  # ← Still works perfectly
```

### **4. Template Compatibility**
Templates can use either format:
```yaml
# Template using portable type
template:
  type: "essay"

# Template using LMS type  
template:
  type: "writing-long"
```

## 🔧 Customization

### **Adding Custom Mappings**
You can extend the type mappings by modifying `type-mapping.go`:

```go
// Add custom mapping
{"my-custom-type", "existing-lms-type", "", "My custom assignment type", false}
```

### **Adding Aliases**
```go
// Add custom alias
atm.aliases["shortcut"] = "target-type"
```

## 🐛 Troubleshooting

### **Problem: Type Not Found**
```bash
❌ Unknown assignment type: mytype
```
**Solution**: Use `assignment-toolkit types` to see available options.

### **Problem: Sync Fails with Type Error**
```bash
❌ API error (400): Invalid assignment type
```
**Solution**: The mapping may be incorrect. Check LMS API documentation.

### **Problem: Unexpected LMS Type**
```bash
# Expected "writing" but got "writing-long"
```
**Solution**: Check the mapping table above. Some types are automatically mapped for clarity.

## 📈 Benefits

### ✅ **For Developers**
- No naming conflicts with existing LMS
- Clear separation between portable and LMS formats
- Extensible mapping system
- Type safety and validation

### ✅ **For Content Creators**
- Intuitive type names ("essay" vs "writing-long")
- Helpful aliases and shortcuts
- Clear error messages and suggestions
- Seamless LMS integration

### ✅ **For System Administrators**
- Backward compatibility maintained
- Existing assignments unaffected
- Clear audit trail of type mappings
- Easy customization and extension

---

## 🎉 Result: Best of Both Worlds

✅ **Portable System**: Use intuitive names like "essay", "quiz", "presentation"  
✅ **LMS Compatibility**: Automatically maps to correct LMS types  
✅ **No Conflicts**: Existing LMS functionality unchanged  
✅ **User-Friendly**: Clear errors, suggestions, and help  
✅ **Extensible**: Easy to add new types and mappings  

**You can now create assignments using natural language while maintaining full compatibility with your existing LMS!**
