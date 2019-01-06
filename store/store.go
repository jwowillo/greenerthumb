package store

import (
	"database/sql"
	"encoding/json"

	_ "github.com/mattn/go-sqlite3" // Driver

	"github.com/jwowillo/greenerthumb"
)

const (
	create = "CREATE TABLE IF NOT EXISTS command (name TEXT primary key, message TEXT)"
	read   = "SELECT message FROM command"
	write  = "REPLACE INTO command (name, message) VALUES (?, ?)"
)

// SQLITEStore stores messages in SQLITE.
type SQLITEStore struct {
	db          *sql.DB
	read, write *sql.Stmt
}

// NewSQLITEStore connects to a SQLITE store at the path.
func NewSQLITEStore(path string) (*SQLITEStore, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	create, err := db.Prepare(create)
	if err != nil {
		return nil, err
	}
	defer create.Close()

	if _, err := create.Exec(); err != nil {
		return nil, err
	}

	write, err := db.Prepare(write)
	if err != nil {
		return nil, err
	}
	read, err := db.Prepare(read)
	if err != nil {
		return nil, err
	}

	return &SQLITEStore{db: db, read: read, write: write}, nil
}

// Write the message to the store.
//
// Returns an error if the message couldn't be written.
func (s *SQLITEStore) Write(msg string) error {
	var x map[string]interface{}
	json.Unmarshal([]byte(msg), &x)

	rawHeader, ok := x["Header"]
	if !ok {
		return greenerthumb.KeyError{Object: x, MissingKey: "Header"}
	}

	xHeader, ok := rawHeader.(map[string]interface{})
	if !ok {
		return greenerthumb.TypeError{
			Value: rawHeader,
			Type:  "map[string]interface{}"}
	}

	rawName, ok := xHeader["Name"]
	if !ok {
		return greenerthumb.KeyError{
			Object:     x,
			MissingKey: "Header/Name"}
	}
	name, ok := rawName.(string)
	if !ok {
		return greenerthumb.TypeError{Value: rawName, Type: "string"}
	}

	_, err := s.write.Exec(name, msg)
	return err
}

// Read all the messages from the store.
//
// Returns an error if the messages couldn't be read.
func (s *SQLITEStore) Read() ([]string, error) {
	rs, err := s.read.Query()
	if err != nil {
		return nil, err
	}
	defer rs.Close()

	var msgs []string
	var msg string
	for rs.Next() {
		if err := rs.Scan(&msg); err != nil {
			return nil, err
		}
		msgs = append(msgs, msg)
	}
	return msgs, rs.Err()
}

// Close the connection.
func (s *SQLITEStore) Close() error {
	s.read.Close()
	s.write.Close()
	return s.db.Close()
}
