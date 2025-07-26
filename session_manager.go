package main

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/mark3labs/mcp-go/server"
)

// CustomSessionIdManager is a session ID manager that accepts both UUID and numeric formats
type CustomSessionIdManager struct{}

const idPrefix = "mcp-session-"

func (s *CustomSessionIdManager) Generate() string {
	return idPrefix + uuid.New().String()
}

func (s *CustomSessionIdManager) Validate(sessionID string) (isTerminated bool, err error) {
	// Check if the session ID has the required prefix
	if !strings.HasPrefix(sessionID, idPrefix) {
		return false, fmt.Errorf("invalid session id: %s", sessionID)
	}

	// Extract the part after the prefix
	idPart := sessionID[len(idPrefix):]

	// Try to parse as UUID (the original format)
	_, uuidErr := uuid.Parse(idPart)
	if uuidErr == nil {
		// It's a valid UUID format
		return false, nil
	}

	// If it's not a UUID, check if it's a numeric format
	for _, c := range idPart {
		if c < '0' || c > '9' {
			// Not a numeric format
			return false, fmt.Errorf("invalid session id: %s", sessionID)
		}
	}

	// It's a valid numeric format
	return false, nil
}

func (s *CustomSessionIdManager) Terminate(sessionID string) (isNotAllowed bool, err error) {
	return false, nil
}

// Ensure CustomSessionIdManager implements the SessionIdManager interface
var _ server.SessionIdManager = (*CustomSessionIdManager)(nil)
