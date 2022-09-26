package base

// BaseModel is the base model for all data models.
type BaseModel interface {
	Validate() error
}

// ValidateAll validates all models.
func ValidateAll(models ...BaseModel) error {
	for _, model := range models {
		if err := model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// BaseModelPtr returns a pointer to the base model.
func BaseModelPtr(base BaseModel) *BaseModel {
	return &base
}
