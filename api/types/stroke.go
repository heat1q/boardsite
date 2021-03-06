package types

import (
	"encoding/json"
)

// Style declares the stoke style.
type Style struct {
	Color   string  `json:"color"`
	Width   float64 `json:"width"`
	Opacity float64 `json:"opacity"`
}

// Stroke declares the structure of most stoke types.
type Stroke struct {
	Type   int       `json:"type"`
	ID     string    `json:"id,omitempty"`
	PageID string    `json:"pageId,omitempty"`
	UserID string    `json:"userId"`
	X      float64   `json:"x"`
	Y      float64   `json:"y"`
	ScaleX float64   `json:"scaleX,omitempty"`
	ScaleY float64   `json:"scaleY,omitempty"`
	Points []float64 `json:"points,omitempty"`
	Style  Style     `json:"style,omitempty"`
}

// StrokeReader defines the set of common function
// to interact with strokes
type StrokeReader interface {
	JSONStringify() ([]byte, error)
	IsDeleted() bool
	GetID() string
	GetUserID() string
	GetPageID() string
}

// JSONStringify return the JSON encoding of Stroke
func (s *Stroke) JSONStringify() ([]byte, error) {
	return json.Marshal(s)
}

// IsDeleted verifies whether stroke is deleted or not
func (s *Stroke) IsDeleted() bool {
	return s.Type == 0
}

// GetID returns the id of the stroke
func (s *Stroke) GetID() string {
	return s.ID
}

// GetUserID returns the userid of the stroke
func (s *Stroke) GetUserID() string {
	return s.UserID
}

// GetPageID returns the page id of the stroke
func (s *Stroke) GetPageID() string {
	return s.PageID
}
