package db

import (
	"github.com/osmanzz/inventory_backend/model"
	"database/sql"
	"sync"
)

var db Repo
var once sync.Once

func GetOrderDB(db2 *sql.DB) Repo {
	once.Do(func() {
		db = &dbRepo{
			db: db2,
		}
	})
	return db
}

type Repo interface {
	SelectUser() ([]*model.User, error)
	SelectUserByUsername(username, password  string) (*model.User,error)
}
type dbRepo struct {
	db *sql.DB
}

func (d dbRepo) SelectUser() ([]*model.User, error) {

	rows, err := d.db.Query("SELECT * from customer")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	users := make([]*model.User, 0)
	for rows.Next() {
		user := new(model.User)
		_ = rows.Scan(&user.Username, &user.Password)
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (d dbRepo) SelectUserByUsername(username , password string)(*model.User,error) {
	queryStatement := "SELECT * from customer where username=$1 and password=$2"

	row:= d.db.QueryRow(queryStatement,username,password)
	user := new(model.User)
	err := row.Scan(&user.Username,&user.Password)
	if err != nil {
		return nil,err
	}

	return user,nil
}
