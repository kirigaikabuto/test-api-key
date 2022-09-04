package endpoints_permission

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/kirigaikabuto/test-api-key/common"
	"log"
	"strconv"
	"strings"
)

var queries = []string{
	`create table if not exists endpoints_permissions (
		id text,
		api_key_id text,
		endpoints json,
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

func (s *store) Create(obj *EndpointsPermission) (*EndpointsPermission, error) {
	obj.Id = uuid.New().String()
	jsonData, err := json.Marshal(obj.Endpoints)
	if err != nil {
		return nil, err
	}
	result, err := s.db.Exec(
		"INSERT INTO endpoints_permissions "+
			"(id, api_key_id, endpoints) "+
			"VALUES ($1, $2, $3)",
		obj.Id, obj.ApiKeyId, jsonData,
	)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrCreateEndpointsPermissionUnknown
	}
	return obj, nil
}

func (s *store) Get(id string) (*EndpointsPermission, error) {
	obj := &EndpointsPermission{}
	var jsonData string
	err := s.db.QueryRow("select id, api_key_id, endpoints from endpoints_permissions where id = $1", id).
		Scan(&obj.Id, &obj.ApiKeyId, &jsonData)
	if err == sql.ErrNoRows {
		return nil, ErrEndpointsPermissionNotFound
	} else if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(jsonData), &obj.Endpoints)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *store) GetByApiKeyId(keyId string) (*EndpointsPermission, error) {
	obj := &EndpointsPermission{}
	var jsonData string
	err := s.db.QueryRow("select id, api_key_id, endpoints from endpoints_permissions where api_key_id = $1", keyId).
		Scan(&obj.Id, &obj.ApiKeyId, &jsonData)
	if err == sql.ErrNoRows {
		return nil, ErrEndpointsPermissionNotFound
	} else if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(jsonData), &obj.Endpoints)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *store) List() ([]EndpointsPermission, error) {
	var objects []EndpointsPermission
	var values []interface{}
	q := "select " +
		"id, api_key_id, endpoints " +
		"from endpoints_permissions"
	rows, err := s.db.Query(q, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var jsonData string
		obj := EndpointsPermission{}
		err = rows.Scan(
			&obj.Id, &obj.ApiKeyId,
			&jsonData)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(jsonData), &obj.Endpoints)
		if err != nil {
			return nil, err
		}
		objects = append(objects, obj)
	}
	return objects, nil
}

func (s *store) Update(obj *EndpointsPermissionUpdate) (*EndpointsPermission, error) {
	q := "update endpoints_permissions set "
	parts := []string{}
	values := []interface{}{}
	cnt := 0
	if obj.ApiKeyId != nil {
		cnt++
		parts = append(parts, "api_key_id = $"+strconv.Itoa(cnt))
		values = append(values, obj.ApiKeyId)
	}
	if obj.Endpoints != nil {
		cnt++
		jsonData, err := json.Marshal(obj.Endpoints)
		if err != nil {
			return nil, err
		}
		parts = append(parts, "endpoints = $"+strconv.Itoa(cnt))
		values = append(values, jsonData)
	}
	if len(parts) <= 0 {
		return nil, ErrNothingToUpdate
	}
	cnt++
	q = q + strings.Join(parts, " , ") + " WHERE id = $" + strconv.Itoa(cnt)
	values = append(values, obj.Id)
	result, err := s.db.Exec(q, values...)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrEndpointsPermissionNotFound
	}
	return s.Get(obj.Id)
}
