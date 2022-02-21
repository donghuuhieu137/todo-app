package handlers

import "database/sql"

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = "dhh13072001"
	dbname   = "todo_app"
)

type Handler struct {
	conn *sql.DB
}

func New() *Handler {
	return &Handler{}
}
