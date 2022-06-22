package mgm

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// IDField struct contains a model's ID field.
type IDField struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
}

// DateFields struct contains the `created_at` and `updated_at`
// fields that autofill when inserting or updating a model.
type DateFields struct {
	CreatedAt int64 `json:"create_time" bson:"create_time"`
	UpdatedAt int64 `json:"last_modify_time" bson:"last_modify_time"`
}

// PrepareID method prepares the ID value to be used for filtering
// e.g convert hex-string ID value to bson.ObjectId
func (f *IDField) PrepareID(id interface{}) (interface{}, error) {
	if idStr, ok := id.(string); ok {
		return primitive.ObjectIDFromHex(idStr)
	}

	// Otherwise id must be ObjectId
	return id, nil
}

// GetID method returns a model's ID
func (f *IDField) GetID() interface{} {
	return f.ID
}

// SetID sets the value of a model's ID field.
func (f *IDField) SetID(id interface{}) {
	f.ID = id.(primitive.ObjectID)
}

//--------------------------------
// DateField methods
//--------------------------------

// Creating hook is used here to set the `created_at` field
// value when inserting a new model into the database.
// TODO: get context as param the next version.
func (f *DateFields) Creating() error {
	f.CreatedAt = time.Now().UnixMilli()
	return nil
}

// Saving hook is used here to set the `updated_at` field
// value when creating or updating a model.
// TODO: get context as param the next version.
func (f *DateFields) Saving() error {
	f.UpdatedAt = time.Now().UnixMilli()
	return nil
}
