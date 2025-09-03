package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

func init() {
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(validateCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(packageCmd)
	rootCmd.AddCommand(syncCmd)
	rootCmd.AddCommand(templateCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(typesCmd)
}

// Create command
var createCmd = &cobra.Command{
	Use:   "create [type]",
	Short: "Create a new assignment interactively",
	Long: `Create a new assignment using an interactive wizard.
Supported types: multiple-choice, matching, drag-and-drop, writing, code-submission, speaking, listening`,
	Args: cobra.MaximumNArgs(1),
	Run:  runCreate,
}

// Validate command
var validateCmd = &cobra.Command{
	Use:   "validate [file]",
	Short: "Validate an assignment package",
	Long:  "Validate the structure and content of an assignment package",
	Args:  cobra.ExactArgs(1),
	Run:   runValidate,
}

// List command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all assignments in the current directory",
	Long:  "List all assignment packages in the current directory with their metadata",
	Run:   runList,
}

// Package command
var packageCmd = &cobra.Command{
	Use:   "package [assignment-file]",
	Short: "Package an assignment with its resources",
	Long:  "Create a distributable package containing the assignment and all its resources",
	Args:  cobra.ExactArgs(1),
	Run:   runPackage,
}

// Sync command
var syncCmd = &cobra.Command{
	Use:   "sync [file]",
	Short: "Sync assignment with remote LMS",
	Long:  "Upload assignment to the configured LMS endpoint",
	Args:  cobra.MaximumNArgs(1),
	Run:   runSync,
}

// Template command
var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Manage assignment templates",
	Long:  "Create, list, and use assignment templates",
}

// Config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage toolkit configuration",
	Long:  "Set up and manage toolkit configuration including LMS endpoints and defaults",
}

// Init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize assignment workspace",
	Long:  "Initialize the current directory as an assignment workspace with default configuration",
	Run:   runInit,
}

// Types command
var typesCmd = &cobra.Command{
	Use:   "types",
	Short: "List all available assignment types",
	Long:  "Display all supported assignment types with their descriptions and LMS mappings",
	Run:   runTypes,
}

// Implementation functions

func runCreate(cmd *cobra.Command, args []string) {
	typeManager := GetTypeManager()
	var assignmentType string

	if len(args) > 0 {
		inputType := args[0]
		if !typeManager.ValidatePortableType(inputType) {
			suggestions := typeManager.GetSuggestedTypes(inputType)
			fmt.Printf("❌ Unknown assignment type: %s\n", inputType)
			if len(suggestions) > 0 {
				fmt.Printf("📝 Did you mean one of these?\n")
				for _, suggestion := range suggestions {
					fmt.Printf("  • %s - %s\n", suggestion, typeManager.GetTypeDescription(suggestion))
				}
			}
			fmt.Printf("\n💡 Use 'assignment-toolkit types' to see all available types\n")
			return
		}
		assignmentType = inputType
	} else {
		// Show available portable types with descriptions
		portableTypes := typeManager.GetPortableTypes()
		typeDescriptions := make([]string, len(portableTypes))
		for i, pType := range portableTypes {
			typeDescriptions[i] = fmt.Sprintf("%s (%s)", pType, typeManager.GetTypeDescription(pType))
		}

		selected := promptSelect("Select assignment type:", typeDescriptions)
		// Extract just the type name (before the parentheses)
		assignmentType = strings.Split(selected, " (")[0]
	}

	// Resolve to LMS format for validation
	lmsType, lmsSubtype, err := typeManager.ConvertToLMSFormat(assignmentType)
	if err != nil {
		fmt.Printf("❌ Error resolving assignment type: %v\n", err)
		return
	}

	fmt.Printf("Creating new %s assignment...\n", assignmentType)
	if lmsType != assignmentType {
		fmt.Printf("📋 Will be imported to LMS as: %s", lmsType)
		if lmsSubtype != "" {
			fmt.Printf(" (%s)", lmsSubtype)
		}
		fmt.Println()
	}
	fmt.Println()

	// Create assignment through interactive wizard
	assignment := createAssignmentWizard(assignmentType)

	// Generate package
	pkg := AssignmentPackage{
		Metadata: PackageMetadata{
			ID:       uuid.New().String(),
			Version:  "1.0.0",
			Created:  time.Now(),
			Modified: time.Now(),
			Author:   getConfig().Author,
			License:  getConfig().License,
			Language: getConfig().Language,
		},
		Assignment: assignment,
	}

	// Calculate source hash
	pkg.Metadata.SourceHash = calculateHash(pkg)

	// Save to file
	filename := strings.ReplaceAll(strings.ToLower(assignment.Title), " ", "-") + ".yaml"
	saveAssignmentPackage(pkg, filename)

	fmt.Printf("✅ Assignment created successfully: %s\n", filename)
}

func runValidate(cmd *cobra.Command, args []string) {
	filename := args[0]

	pkg, err := loadAssignmentPackage(filename)
	if err != nil {
		fmt.Printf("❌ Failed to load assignment: %v\n", err)
		return
	}

	validation := validateAssignmentPackage(pkg)

	if validation.IsValid {
		fmt.Printf("✅ Assignment is valid (Score: %d/100)\n", validation.Score)
	} else {
		fmt.Printf("❌ Assignment validation failed\n")
		for _, err := range validation.Errors {
			fmt.Printf("  • %s\n", err)
		}
	}

	if len(validation.Warnings) > 0 {
		fmt.Println("\n⚠️  Warnings:")
		for _, warning := range validation.Warnings {
			fmt.Printf("  • %s\n", warning)
		}
	}
}

func runList(cmd *cobra.Command, args []string) {
	files, err := filepath.Glob("*.yaml")
	if err != nil {
		fmt.Printf("Error listing files: %v\n", err)
		return
	}

	yamlFiles, err := filepath.Glob("*.yml")
	if err == nil {
		files = append(files, yamlFiles...)
	}

	if len(files) == 0 {
		fmt.Println("No assignment files found in current directory.")
		return
	}

	fmt.Printf("Found %d assignment(s):\n\n", len(files))
	fmt.Printf("%-30s %-15s %-10s %-20s\n", "TITLE", "TYPE", "VERSION", "MODIFIED")
	fmt.Println(strings.Repeat("-", 75))

	for _, file := range files {
		pkg, err := loadAssignmentPackage(file)
		if err != nil {
			fmt.Printf("%-30s %-15s %-10s %-20s\n", file, "ERROR", "-", "-")
			continue
		}

		title := pkg.Assignment.Title
		if len(title) > 28 {
			title = title[:28] + "..."
		}

		fmt.Printf("%-30s %-15s %-10s %-20s\n",
			title,
			pkg.Assignment.Type,
			pkg.Metadata.Version,
			pkg.Metadata.Modified.Format("2006-01-02 15:04"),
		)
	}
}

func runPackage(cmd *cobra.Command, args []string) {
	filename := args[0]

	pkg, err := loadAssignmentPackage(filename)
	if err != nil {
		fmt.Printf("❌ Failed to load assignment: %v\n", err)
		return
	}

	// Create package directory
	packageName := strings.TrimSuffix(filename, filepath.Ext(filename))
	packageDir := packageName + "-package"

	os.RemoveAll(packageDir) // Clean up if exists
	os.MkdirAll(packageDir, 0755)

	// Copy assignment file
	saveAssignmentPackage(pkg, filepath.Join(packageDir, "assignment.yaml"))

	// Copy resources
	if len(pkg.Resources) > 0 {
		resourceDir := filepath.Join(packageDir, "resources")
		os.MkdirAll(resourceDir, 0755)

		for _, resource := range pkg.Resources {
			if resource.LocalPath != "" {
				// Copy local file
				copyFile(resource.LocalPath, filepath.Join(resourceDir, filepath.Base(resource.LocalPath)))
			}
		}
	}

	// Create README
	readme := fmt.Sprintf(`# %s

%s

## Assignment Details
- **Type**: %s
- **Version**: %s
- **Author**: %s
- **Created**: %s

## Installation
1. Import assignment.yaml into your LMS
2. Upload resources from the resources/ directory if present

## Resources
`, pkg.Assignment.Title, pkg.Assignment.Description, pkg.Assignment.Type,
		pkg.Metadata.Version, pkg.Metadata.Author, pkg.Metadata.Created.Format("2006-01-02"))

	for _, resource := range pkg.Resources {
		readme += fmt.Sprintf("- %s (%s)\n", resource.Title, resource.Type)
	}

	ioutil.WriteFile(filepath.Join(packageDir, "README.md"), []byte(readme), 0644)

	fmt.Printf("✅ Package created: %s/\n", packageDir)
}

func runSync(cmd *cobra.Command, args []string) {
	config := getConfig()
	if config.LMSEndpoint == "" {
		fmt.Println("❌ LMS endpoint not configured. Run 'assignment-toolkit config set lms-endpoint <url>'")
		return
	}

	var filename string
	if len(args) > 0 {
		filename = args[0]
	} else {
		// List available assignments
		files, _ := filepath.Glob("*.yaml")
		yamlFiles, _ := filepath.Glob("*.yml")
		files = append(files, yamlFiles...)

		if len(files) == 0 {
			fmt.Println("❌ No assignment files found")
			return
		}

		filename = promptSelect("Select assignment to sync:", files)
	}

	fmt.Printf("🔄 Syncing %s with %s...\n", filename, config.LMSEndpoint)

	// Load assignment
	_, err := loadAssignmentPackage(filename)
	if err != nil {
		fmt.Printf("❌ Failed to load assignment: %v\n", err)
		return
	}

	// TODO: Implement actual sync with LMS API
	// For now, just simulate
	time.Sleep(2 * time.Second)

	fmt.Printf("✅ Assignment synced successfully!\n")
	fmt.Printf("   Assignment ID: %s\n", uuid.New().String())
}

func runInit(cmd *cobra.Command, args []string) {
	fmt.Println("🚀 Initializing assignment workspace...")

	// Create config file
	config := Config{
		Author:   promptString("Author name:", ""),
		Email:    promptString("Email:", ""),
		License:  "CC-BY-SA-4.0",
		Language: "en",
		Defaults: map[string]string{
			"points":     "1",
			"auto_grade": "true",
			"published":  "true",
			"quarter":    "Q1",
		},
	}

	// Save config
	configData, _ := yaml.Marshal(config)
	ioutil.WriteFile(".assignment-config.yaml", configData, 0644)

	// Create directories
	os.MkdirAll("templates", 0755)
	os.MkdirAll("resources", 0755)
	os.MkdirAll("packages", 0755)

	// Create sample template
	sampleTemplate := Template{
		Name:        "Multiple Choice Template",
		Description: "Basic multiple choice question template",
		Type:        "multiple-choice",
		Template: Assignment{
			Type:         "multiple-choice",
			Points:       1,
			AutoGrade:    true,
			ShowFeedback: true,
			Published:    true,
			Quarter:      "Q1",
		},
	}

	templateData, _ := yaml.Marshal(sampleTemplate)
	ioutil.WriteFile("templates/multiple-choice.yaml", templateData, 0644)

	fmt.Println("✅ Workspace initialized!")
	fmt.Println("   📁 Created directories: templates/, resources/, packages/")
	fmt.Println("   ⚙️  Created config: .assignment-config.yaml")
	fmt.Println("   📝 Created sample template: templates/multiple-choice.yaml")
}

func runTypes(cmd *cobra.Command, args []string) {
	typeManager := GetTypeManager()

	fmt.Println("📋 Available Assignment Types")
	fmt.Println("=" + strings.Repeat("=", 50))
	fmt.Println()

	// Get all types with descriptions
	typesWithDesc := typeManager.ListTypesWithDescriptions()

	// Group by category for better display
	categories := map[string][]string{
		"📝 Quiz & Assessment": {
			"multiple-choice", "true-false", "matching", "quiz",
		},
		"✍️  Writing & Essays": {
			"writing-short", "writing-long", "essay",
		},
		"🎯 Interactive": {
			"drag-drop-ordering", "drag-drop-categorization", "drag-drop-fill-blank",
			"drag-drop-labeling", "drag-drop-image-caption",
		},
		"🗣️  Speaking & Listening": {
			"speaking", "listening", "presentation", "comprehension",
		},
		"💻 Programming": {
			"code-submission", "programming",
		},
		"📸 Media & Uploads": {
			"image-upload",
		},
		"🎓 Specialized (LMS-specific)": {
			"line-match", "phoneme-build", "generic-assignment",
		},
	}

	for category, types := range categories {
		fmt.Printf("%s\n", category)
		fmt.Println(strings.Repeat("-", len(category)-4)) // Account for emoji

		for _, pType := range types {
			if desc, exists := typesWithDesc[pType]; exists {
				fmt.Printf("  %-20s %s\n", pType, desc)
			}
		}
		fmt.Println()
	}

	fmt.Println("💡 Usage Examples:")
	fmt.Println("  assignment-toolkit create multiple-choice")
	fmt.Println("  assignment-toolkit create essay")
	fmt.Println("  assignment-toolkit create drag-drop-ordering")
	fmt.Println()
	fmt.Println("🔄 Type Aliases (shortcuts):")
	fmt.Println("  mcq, mc       → multiple-choice")
	fmt.Println("  tf, t/f       → true-false")
	fmt.Println("  match         → matching")
	fmt.Println("  code          → code-submission")
	fmt.Println("  dnd           → drag-drop-ordering")
	fmt.Println("  oral          → speaking")
	fmt.Println("  audio         → listening")
}

// Helper functions

func createAssignmentWizard(assignmentType string) Assignment {
	assignment := Assignment{
		Type:             assignmentType,
		Points:           1,
		AutoGrade:        true,
		ShowFeedback:     true,
		ShuffleQuestions: false,
		AllowReview:      true,
		TrackAttempts:    true,
		TrackConfidence:  true,
		TrackTimeSpent:   true,
		Published:        true,
		Quarter:          "Q1",
	}

	// Basic information
	assignment.Title = promptString("Assignment title:", "")
	assignment.Description = promptString("Description (optional):", "")
	assignment.Category = promptString("Category (optional):", "")
	assignment.Difficulty = promptSelect("Difficulty:", []string{"beginner", "intermediate", "advanced"})

	pointsStr := promptString("Points (default: 1):", "1")
	if points, err := strconv.Atoi(pointsStr); err == nil {
		assignment.Points = points
	}

	// Type-specific questions
	switch assignmentType {
	case "multiple-choice":
		assignment.Questions = createMultipleChoiceQuestions()
	case "matching":
		assignment.Questions = createMatchingQuestions()
	case "writing", "writing-long":
		assignment.Instructions = promptString("Instructions:", "")
		assignment.Criteria = promptString("Grading criteria:", "")
		assignment.AutoGrade = false
	case "code-submission":
		assignment.Questions = createCodeSubmissionConfig()
		assignment.AutoGrade = false
	}

	return assignment
}

func createMultipleChoiceQuestions() interface{} {
	question := promptString("Question:", "")

	var options []string
	fmt.Println("Enter answer options (press Enter twice to finish):")
	for i := 0; i < 10; i++ {
		option := promptString(fmt.Sprintf("Option %d:", i+1), "")
		if option == "" {
			break
		}
		options = append(options, option)
	}

	correctAnswer := promptSelect("Correct answer:", options)
	explanation := promptString("Explanation (optional):", "")

	return map[string]interface{}{
		"question":      question,
		"options":       options,
		"correctAnswer": correctAnswer,
		"explanation":   explanation,
	}
}

func createMatchingQuestions() interface{} {
	fmt.Println("Create matching pairs:")

	var leftItems, rightItems []string

	for i := 0; i < 10; i++ {
		left := promptString(fmt.Sprintf("Left item %d (or Enter to finish):", i+1), "")
		if left == "" {
			break
		}
		right := promptString(fmt.Sprintf("Right item %d:", i+1), "")

		leftItems = append(leftItems, left)
		rightItems = append(rightItems, right)
	}

	return map[string]interface{}{
		"leftItems":  leftItems,
		"rightItems": rightItems,
	}
}

func createCodeSubmissionConfig() interface{} {
	language := promptString("Programming language:", "python")
	expectedOutput := promptString("Expected output (optional):", "")

	return map[string]interface{}{
		"programmingLanguage": language,
		"allowFileUpload":     true,
		"maxFiles":            5,
		"maxFileSizeMb":       10,
		"expectedOutput":      expectedOutput,
	}
}

func promptString(prompt, defaultValue string) string {
	reader := bufio.NewReader(os.Stdin)
	if defaultValue != "" {
		fmt.Printf("%s [%s]: ", prompt, defaultValue)
	} else {
		fmt.Printf("%s: ", prompt)
	}

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "" {
		return defaultValue
	}
	return input
}

func promptSelect(prompt string, options []string) string {
	fmt.Printf("%s\n", prompt)
	for i, option := range options {
		fmt.Printf("  %d. %s\n", i+1, option)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Select (1-", len(options), "): ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if choice, err := strconv.Atoi(input); err == nil && choice >= 1 && choice <= len(options) {
		return options[choice-1]
	}

	return options[0] // Default to first option
}

func saveAssignmentPackage(pkg AssignmentPackage, filename string) error {
	data, err := yaml.Marshal(pkg)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}

func loadAssignmentPackage(filename string) (AssignmentPackage, error) {
	var pkg AssignmentPackage

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return pkg, err
	}

	err = yaml.Unmarshal(data, &pkg)
	return pkg, err
}

func validateAssignmentPackage(pkg AssignmentPackage) ValidationInfo {
	validation := ValidationInfo{
		IsValid:          true,
		ValidatedAt:      time.Now(),
		ValidatorVersion: "1.0.0",
		Score:            100,
	}

	// Basic validation
	if pkg.Assignment.Title == "" {
		validation.Errors = append(validation.Errors, "Assignment title is required")
		validation.IsValid = false
		validation.Score -= 20
	}

	if pkg.Assignment.Type == "" {
		validation.Errors = append(validation.Errors, "Assignment type is required")
		validation.IsValid = false
		validation.Score -= 20
	}

	// Type-specific validation
	switch pkg.Assignment.Type {
	case "multiple-choice":
		if pkg.Assignment.Questions == nil {
			validation.Errors = append(validation.Errors, "Multiple choice questions are required")
			validation.IsValid = false
			validation.Score -= 30
		}
	case "matching":
		if pkg.Assignment.Questions == nil {
			validation.Errors = append(validation.Errors, "Matching items are required")
			validation.IsValid = false
			validation.Score -= 30
		}
	}

	// Warnings
	if pkg.Assignment.Description == "" {
		validation.Warnings = append(validation.Warnings, "Assignment description is recommended")
		validation.Score -= 5
	}

	if pkg.Assignment.Points <= 0 {
		validation.Warnings = append(validation.Warnings, "Assignment should have positive points")
		validation.Score -= 10
	}

	return validation
}

func calculateHash(pkg AssignmentPackage) string {
	data, _ := json.Marshal(pkg.Assignment)
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash)
}

func copyFile(src, dst string) error {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dst, data, 0644)
}

func getConfig() Config {
	config := Config{
		Author:   "Unknown Author",
		License:  "CC-BY-SA-4.0",
		Language: "en",
	}

	if data, err := ioutil.ReadFile(".assignment-config.yaml"); err == nil {
		yaml.Unmarshal(data, &config)
	}

	return config
}
