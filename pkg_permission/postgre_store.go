package pkg_permission

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/kirigaikabuto/test-api-key/common"
	"log"
	"strconv"
	"strings"
)

var queries = []string{
	`create table if not exists package_permissions (
		id text,
		access_zone text,
		api_key_id text,
		action text,
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

func (s *store) Create(obj *PackagePermission) (*PackagePermission, error) {
	obj.Id = uuid.New().String()
	result, err := s.db.Exec(
		"INSERT INTO package_permissions "+
			"(id, access_zone, api_key_id, action) "+
			"VALUES ($1, $2, $3, $4)",
		obj.Id, obj.AccessZone, obj.ApiKeyId, obj.Action,
	)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrCreatePkgPermissionUnknown
	}
	return obj, nil
}

func (s *store) GetByApiKeyId(apiKeyId string) ([]PackagePermission, error) {
	var objects []PackagePermission
	var values []interface{}
	q := "select " +
		"id, access_zone, api_key_id, action " +
		"from package_permissions where api_key_id = $1"
	values = append(values, apiKeyId)
	rows, err := s.db.Query(q, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		obj := PackagePermission{}
		err = rows.Scan(
			&obj.Id, &obj.AccessZone,
			&obj.ApiKeyId, &obj.Action)
		if err != nil {
			return nil, err
		}
		objects = append(objects, obj)
	}
	return objects, nil
}

func (s *store) List() ([]PackagePermission, error) {
	var objects []PackagePermission
	var values []interface{}
	q := "select " +
		"id, access_zone, api_key_id, action " +
		"from package_permissions where api_key_id = $1"
	rows, err := s.db.Query(q, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		obj := PackagePermission{}
		err = rows.Scan(
			&obj.Id, &obj.AccessZone,
			&obj.ApiKeyId, &obj.Action)
		if err != nil {
			return nil, err
		}
		objects = append(objects, obj)
	}
	return objects, nil
}

func (s *store) GetById(id string) (*PackagePermission, error) {
	obj := &PackagePermission{}
	err := s.db.QueryRow("select id, access_zone, api_key_id, action from package_permissions where id = $1", id).
		Scan(&obj.Id, &obj.AccessZone, &obj.ApiKeyId, &obj.Action)
	if err == sql.ErrNoRows {
		return nil, ErrPkgPermissionNotFound
	} else if err != nil {
		return nil, err
	}
	return obj, nil
}

func (s *store) Update(obj *UpdatePackagePermission) (*PackagePermission, error) {
	q := "update package_permissions set "
	parts := []string{}
	values := []interface{}{}
	cnt := 0
	if obj.ApiKeyId != nil {
		cnt++
		parts = append(parts, "api_key_id = $"+strconv.Itoa(cnt))
		values = append(values, obj.ApiKeyId)
	}
	if obj.AccessZone != nil {
		cnt++
		parts = append(parts, "access_zone = $"+strconv.Itoa(cnt))
		values = append(values, obj.AccessZone)
	}
	if obj.Action != nil {
		cnt++
		parts = append(parts, "action = $"+strconv.Itoa(cnt))
		values = append(values, obj.Action)
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
		return nil, ErrPkgPermissionNotFound
	}
	return s.GetById(obj.Id)
}
