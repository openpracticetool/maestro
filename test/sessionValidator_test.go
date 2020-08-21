package test

import (
	"testing"
	"time"

	"github.com/go-playground/validator"
	"github.com/openpracticetool/maestro/model"
	"github.com/openpracticetool/maestro/validators"
)

//TestSessionValidator :: test if model pass in all validations
func TestSessionValidator(t *testing.T) {

	var model = model.SessionModel{
		ID:          123456,
		IDWorkspace: 123456789,
		Description: "Este é um exmplo de teste para criação de uma seção",
		Name:        "Lean Coffee table",
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
		UpdatedBY:   "lhsribas",
		CreatedBy:   "lhsribas",
	}

	err := validators.ValidateStruct(model)

	if err != nil {
		t.Errorf("Error ::: %s", err)
	}
}

//TestSessionDescriptionGreatherThan255Chars :: test with send greather than 255 characteres to return a message error
func TestSessionDescriptionGreatherThan255Chars(t *testing.T) {
	var model = model.SessionModel{
		ID:          123456,
		IDWorkspace: 123456789,
		Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Vitae congue eu consequat ac felis donec et. Massa massa ultricies mi quis hendrerit dolor. Rhoncus aenean vel elit scelerisque. Imperdiet nulla malesuada pellentesque elit. Duis at tellus at urna condimentum mattis pellentesque id nibh. Elementum facilisis leo vel fringilla est ullamcorper eget nulla facilisi. Ornare arcu dui vivamus arcu. Nulla pellentesque dignissim enim sit amet venenatis. Nam libero justo laoreet sit amet cursus sit amet dictum. Sit amet aliquam id diam maecenas. Metus aliquam eleifend mi in nulla. Pellentesque elit eget gravida cum sociis. Risus pretium quam vulputate dignissim suspendisse in est ante.",
		Name:        "Lean Coffee table",
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
		UpdatedBY:   "lhsribas",
		CreatedBy:   "lhsribas",
	}

	if err := validators.ValidateStruct(model); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Description" {
				return
			} else {
				t.Error("The test dont pass in the requirements")
			}
		}
	}
}

//TestSessionDescriptionLessThan30Chars :: test with send less than 50 characteres to return a message error
func TestSessionDescriptionLessThan30Chars(t *testing.T) {
	var model = model.SessionModel{
		ID:          123456,
		IDWorkspace: 123456789,
		Description: "Este é um exemplo",
		Name:        "Lean Coffee table",
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
		UpdatedBY:   "lhsribas",
		CreatedBy:   "lhsribas",
	}

	if err := validators.ValidateStruct(model); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Description" {
				return
			} else {
				t.Error("The test dont pass in the requirements")
			}
		}
	}
}

//TestSessionNameGreatherThan50Chars :: test with send greather than 50 characteres to return a message error
func TestSessionNameGreatherThan50Chars(t *testing.T) {
	var model = model.SessionModel{
		ID:          123456,
		IDWorkspace: 123456789,
		Description: "Este é um exmplo de teste para criação de uma seção",
		Name:        "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
		UpdatedBY:   "lhsribas",
		CreatedBy:   "lhsribas",
	}

	if err := validators.ValidateStruct(model); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Name" {
				return
			} else {
				t.Error("The test dont pass in the requirements")
			}
		}
	}
}

//TestSessionNameLessThan10Chars :: test with send less than 10 characteres to return a message error
func TestSessionNameLessThan10Chars(t *testing.T) {
	var model = model.SessionModel{
		ID:          123456,
		IDWorkspace: 123456789,
		Description: "Este é um exmplo de teste para criação de uma seção",
		Name:        "Lorem",
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
		UpdatedBY:   "lhsribas",
		CreatedBy:   "lhsribas",
	}

	if err := validators.ValidateStruct(model); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if err.Field() == "Name" {
				return
			} else {
				t.Error("The test dont pass in the requirements")
			}
		}
	}
}
