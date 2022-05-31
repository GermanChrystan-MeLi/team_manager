package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/GermanChrystan-MeLi/team_manager/internal/domain"
	"github.com/GermanChrystan-MeLi/team_manager/internal/dto"
	"github.com/GermanChrystan-MeLi/team_manager/utils/error_vars"
)

const INITIAL_ACCOUNT_AMOUNT = 100000

//=================================================================================//
type UserRepository interface {
	Get(ctx context.Context, id string) (dto.UserEdit, error)
	GetProfile(ctx context.Context, id string) (dto.FullProfile, error)
	Register(ctx context.Context, user dto.UserRegister) (string, error)
	Login(ctx context.Context, user dto.UserLogin, isEmail bool) (string, error)
	EditUser(ctx context.Context, user dto.UserEdit) (dto.UserEdit, error)
	DeleteUser(ctx context.Context, id string) error
	IsUserUnique(ctx context.Context, user dto.UserRegister) error
}

//=================================================================================//
type repository struct {
	db *sql.DB
}

//=================================================================================//
func NewRepository(db *sql.DB) UserRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) Get(ctx context.Context, id string) (dto.UserEdit, error) {
	var user dto.UserEdit
	getUserQuery := "SELECT (ldap_user, first_name, last_name, email, meli_site) FROM users WHERE id=?"
	row := r.db.QueryRow(getUserQuery, id)
	err := row.Scan(&user.LdapUser, &user.FirstName, &user.LastName, &user.Email, &user.MeliSite)
	if err != nil {
		return dto.UserEdit{}, error_vars.ErrActionNotPerformed
	}
	return user, nil
}

//=================================================================================//
func (r *repository) GetProfile(ctx context.Context, id string) (dto.FullProfile, error) {
	// Get User Data
	var userData dto.UserData
	getUserDataQuery := "SELECT (ldap_user, first_name, last_name, email, meli_site) FROM users WHERE id = ?"

	row := r.db.QueryRow(getUserDataQuery, id)
	err := row.Scan(&userData.LdapUser, &userData.FirstName, &userData.LastName, &userData.Email, &userData.MeliSite)
	if err != nil {
		return dto.FullProfile{}, error_vars.ErrRetrievingData
	}

	// Get User Team

}

//=================================================================================//
func (r *repository) Register(ctx context.Context, user dto.UserRegister) (string, error) {
	userCreateQuery := "INSERT INTO users(ldap_user, first_name, last_name, email, meli_site, password, account) VALUES (?,?,?,?,?,?,?)"
	stmt, err := r.db.Prepare(userCreateQuery)
	if err != nil {
		return "", error_vars.ErrDatabase
	}

	res, err := stmt.Exec(
		user.LdapUser,
		user.FirstName,
		user.LastName,
		user.Email,
		user.MeliSite,
		user.Password,
		INITIAL_ACCOUNT_AMOUNT,
	)
	if err != nil {
		return "", error_vars.ErrActionNotPerformed
	}

	id, err := res.LastInsertId()
	if err != nil {
		return "", error_vars.XNotFound("user")
	}
	return r.createSession(ctx, int(id))
}

//=================================================================================//
func (r *repository) Login(ctx context.Context, userLogin dto.UserLogin, isEmail bool) (string, error) {
	var firstValue string
	if isEmail {
		firstValue = "email"
	} else {
		firstValue = "username"
	}

	getUserQuery := fmt.Sprintf("SELECT (id, hashed_password) FROM users WHERE %s=?", firstValue)
	row := r.db.QueryRow(getUserQuery, userLogin.EmailOrUsername)

	var id int
	var password string
	err := row.Scan(
		&id,
		&password,
	)

	if err != nil {
		return "", error_vars.ErrDatabase
	}

	// Checking password
	passwordError := bcrypt.CompareHashAndPassword([]byte(userLogin.Password), []byte(password))
	if passwordError != nil {
		return "", error_vars.ErrWrongCredentials
	}
	return r.createSession(ctx, id)
}

//=================================================================================//
func (r *repository) EditUser(ctx context.Context, user dto.UserEdit) (dto.UserEdit, error) {
	query := "UPDATE users SET ldap_user=?, first_name=?, last_name=?, email=?, meli_site=? WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return dto.UserEdit{}, error_vars.ErrDatabase
	}

	res, err := stmt.Exec(user.LdapUser, user.FirstName, user.LastName, user.Email, user.MeliSite, user.ID)
	if err != nil {
		return dto.UserEdit{}, error_vars.ErrActionNotPerformed
	}

	_, err = res.RowsAffected()
	if err != nil {
		return dto.UserEdit{}, error_vars.ErrNotFound
	}

	return user, nil
}

//=================================================================================//
func (r *repository) DeleteUser(ctx context.Context, id string) error {
	// DELETE USER
	deleteUser := "DELETE FROM users WHERE id=?"
	stmt, err := r.db.Prepare(deleteUser)
	if err != nil {
		return error_vars.ErrDatabase
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return error_vars.ErrActionNotPerformed
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return error_vars.ErrRetrievingData
	}

	if affected == 0 {
		return error_vars.XNotFound("user")
	}
	// CASCADING SIDE - EFFECTS
	// DELETE CONTRACTS
	// ...
	return nil
}

//=================================================================================//
func (r *repository) IsUserUnique(ctx context.Context, user dto.UserRegister) error {
	var amountUsers int

	checkUniqueQuery := "SELECT COUNT(*) FROM users WHERE email = ? OR ldap_user = ?"
	row := r.db.QueryRow(checkUniqueQuery, user.Email, user.LdapUser)
	_ = row.Scan(&amountUsers)
	if amountUsers > 0 {
		return errors.New("User already has an account")
	}
	return nil
}

//=================================================================================//
func (r *repository) createSession(ctx context.Context, user_id int) (string, error) {
	//Create session
	newId := uuid.New().String()
	newSession := domain.Session{
		ID:        newId,
		UserID:    user_id,
		CreatedAt: time.Now().Unix(),
	}
	sessionCreationQuery := "INSERT INTO sessions(id, user_id, created_at) VALUES (?, ?, ?)"
	stmt, err := r.db.Prepare(sessionCreationQuery)
	if err != nil {
		return "", error_vars.ErrDatabase
	}

	_, err = stmt.Exec(newSession.ID, newSession.UserID, newSession.CreatedAt)
	if err != nil {
		return "", error_vars.ErrActionNotPerformed
	}

	return newId, nil
}
