package orm

import (
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Uid      int    `db:uid`
	Mid      string `db:"m_id"`
	Name     string `db:"name"`
	Username string `db:"username"`
	Password string `db:"password"`
	N        int    `db:"n"`
	Y        int    `db:"y"`
	R        int    `db:"r"`
	S        int    `db:"s"`
	F        int    `db:"f"`
	M        int    `db:"m"`
	Ma       string `db:"ma"`
	Aqm      string `db:"aqm"`
}

//
//func OpenDB(t *testing.T) *Engine {
//	t.Helper()
//	dsn := GetDbConnectionString()
//	err := NewEngine("mysql", dsn)
//	if err != nil {
//		panic(err)
//	}
//	return engine
//}
//
//func TestNewEngine(t *testing.T) {
//	engine := OpenDB(t)
//	defer engine.Close()
//}
//
//func TestEngine_Migrate(t *testing.T) {
//
//	engine := OpenDB(t)
//
//	defer engine.Close()
//
//	s := engine.NewSession()
//	//_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
//	//_, _ = s.Raw("CREATE TABLE User(Name text PRIMARY KEY, XXX integer);").Exec()
//	//_, _ = s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
//	//engine.Migrate(&User{})
//	rows, _ := s.Raw("SELECT %v FROM o_user_list", GetAllFieldsAsStringWithTableName(User{}, "o_user_list")).QueryRows()
//	//columns, _ := rows.Columns()
//	//if !reflect.DeepEqual(columns, []string{"name", "username"}) {
//	//	t.Fatal("Failed to migrate table User, got columns", columns)
//	//}
//	for rows.Next() {
//
//		user := User{}
//
//		err := rows.Scan(&user.Uid, &user.Name, &user.Username, &user.Password)
//		if err != nil {
//			// handle this error
//			panic(err)
//		}
//		fmt.Println(user)
//	}
//
//	fmt.Println("ababbbbbbb66666")
//
//}

//func transactionRollback(t *testing.T) {
//	engine := OpenDB(t)
//	defer engine.Close()
//	s := engine.NewSession()
//	_ = s.Model(&User{}).DropTable()
//	_, err := engine.Transaction(func(s *session.Session) (result interface{}, err error) {
//		_ = s.Model(&User{}).CreateTable()
//		_, err = s.Insert(&User{"Tom", 18})
//		return nil, errors.New("Error")
//	})
//	if err == nil || s.HasTable() {
//		t.Fatal("failed to rollback")
//	}
//}
//
//func transactionCommit(t *testing.T) {
//	engine := OpenDB(t)
//	defer engine.Close()
//	s := engine.NewSession()
//	_ = s.Model(&User{}).DropTable()
//	_, err := engine.Transaction(func(s *session.Session) (result interface{}, err error) {
//		_ = s.Model(&User{}).CreateTable()
//		_, err = s.Insert(&User{"Tom", 18})
//		return
//	})
//	u := &User{}
//	_ = s.First(u)
//	if err != nil || u.Name != "Tom" {
//		t.Fatal("failed to commit")
//	}
//}
//
//func TestEngine_Transaction(t *testing.T) {
//	t.Run("rollback", func(t *testing.T) {
//		transactionRollback(t)
//	})
//	t.Run("commit", func(t *testing.T) {
//		transactionCommit(t)
//	})
//}
