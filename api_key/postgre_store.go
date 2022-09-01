package api_key

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/kirigaikabuto/test-api-key/common"
	_ "github.com/lib/pq"
	"log"
)

var queries = []string{
	`create table if not exists api_keys (
		id text,
		key text,
		name text,
		primary key(id)
	);`,
}

type store struct {
	db *sql.DB
}

func NewPostgresStore(cfg common.PostgresConfig) (Store, error) {
	db, err := common.GetDbConn(common.GetConnString(cfg))
	if err != nil {
		return nil, err
	}
	for _, q := range queries {
		_, err = db.Exec(q)
		if err != nil {
			log.Println(err)
		}
	}
	db.SetMaxOpenConns(10)
	s := &store{db: db}
	return s, nil
}

func (s *store) Create(obj *ApiKey) (*ApiKey, error) {
	obj.Id = uuid.New().String()
	obj.Key = uuid.New().String()
	result, err := s.db.Exec(
		"INSERT INTO api_keys "+
			"(id, key, name) "+
			"VALUES ($1, $2, $3)",
		obj.Id, obj.Key, obj.Name,
	)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrCreateApiKeyUnknown
	}
	return obj, nil
}

func (s *store) Get(id string) (*ApiKey, error) {
	obj := &ApiKey{}
	err := s.db.QueryRow("select id, key, name from api_keys where id = $1", id).
		Scan(&obj.Id, &obj.Key, &obj.Name)
	if err == sql.ErrNoRows {
		return nil, ErrApiKeyNotFound
	} else if err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *store) GetByKey(key string) (*ApiKey, error) {
	obj := &ApiKey{}
	err := s.db.QueryRow("select id, key, name from api_keys where key = $1", key).
		Scan(&obj.Id, &obj.Key, &obj.Name)
	if err == sql.ErrNoRows {
		return nil, ErrApiKeyNotFound
	} else if err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *store) List() ([]ApiKey, error) {
	var objects []ApiKey
	var values []interface{}
	q := "select " +
		"id, key, name " +
		"from api_keys"
	rows, err := s.db.Query(q, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		obj := ApiKey{}
		err = rows.Scan(
			&obj.Id, &obj.Key,
			&obj.Name)
		if err != nil {
			return nil, err
		}
		objects = append(objects, obj)
	}
	return objects, nil
}
