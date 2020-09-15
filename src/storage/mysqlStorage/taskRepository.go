package mysqlStorage

import (
	"database/sql"
	"fmt"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"
	"github.com/jmoiron/sqlx"
)

const DateFormat = "%Y-%m-%d %H:%i:%s"

type TaskRepository struct {
	db *sqlx.DB
}

func (repo TaskRepository) Create(task *dto.Task) error {
	insertStatement := fmt.Sprintf("INSERT INTO tasks (user_id, title, text_content, date_create, date_target, is_completed) VALUES (?, ?, ?, ?, STR_TO_DATE(?, '%s'), STR_TO_DATE(?, '%s'), false)", DateFormat, DateFormat)
	res, err := repo.db.Exec(insertStatement, task.UserId, task.Title, task.TextContent, task.DateCreate, task.DateTarget)
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
		"id, user_id, title, text_content, date_create, date_close, date_target, is_completed " +
		"FROM tasks " +
		"WHERE id = ?"
	task := &dto.Task{}
	err := repo.db.Get(task, selectStatement, id)
	if err != nil {
		return nil, err
	}
	return task, err
}

func (repo TaskRepository) GetAll() (*[]dto.Task, error) {
	selectStatement := "SELECT " +
		"id, user_id, title, text_content, date_create, date_close, date_target, is_completed " +
		"FROM tasks"
	tasks := &[]dto.Task{}
	err := repo.db.Select(tasks, selectStatement)
	if err != nil {
		return nil, err
	}
	return tasks, err
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
	updateStatement := fmt.Sprintf("UPDATE `tasks` SET user_id = ?, title = ?, text_content = ?, date_create = STR_TO_DATE(?, '%s'), date_close = STR_TO_DATE(?, '%s'), date_target = STR_TO_DATE(?, '%s'), is_completed = ? WHERE id = ?", DateFormat, DateFormat, DateFormat)
	res, err := repo.db.Exec(updateStatement, task.UserId, task.Title, task.TextContent, task.DateCreate, task.DateClose, task.DateTarget, task.IsCompleted, task.Id)
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
