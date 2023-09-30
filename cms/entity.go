package cms

type Mock struct {
	tableName    struct{} `pg:"mock"`
	Id           int
	Name         string
	Method       string
	Path         string
	ResponseCode int
	Request      interface{}
	Response     interface{}
	CreatedAt    string
	UpdatedAt    string
}
