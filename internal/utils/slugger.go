package utils

import (
	"fmt"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"regexp"
	"strings"
)

var (
	reSpecials = regexp.MustCompile(`[^a-z0-9\s-]`)
	reSpaces   = regexp.MustCompile(`\s+`)
	reHyphens  = regexp.MustCompile(`-+`)
)

// GenerateSlug creates a basic clean slug
func GenerateSlug(input string) string {
	slug := strings.ToLower(input)
	slug = reSpecials.ReplaceAllString(slug, "")
	slug = reSpaces.ReplaceAllString(slug, "-")
	slug = reHyphens.ReplaceAllString(slug, "-")
	return strings.Trim(slug, "-")
}

// GenerateUniqueSlug creates a slug and appends a short, random NanoID
func GenerateUniqueSlug(input string) (string, error) {
	base := GenerateSlug(input)

	// Generate a 6-character random ID (alphanumeric)
	// We use 6 chars to keep the URL short but unique enough for SaaS
	id, err := gonanoid.New(6)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s-%s", base, id), nil
}
