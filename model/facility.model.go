package model

import (
	"time"

	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
)

type Facility struct {
	ID            uuid.UUID   `json:"id"`
	Name          string      `json:"name"`
	Description   null.String `json:"description"`
	FacilityType  interface{} `json:"facility_type"`
	FacilityImg   string      `json:"facility_img"`
	FacilityImgID string      `json:"facility_img_id"`
	SerialNo      string      `json:"serial_no"`
	CoverImg      string      `json:"cover_img"`
	CoverImgID    string      `json:"cover_img_id"`
	Balance       float64     `json:"balance"`
	Deleted       bool        `json:"deleted"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     null.Time   `json:"updated_at"`
}

type FacilityBranch struct {
	ID         uuid.UUID `json:"id"`
	FacilityID uuid.UUID `json:"facility_id"`
	City       string    `json:"city"`
	State      string    `json:"state"`
	Country    string    `json:"country"`
	Address    string    `json:"address"`
	Headquater bool      `json:"headquater"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  null.Time `json:"updated_at"`
}

type FacilityMetum struct {
	ID          uuid.UUID `json:"id"`
	FacilityID  uuid.UUID `json:"facility_id"`
	Services    []string  `json:"services"`
	ApFee       float64   `json:"ap_fee"`
	RegFee      float64   `json:"reg_fee"`
	BaselineFee float64   `json:"baseline_fee"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   null.Time `json:"updated_at"`
}

type FacilityStaff struct {
	ID                    uuid.UUID   `json:"id"`
	Mid                   string      `json:"mid"`
	FacilityID            uuid.UUID   `json:"facility_id"`
	BranchID              uuid.UUID   `json:"branch_id"`
	Specializations       []string    `json:"specializations"`
	Role                  string      `json:"role"`
	Comment               null.String `json:"comment"`
	InventoryManagement   string      `json:"inventory_management"`
	StaffsManagement      string      `json:"staffs_management"`
	SessionManagement     string      `json:"session_management"`
	CanCreateSession      string      `json:"can_create_session"`
	TestManagement        string      `json:"test_management"`
	DepartmentsManagement string      `json:"departments_management"`
	UnitManagement        string      `json:"unit_management"`
	EncounterManagement   string      `json:"encounter_management"`
	AppointmentManagement string      `json:"appointment_management"`
	ExpenseTracker        string      `json:"expense_tracker"`
	Settings              string      `json:"settings"`
	CanBill               string      `json:"can_bill"`
	ViewExpenseTracker    string      `json:"view_expense_tracker"`
	ViewExpenseHistory    string      `json:"view_expense_history"`
	CreatedAt             time.Time   `json:"created_at"`
	UpdatedAt             null.Time   `json:"updated_at"`
}
