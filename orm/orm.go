package orm

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"hxsg/config"
	"hxsg/logger"
	"hxsg/orm/dialect"
	"hxsg/orm/session"
	"reflect"
	"strings"
)

// Engine is the main struct of geeorm, manages all db sessions and transactions.
type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func GetDbConnectionString(host, username, password, dbname string, port int) string {
	str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, dbname)
	str += "?charset=utf8&timeout=5s&parseTime=true&loc=Asia%2FShanghai&allowNativePasswords=true"
	return str
}

var Orm *Engine

func Init() {
	driver, _ := config.Cfg.GetString("database:driver")
	host, _ := config.Cfg.GetString("database:host")
	username, _ := config.Cfg.GetString("database:username")
	password, _ := config.Cfg.GetString("database:password")
	dbname, _ := config.Cfg.GetString("database:dbname")
	port, _ := config.Cfg.GetInt("database:port")
	dsn := GetDbConnectionString(host, username, password, dbname, port)
	err := NewEngine(driver, dsn)
	if err != nil {
		panic(err)
	}
}

// NewEngine create a instance of Engine
// connect database and ping it to test whether it's alive
func NewEngine(driver, source string) (err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		logger.Error("dialect %s Not Found", err)
		return
	}
	// Send a ping to make sure the database connection is alive.
	if err = db.Ping(); err != nil {
		logger.Error("dialect %s Not Found", err)
		return
	}
	// make sure the specific dialect exists
	dial, ok := dialect.GetDialect(driver)
	if !ok {
		logger.Error("dialect %s Not Found", driver)
		return
	}
	Orm = &Engine{db: db, dialect: dial}
	logger.Info("Connect database success")
	return
}

// Close database connection
func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		logger.Error("Failed to close database")
		return
	}
	logger.Info("Close database success")
}

// NewSession creates a new session for next operations
func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db, engine.dialect)
}

// TxFunc will be called between tx.Begin() and tx.Commit()
// https://stackoverflow.com/questions/16184238/database-sql-tx-detecting-commit-or-rollback
type TxFunc func(*session.Session) (interface{}, error)

// Transaction executes sql wrapped in a transaction, then automatically commit if no error occurs
func (engine *Engine) Transaction(f TxFunc) (result interface{}, err error) {
	s := engine.NewSession()
	if err := s.Begin(); err != nil {
		return nil, err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = s.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			_ = s.Rollback() // err is non-nil; don't change it
		} else {
			err = s.Commit() // err is nil; if Commit returns error update err
		}
	}()

	return f(s)
}

// difference returns a - b
func difference(a []string, b []string) (diff []string) {
	mapB := make(map[string]bool)
	for _, v := range b {
		mapB[v] = true
	}
	for _, v := range a {
		if _, ok := mapB[v]; !ok {
			diff = append(diff, v)
		}
	}
	return
}

// Migrate table
func (engine *Engine) Migrate(value interface{}) error {
	_, err := engine.Transaction(func(s *session.Session) (result interface{}, err error) {
		if !s.Model(value).HasTable() {
			logger.Info("table %s doesn't exist", s.RefTable().Name)
			return nil, s.CreateTable()
		}
		table := s.RefTable()
		rows, _ := s.Raw(fmt.Sprintf("SELECT * FROM %s LIMIT 1", table.Name)).QueryRows()
		columns, _ := rows.Columns()
		addCols := difference(table.FieldNames, columns)
		delCols := difference(columns, table.FieldNames)
		logger.Info("added cols %v, deleted cols %v", addCols, delCols)

		for _, col := range addCols {
			f := table.GetField(col)
			sqlStr := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s;", table.Name, f.Name, f.Type)
			if _, err = s.Raw(sqlStr).Exec(); err != nil {
				return
			}
		}

		if len(delCols) == 0 {
			return
		}
		tmp := "tmp_" + table.Name
		fieldStr := strings.Join(table.FieldNames, ", ")
		s.Raw(fmt.Sprintf("CREATE TABLE %s AS SELECT %s from %s;", tmp, fieldStr, table.Name))
		s.Raw(fmt.Sprintf("DROP TABLE %s;", table.Name))
		s.Raw(fmt.Sprintf("ALTER TABLE %s RENAME TO %s;", tmp, table.Name))
		_, err = s.Exec()
		return
	})
	return err
}

func GetAllFieldsAsStringWithTableName(obj interface{}, tableName string) string {
	objT := reflect.TypeOf(obj)
	var fields []string
	for i := 0; i < objT.NumField(); i++ {
		fieldT := objT.Field(i)
		tag := fieldT.Tag.Get("db")
		if tag == "" {
			continue
		}
		oneFileName := fmt.Sprintf("`%s`", tag)
		if tableName != "" {
			oneFileName = fmt.Sprintf("%s.`%s`", tableName, tag)
		}
		fields = append(fields, oneFileName)
	}
	return strings.Join(fields, ",")
}
