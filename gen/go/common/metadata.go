package common

import "github.com/google/uuid"

// Create creates a new Metadata object with a new UUID and generation 1.
func (x *MetadataFields) Create() error {
	if x == nil {
		return nil
	}
	if len(x.GetManagedFields().GetFields()) > 0 {
		return ErrManagedFieldsMustBeEmpty
	}
	new_id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	x.SetId(new_id.String())
	x.SetGeneration(1)
	return nil
}

// IncGeneration increments the generation of the object.
func (x *MetadataFields) IncGeneration() {
	x.SetGeneration(x.GetGeneration() + 1)
}

// Id returns the UUID of the object or nil if the object is nil or the ID is invalid.
func (x *MetadataFields) Id() uuid.UUID {
	if x == nil {
		return uuid.Nil
	}
	id, err := uuid.Parse(x.GetId())
	if err != nil {
		return uuid.Nil
	}
	return id
}
