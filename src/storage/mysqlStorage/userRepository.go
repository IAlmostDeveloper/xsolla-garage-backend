package mysqlStorage

import (
	"database/sql"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func (repo *UserRepository) Create(user *dto.User) error {
	insertStatement := "INSERT INTO `users` (`id`, `email`, `name`, `given_name`, `family_name`, `locale`, `picture`) VALUES (:id, :email, :name, :given_name, :family_name, :locale, :picture)"
	if _, err := repo.db.NamedExec(insertStatement, user); err != nil {
		return err
	}

	return nil
}

// if user not found return nil
func (repo *UserRepository) GetById(id string) (*dto.User, error) {
	selectStatement := "SELECT * FROM `users` WHERE id = :id"
	user := &dto.User{}
	if err := repo.db.Get(user, selectStatement, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) RemoveById(id string) error {
	deleteStatement := "DELETE FROM `users` WHERE id = ?"
	if _, err := repo.db.Exec(deleteStatement, id); err != nil {
		return err
	}
	return nil
}
