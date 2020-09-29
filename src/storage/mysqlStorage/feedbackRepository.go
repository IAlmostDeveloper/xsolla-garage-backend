package mysqlStorage

import (
	"github.com/IAlmostDeveloper/xsolla-garage-backend/src/dto"
	"github.com/jmoiron/sqlx"
)

type FeedbackRepository struct{
	db *sqlx.DB
}

func (repo *FeedbackRepository) AddFeedback(feedback *dto.Feedback) error {
	insertStatement := "INSERT INTO feedback (date, content) VALUES (:date_create, :content)"
	res, err := repo.db.NamedExec(insertStatement, feedback)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	feedback.Id = int(id)

	return nil
}

func (repo *FeedbackRepository) GetAllFeedback() ([]*dto.Feedback, error){
	selectStatement := "SELECT id, date_create, content FROM feedback"
	allFeedback := &[]*dto.Feedback{}
	if err := repo.db.Select(allFeedback, selectStatement); err != nil {
		return nil, err
	}
	return *allFeedback, nil
}
