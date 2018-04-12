package store

import (
	"database/sql"
	"io/ioutil"

	_ "github.com/lib/pq"

	"github.com/rs/zerolog/log"
)

type Store struct {
	ConnStr string
	db      *sql.DB
}

func (s *Store) Connect() {
	db, err := sql.Open("postgres", s.ConnStr)
	if err != nil {
		log.Error().Err(err).Msg("could not connect to postgres")
	}
	s.db = db

	err = db.Ping()
	if err != nil {
		log.Error().Err(err).Msg("could not ping postgres")
	}
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) Setup() error {
	_, err := s.db.Exec("CREATE TABLE files ( name varchar(20), data BYTEA )")
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) Read(filename string) ([]byte, error) {
	var data []byte
	err := s.db.QueryRow("SELECT data FROM files WHERE name = $1", filename).Scan(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *Store) Save(filename string, path string) error {
	// read file
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// prepare statement
	stmt, err := s.db.Prepare("INSERT INTO files (name, data) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(filename, dat)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) List() ([]string, error) {
	entries := make([]string, 0)

	rows, err := s.db.Query("SELECT name FROM files")
	if err != nil {
		log.Error().Err(err).Msg("could not fetch files")
		return entries, err
	}
	defer rows.Close()

	var (
		name string
	)
	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			log.Error().Err(err)
			return entries, err
		}
		entries = append(entries, name)
	}
	err = rows.Err()
	if err != nil {
		log.Error().Err(err)
		return entries, err
	}
	return entries, nil
}

func (s *Store) Delete(filename string) error {
	stmt, err := s.db.Prepare("DELETE FROM files WHERE name=$1")
	if err != nil {
		log.Error().Err(err).Msg("could not prepare delete statement")
		return err
	}
	_, err = stmt.Exec(filename)
	if err != nil {
		log.Error().Err(err).Msg("could not delete file")
		return err
	}
	return nil
}
