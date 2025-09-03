package main

import (
	"time"
)

// AssignmentPackage represents a complete portable assignment
type AssignmentPackage struct {
	Metadata     PackageMetadata `json:"metadata" yaml:"metadata"`
	Assignment   Assignment      `json:"assignment" yaml:"assignment"`
	Resources    []Resource      `json:"resources,omitempty" yaml:"resources,omitempty"`
	Dependencies Dependencies    `json:"dependencies,omitempty" yaml:"dependencies,omitempty"`
	Validation   ValidationInfo  `json:"validation,omitempty" yaml:"validation,omitempty"`
}

// PackageMetadata contains package-level information
type PackageMetadata struct {
	ID          string            `json:"id" yaml:"id"`
	Version     string            `json:"version" yaml:"version"`
	Created     time.Time         `json:"created" yaml:"created"`
	Modified    time.Time         `json:"modified" yaml:"modified"`
	Author      string            `json:"author" yaml:"author"`
	Email       string            `json:"email,omitempty" yaml:"email,omitempty"`
	License     string            `json:"license,omitempty" yaml:"license,omitempty"`
	Tags        []string          `json:"tags,omitempty" yaml:"tags,omitempty"`
	Description string            `json:"description,omitempty" yaml:"description,omitempty"`
	Language    string            `json:"language,omitempty" yaml:"language,omitempty"`
	SourceHash  string            `json:"source_hash" yaml:"source_hash"`
	Custom      map[string]string `json:"custom,omitempty" yaml:"custom,omitempty"`
}

// Assignment represents the core assignment data
type Assignment struct {
	// Basic Information
	Title       string `json:"title" yaml:"title"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Type        string `json:"type" yaml:"type"`
	Subtype     string `json:"subtype,omitempty" yaml:"subtype,omitempty"`
	Category    string `json:"category,omitempty" yaml:"category,omitempty"`
	Difficulty  string `json:"difficulty,omitempty" yaml:"difficulty,omitempty"`

	// Content
	Questions           interface{} `json:"questions,omitempty" yaml:"questions,omitempty"`
	Instructions        string      `json:"instructions,omitempty" yaml:"instructions,omitempty"`
	Criteria            string      `json:"criteria,omitempty" yaml:"criteria,omitempty"`
	CodeSubmissionConfig interface{} `json:"code_submission_config,omitempty" yaml:"code_submission_config,omitempty"`

	// Scoring & Behavior
	Points           int  `json:"points" yaml:"points"`
	TimeLimit        *int `json:"time_limit,omitempty" yaml:"time_limit,omitempty"`
	MaxAttempts      *int `json:"max_attempts,omitempty" yaml:"max_attempts,omitempty"`
	AutoGrade        bool `json:"auto_grade" yaml:"auto_grade"`
	ShowFeedback     bool `json:"show_feedback" yaml:"show_feedback"`
	ShuffleQuestions bool `json:"shuffle_questions" yaml:"shuffle_questions"`
	AllowReview      bool `json:"allow_review" yaml:"allow_review"`

	// Scheduling
	DueDate       *time.Time `json:"due_date,omitempty" yaml:"due_date,omitempty"`
	AvailableFrom *time.Time `json:"available_from,omitempty" yaml:"available_from,omitempty"`
	AvailableTo   *time.Time `json:"available_to,omitempty" yaml:"available_to,omitempty"`
	Quarter       string     `json:"quarter,omitempty" yaml:"quarter,omitempty"`

	// Tracking
	TrackAttempts    bool `json:"track_attempts" yaml:"track_attempts"`
	TrackConfidence  bool `json:"track_confidence" yaml:"track_confidence"`
	TrackTimeSpent   bool `json:"track_time_spent" yaml:"track_time_spent"`

	// Educational
	LearningObjectives []string `json:"learning_objectives,omitempty" yaml:"learning_objectives,omitempty"`
	Prerequisites      []string `json:"prerequisites,omitempty" yaml:"prerequisites,omitempty"`
	RecommendedCourses []string `json:"recommended_courses,omitempty" yaml:"recommended_courses,omitempty"`
	Tags               []string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Publishing
	Published bool `json:"published" yaml:"published"`
}

// Resource represents a learning resource attached to an assignment
type Resource struct {
	ID          string            `json:"id" yaml:"id"`
	Title       string            `json:"title" yaml:"title"`
	Description string            `json:"description,omitempty" yaml:"description,omitempty"`
	Type        string            `json:"type" yaml:"type"`
	LocalPath   string            `json:"local_path,omitempty" yaml:"local_path,omitempty"`
	URL         string            `json:"url,omitempty" yaml:"url,omitempty"`
	FileSize    int64             `json:"file_size,omitempty" yaml:"file_size,omitempty"`
	MimeType    string            `json:"mime_type,omitempty" yaml:"mime_type,omitempty"`
	Checksum    string            `json:"checksum,omitempty" yaml:"checksum,omitempty"`
	Tags        []string          `json:"tags,omitempty" yaml:"tags,omitempty"`
	Order       int               `json:"order,omitempty" yaml:"order,omitempty"`
	IsPublic    bool              `json:"is_public" yaml:"is_public"`
	Metadata    map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`
}

// Dependencies represents assignment dependencies and relationships
type Dependencies struct {
	Subjects           []string `json:"subjects,omitempty" yaml:"subjects,omitempty"`
	Prerequisites      []string `json:"prerequisites,omitempty" yaml:"prerequisites,omitempty"`
	RecommendedCourses []string `json:"recommended_courses,omitempty" yaml:"recommended_courses,omitempty"`
	RequiredResources  []string `json:"required_resources,omitempty" yaml:"required_resources,omitempty"`
	SoftwareRequirements []SoftwareRequirement `json:"software_requirements,omitempty" yaml:"software_requirements,omitempty"`
}

// SoftwareRequirement represents required software/tools
type SoftwareRequirement struct {
	Name        string `json:"name" yaml:"name"`
	Version     string `json:"version,omitempty" yaml:"version,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Required    bool   `json:"required" yaml:"required"`
}

// ValidationInfo contains validation results and metadata
type ValidationInfo struct {
	IsValid       bool      `json:"is_valid" yaml:"is_valid"`
	ValidatedAt   time.Time `json:"validated_at" yaml:"validated_at"`
	ValidatorVersion string `json:"validator_version" yaml:"validator_version"`
	Errors        []string  `json:"errors,omitempty" yaml:"errors,omitempty"`
	Warnings      []string  `json:"warnings,omitempty" yaml:"warnings,omitempty"`
	Score         int       `json:"score,omitempty" yaml:"score,omitempty"` // Quality score 0-100
}

// ImportResult represents the result of importing assignments to LMS
type ImportResult struct {
	AssignmentID string            `json:"assignment_id,omitempty"`
	ResourceIDs  []string          `json:"resource_ids,omitempty"`
	Conflicts    []string          `json:"conflicts,omitempty"`
	Status       string            `json:"status"`
	Message      string            `json:"message,omitempty"`
	Metadata     map[string]string `json:"metadata,omitempty"`
}

// BatchImportResult represents results from batch import
type BatchImportResult struct {
	BatchID      string         `json:"batch_id"`
	TotalCount   int            `json:"total_count"`
	SuccessCount int            `json:"success_count"`
	FailureCount int            `json:"failure_count"`
	Results      []ImportResult `json:"results"`
	StartedAt    time.Time      `json:"started_at"`
	CompletedAt  time.Time      `json:"completed_at"`
}

// Config represents the toolkit configuration
type Config struct {
	Author      string            `json:"author" yaml:"author"`
	Email       string            `json:"email" yaml:"email"`
	License     string            `json:"license" yaml:"license"`
	Language    string            `json:"language" yaml:"language"`
	LMSEndpoint string            `json:"lms_endpoint" yaml:"lms_endpoint"`
	APIKey      string            `json:"api_key,omitempty" yaml:"api_key,omitempty"`
	Templates   map[string]string `json:"templates" yaml:"templates"`
	Defaults    map[string]string `json:"defaults" yaml:"defaults"`
}

// Template represents an assignment template
type Template struct {
	Name        string            `json:"name" yaml:"name"`
	Description string            `json:"description" yaml:"description"`
	Type        string            `json:"type" yaml:"type"`
	Category    string            `json:"category,omitempty" yaml:"category,omitempty"`
	Template    Assignment        `json:"template" yaml:"template"`
	Fields      []TemplateField   `json:"fields,omitempty" yaml:"fields,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`
}

// TemplateField represents configurable fields in templates
type TemplateField struct {
	Name        string      `json:"name" yaml:"name"`
	Type        string      `json:"type" yaml:"type"` // string, int, bool, select, multiselect
	Label       string      `json:"label" yaml:"label"`
	Description string      `json:"description,omitempty" yaml:"description,omitempty"`
	Required    bool        `json:"required" yaml:"required"`
	Default     interface{} `json:"default,omitempty" yaml:"default,omitempty"`
	Options     []string    `json:"options,omitempty" yaml:"options,omitempty"`
	Validation  string      `json:"validation,omitempty" yaml:"validation,omitempty"`
}
