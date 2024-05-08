package model

// RecordID defines a record id, Together with RecordType
// identifies unique records across all types.

type RecordID string

// RecordType defines a record type. Together with RecordID
// identifies unique records across all types.
type RecordType string

// RecordTypeMovie Existing record types.
const (
	RecordTypeMovie = RecordType("movie")
)

// UserID defines a user id.
type UserID string
type (
	// RatingValue Rating Value defines a value of a rating record.
	RatingValue int
)

type Rating struct {
	RecordID   string      `json:"recordId"`
	RecordType string      `json:"recordType"`
	UserID     UserID      `json:"userId"`
	Value      RatingValue `json:"value"`
}
