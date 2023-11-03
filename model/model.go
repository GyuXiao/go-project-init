package model

import (
	"GyuBlog/global"
	"GyuBlog/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

// 每张表的公共字段

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

// NewDBEngine
/**
 * @Description: 创建 DB 实例，同时增加 gorm 库的引入和 MySql 驱动库的初始化
 * @param databaseSetting
 * @return *gorm.DB
 * @return error
 */
func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf(s,
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	))

	// 这里只是打印了 err，没有返回 err，不知是否正确
	if err != nil {
		log.Printf("init mysql failed, err: %v", err)
	}

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}

	db.SingularTable(true)
	// 注册回调行为
	db.Callback().Create().Replace("gorm:update_time-stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time-stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}

// 采用 model callback 的方式去处理表的公共字段

// 新增行为回调
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		// 通过 `scope.FieldByName` 判断当前是否包含所需的字段，比如"CreatedOn"
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			// 判断 `Field.IsBlank` 判断该字段的值是否为空
			if createTimeField.IsBlank {
				_ = createTimeField.Set(nowTime)
			}
		}
		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				_ = modifyTimeField.Set(nowTime)
			}
		}
	}
}

// 更新行为回调
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	// 通过 `scope.Get()` 去获取当前设置了标识为 `gorm:update_column` 的字段属性
	// 如果不存在，则将会在更新回调内设置默认字段 ModifiedOn 的值为当前的时间戳。
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

// 删除行为回调
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}
		// 判断是否存在 DeletedOn 和 IsDel 字段，
		// 若存在则调整为执行 UPDATE 操作进行软删除（修改 DeletedOn 和 IsDel 的值），
		// 否则执行 DELETE 进行硬删除（多少觉得有点问题，delete 语句是删了整张表？）
		deleteOnField, hasDeleteOnField := scope.FieldByName("DeleteOn")
		isDelField, hasIsDelField := scope.FieldByName("IsDel")
		if !scope.Search.Unscoped && hasDeleteOnField && hasIsDelField {
			now := time.Now().Unix()
			// 调用 scope.QuotedTableName 方法获取当前所引用的表名，
			// 并调用一系列方法针对 SQL 语句的组成部分进行处理，转移，参数设置
			// 最后调用 scope.CombinedConditionSql 方法完成 SQL 语句的组装
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v, %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deleteOnField.DBName),
				scope.AddToVars(now),
				scope.Quote(isDelField.DBName),
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}

func CloseDBEngine() {
	_ = global.DBEngine.Close()
}
