package dtos

import "time"

type GreatestUserDTO struct {
	ID            string    `json:"id"`
	LegacyLogin   string    `json:"legacy_login"`
	LegacyID      int       `json:"legacy_id"`
	LegacyNodeID  string    `json:"legacy_node_id"`
	LegacyURL     string    `json:"legacy_url"`
	LegacyHTMLURL string    `json:"legacy_html_url"`
	NewEmail      string    `json:"new_email"`
	CreatedAt     time.Time `json:"created_at"`
}
