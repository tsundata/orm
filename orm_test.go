package orm

import (
	_ "github.com/mattn/go-sqlite3"
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

func TestNewEngine(t *testing.T) {
	e := OpenDB(t)
	defer e.Close()
}
