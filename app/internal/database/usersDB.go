package database

import (
	"context"
	"fmt"
	"gopkg.in/hedzr/errors.v3"
	"tgBotIntern/app/internal/entity"
	"tgBotIntern/app/internal/helpers/encoder"
)

type UsersDB interface {
	AddUser(ctx context.Context, password, username string, role int) error
	GetUser(ctx context.Context, username string) (*entity.User, error)
	GetUserRoleID(ctx context.Context, username string) (int, error)
	IsExist(ctx context.Context, username, password string) (bool, error)
}

// AddUser method creates new entry in the users table of the database
func (db *BotDatabase) AddUser(ctx context.Context, password, username string, role int) error {

	query := `
		insert into users(
			username,
			password,
			role
		)values (
				 $1,$2,$3
		)
	`
	password, err := encoder.EncodePassword(password)
	if err != nil {
		return errors.New("error of adding the user:%v", err)
	}
	_, err = db.db.Query(ctx, query,
		username,
		password,
		role,
	)
	if err != nil {
		return errors.New("failed adding user:%v", err)
	}
	return nil
}

// GetUser method returns an object of user from the database
func (db *BotDatabase) GetUser(ctx context.Context, username string) (*entity.User, error) {

	query := `
		select username,password,role from users
		where username=$1
`
	var user entity.User
	row := db.db.QueryRow(ctx, query, username)
	fmt.Println()
	if err := row.Scan(&user.Username, &user.Password, &user.Role); err != nil {
		return nil, errors.New("failed getting user:%v", err)
	}
	return &user, nil
}

// GetUserRole returns string role of the user with given username
func (db *BotDatabase) GetUserRoleID(ctx context.Context, username string) (int, error) {

	//query := `
	//	select r.role from users
	//	inner join roles r on r.role_id = users.role
	//	where users.username=$1
	//`
	query := `
		select r.role_id from users 
		inner join roles r on r.role_id = users.role
		where users.username=$1
	`
	var role int
	row := db.db.QueryRow(ctx, query, username)
	if err := row.Scan(&role); err != nil {
		return -1, errors.New("failed to get user role:%v", err)
	}
	return role, nil
}

func (db *BotDatabase) IsExist(ctx context.Context, username, password string) (bool, error) {
	user, err := db.GetUser(ctx, username)
	if err != nil {
		return false, errors.New("failed to check existence of the user:%v", err)
	}
	matches, err := encoder.IsMatch(user.Password, password)
	if err != nil {
		return false, errors.New("failed to check matching of the password:%v", err)
	}
	return matches, nil
}
