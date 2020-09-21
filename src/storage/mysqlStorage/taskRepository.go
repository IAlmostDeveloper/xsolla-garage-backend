package mysqlStorage

import (
	"database/sql"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"
	"github.com/jmoiron/sqlx"
)

type TaskRepository struct {
	db *sqlx.DB
}

func (repo TaskRepository) Create(task *dto.Task) error {
	insertStatement := "INSERT INTO tasks (user_id, title, text_content, date_create, date_target, is_completed, is_important, is_urgent) VALUES (:user_id, :title, :text_content, :date_create, :date_target, false, :is_important, :is_urgent)"
	res, err := repo.db.NamedExec(insertStatement, task)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	task.Id = int(id)

	return nil
}

func (repo TaskRepository) GetByID(id int) (*dto.Task, error) {
	selectStatement := "SELECT " +
		"id, user_id, title, text_content, date_create, date_close, date_target, is_completed, is_important, is_urgent " +
		"FROM tasks " +
		"WHERE id = ?"
	task := &dto.Task{}
	err := repo.db.Get(task, selectStatement, id)
	if err != nil {
		return nil, err
	}
	return task, err
}

func (repo TaskRepository) GetAll() ([]*dto.Task, error) {
	selectStatement := "SELECT " +
		"id, user_id, title, text_content, date_create, date_close, date_target, is_completed, is_important, is_urgent " +
		"FROM tasks"
	tasks := &[]*dto.Task{}
	err := repo.db.Select(tasks, selectStatement)
	if err != nil {
		return nil, err
	}
	return *tasks, err
}

func (repo TaskRepository) RemoveByID(id int) error {
	removeStatement := "DELETE FROM `tasks` WHERE id = ?"
	res, err := repo.db.Exec(removeStatement, id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (repo TaskRepository) Update(task *dto.Task) error {
	updateStatement := "UPDATE `tasks` SET user_id = :user_id, title = :title, text_content = :text_content, date_create = :date_create, date_close = :date_close, date_target = :date_target, is_completed = :is_completed, is_important = :is_important, is_urgent = :is_urgent WHERE id = :id"
	res, err := repo.db.NamedExec(updateStatement, task)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}
