package orm

import (
	"context"
	"errors"

	"github.com/xuanbo/pig/client/orm/condition"
	"github.com/xuanbo/pig/model"

	"gorm.io/gorm"
)

// API 数据库操作通用API
type API struct {
	*gorm.DB
	dialect Dialect
}

// AutoMigrate 迁移数据
// 使用例子：AutoMigrate(ctx, &email, &user)
func (a *API) AutoMigrate(ctx context.Context, entities ...interface{}) error {
	return a.WithContext(ctx).AutoMigrate(entities)
}

// Insert 创建
// 使用例子：Insert(ctx, &email)
func (a *API) Insert(ctx context.Context, entity interface{}) error {
	return a.WithContext(ctx).Create(entity).Error
}

// Update 更新
// 使用例子：Update(ctx, &email)
func (a *API) Update(ctx context.Context, entity interface{}) error {
	return a.WithContext(ctx).Save(entity).Error
}

// FindByID 主键查询
// 使用例子：FindByID(ctx, "1", &email)
func (a *API) FindByID(ctx context.Context, id string, entity interface{}) error {
	return a.WithContext(ctx).Find(entity, id).Error
}

// FindByIDs 主键批量查询
// 使用例子：FindByIDs(ctx, []string{"1", "2"}, &emails)
func (a *API) FindByIDs(ctx context.Context, ids []string, entities interface{}) error {
	return a.WithContext(ctx).Find(entities, "id IN ?", ids).Error
}

// Query 查询
// 使用例子：
// type EmailCondition struct {
// 		Name string `condition:"name:name;op:like"`
// }
// 使用：Query(ctx, &emailCondition, &emails)
// 使用：Query(ctx, clause, &emails)
func (a *API) Query(ctx context.Context, cond interface{}, entities interface{}) error {
	if cond == nil {
		return a.WithContext(ctx).Find(entities).Error
	}
	var (
		clause condition.Clause
		ok     bool
	)
	if clause, ok = cond.(condition.Clause); !ok {
		clause = condition.ParseCondition(cond)
	}
	s, v := a.dialect.ParseClause(clause)
	if s == "" {
		return a.WithContext(ctx).Find(entities).Error
	}
	switch clause.Type() {
	case "single":
		return a.WithContext(ctx).Find(entities, s, v).Error
	case "combine":
		return a.WithContext(ctx).Where(s, v.([]interface{})...).Find(entities).Error
	default:
		return errors.New("orm: unsupport clause")
	}
}

// Count count
// 使用例子：
// type EmailCondition struct {
// 		Name string `condition:"name:name;op:like"`
// }
// 使用：Count(ctx, &emailCondition, &Email{})
// 使用：Count(ctx, clause, &Email{})
func (a *API) Count(ctx context.Context, cond interface{}, entity interface{}) (int64, error) {
	var (
		total int64
		err   error
	)
	if cond == nil {
		err = a.WithContext(ctx).Model(entity).Count(&total).Error
		return total, err
	}
	var (
		clause condition.Clause
		ok     bool
	)
	if clause, ok = cond.(condition.Clause); !ok {
		clause = condition.ParseCondition(cond)
	}
	s, v := a.dialect.ParseClause(clause)
	if s == "" {
		err = a.WithContext(ctx).Model(entity).Count(&total).Error
		return total, err
	}
	switch clause.Type() {
	case "single":
		err = a.WithContext(ctx).Model(entity).Where(s, v).Count(&total).Error
	case "combine":
		err = a.WithContext(ctx).Model(entity).Where(s, v.([]interface{})...).Count(&total).Error
	default:
		err = errors.New("orm: unsupport clause")
	}
	return total, err
}

// Page 分页查询
// 使用例子：
// type EmailCondition struct {
// 		Name string `condition:"name:name;op:like"`
// }
// 使用：Page(ctx, &emailCondition, &pagination, &emails)
// 使用：Page(ctx, clause, &pagination, &emails)
func (a *API) Page(ctx context.Context, cond interface{}, pagination *model.Pagination, entities interface{}) error {
	var (
		offset = pagination.Offset
		size   = pagination.Size
		total  int64
		err    error
	)
	if cond == nil {
		if err = a.WithContext(ctx).Model(entities).Count(&total).Error; err != nil {
			return err
		}
		if total == 0 {
			pagination.Total = 0
			return nil
		}
		if err := a.WithContext(ctx).Offset(offset).Limit(size).Find(entities).Error; err != nil {
			return err
		}
		pagination.Set(total, entities)
		return nil
	}
	var (
		clause condition.Clause
		ok     bool
	)
	if clause, ok = cond.(condition.Clause); !ok {
		clause = condition.ParseCondition(cond)
	}
	s, v := a.dialect.ParseClause(clause)
	if s == "" {
		if err = a.WithContext(ctx).Model(entities).Count(&total).Error; err != nil {
			return err
		}
		if total == 0 {
			pagination.Total = 0
			return nil
		}
		if err = a.WithContext(ctx).Offset(offset).Limit(size).Find(entities).Error; err != nil {
			return err
		}
		pagination.Set(total, entities)
		return nil
	}
	switch clause.Type() {
	case "single":
		if err = a.WithContext(ctx).Model(entities).Where(s, v).Count(&total).Error; err != nil {
			return err
		}
		if total == 0 {
			pagination.Total = 0
			return nil
		}
		if err = a.WithContext(ctx).Where(s, v).Offset(offset).Limit(size).Find(entities).Error; err != nil {
			return err
		}
		pagination.Set(total, entities)
		return nil
	case "combine":
		args := v.([]interface{})
		if err = a.WithContext(ctx).Model(entities).Where(s, args...).Count(&total).Error; err != nil {
			return err
		}
		if total == 0 {
			pagination.Total = 0
			return nil
		}
		if err = a.WithContext(ctx).Where(s, args...).Offset(offset).Limit(size).Find(entities).Error; err != nil {
			return err
		}
		pagination.Set(total, entities)
		return nil
	default:
		return errors.New("orm: unsupport clause")
	}
}

// DeleteByID 主键删除
// 使用例子：DeleteByID(ctx, "1", &Email{})
func (a *API) DeleteByID(ctx context.Context, id string, entity interface{}) error {
	return a.WithContext(ctx).Delete(entity, id).Error
}

// DeleteByIDs 主键批量查询
// 使用例子：DeleteByIDs(ctx, []string{"1", "2"}, &Email{})
func (a *API) DeleteByIDs(ctx context.Context, ids []string, entities interface{}) error {
	return a.WithContext(ctx).Delete(entities, "id IN ?", ids).Error
}
