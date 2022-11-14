package model

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type UserWithRelations struct {
	User
	Details *Detail `json:"details"`
	Nok     *Nok    `json:"nok"`
}

type User struct {
	ID         uuid.UUID   `json:"id"`
	Mid        string      `json:"mid"`
	FirstName  string      `json:"first_name"`
	MiddleName null.String `json:"middle_name"`
	LastName   string      `json:"last_name"`
	Email      string      `json:"email"`
	Phone      string      `json:"phone"`
	ProfileImg string      `json:"profile_img"`
	Verified   bool        `json:"verified"`
	Suspended  bool        `json:"suspended"`
	Tos        bool        `json:"tos"`
	Role       interface{} `json:"role"`
	Session    null.String `json:"session"`
	Balance    float64     `json:"balance"`
	Deleted    bool        `json:"deleted"`
	CreatedAt  time.Time   `json:"created_at"`
	Developer  bool        `json:"developer"`
	UpdatedAt  null.Time   `json:"updated_at"`
}

type Detail struct {
	ID            uuid.NullUUID `json:"id"`
	Mid           string        `json:"mid"`
	Country       null.String   `json:"country"`
	City          null.String   `json:"city"`
	Address       null.String   `json:"address"`
	Gender        null.String   `json:"gender"`
	Dob           null.Time     `json:"dob"`
	MaritalStatus null.String   `json:"marital_status"`
	State         null.String   `json:"state"`
}

type Nok struct {
	ID        uuid.NullUUID `json:"id"`
	Mid       string        `json:"mid"`
	NokMid    string        `json:"nok_mid"`
	UpdatedAt null.Time     `json:"updated_at"`
}
