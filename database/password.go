package database

import (
	"github.com/go-pg/pg"
)

type Password struct {
	DB *pg.DB
}