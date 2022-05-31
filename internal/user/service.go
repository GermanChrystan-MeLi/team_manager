package user

import (
	"context"
	"errors"
	"reflect"

	"github.com/GermanChrystan-MeLi/team_manager/internal/dto"
	"github.com/GermanChrystan-MeLi/team_manager/utils/check"
	"golang.org/x/crypto/bcrypt"
)

//=================================================================================//
type UserService interface {
	Get(ctx context.Context, id string) (dto.UserEdit, error)
	GetProfile(ctx context.Context, id string) (dto.FullProfile, error)
	Register(ctx context.Context, userRegister dto.UserRegister) (string, error)
	Login(ctx context.Context, user dto.UserLogin) (string, error)
	EditUser(ctx context.Context, user dto.UserEdit) (dto.UserEdit, error)
	DeleteUser(ctx context.Context, id string) error
}

//=================================================================================//

type service struct {
	repository UserRepository
}

//=================================================================================//

func NewService(repository UserRepository) UserService {
	return &service{
		repository: repository,
	}
}

//=================================================================================//

func (s *service) Get(ctx context.Context, id string) (dto.UserEdit, error) {
	return s.repository.Get(ctx, id)
}

//=================================================================================//
func (s *service) GetProfile(ctx context.Context, id string) (dto.FullProfile, error) {
	return s.repository.GetProfile(ctx, id)
}

//=================================================================================//
func (s *service) Register(ctx context.Context, userRegister dto.UserRegister) (string, error) {
	// Checking email is valid
	if isEmailValid := check.IsEmail(userRegister.Email); !isEmailValid {
		return "", errors.New("Email not valid")
	}
	// Checking user is unique
	if isUserUnique := s.repository.IsUserUnique(ctx, userRegister); isUserUnique != nil {
		return "", isUserUnique
	}

	// Hashing Password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userRegister.Password), 10)
	userRegister.Password = string(hashedPassword)

	// Returns session
	return s.repository.Register(ctx, userRegister)
}

//=================================================================================//
func (s *service) Login(ctx context.Context, user dto.UserLogin) (string, error) {
	return s.repository.Login(ctx, user, check.IsEmail(user.EmailOrUsername))
}

//=================================================================================//
func (s *service) EditUser(ctx context.Context, user dto.UserEdit) (dto.UserEdit, error) {
	/*
		Values to edit
			LdapUser
			FirstName
			LastName
			Email
			MeliSite
	*/
	previousUser, notFoundErr := s.Get(ctx, user.ID)
	previousUserValue := reflect.ValueOf(previousUser)

	if notFoundErr != nil {
		return dto.UserEdit{}, errors.New("Could not access user")
	}

	userType := reflect.TypeOf(user)
	referenceOfUser := reflect.ValueOf(&user).Elem()

	pNumFields := userType.NumField()
	for i := 0; i < pNumFields; i++ {
		fieldType := referenceOfUser.Field(i)
		pFieldValue := referenceOfUser.Field(i)
		previousValue := previousUserValue.Field(i)

		if pFieldValue.IsZero() {
			if fieldType.Kind() == reflect.Int {
				var PrevIntValue int = int(reflect.ValueOf(previousValue.Int()).Int())
				pFieldValue.Set(reflect.ValueOf(PrevIntValue))
			} else if fieldType.Kind() == reflect.Float32 {
				var PrevFloatValue float32 = float32(reflect.ValueOf(previousValue.Float()).Float())
				pFieldValue.Set(reflect.ValueOf(PrevFloatValue))
			} else if fieldType.Kind() == reflect.String {
				pFieldValue.Set(reflect.ValueOf(previousValue.String()))
			}
		}
	}

	userEdited, updateErr := s.repository.EditUser(ctx, user)
	if updateErr != nil {
		return dto.UserEdit{}, errors.New("Could not update user")
	} else {
		return userEdited, nil
	}
}

//=================================================================================//
func (s *service) DeleteUser(ctx context.Context, id string) error {
	return s.repository.DeleteUser(ctx, id)
}

//=================================================================================//
