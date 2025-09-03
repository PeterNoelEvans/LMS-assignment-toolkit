# Assignment Toolkit

A powerful CLI tool for creating, managing, and syncing portable assignments with your LMS while traveling or working offline.

## 🚀 Features

- **Interactive Assignment Creation**: Create assignments using guided wizards
- **Multiple Assignment Types**: Support for all major assignment types
- **Offline Development**: Work without internet connectivity
- **Portable Format**: YAML-based format for easy version control
- **Resource Management**: Attach and manage learning resources
- **Validation**: Built-in validation for assignment quality
- **LMS Sync**: Seamless synchronization with your LMS
- **Template System**: Reusable templates for common assignment patterns
- **Batch Operations**: Handle multiple assignments efficiently

## 📋 Supported Assignment Types

- **multiple-choice**: Traditional multiple choice questions
- **true-false**: True/false questions
- **matching**: Match items from two lists
- **drag-and-drop**: Interactive drag and drop exercises
- **writing**: Short writing assignments
- **writing-long**: Extended writing assignments
- **speaking**: Oral presentation assignments
- **listening**: Audio comprehension exercises
- **code-submission**: Programming assignments

## 🛠 Installation

```bash
# Build from source
go build -o assignment-toolkit

# Make it executable globally (optional)
sudo mv assignment-toolkit /usr/local/bin/
```

## 🚀 Quick Start

### 1. Initialize Workspace

```bash
assignment-toolkit init
```

This creates:
- `.assignment-config.yaml` - Your configuration
- `templates/` - Assignment templates
- `resources/` - Resource files
- `packages/` - Packaged assignments

### 2. Configure LMS Connection

```bash
# Set your LMS endpoint
assignment-toolkit config set lms-endpoint https://your-lms.com

# Set your API key (if required)
assignment-toolkit config set api-key your-api-key
```

### 3. Create Your First Assignment

```bash
assignment-toolkit create multiple-choice
```

Follow the interactive wizard to create your assignment.

### 4. Validate Assignment

```bash
assignment-toolkit validate my-assignment.yaml
```

### 5. Sync with LMS

```bash
assignment-toolkit sync my-assignment.yaml
```

## 📝 Assignment Format

Assignments are stored in YAML format for easy editing and version control:

```yaml
metadata:
  id: "550e8400-e29b-41d4-a716-446655440000"
  version: "1.0.0"
  created: 2024-01-01T00:00:00Z
  modified: 2024-01-01T00:00:00Z
  author: "Your Name"
  license: "CC-BY-SA-4.0"
  language: "en"
  source_hash: "abc123..."

assignment:
  title: "Capital Cities Quiz"
  description: "Test your knowledge of world capitals"
  type: "multiple-choice"
  category: "Geography"
  difficulty: "beginner"
  points: 5
  auto_grade: true
  show_feedback: true
  published: true
  
  questions:
    question: "What is the capital of France?"
    options:
      - "London"
      - "Paris"
      - "Berlin"
      - "Madrid"
    correctAnswer: "Paris"
    explanation: "Paris is the capital and largest city of France."

resources:
  - id: "resource-1"
    title: "World Map"
    description: "Reference map for geography questions"
    type: "image"
    local_path: "./resources/world-map.png"
    is_public: true

dependencies:
  subjects: ["geography"]
  prerequisites: []
  recommended_courses: ["world-geography-101"]
```

## 🎯 Assignment Types Examples

### Multiple Choice

```yaml
assignment:
  type: "multiple-choice"
  questions:
    question: "What is 2 + 2?"
    options: ["3", "4", "5", "6"]
    correctAnswer: "4"
    explanation: "Basic addition"
```

### Matching

```yaml
assignment:
  type: "matching"
  questions:
    leftItems: ["France", "Germany", "Spain"]
    rightItems: ["Paris", "Berlin", "Madrid"]
```

### Code Submission

```yaml
assignment:
  type: "code-submission"
  instructions: "Write a function that calculates factorial"
  code_submission_config:
    programmingLanguage: "python"
    allowFileUpload: true
    maxFiles: 3
    maxFileSizeMb: 5
    expectedOutput: "factorial(5) = 120"
```

### Writing Assignment

```yaml
assignment:
  type: "writing"
  instructions: "Write a 500-word essay about climate change"
  criteria: "Grammar, structure, and argument quality will be evaluated"
  auto_grade: false
  time_limit: 3600  # 1 hour in seconds
```

## 📦 Commands Reference

### Core Commands

- `init` - Initialize assignment workspace
- `create [type]` - Create new assignment interactively
- `validate [file]` - Validate assignment package
- `list` - List all assignments in directory
- `sync [file]` - Sync assignment with LMS
- `package [file]` - Create distributable package

### Template Commands

- `template list` - List available templates
- `template create` - Create new template
- `template use [name]` - Create assignment from template

### Configuration Commands

- `config set [key] [value]` - Set configuration value
- `config get [key]` - Get configuration value
- `config list` - List all configuration

## 🔧 Configuration

The toolkit uses `.assignment-config.yaml` for configuration:

```yaml
author: "Your Name"
email: "your.email@example.com"
license: "CC-BY-SA-4.0"
language: "en"
lms_endpoint: "https://your-lms.com"
api_key: "your-api-key"

defaults:
  points: "1"
  auto_grade: "true"
  published: "true"
  quarter: "Q1"

templates:
  multiple-choice: "./templates/multiple-choice.yaml"
  writing: "./templates/writing.yaml"
```

## 🔄 Workflow Examples

### Offline Assignment Creation

```bash
# Initialize workspace
assignment-toolkit init

# Create multiple assignments offline
assignment-toolkit create multiple-choice
assignment-toolkit create writing
assignment-toolkit create code-submission

# Validate all assignments
for file in *.yaml; do
  assignment-toolkit validate "$file"
done

# When online, sync all assignments
for file in *.yaml; do
  assignment-toolkit sync "$file"
done
```

### Template-Based Development

```bash
# Create a template
assignment-toolkit template create

# Use template for new assignments
assignment-toolkit template use multiple-choice
assignment-toolkit template use writing-essay
```

### Batch Operations

```bash
# Package all assignments for distribution
for file in *.yaml; do
  assignment-toolkit package "$file"
done

# Validate entire directory
assignment-toolkit validate-all
```

## 🔒 Security Considerations

- API keys are stored in local configuration files
- Use environment variables for sensitive data in CI/CD
- Assignment files can be version controlled safely
- Resource files should be managed separately for large files

## 🤝 Integration with LMS

The toolkit integrates with your LMS through REST API:

- **Authentication**: Bearer token or API key
- **Assignment Import**: POST `/api/assignments`
- **Resource Upload**: POST `/api/resources`
- **Batch Operations**: POST `/api/assignments/batch`

### Required LMS Endpoints

Your LMS should support these endpoints:

```
GET  /api/auth/me                    # Test authentication
POST /api/assignments                # Create assignment
POST /api/resources                  # Upload resource
GET  /api/assignments?sourceHash=X   # Check for duplicates
```

## 📊 Quality Scoring

The validator assigns quality scores (0-100) based on:

- **Required Fields** (40 points): Title, type, questions
- **Educational Content** (30 points): Description, learning objectives
- **Technical Quality** (20 points): Proper formatting, validation
- **Completeness** (10 points): Resources, metadata

## 🐛 Troubleshooting

### Common Issues

**Assignment validation fails**
- Check required fields are present
- Verify question format matches assignment type
- Ensure YAML syntax is correct

**Sync fails with authentication error**
- Verify LMS endpoint is correct
- Check API key is valid and has proper permissions
- Test connection with `assignment-toolkit config test`

**Resource upload fails**
- Check file paths are correct
- Verify file sizes are within limits
- Ensure file types are supported by LMS

### Debug Mode

Enable debug output:
```bash
export DEBUG=true
assignment-toolkit sync my-assignment.yaml
```

## 📚 Examples Directory

Check the `examples/` directory for sample assignments:

- `examples/multiple-choice-basic.yaml`
- `examples/matching-geography.yaml`
- `examples/code-submission-python.yaml`
- `examples/writing-essay.yaml`

## 🔮 Roadmap

- [ ] GUI interface for non-technical users
- [ ] Advanced question types (drag-and-drop, hotspot)
- [ ] Integration with popular LMS platforms
- [ ] Assignment analytics and reporting
- [ ] Collaborative editing features
- [ ] Mobile app for assignment creation

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📞 Support

- Documentation: [Wiki](wiki)
- Issues: [GitHub Issues](issues)
- Discussions: [GitHub Discussions](discussions)
- Email: support@assignment-toolkit.com
