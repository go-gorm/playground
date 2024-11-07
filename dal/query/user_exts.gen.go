// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"playground/model"
)

func newUserExt(db *gorm.DB, opts ...gen.DOOption) userExt {
	_userExt := userExt{}

	_userExt.userExtDo.UseDB(db, opts...)
	_userExt.userExtDo.UseModel(&model.UserExt{})

	tableName := _userExt.userExtDo.TableName()
	_userExt.ALL = field.NewAsterisk(tableName)
	_userExt.ID = field.NewUint(tableName, "id")
	_userExt.CreatedAt = field.NewTime(tableName, "created_at")
	_userExt.UpdatedAt = field.NewTime(tableName, "updated_at")
	_userExt.DeletedAt = field.NewField(tableName, "deleted_at")
	_userExt.UserID = field.NewUint(tableName, "user_id")
	_userExt.UserInfo = userExtBelongsToUserInfo{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("UserInfo", "model.User"),
		UserExtInfo: struct {
			field.RelationField
			UserInfo struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("UserInfo.UserExtInfo", "model.UserExt"),
			UserInfo: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("UserInfo.UserExtInfo.UserInfo", "model.User"),
			},
		},
		UserAccountRelationInfo: struct {
			field.RelationField
			UserInfo struct {
				field.RelationField
			}
			AccountInfo struct {
				field.RelationField
				CompanyInfo struct {
					field.RelationField
				}
			}
		}{
			RelationField: field.NewRelation("UserInfo.UserAccountRelationInfo", "model.UserAccountRelation"),
			UserInfo: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("UserInfo.UserAccountRelationInfo.UserInfo", "model.User"),
			},
			AccountInfo: struct {
				field.RelationField
				CompanyInfo struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("UserInfo.UserAccountRelationInfo.AccountInfo", "model.Account"),
				CompanyInfo: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("UserInfo.UserAccountRelationInfo.AccountInfo.CompanyInfo", "model.Company"),
				},
			},
		},
	}

	_userExt.fillFieldMap()

	return _userExt
}

type userExt struct {
	userExtDo userExtDo

	ALL       field.Asterisk
	ID        field.Uint
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	UserID    field.Uint
	UserInfo  userExtBelongsToUserInfo

	fieldMap map[string]field.Expr
}

func (u userExt) Table(newTableName string) *userExt {
	u.userExtDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userExt) As(alias string) *userExt {
	u.userExtDo.DO = *(u.userExtDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userExt) updateTableName(table string) *userExt {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewUint(table, "id")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")
	u.DeletedAt = field.NewField(table, "deleted_at")
	u.UserID = field.NewUint(table, "user_id")

	u.fillFieldMap()

	return u
}

func (u *userExt) WithContext(ctx context.Context) *userExtDo { return u.userExtDo.WithContext(ctx) }

func (u userExt) TableName() string { return u.userExtDo.TableName() }

func (u userExt) Alias() string { return u.userExtDo.Alias() }

func (u userExt) Columns(cols ...field.Expr) gen.Columns { return u.userExtDo.Columns(cols...) }

func (u *userExt) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userExt) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 6)
	u.fieldMap["id"] = u.ID
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
	u.fieldMap["user_id"] = u.UserID

}

func (u userExt) clone(db *gorm.DB) userExt {
	u.userExtDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userExt) replaceDB(db *gorm.DB) userExt {
	u.userExtDo.ReplaceDB(db)
	return u
}

type userExtBelongsToUserInfo struct {
	db *gorm.DB

	field.RelationField

	UserExtInfo struct {
		field.RelationField
		UserInfo struct {
			field.RelationField
		}
	}
	UserAccountRelationInfo struct {
		field.RelationField
		UserInfo struct {
			field.RelationField
		}
		AccountInfo struct {
			field.RelationField
			CompanyInfo struct {
				field.RelationField
			}
		}
	}
}

func (a userExtBelongsToUserInfo) Where(conds ...field.Expr) *userExtBelongsToUserInfo {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a userExtBelongsToUserInfo) WithContext(ctx context.Context) *userExtBelongsToUserInfo {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a userExtBelongsToUserInfo) Session(session *gorm.Session) *userExtBelongsToUserInfo {
	a.db = a.db.Session(session)
	return &a
}

func (a userExtBelongsToUserInfo) Model(m *model.UserExt) *userExtBelongsToUserInfoTx {
	return &userExtBelongsToUserInfoTx{a.db.Model(m).Association(a.Name())}
}

type userExtBelongsToUserInfoTx struct{ tx *gorm.Association }

func (a userExtBelongsToUserInfoTx) Find() (result *model.User, err error) {
	return result, a.tx.Find(&result)
}

func (a userExtBelongsToUserInfoTx) Append(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a userExtBelongsToUserInfoTx) Replace(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a userExtBelongsToUserInfoTx) Delete(values ...*model.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a userExtBelongsToUserInfoTx) Clear() error {
	return a.tx.Clear()
}

func (a userExtBelongsToUserInfoTx) Count() int64 {
	return a.tx.Count()
}

type userExtDo struct{ gen.DO }

func (u userExtDo) Debug() *userExtDo {
	return u.withDO(u.DO.Debug())
}

func (u userExtDo) WithContext(ctx context.Context) *userExtDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userExtDo) ReadDB() *userExtDo {
	return u.Clauses(dbresolver.Read)
}

func (u userExtDo) WriteDB() *userExtDo {
	return u.Clauses(dbresolver.Write)
}

func (u userExtDo) Session(config *gorm.Session) *userExtDo {
	return u.withDO(u.DO.Session(config))
}

func (u userExtDo) Clauses(conds ...clause.Expression) *userExtDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userExtDo) Returning(value interface{}, columns ...string) *userExtDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userExtDo) Not(conds ...gen.Condition) *userExtDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userExtDo) Or(conds ...gen.Condition) *userExtDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userExtDo) Select(conds ...field.Expr) *userExtDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userExtDo) Where(conds ...gen.Condition) *userExtDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userExtDo) Order(conds ...field.Expr) *userExtDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userExtDo) Distinct(cols ...field.Expr) *userExtDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userExtDo) Omit(cols ...field.Expr) *userExtDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userExtDo) Join(table schema.Tabler, on ...field.Expr) *userExtDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userExtDo) LeftJoin(table schema.Tabler, on ...field.Expr) *userExtDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userExtDo) RightJoin(table schema.Tabler, on ...field.Expr) *userExtDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userExtDo) Group(cols ...field.Expr) *userExtDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userExtDo) Having(conds ...gen.Condition) *userExtDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userExtDo) Limit(limit int) *userExtDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userExtDo) Offset(offset int) *userExtDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userExtDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *userExtDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userExtDo) Unscoped() *userExtDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userExtDo) Create(values ...*model.UserExt) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userExtDo) CreateInBatches(values []*model.UserExt, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userExtDo) Save(values ...*model.UserExt) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userExtDo) First() (*model.UserExt, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserExt), nil
	}
}

func (u userExtDo) Take() (*model.UserExt, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserExt), nil
	}
}

func (u userExtDo) Last() (*model.UserExt, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserExt), nil
	}
}

func (u userExtDo) Find() ([]*model.UserExt, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserExt), err
}

func (u userExtDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserExt, err error) {
	buf := make([]*model.UserExt, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userExtDo) FindInBatches(result *[]*model.UserExt, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userExtDo) Attrs(attrs ...field.AssignExpr) *userExtDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userExtDo) Assign(attrs ...field.AssignExpr) *userExtDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userExtDo) Joins(fields ...field.RelationField) *userExtDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userExtDo) Preload(fields ...field.RelationField) *userExtDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userExtDo) FirstOrInit() (*model.UserExt, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserExt), nil
	}
}

func (u userExtDo) FirstOrCreate() (*model.UserExt, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserExt), nil
	}
}

func (u userExtDo) FindByPage(offset int, limit int) (result []*model.UserExt, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userExtDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userExtDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userExtDo) Delete(models ...*model.UserExt) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userExtDo) withDO(do gen.Dao) *userExtDo {
	u.DO = *do.(*gen.DO)
	return u
}
