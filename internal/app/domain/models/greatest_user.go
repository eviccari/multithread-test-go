package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type GreatestUser struct {
	ID            string
	LegacyLogin   string
	LegacyID      int
	LegacyNodeID  string
	LegacyURL     string
	LegacyHTMLURL string
	NewEmail      string
	CreatedAt     time.Time
}

func (gu *GreatestUser) SetAsNew() {
	gu.ID = uuid.New().String()
	gu.CreatedAt = time.Now()
	gu.NewEmail = fmt.Sprintf("%s@greatestuser.com", gu.ID)
}

func (gu GreatestUser) Validate() (errorsList []error) {
	pm := NewGreatestUserPoliciesManager(gu)
	return pm.Apply()
}
