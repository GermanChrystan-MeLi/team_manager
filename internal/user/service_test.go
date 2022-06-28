package user

import (
	"context"
	"testing"

	"github.com/GermanChrystan-MeLi/team_manager/internal/domain"
	"github.com/GermanChrystan-MeLi/team_manager/internal/dto"
	"github.com/stretchr/testify/assert"
)

var mockDB = []domain.User{
	{
		ID:             "ID",
		LdapUser:       "LdapUser",
		FirstName:      "FirstName",
		LastName:       "LastName",
		Email:          "someemail@mercadolibre.com",
		MeliSite:       "site",
		IsAdmin:        false,
		HashedPassword: "password",
		Account:        1000000,
	},
}

var r = NewMockRepository(&mockDB)
var s = NewService(r)

//=================================================================================//
func TestGet_Ok(t *testing.T) {
	ctx := context.Background()
	user, err := s.Get(ctx, "ID")
	assert.Zero(t, err)
	assert.Equal(t, "FirstName", user.FirstName)
}

//=================================================================================//
func TestGet_Error(t *testing.T) {
	ctx := context.Background()
	user, err := s.Get(ctx, "id")
	assert.Error(t, err)
	assert.Equal(t, "", user.FirstName)
}

//=================================================================================//
func TestGetProfile_Ok(t *testing.T) {
	ctx := context.Background()
	profile, err := s.GetProfile(ctx, "ID")
	assert.Zero(t, err)
	assert.Equal(t, "FirstName", profile.UserData.FirstName)
}

//=================================================================================//
func TestGetProfile_Error(t *testing.T) {
	ctx := context.Background()
	profile, err := s.GetProfile(ctx, "id")
	assert.Error(t, err)
	assert.Equal(t, "", profile.UserData.FirstName)
}

//=================================================================================//
func TestRegister_Ok(t *testing.T) {
	ctx := context.Background()

	userRegister := dto.UserRegister{
		LdapUser:  "NewLdap",
		FirstName: "NewFirstName",
		LastName:  "NewLastName",
		Email:     "another_email@mercadolibre.com",
		MeliSite:  "site",
		Password:  "password",
	}

	session, err := s.Register(ctx, userRegister)

	assert.Zero(t, err)
	assert.NotEqual(t, "", session)
}

//=================================================================================//
func TestRegister_EmailError(t *testing.T) {
	ctx := context.Background()

	userRegister := dto.UserRegister{
		LdapUser:  "NewLdap",
		FirstName: "NewFirstName",
		LastName:  "NewLastName",
		Email:     "someemail@mercadolibre.com",
		MeliSite:  "site",
		Password:  "password",
	}

	session, err := s.Register(ctx, userRegister)

	assert.Error(t, err)
	assert.Equal(t, "", session)
}

//=================================================================================//
func TestRegister_NotUniqueError(t *testing.T) {
	ctx := context.Background()

	userRegister := dto.UserRegister{
		LdapUser:  "LdapUser",
		FirstName: "NewFirstName",
		LastName:  "NewLastName",
		Email:     "another_email@mercadolibre.com",
		MeliSite:  "site",
		Password:  "password",
	}

	session, err := s.Register(ctx, userRegister)

	assert.Error(t, err)
	assert.Equal(t, "", session)
}

//=================================================================================//
func TestLogin_Ok(t *testing.T) {
	ctx := context.Background()

	userLogin := dto.UserLogin{
		EmailOrUsername: "LdapUser",
		Password:        "password",
	}
	session, err := s.Login(ctx, userLogin)

	assert.Zero(t, err)
	assert.NotEqual(t, "", session)
}

//=================================================================================//
func TestLogin_Error(t *testing.T) {
	ctx := context.Background()

	userLogin := dto.UserLogin{
		EmailOrUsername: "not an user",
		Password:        "password",
	}
	session, err := s.Login(ctx, userLogin)

	assert.Error(t, err)
	assert.Equal(t, "", session)
}

//=================================================================================//
func TestEditUser_Ok(t *testing.T) {
	ctx := context.Background()

	userEdit := dto.UserEdit{
		ID:        "ID",
		LdapUser:  "NewLDAP",
		FirstName: "NewName",
	}
	userEdited, err := s.EditUser(ctx, userEdit)
	assert.Zero(t, err)
	assert.Equal(t, "NewLDAP", userEdited.LdapUser)
}

//=================================================================================//
func TestEditUser_Error(t *testing.T) {
	ctx := context.Background()

	userEdit := dto.UserEdit{
		ID:        "id",
		LdapUser:  "NewLDAP",
		FirstName: "NewName",
	}
	userEdited, err := s.EditUser(ctx, userEdit)
	assert.Error(t, err)
	assert.Equal(t, "", userEdited.LdapUser)
}

//=================================================================================//
func TestDeleteUser_Ok(t *testing.T) {
	ctx := context.Background()

	err := s.DeleteUser(ctx, "ID")
	assert.Zero(t, err)
}

//=================================================================================//
func TestDeleteUser_Error(t *testing.T) {
	ctx := context.Background()

	err := s.DeleteUser(ctx, "id")
	assert.Error(t, err)
}

//=================================================================================//
