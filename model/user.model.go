package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type UserWithRelations struct {
	User
	Details *Detail `json:"details"`
	Nok     *Nok    `json:"nok"`
}

type User struct {
	ID         uuid.UUID      `json:"id"`
	Mid        string         `json:"mid"`
	FirstName  string         `json:"first_name"`
	MiddleName sql.NullString `json:"middle_name"`
	LastName   string         `json:"last_name"`
	Email      string         `json:"email"`
	Phone      string         `json:"phone"`
	ProfileImg string         `json:"profile_img"`
	Verified   bool           `json:"verified"`
	Suspended  bool           `json:"suspended"`
	Tos        bool           `json:"tos"`
	Role       interface{}    `json:"role"`
	Session    sql.NullString `json:"session"`
	Deleted    bool           `json:"deleted"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  sql.NullTime   `json:"updated_at"`
}

type Detail struct {
	ID            uuid.NullUUID  `json:"id"`
	Mid           string         `json:"mid"`
	Country       sql.NullString `json:"country"`
	City          sql.NullString `json:"city"`
	Address       sql.NullString `json:"address"`
	Gender        sql.NullString `json:"gender"`
	Dob           sql.NullTime   `json:"dob"`
	MaritalStatus sql.NullString `json:"marital_status"`
	State         sql.NullString `json:"state"`
}

type Nok struct {
	ID        uuid.NullUUID `json:"id"`
	Mid       string        `json:"mid"`
	NokMid    string        `json:"nok_mid"`
	UpdatedAt sql.NullTime  `json:"updated_at"`
}
