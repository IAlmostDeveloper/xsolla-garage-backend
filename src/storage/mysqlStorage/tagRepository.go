package mysqlStorage

import (
	"database/sql"
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"
	"github.com/jmoiron/sqlx"
)

type TagRepository struct {
	db *sqlx.DB
}

func (repo TagRepository) AddToTask(taskId int, tag *dto.Tag) error {
	if err := repo.insertOrSelectTag(tag); err != nil {
		return err
	}
	insertStatement := "INSERT INTO `task_tags` (task_id, tag_id) VALUES (?, ?)"
	if _, err := repo.db.Exec(insertStatement, taskId, tag.Id); err != nil {
		return err
	}
	return nil
}

func (repo TagRepository) RemoveFromTask(taskId int, tagId int) error {
	deleteStatement := "DELETE FROM task_tags WHERE task_id = ? AND  tag_id = ?"
	res, err := repo.db.Exec(deleteStatement, taskId, tagId)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}
	if err := repo.removeIfNotUsed(tagId); err != nil {
		return err
	}
	return nil

}

func (repo TagRepository) GetByTaskId(taskId int) ([]*dto.Tag, error) {
	selectStatement := "SELECT t.id, t.name FROM task_tags tt JOIN tags t on t.id = tt.tag_id WHERE tt.task_id = ?"
	tags := &[]*dto.Tag{}
	if err := repo.db.Select(tags, selectStatement, taskId); err != nil {
		return nil, err
	}
	return *tags, nil
}

// inserts tag into table and writes its id
// if tag already exists writes its id
func (repo TagRepository) insertOrSelectTag(tag *dto.Tag) error {
	selectStatement := "SELECT id, name FROM `tags` WHERE LOWER(name) = LOWER(?)"
	if err := repo.db.Get(tag, selectStatement, tag.Name); err != nil {
		if err == sql.ErrNoRows {
			insertStatement := "INSERT INTO `tags` (name) values (?)"
			res, err := repo.db.Exec(insertStatement, tag.Name)
			if err != nil {
				return err
			}
			insertedId, err := res.LastInsertId()
			if err != nil {
				return err
			}

			tag.Id = int(insertedId)
		} else {
			return err
		}
	}
	return nil
}

func (repo TagRepository) removeIfNotUsed(tagId int) error {
	selectStatement := "SELECT COUNT(*) FROM task_tags WHERE tag_id = ?"
	useCount := 0 // how many tasks contain this tag

	if err := repo.db.Get(&useCount, selectStatement, tagId); err != nil {
		return err
	}
	if useCount == 0 {
		deleteStatement := "DELETE FROM	tags WHERE id = ?"
		if _, err := repo.db.Exec(deleteStatement, tagId); err != nil {
			return err
		}
	}
	return nil
}
