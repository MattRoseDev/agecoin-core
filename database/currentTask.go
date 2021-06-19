package database

import (
	"github.com/go-pg/pg"
)

type CurrentTask struct {
	DB *pg.DB
}
