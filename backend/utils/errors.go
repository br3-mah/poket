package utils

import (
	"strings"
	
	"gorm.io/gorm"
)

// IsDuplicateKeyError checks if the error is a database duplicate key violation
func IsDuplicateKeyError(err error) bool {
	if err == nil {
		return false
	}
	
	// PostgreSQL error
	if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		return true
	}
	
	// MySQL/MariaDB error (if you're using these)
	if strings.Contains(err.Error(), "Error 1062: Duplicate entry") {
		return true
	}
	
	// SQLite error
	if strings.Contains(err.Error(), "UNIQUE constraint failed") {
		return true
	}
	
	// Generic GORM duplicate key error
	if err == gorm.ErrDuplicatedKey {
		return true
	}
	
	return false
}