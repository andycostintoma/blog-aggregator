package internal

import (
	"github.com/andycostintoma/blog-aggregator/internal/database"
)

type State struct {
	Db  *database.Queries
	Cfg *Config
}
