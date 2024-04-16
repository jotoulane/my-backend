package global

import "gorm.io/gorm"

// PageInfo 分页查询信息
type PageInfo struct {
	PageNum   int   `json:"pageNum" binding:"required,min=1"`
	PageSize  int   `json:"pageSize" binding:"required"`
	RecordNum int64 `json:"-"` // 不从请求读入
}

// Pagination 数据库内分页，若pageSize为0，则返回全部
func Pagination(pageInfo *PageInfo) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		limit := pageInfo.PageSize
		offset := (pageInfo.PageNum - 1) * pageInfo.PageSize
		if pageInfo.PageSize == -1 {
			return db
		} else {
			return db.Limit(limit).Offset(offset)
		}
	}
}
