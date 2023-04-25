package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"gopkg.in/hedzr/errors.v3"
	"tgBotIntern/app/internal/entity"
	"tgBotIntern/app/internal/helpers/encoder"
)

// UsersDatabase interface is an interface that provides contract to
// a structure that implements the functionality of a users table in the database
// The AddUser method is used to add new user into the user table
// The GetUser method is used to get user info from the table
// The GetUserRoleID method is used to get user's role identification from db
// The IsExist method is used to check existence of user with given username and password
// The GetSlavesList method is used to get all the slaves of user with given id's
// The GetUserID method is used to get identification of user with given username
// THe UpdatePassword method is used to reset old password of user
type UsersDatabase interface {
	AddUser(ctx context.Context, password, username string, role int) error
	GetUser(ctx context.Context, username string) (*entity.User, error)
	GetUserRoleID(ctx context.Context, username string) (int, error)
	IsExist(ctx context.Context, username, password string) (*entity.User, error)
	GetSlavesList(ctx context.Context, masterID int, slaveRoleID int) ([]entity.User, error)
	GetUserID(ctx context.Context, username string) (int, error)
	UpdatePassword(ctx context.Context, name string, password string) error
}
type UsersRepository struct {
	db *pgxpool.Pool
}

func NewUsersDB(db *pgxpool.Pool) *UsersRepository {
	return &UsersRepository{db: db}
}

func (db *UsersRepository) UpdatePassword(ctx context.Context, name string, password string) error {
	hashPassword, err := encoder.Encode(password)
	if err != nil {
		return err
	}
	query := `
	update users 
	set password=$1
	where username=$2
`
	_, err = db.db.Query(ctx, query, hashPassword, name)
	if err != nil {
		return errors.New("failed to reset your password")
	}
	return nil
}
func (db *UsersRepository) GetUserID(ctx context.Context, username string) (int, error) {
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

func (u *UsersRepository) GetSlavesList(ctx context.Context, masterID int, slaveRoleID int) ([]entity.User, error) {
	query := `
	select username,role,password from users u
	inner join relation r on u.id = r.slave_id
	where r.master_id=$1 and u.role=$2
`
	rows, err := u.db.Query(ctx, query, masterID, slaveRoleID)
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
func (db *UsersRepository) AddUser(ctx context.Context, password, username string, role int) error {

	query := `
		insert into users(
			username,
			password,
			role
		)values (
				 $1,$2,$3
		)
	`
	password, err := encoder.Encode(password)
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
func (db *UsersRepository) GetUser(ctx context.Context, username string) (*entity.User, error) {

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
func (db *UsersRepository) GetUserRoleID(ctx context.Context, username string) (int, error) {

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

func (db *UsersRepository) IsExist(ctx context.Context, username, password string) (*entity.User, error) {
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
