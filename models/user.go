package models

import (
	"database/sql"
)

type User struct {
	Id       int            `json:"id" xorm:"pk autoincr id"`
	Username sql.NullString `json:"username" xorm:"username"`
	Password sql.NullString `json:"password" xorm:"password"`
}
