package themes

import (
	"time"
)

// ProviderName type.
type ProviderName string

// Theme represents metadata about a theme.
type Theme struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Hash        string `json:"hash"`
	Light       bool   `json:"light"`
	Version     string `json:"version,omitempty"`

	ProjectRepoID string       `json:"projectRepoId"`
	ProjectRepo   string       `json:"projectRepo"`
	Readme        string       `json:"readme"`
	License       string       `json:"license,omitempty"`
	Provider      ProviderName `json:"provider,omitempty"`
	LastUpdate    time.Time    `json:"updatedAt,omitempty"`
}

// Gallery represents a collection of themes.
type Gallery []Theme

// Provider seeks theme metadata from a source.
type Provider interface {
	GetGallery() (Gallery, error)
}