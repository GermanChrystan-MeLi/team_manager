package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/GermanChrystan-MeLi/team_manager/internal/domain"
	"github.com/GermanChrystan-MeLi/team_manager/internal/dto"
	"github.com/GermanChrystan-MeLi/team_manager/pkg/utils/error_vars"
)

type mockRepository struct {
	db *[]domain.User
}

//=================================================================================//
func NewMockRepository(db *[]domain.User) UserRepository {
	return &mockRepository{
		db: db,
	}
}

//=================================================================================//
func (r *mockRepository) Get(ctx context.Context, id string) (dto.UserEdit, error) {
	var result dto.UserEdit

	for _, u := range *r.db {
		if u.ID == id {
			result.ID = u.ID
			result.FirstName = u.FirstName
			result.LastName = u.LastName
			result.LdapUser = u.LdapUser
			result.Email = u.Email
			result.MeliSite = u.MeliSite

			return result, nil
		}
	}
	return dto.UserEdit{}, error_vars.XNotFound("user")
}

//=================================================================================//
func (r *mockRepository) GetProfile(ctx context.Context, id string) (dto.FullProfile, error) {
	var userData dto.UserData
	for _, u := range *r.db {
		if u.ID == id {
			userData.FirstName = u.FirstName
			userData.LastName = u.LastName
			userData.LdapUser = u.LdapUser
			userData.Email = u.Email
			userData.MeliSite = u.MeliSite
			userData.Account = u.Account

			return dto.FullProfile{
				UserData: userData,
				Team:     dto.Team{},
			}, nil
		}
	}
	return dto.FullProfile{}, error_vars.XNotFound("user")
}

//=================================================================================//
func (r *mockRepository) Register(ctx context.Context, user dto.UserRegister) (string, error) {
	register := domain.User{
		ID:             "new_id",
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		LdapUser:       user.LdapUser,
		Email:          user.Email,
		MeliSite:       user.MeliSite,
		HashedPassword: user.Password,
	}
	*r.db = append(*r.db, register)
	return "new_session", nil
}

//=================================================================================//
func (r *mockRepository) Login(ctx context.Context, user dto.UserLogin, isEmail bool) (string, error) {
	if isEmail {
		for _, u := range *r.db {
			if u.Email == user.EmailOrUsername {
				if u.HashedPassword == user.Password {
					return "new session", nil
				} else {
					return "", error_vars.ErrWrongCredentials
				}
			}
		}
		return "", error_vars.ErrWrongCredentials
	} else {
		for _, u := range *r.db {
			if u.LdapUser == user.EmailOrUsername {
				if u.HashedPassword == user.Password {
					return "new session", nil
				} else {
					return "", error_vars.ErrWrongCredentials
				}
			}
		}
		return "", error_vars.ErrWrongCredentials
	}
}

//=================================================================================//
func (r *mockRepository) EditUser(ctx context.Context, user dto.UserEdit) (dto.UserEdit, error) {
	for i, u := range *r.db {
		if u.ID == user.ID {
			(*(*r).db)[i] = domain.User{
				ID:        user.ID,
				FirstName: u.FirstName,
				LastName:  u.LastName,
				LdapUser:  u.LdapUser,
				Email:     u.Email,
				MeliSite:  u.MeliSite,

				IsAdmin:        u.IsAdmin,
				HashedPassword: u.HashedPassword,
				Account:        u.Account,
			}
			return user, nil
		}
	}
	return dto.UserEdit{}, error_vars.ErrNotFound
}

//=================================================================================//
func (r *mockRepository) DeleteUser(ctx context.Context, id string) error {
	prevLen := len(*(*r).db)
	index := 0
	userFound := false
	for i, p := range *r.db {
		if p.ID == id {
			index = i
			userFound = true
		}
	}
	if !userFound {
		msg := fmt.Sprintf("User with id %s not found", id)
		return errors.New(msg)
	} else {
		*r.db = append((*(*r).db)[:index], (*(*r).db)[index+1:]...)

		if len(*(*r).db) < prevLen {
			return nil
		}
		return errors.New("SQL error")
	}
}

//=================================================================================//
func (r *mockRepository) IsUserUnique(ctx context.Context, user dto.UserRegister) error {
	amountUsers := 0

	for _, u := range *r.db {
		if u.Email == user.Email || u.LdapUser == user.LdapUser {
			amountUsers++
		}
	}
	if amountUsers > 0 {
		return nil
	}
	return errors.New("user already has an account")
}

//=================================================================================//
