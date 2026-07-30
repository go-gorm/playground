package main

import (
	"database/sql"
	"database/sql/driver"
	"io"
	"sync/atomic"
)

// A database/sql driver that counts prepared statements and their closes.
// database/sql exposes no open-statement count, so counting in the driver is
// the only way to observe a *sql.Stmt that GORM prepared and never closed.

const countingDriverName = "gorm-playground-counting"

var stmtsPrepared, stmtsClosed atomic.Int64

func init() { sql.Register(countingDriverName, countingDriver{}) }

type countingDriver struct{}

func (countingDriver) Open(string) (driver.Conn, error) { return countingConn{}, nil }

type countingConn struct{}

func (countingConn) Prepare(string) (driver.Stmt, error) {
	stmtsPrepared.Add(1)
	return &countingStmt{}, nil
}
func (countingConn) Close() error              { return nil }
func (countingConn) Begin() (driver.Tx, error) { return nil, driver.ErrSkip }

type countingStmt struct{ closed atomic.Bool }

func (s *countingStmt) Close() error {
	if s.closed.CompareAndSwap(false, true) {
		stmtsClosed.Add(1)
	}
	return nil
}
func (*countingStmt) NumInput() int                              { return 0 }
func (*countingStmt) Exec([]driver.Value) (driver.Result, error) { return driver.RowsAffected(0), nil }
func (*countingStmt) Query([]driver.Value) (driver.Rows, error)  { return countingRows{}, nil }

type countingRows struct{}

func (countingRows) Columns() []string         { return nil }
func (countingRows) Close() error              { return nil }
func (countingRows) Next([]driver.Value) error { return io.EOF }
