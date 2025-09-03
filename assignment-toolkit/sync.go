package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

// LMSClient handles communication with the LMS API
type LMSClient struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

// NewLMSClient creates a new LMS client
func NewLMSClient(baseURL, apiKey string) *LMSClient {
	return &LMSClient{
		BaseURL:    strings.TrimSuffix(baseURL, "/"),
		APIKey:     apiKey,
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
	}
}

// SyncAssignment uploads an assignment to the LMS
func (c *LMSClient) SyncAssignment(pkg AssignmentPackage) (*ImportResult, error) {
	// Convert assignment to LMS format
	lmsAssignment := convertToLMSFormat(pkg)

	// Create JSON payload
	jsonData, err := json.Marshal(lmsAssignment)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal assignment: %v", err)
	}

	// Create HTTP request
	url := fmt.Sprintf("%s/api/assignments", c.BaseURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))

	// Send request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("API error (%d): %s", resp.StatusCode, string(body))
	}

	// Parse response
	var response struct {
		Assignment struct {
			ID string `json:"id"`
		} `json:"assignment"`
		Message string `json:"message"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	result := &ImportResult{
		AssignmentID: response.Assignment.ID,
		Status:       "success",
		Message:      response.Message,
	}

	// Upload resources if any
	if len(pkg.Resources) > 0 {
		resourceIDs, err := c.uploadResources(response.Assignment.ID, pkg.Resources)
		if err != nil {
			result.Status = "partial"
			result.Message += fmt.Sprintf(" Warning: Resource upload failed: %v", err)
		} else {
			result.ResourceIDs = resourceIDs
		}
	}

	return result, nil
}

// BatchSyncAssignments uploads multiple assignments
func (c *LMSClient) BatchSyncAssignments(packages []AssignmentPackage) (*BatchImportResult, error) {
	result := &BatchImportResult{
		BatchID:      uuid.New().String(),
		TotalCount:   len(packages),
		SuccessCount: 0,
		FailureCount: 0,
		Results:      make([]ImportResult, 0, len(packages)),
		StartedAt:    time.Now(),
	}

	for _, pkg := range packages {
		importResult, err := c.SyncAssignment(pkg)
		if err != nil {
			result.FailureCount++
			result.Results = append(result.Results, ImportResult{
				Status:  "failed",
				Message: err.Error(),
			})
		} else {
			result.SuccessCount++
			result.Results = append(result.Results, *importResult)
		}
	}

	result.CompletedAt = time.Now()
	return result, nil
}

// uploadResources uploads resource files to the LMS
func (c *LMSClient) uploadResources(assignmentID string, resources []Resource) ([]string, error) {
	var resourceIDs []string

	for _, resource := range resources {
		if resource.LocalPath == "" {
			continue // Skip resources without local files
		}

		resourceID, err := c.uploadResource(assignmentID, resource)
		if err != nil {
			return resourceIDs, fmt.Errorf("failed to upload %s: %v", resource.Title, err)
		}
		resourceIDs = append(resourceIDs, resourceID)
	}

	return resourceIDs, nil
}

// uploadResource uploads a single resource file
func (c *LMSClient) uploadResource(assignmentID string, resource Resource) (string, error) {
	// Open file
	file, err := os.Open(resource.LocalPath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Create multipart form
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Add file field
	part, err := writer.CreateFormFile("file", filepath.Base(resource.LocalPath))
	if err != nil {
		return "", fmt.Errorf("failed to create form file: %v", err)
	}

	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	part.Write(fileContent)

	// Add metadata fields
	writer.WriteField("title", resource.Title)
	writer.WriteField("description", resource.Description)
	writer.WriteField("type", resource.Type)
	writer.WriteField("assignmentId", assignmentID)

	writer.Close()

	// Create request
	url := fmt.Sprintf("%s/api/resources", c.BaseURL)
	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))

	// Send request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("API error (%d): %s", resp.StatusCode, string(body))
	}

	// Parse response
	var response struct {
		Resource struct {
			ID string `json:"id"`
		} `json:"resource"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("failed to parse response: %v", err)
	}

	return response.Resource.ID, nil
}

// convertToLMSFormat converts our assignment format to LMS API format
func convertToLMSFormat(pkg AssignmentPackage) map[string]interface{} {
	assignment := pkg.Assignment

	// Resolve portable type to LMS type
	typeManager := GetTypeManager()
	lmsType, lmsSubtype, err := typeManager.ConvertToLMSFormat(assignment.Type)
	if err != nil {
		// Fallback to original type if mapping fails
		lmsType = assignment.Type
		lmsSubtype = assignment.Subtype
	}

	lmsAssignment := map[string]interface{}{
		"title":                assignment.Title,
		"description":          assignment.Description,
		"type":                 lmsType,
		"subtype":              lmsSubtype,
		"category":             assignment.Category,
		"difficulty":           assignment.Difficulty,
		"points":               assignment.Points,
		"instructions":         assignment.Instructions,
		"criteria":             assignment.Criteria,
		"autoGrade":            assignment.AutoGrade,
		"showFeedback":         assignment.ShowFeedback,
		"shuffleQuestions":     assignment.ShuffleQuestions,
		"allowReview":          assignment.AllowReview,
		"published":            assignment.Published,
		"quarter":              assignment.Quarter,
		"trackAttempts":        assignment.TrackAttempts,
		"trackConfidence":      assignment.TrackConfidence,
		"trackTimeSpent":       assignment.TrackTimeSpent,
		"learningObjectives":   assignment.LearningObjectives,
		"prerequisites":        assignment.Prerequisites,
		"recommendedCourses":   assignment.RecommendedCourses,
		"tags":                 assignment.Tags,
		"questions":            assignment.Questions,
		"codeSubmissionConfig": assignment.CodeSubmissionConfig,

		// Portable assignment metadata
		"templateId":   pkg.Metadata.ID,
		"version":      pkg.Metadata.Version,
		"sourceHash":   pkg.Metadata.SourceHash,
		"importedFrom": "assignment-toolkit",
		"importedAt":   time.Now(),
	}

	// Handle time fields
	if assignment.DueDate != nil {
		lmsAssignment["dueDate"] = assignment.DueDate.Format(time.RFC3339)
	}
	if assignment.AvailableFrom != nil {
		lmsAssignment["availableFrom"] = assignment.AvailableFrom.Format(time.RFC3339)
	}
	if assignment.AvailableTo != nil {
		lmsAssignment["availableTo"] = assignment.AvailableTo.Format(time.RFC3339)
	}

	// Handle optional integer fields
	if assignment.TimeLimit != nil {
		lmsAssignment["timeLimit"] = *assignment.TimeLimit
	}
	if assignment.MaxAttempts != nil {
		lmsAssignment["maxAttempts"] = *assignment.MaxAttempts
	}

	return lmsAssignment
}

// TestConnection tests the connection to the LMS
func (c *LMSClient) TestConnection() error {
	url := fmt.Sprintf("%s/api/auth/me", c.BaseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to connect to LMS: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return fmt.Errorf("authentication failed - check your API key")
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("LMS returned error (%d): %s", resp.StatusCode, string(body))
	}

	return nil
}

// GetAssignmentByHash checks if an assignment with the given hash already exists
func (c *LMSClient) GetAssignmentByHash(hash string) (*ImportResult, error) {
	url := fmt.Sprintf("%s/api/assignments?sourceHash=%s", c.BaseURL, hash)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil // Assignment doesn't exist
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error (%d): %s", resp.StatusCode, string(body))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	var response struct {
		Assignment struct {
			ID string `json:"id"`
		} `json:"assignment"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	return &ImportResult{
		AssignmentID: response.Assignment.ID,
		Status:       "exists",
		Message:      "Assignment already exists",
	}, nil
}
