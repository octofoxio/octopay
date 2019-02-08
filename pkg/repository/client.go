package repository

import "github.com/jinzhu/gorm"

type ClientApplicationInformation struct {
	Name        string
	ID          string `gorm:"unique,primary_key"`
	Secret      string
	CallbackURL string
}

type ValidateCredentialInput struct {
	ID     string
	Secret string
}
type ValidateCredentialOutput struct {
	ID   string
	Name string
}
type RegisterCallbackURLInput struct {
	ID  string
	Url string
}
type ClientApplicationInformationRepository interface {
	ValidateCredential(input *ValidateCredentialInput) (*ValidateCredentialOutput, error)
	RegisterCallbackURL(input *RegisterCallbackURLInput) error
}

type DefaultClientApplicationInformationRepository struct {
	DB *gorm.DB
}

func (d *DefaultClientApplicationInformationRepository) RegisterCallbackURL(input *RegisterCallbackURLInput) error {
	return d.DB.Model(&ClientApplicationInformation{}).Update(&ClientApplicationInformation{
		ID:          input.ID,
		CallbackURL: input.Url,
	}).Error
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
