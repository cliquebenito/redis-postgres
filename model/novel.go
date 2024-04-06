package model

import "database/sql"

type Novel struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	*sql.DB
}
