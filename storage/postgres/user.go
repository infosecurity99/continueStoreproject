package postgres

import (
	"database/sql"
	"fmt"
	"test/models"
	"test/storage"

	"github.com/google/uuid"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) storage.IUserStorage {
	return userRepo{
		db: db,
	}
}

//create
func (u userRepo) Create(createUser models.CreateUser) (models.User, error) {

	uid := uuid.New()

	if _, err := u.db.Exec(`insert into 
			users values ($1, $2, $3, $4, $5, $6)
			`,
		uid,
		createUser.FullName,
		createUser.Phone,
		createUser.Password,
		createUser.UserType,
		createUser.Cash,
	); err != nil {
		fmt.Println("error while inserting data", err.Error())
		return models.User{}, err
	}

	return models.User{}, nil
}

//getbyid
func (u userRepo) GetByID(rex models.PrimaryKey) (models.User, error) {
	id := rex.ID
	row := u.db.QueryRow(`select * from users where id=$1`, id)
	users := models.User{}
	if err := row.Scan(
		&users.ID,
		&users.FullName,
		&users.Phone,
		&users.Password,
		&users.UserType,
		&users.Cash,
	); err != nil {
		fmt.Println("error while getbyid data", err.Error())
		return models.User{}, err
	}

	return models.User{}, nil
}

func (u userRepo) GetList(models.GetListRequest) (models.UsersResponse, error) {
	return models.UsersResponse{}, nil
}

func (u userRepo) Update(models.UpdateUser) (models.User, error) {

	return models.User{}, nil
}

//delete
func (u userRepo) Delete(rex models.PrimaryKey) error {
	id := rex.ID
	fmt.Println(id)
	_, err := u.db.Exec(`delete from users where id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}
