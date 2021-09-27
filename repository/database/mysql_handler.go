package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"gobackend/core/config"
	"gobackend/core/errors"
)

type SQLHandler struct {
	db *gorm.DB
}

var db *gorm.DB

func init() {
	db = newDB()
}

func newDB() *gorm.DB {
	config := config.GetDBConfig()

	USER := config.User
	PASS := config.Pass
	PROTOCOL := config.Protocol
	DATABASE := config.Database
	dsn := USER + ":" + PASS + "@" + PROTOCOL + "/" + DATABASE + "?parseTime=true"
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	log.Print("DB connected !")
	return _db
}

func getDB() *gorm.DB {
	return db
}

func NewSQLHandler() *SQLHandler {
	sqlHandler := new(SQLHandler)
	sqlHandler.db = getDB()
	return sqlHandler
}


func (h *SQLHandler) Find(model interface{}) error {
	err := h.db.Find(model).Error
  return BuildDBError(err,"repo.Find")
}

func (h *SQLHandler) FindByModel(m interface{},q interface{}) error {
  err := h.db.Where(q).First(m).Error
  return BuildDBError(err,"repo.FindByModel")
}

func (h *SQLHandler) FindWithPreload(model interface{}) error {
	err := h.db.Preload(clause.Associations).First(model).Error;
  return BuildDBError(err,"repo.FindWithPreload")
}

//FindByConditionWithPreload with preload
func (h *SQLHandler) FindByConditionWithPreload(model interface{}, c string, p []interface{}) error {
	err := h.db.Preload(clause.Associations).Where(c, p...).First(model).Error
  return BuildDBError(err,"repo.FindByConditionWithPreload")
}

//FindAll
func (repo *Repository) FindAllWithPreload(model interface{}) (error) {
  err := repo.db.Set("gorm:auto_preload", true).Find(model).Error
  return BuildDBError(err,"repo.Find")
}

func (h *SQLHandler) FindAllByConditionWithPreload(model interface{}, c string, p []interface{}, offset int, limit int) ( error) {
	err := h.db.Limit(limit).Offset(offset).Preload(clause.Associations).Where(c, p...).Find(model).Error
  return BuildDBError(err,"repo.FindAllByConditionWithPreload")
}

func (h *SQLHandler) FindAllWithPreloadAndJoin(model interface{}, c string, p []interface{}, offset int, limit int, joins string) (error) {
  err := h.db.Joins(joins).Limit(limit).Offset(offset).Preload(clause.Associations).Where(c, p...).Find(model).Error
  return BuildDBError(err,"repo.FindAllWithPreloadAndJoin")
}



func BuildDBError (err error, methodName string) (error) {
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.Wrapf(err, "record is not found, method: "+ methodName ).NotFound()
		}
		return errors.Wrapf(err, "database error, method: "+ methodName).DatabaseError()
	}
	return nil
}




