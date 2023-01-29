package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Specialist struct {
	ID                    uuid.UUID      `json:"id"`
	Mid                   string         `json:"mid"`
	Bio                   sql.NullString `json:"bio"`
	Specialisation        string         `json:"specialisation"`
	DaysAvailable         []string       `json:"days_available"`
	AvailabilityStartTime sql.NullTime   `json:"availability_start_time"`
	AvailabilityEndTime   sql.NullTime   `json:"availability_end_time"`
	CertificationImg      string         `json:"certification_img"`
	CertificationImgID    string         `json:"certification_img_id"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             null.Time      `json:"updated_at"`
}
