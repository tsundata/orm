package orm

import (
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sysatom/orm/session"
	"testing"
)

func OpenDB(t *testing.T) *Engine {
	t.Helper()
	e, err := NewEngine("sqlite3", "demo.db")
	if err != nil {
		t.Fatal("failed to connect", err)
	}
	return e
}

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

func TestNewEngine(t *testing.T) {
	e := OpenDB(t)
	defer e.Close()
}

func TestEngine_Transaction(t *testing.T) {
	t.Run("rollback", func(t *testing.T) {
		transactionRollback(t)
	})
	t.Run("commit", func(t *testing.T) {
		transactionCommit(t)
	})
}

func transactionRollback(t *testing.T) {
	e := OpenDB(t)
	defer e.Close()
	s := e.NewSession()
	_ = s.Model(&User{}).DropTable()
	_, err := e.Transaction(func(s *session.Session) (result interface{}, err error) {
		_ = s.Model(&User{}).CreateTable()
		_, err = s.Insert(&User{"Tom", 18})
		return nil, errors.New("error")
	})
	if err == nil || s.HasTable() {
		t.Fatal("failed to rollback")
	}
}
func transactionCommit(t *testing.T) {
	e := OpenDB(t)
	defer e.Close()
	s := e.NewSession()
	_ = s.Model(&User{}).DropTable()
	_, err := e.Transaction(func(s *session.Session) (result interface{}, err error) {
		_ = s.Model(&User{}).CreateTable()
		_, err = s.Insert(&User{"Tom", 18})
		return
	})
	u := &User{}
	_ = s.First(u)
	if err != nil || u.Name != "Tom" {
		t.Fatal("failed to commit")
	}
}
