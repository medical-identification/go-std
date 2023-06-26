package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Message struct {
	ID                uuid.UUID      `json:"id"`
	Mid               null.String    `json:"mid"`
	FacilityID        uuid.NullUUID  `json:"facility_id"`
	FacilityAuthor    uuid.NullUUID  `json:"facility_author"`
	Specialist        null.String    `json:"specialist"`
	Message           sql.NullString `json:"message"`
	Type              null.String    `json:"type"`
	Support           null.Bool      `json:"support"`
	SupportAuthor     uuid.NullUUID  `json:"support_author"`
	Recieved          bool           `json:"recieved"`
	Read              bool           `json:"read"`
	SecondaryRecieved bool           `json:"secondary_recieved"`
	SecondaryRead     bool           `json:"secondary_read"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         null.Time      `json:"updated_at"`
}

type Notification struct {
	ID          uuid.UUID   `json:"id"`
	Mid         null.String `json:"mid"`
	Type        string      `json:"type"`
	Content     string      `json:"content"`
	Title       null.String `json:"title"`
	Received    bool        `json:"received"`
	Read        bool        `json:"read"`
	FacilityID  null.String `json:"facility_id"`
	SecondaryID null.String `json:"secondary_id"`
	CreatedAt   time.Time   `json:"created_at"`
}

type Billing struct {
	ID                    uuid.UUID     `json:"id"`
	SessionID             uuid.NullUUID `json:"session_id"`
	Status                interface{}   `json:"status"`
	Amount                float64       `json:"amount"`
	BillingSource         string        `json:"billing_source"`
	BillingSourceType     string        `json:"billing_source_type"`
	BilledRecipientSource string        `json:"billed_recipient_source"`
	BilledRecipientType   string        `json:"billed_recipient_type"`
	Author                null.String   `json:"author"`
	CreatedAt             time.Time     `json:"created_at"`
	BranchID              uuid.NullUUID `json:"branch_id"`
	Discount              bool          `json:"discount"`
	DiscountPercentage    int32         `json:"discount_percentage"`
	TransactionType       interface{}   `json:"transaction_type"`
	UpdatedAt             null.Time     `json:"updated_at"`
}

type ExpenseHistory struct {
	ID         uuid.UUID     `json:"id"`
	Reference  string        `json:"reference"`
	Amount     float64       `json:"amount"`
	Type       interface{}   `json:"type"`
	Status     interface{}   `json:"status"`
	Comment    null.String   `json:"comment"`
	From       null.String   `json:"from"`
	FromSource string        `json:"from_source"`
	FromTarget null.String   `json:"from_target"`
	SourceID   string        `json:"source_id"`
	SourceType string        `json:"source_type"`
	CreatedAt  time.Time     `json:"created_at"`
	UpdatedAt  null.Time     `json:"updated_at"`
	BranchID   uuid.NullUUID `json:"branch_id"`
	Author     null.String   `json:"author"`
}
