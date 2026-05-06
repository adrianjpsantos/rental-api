package category

import (
	"regexp"
	"strings"
)

// generateSlug gera um slug a partir do nome
func generateSlug(name string) string {
	slug := strings.ToLower(strings.TrimSpace(name))
	slug = strings.ReplaceAll(slug, " ", "-")
	// Remove caracteres especiais
	slug = regexp.MustCompile(`[^a-z0-9-]`).ReplaceAllString(slug, "")
	// Remove hífens duplicados
	slug = regexp.MustCompile(`-+`).ReplaceAllString(slug, "-")
	return strings.Trim(slug, "-")
}

// isValidSlug verifica se o slug está no formato correto
func isValidSlug(slug string) bool {
	match, _ := regexp.MatchString(`^[a-z0-9]+(?:-[a-z0-9]+)*$`, slug)
	return match
}
