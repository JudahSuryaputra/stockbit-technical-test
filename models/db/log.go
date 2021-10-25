package db

type Log struct {
	Method    string `db:"method" json:"method"`
	URL       string `db:"url" json:"url"`
	Host      string `db:"host" json:"host"`
	UserAgent string `db:"user_agent" json:"user_agent"`
}

func (c Log) TableName() string {
	return "logs"
}
