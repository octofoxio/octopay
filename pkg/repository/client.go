package repository

import "github.com/jinzhu/gorm"

type ClientApplicationInformation struct {
	Name   string
	ID     string `gorm:"unique,primary_key"`
	Secret string
}

type ValidateCredentialInput struct {
	ID     string
	Secret string
}
type ValidateCredentialOutput struct {
	ID   string
	Name string
}
type ClientApplicationInformationRepository interface {
	ValidateCredential(input *ValidateCredentialInput) (*ValidateCredentialOutput, error)
}

type DefaultClientApplicationInformationRepository struct {
	DB *gorm.DB
}

func (d *DefaultClientApplicationInformationRepository) ValidateCredential(input *ValidateCredentialInput) (*ValidateCredentialOutput, error) {
	var result ClientApplicationInformation
	st := d.DB.Model(&ClientApplicationInformation{}).
		Find(&result, &ClientApplicationInformation{
			Secret: input.Secret,
			ID:     input.ID,
		})
	return &ValidateCredentialOutput{
		ID:   result.ID,
		Name: result.Name,
	}, st.Error
}
