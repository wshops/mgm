package mgm

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// CollectionGetter interface contains a method to return
// a model's custom collection.
type CollectionGetter interface {
	// Collection method return collection
	Collection() *Collection
}

// CollectionNameGetter interface contains a method to return
// the collection name of a model.
type CollectionNameGetter interface {
	// CollectionName method return model collection's name.
	CollectionName() string
}

// Model interface contains base methods that must be implemented by
// each model. If you're using the `DefaultModel` struct in your model,
// you don't need to implement any of these methods.
type Model interface {
	// PrepareID converts the id value if needed, then
	// returns it (e.g convert string to objectId).
	PrepareID(id any) (any, error)
	GetID() primitive.ObjectID
	SetID(id any)
	GetIdStr() string
}

//go:generate msgp
// DefaultModel struct contains a model's default fields.
type DefaultModel struct {
	ObjectId  primitive.ObjectID `json:"-" bson:"_id,omitempty" msg:"-"`
	Id        string             `json:"id" bson:"-" msg:"id"`
	CreatedAt int64              `json:"create_time" bson:"create_time" msg:"create_time"`
	UpdatedAt int64              `json:"last_modify_time" bson:"last_modify_time" msg:"last_modify_time"`
}

// Creating function calls the inner fields' defined hooks
func (model *DefaultModel) Creating() error {
	model.CreatedAt = time.Now().UnixMilli()
	return nil
}

// Saving function calls the inner fields' defined hooks
func (model *DefaultModel) Saving() error {
	model.UpdatedAt = time.Now().UnixMilli()
	return nil
}

// PrepareID method prepares the ID value to be used for filtering
// e.g convert hex-string ID value to bson.ObjectId
func (model *DefaultModel) PrepareID(id any) (any, error) {

	if idStr, ok := id.(string); ok {
		return primitive.ObjectIDFromHex(idStr)
	}

	// Otherwise id must be ObjectId
	return id, nil
}

// GetID method returns a model's ID
func (model *DefaultModel) GetID() primitive.ObjectID {
	return model.ObjectId
}

// GetIdStr method returns a model's ID
func (model *DefaultModel) GetIdStr() string {
	return model.Id
}

// SetID sets the value of a model's ID field.
func (model *DefaultModel) SetID(id any) {
	model.ObjectId = id.(primitive.ObjectID)
	model.Id = id.(primitive.ObjectID).Hex()
}
