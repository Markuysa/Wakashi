package database

import (
	"context"
	"gopkg.in/hedzr/errors.v3"
	"tgBotIntern/app/internal/entity"
	"tgBotIntern/app/internal/helpers/encoder"
)

type UsersDB interface {
	AddUser(ctx context.Context, password, username string, role int) error
	GetUser(ctx context.Context, username string) (*entity.User, error)
	GetUserRoleID(ctx context.Context, username string) (int, error)
	IsExist(ctx context.Context, username, password string) (*entity.User, error)
	GetSlavesList(ctx context.Context, masterUsername string, slaveRoleID int) ([]entity.User, error)
	GetUserID(ctx context.Context, username string) (int, error)
}

func (db *BotDatabase) GetUserID(ctx context.Context, username string) (int, error) {
	query := `
		select id from users
		where username=$1
`
	var userID int
	row := db.db.QueryRow(ctx, query, username)
	if err := row.Scan(&userID); err != nil {
		return -1, errors.New("failed getting userID:%v", err)
	}
	return userID, nil
}

func (u *BotDatabase) GetSlavesList(ctx context.Context, masterUsername string, slaveRoleID int) ([]entity.User, error) {
	query := `
	select username,role,password from users 
	inner join relation r on users.id = r.slave_id
	where users.role=$1 and users.id=$2
`
	rows, err := u.db.Query(ctx, query, slaveRoleID, masterUsername)
	if err != nil {
		return nil, errors.New("failed to get daimyo list:%v", err)
	}
	var users []entity.User
	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.Username, &user.Role, &user.Password); err != nil {
			return nil, errors.New("failed to scan daimyo:%v", err)
		}
		users = append(users, user)
	}
	return users, nil
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
	if err := row.Scan(&user.Username, &user.Password, &user.Role); err != nil {
		return nil, errors.New("failed getting user:%v", err)
	}
	return &user, nil
}

// GetUserRoleID returns string role of the user with given username
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

func (db *BotDatabase) IsExist(ctx context.Context, username, password string) (*entity.User, error) {
	user, err := db.GetUser(ctx, username)
	if err != nil {
		return nil, errors.New(" incorrect username")
	}
	matches, err := encoder.IsMatch(user.Password, password)
	if err != nil {
		return nil, errors.New("failed to check matching of the password:%v", err)
	}
	if matches {
		return user, nil
	}
	return nil, errors.New("user cannot be found: incorrect password")
}
