package constants

import (
	"go-minimal/internal/config"
	"strconv"
)

// We use variables (var) instead of constants (const)
// because these are loaded at runtime.
var (
	AdminID   int
	ManagerID int
	UserID    int
)

// LoadRoles initializes the IDs from the environment.
// Call this once in your main.go after loading the config.
func LoadRoles() {
	AdminID = parseID(config.GetAdminID())
	ManagerID = parseID(config.GetManagerID())
	UserID = parseID(config.GetUserID())
}

func parseID(s string) int {
	id, err := strconv.Atoi(s)
	if err != nil {
		// If the env variable is missing or wrong, we want to know immediately
		panic("invalid role ID in environment: " + s)
	}
	return id
}
