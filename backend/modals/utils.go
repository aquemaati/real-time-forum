package modals

import "database/sql"

const DbFilePathEnv = ""

type QueryExecutor func(*sql.Rows) (interface{}, error)
