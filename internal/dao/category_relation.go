// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"myblog-gf/internal/dao/internal"
)

// internalCategoryRelationDao is internal type for wrapping internal DAO implements.
type internalCategoryRelationDao = *internal.CategoryRelationDao

// categoryRelationDao is the data access object for table category_relation.
// You can define custom methods on it to extend its functionality as you wish.
type categoryRelationDao struct {
	internalCategoryRelationDao
}

var (
	// CategoryRelation is globally public accessible object for table category_relation operations.
	CategoryRelation = categoryRelationDao{
		internal.NewCategoryRelationDao(),
	}
)

// Fill with you ideas below.
