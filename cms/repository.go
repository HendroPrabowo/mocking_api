package cms

import "github.com/go-pg/pg/v10"

type repository struct {
	postgres *pg.DB
}

func newRepository(postgres *pg.DB) repository {
	return repository{
		postgres: postgres,
	}
}

func (r repository) insert(mock Mock) error {
	_, err := r.postgres.Model(&mock).Insert()
	return err
}

func (r repository) get(queryParam MockQueryDto) ([]Mock, error) {
	var mocks []Mock
	offset := queryParam.Page * queryParam.Limit
	query := r.postgres.Model(&mocks).Limit(queryParam.Limit).Offset(offset)
	if queryParam.Name != "" {
		query.Where("name ilike ?", queryParam.Name+"%")
	}
	if queryParam.Method != "" {
		query.Where("method ilike ?", queryParam.Method+"%")
	}
	if queryParam.Path != "" {
		query.Where("path ilike ?", queryParam.Path+"%")
	}
	if queryParam.ResponseCode != 0 {
		query.Where("response_code = ?", queryParam.ResponseCode)
	}
	err := query.Select()
	return mocks, err
}

func (r repository) count(queryParam MockQueryDto) (int, error) {
	var mocks []Mock
	offset := queryParam.Page * queryParam.Limit
	query := r.postgres.Model(&mocks).Limit(queryParam.Limit).Offset(offset)
	if queryParam.Name != "" {
		query.Where("name ilike ?", queryParam.Name+"%")
	}
	if queryParam.Method != "" {
		query.Where("method ilike ?", queryParam.Method+"%")
	}
	if queryParam.Path != "" {
		query.Where("path ilike ?", queryParam.Path+"%")
	}
	if queryParam.ResponseCode != 0 {
		query.Where("response_code = ?", queryParam.ResponseCode)
	}
	count, err := query.Count()
	return count, err
}

func (r repository) getOneByMethodAndPath(method string, path string) (Mock, error) {
	var mock Mock
	err := r.postgres.Model(&mock).Where("method = ?", method).Where("path = ?", path).Limit(1).Select()
	return mock, err
}

func (r repository) update(mock Mock) error {
	_, err := r.postgres.Model(&mock).Column("name", "method", "path", "response_code", "request", "response", "updated_at").WherePK().Update()
	return err
}

func (r repository) delete(id int) error {
	_, err := r.postgres.Model((*Mock)(nil)).Where("id = ?", id).Delete()
	return err
}
