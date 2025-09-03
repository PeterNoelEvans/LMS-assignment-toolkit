package main

import (
	"fmt"
	"strings"
)

// TypeMapping handles assignment type conflicts and transformations
type TypeMapping struct {
	PortableType string `json:"portable_type"`
	LMSType      string `json:"lms_type"`
	LMSSubtype   string `json:"lms_subtype,omitempty"`
	Description  string `json:"description"`
	Deprecated   bool   `json:"deprecated,omitempty"`
}

// AssignmentTypeManager manages type mappings and conflicts
type AssignmentTypeManager struct {
	mappings map[string]TypeMapping
	aliases  map[string]string
}

// NewAssignmentTypeManager creates a new type manager with default mappings
func NewAssignmentTypeManager() *AssignmentTypeManager {
	manager := &AssignmentTypeManager{
		mappings: make(map[string]TypeMapping),
		aliases:  make(map[string]string),
	}

	// Initialize default mappings
	manager.initializeDefaultMappings()
	return manager
}

// initializeDefaultMappings sets up the default type mappings
func (atm *AssignmentTypeManager) initializeDefaultMappings() {
	// Direct mappings (no conflicts)
	directMappings := []TypeMapping{
		{"multiple-choice", "multiple-choice", "", "Multiple choice questions", false},
		{"true-false", "true-false", "", "True/false questions", false},
		{"matching", "matching", "", "Match items from two lists", false},
		{"writing-short", "writing", "", "Short writing assignments", false},
		{"writing-long", "writing-long", "", "Extended writing assignments", false},
		{"speaking", "speaking", "", "Oral presentation assignments", false},
		{"listening", "listening", "", "Audio comprehension exercises", false},
		{"code-submission", "code-submission", "", "Programming assignments", false},
		{"image-upload", "image-upload", "", "Image upload assignments", false},

		// Specialized LMS types
		{"line-match", "line-match", "", "Line matching exercises (LMS specific)", false},
		{"phoneme-build", "phoneme-build", "", "Phoneme building exercises (LMS specific)", false},

		// Drag-and-drop with subtypes
		{"drag-drop-ordering", "drag-and-drop", "ordering", "Drag and drop ordering", false},
		{"drag-drop-categorization", "drag-and-drop", "categorization", "Drag and drop categorization", false},
		{"drag-drop-fill-blank", "drag-and-drop", "fill-blank", "Drag and drop fill in blanks", false},
		{"drag-drop-labeling", "drag-and-drop", "labeling", "Drag and drop labeling", false},
		{"drag-drop-image-caption", "drag-and-drop", "image-caption", "Drag and drop image captions", false},

		// Generic assignment (conflict resolution)
		{"generic-assignment", "assignment", "", "Generic assignment type", false},

		// Portable-specific types (not in LMS)
		{"essay", "writing-long", "", "Essay assignment (mapped to writing-long)", false},
		{"quiz", "multiple-choice", "", "Quiz assignment (mapped to multiple-choice)", false},
		{"presentation", "speaking", "", "Presentation (mapped to speaking)", false},
		{"comprehension", "listening", "", "Comprehension exercise (mapped to listening)", false},
	}

	for _, mapping := range directMappings {
		atm.mappings[mapping.PortableType] = mapping
	}

	// Set up aliases for common terms
	atm.aliases = map[string]string{
		"mcq":         "multiple-choice",
		"mc":          "multiple-choice",
		"tf":          "true-false",
		"t/f":         "true-false",
		"match":       "matching",
		"essay":       "writing-long",
		"short-essay": "writing-short",
		"code":        "code-submission",
		"programming": "code-submission",
		"drag-drop":   "drag-drop-ordering",
		"dnd":         "drag-drop-ordering",
		"oral":        "speaking",
		"audio":       "listening",
		"image":       "image-upload",
		"upload":      "image-upload",
	}
}

// ResolveType resolves a portable type to LMS format
func (atm *AssignmentTypeManager) ResolveType(portableType string) (TypeMapping, error) {
	// Normalize input
	normalizedType := strings.ToLower(strings.TrimSpace(portableType))

	// Check direct mapping first
	if mapping, exists := atm.mappings[normalizedType]; exists {
		return mapping, nil
	}

	// Check aliases
	if aliasTarget, exists := atm.aliases[normalizedType]; exists {
		if mapping, exists := atm.mappings[aliasTarget]; exists {
			return mapping, nil
		}
	}

	// Return error for unknown types
	return TypeMapping{}, fmt.Errorf("unknown assignment type: %s", portableType)
}

// GetPortableTypes returns all available portable types
func (atm *AssignmentTypeManager) GetPortableTypes() []string {
	var types []string
	for portableType := range atm.mappings {
		types = append(types, portableType)
	}
	return types
}

// GetLMSTypes returns all LMS types
func (atm *AssignmentTypeManager) GetLMSTypes() []string {
	lmsTypes := make(map[string]bool)
	for _, mapping := range atm.mappings {
		lmsTypes[mapping.LMSType] = true
	}

	var types []string
	for lmsType := range lmsTypes {
		types = append(types, lmsType)
	}
	return types
}

// ValidatePortableType checks if a portable type is valid
func (atm *AssignmentTypeManager) ValidatePortableType(portableType string) bool {
	_, err := atm.ResolveType(portableType)
	return err == nil
}

// GetTypeDescription returns a description for a type
func (atm *AssignmentTypeManager) GetTypeDescription(portableType string) string {
	if mapping, err := atm.ResolveType(portableType); err == nil {
		return mapping.Description
	}
	return "Unknown assignment type"
}

// ConvertToLMSFormat converts a portable assignment type to LMS format
func (atm *AssignmentTypeManager) ConvertToLMSFormat(portableType string) (string, string, error) {
	mapping, err := atm.ResolveType(portableType)
	if err != nil {
		return "", "", err
	}

	return mapping.LMSType, mapping.LMSSubtype, nil
}

// GetSuggestedTypes returns type suggestions for invalid input
func (atm *AssignmentTypeManager) GetSuggestedTypes(input string) []string {
	input = strings.ToLower(input)
	var suggestions []string

	// Check for partial matches
	for portableType := range atm.mappings {
		if strings.Contains(portableType, input) || strings.Contains(input, portableType) {
			suggestions = append(suggestions, portableType)
		}
	}

	// Check aliases
	for alias, target := range atm.aliases {
		if strings.Contains(alias, input) || strings.Contains(input, alias) {
			suggestions = append(suggestions, target)
		}
	}

	// Remove duplicates
	seen := make(map[string]bool)
	var unique []string
	for _, suggestion := range suggestions {
		if !seen[suggestion] {
			seen[suggestion] = true
			unique = append(unique, suggestion)
		}
	}

	return unique
}

// ListTypesWithDescriptions returns a formatted list of all types
func (atm *AssignmentTypeManager) ListTypesWithDescriptions() map[string]string {
	result := make(map[string]string)
	for portableType, mapping := range atm.mappings {
		lmsInfo := mapping.LMSType
		if mapping.LMSSubtype != "" {
			lmsInfo += " (" + mapping.LMSSubtype + ")"
		}
		result[portableType] = fmt.Sprintf("%s → %s", mapping.Description, lmsInfo)
	}
	return result
}

// Global type manager instance
var globalTypeManager *AssignmentTypeManager

// GetTypeManager returns the global type manager instance
func GetTypeManager() *AssignmentTypeManager {
	if globalTypeManager == nil {
		globalTypeManager = NewAssignmentTypeManager()
	}
	return globalTypeManager
}
