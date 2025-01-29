package cms

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type repository struct {
	sqliteDB *gorm.DB
}

func newRepository(sqliteDB *gorm.DB) repository {
	if err := sqliteDB.AutoMigrate(&Mock{}); err != nil {
		log.Fatal(err)
	}
	return repository{
		sqliteDB: sqliteDB,
	}
}

func (r repository) insert(mock Mock) error {
	return r.sqliteDB.Create(&mock).Error
}

func (r repository) get(queryParam MockQueryDto) ([]Mock, error) {
	var mocks []Mock
	offset := queryParam.Page * queryParam.Limit
	query := r.sqliteDB.Model(&Mock{}).Limit(queryParam.Limit).Offset(offset)
	if queryParam.Name != "" {
		query = query.Where("name LIKE ?", queryParam.Name+"%")
	}
	if queryParam.Method != "" {
		query = query.Where("method LIKE ?", queryParam.Method+"%")
	}
	if queryParam.Path != "" {
		query = query.Where("path LIKE ?", queryParam.Path+"%")
	}
	if queryParam.ResponseCode != 0 {
		query = query.Where("response_code = ?", queryParam.ResponseCode)
	}
	err := query.Find(&mocks).Error
	return mocks, err
}

func (r repository) count(queryParam MockQueryDto) (int, error) {
	var count int64
	query := r.sqliteDB.Model(&Mock{})
	if queryParam.Name != "" {
		query = query.Where("name LIKE ?", queryParam.Name+"%")
	}
	if queryParam.Method != "" {
		query = query.Where("method LIKE ?", queryParam.Method+"%")
	}
	if queryParam.Path != "" {
		query = query.Where("path LIKE ?", queryParam.Path+"%")
	}
	if queryParam.ResponseCode != 0 {
		query = query.Where("response_code = ?", queryParam.ResponseCode)
	}
	err := query.Count(&count).Error
	return int(count), err
}

func (r repository) getOneByMethodAndPath(method string, path string) (mock Mock, err error) {
	err = r.sqliteDB.Where("method = ? AND path = ?", method, path).First(&mock).Error
	return mock, err
}

func (r repository) update(mock Mock) error {
	var entity Mock
	if err := r.sqliteDB.First(&entity, "id = ?", mock.ID).Error; err != nil {
		return err
	}

	entity.Name = mock.Name
	entity.Method = mock.Method
	entity.Path = mock.Path
	entity.ResponseCode = mock.ResponseCode
	entity.Request = mock.Request
	entity.Response = mock.Response
	return r.sqliteDB.Save(entity).Error
}

func (r repository) delete(id int) error {
	return r.sqliteDB.Unscoped().Delete(&Mock{}, id).Error
}
