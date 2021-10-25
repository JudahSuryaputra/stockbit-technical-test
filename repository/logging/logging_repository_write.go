package logging

import (
	"stockbit-backend/models/db"
	"stockbit-backend/models/requests"

	"github.com/gocraft/dbr"
)

func WriteLog(sess *dbr.Session, r requests.WriteLogRequest) error {
	log := db.Log{
		Method:    r.Method,
		URL:       r.URL,
		Host:      r.Host,
		UserAgent: r.UserAgent,
	}

	columns := []string{
		"method",
		"url",
		"host",
		"user_agent",
	}

	_, err := sess.InsertInto(db.Log{}.TableName()).
		Columns(columns...).
		Record(log).
		Exec()
	if err != nil {
		return err
	}

	return nil
}
