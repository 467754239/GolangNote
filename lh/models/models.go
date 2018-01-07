package models

import (
	"fmt"

	"github.com/467754239/GolangNote/lh/common"
	"github.com/467754239/GolangNote/lh/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type NewMysql struct {
	conn *xorm.Engine
}

func (this *NewMysql) Connect() error {
	engine, err := xorm.NewEngine(
		config.VKLH_DB_TYPE,
		config.DatabaseDialString())
	if err != nil {
		return fmt.Errorf("models.init(fail to conntect database): %v", err)
	}
	this.conn = engine
	return nil
}

/*
func (this *NewMysql) Sync() error {
	err := this.conn.Sync2(new(UserAudit))
	return err
}
*/

func (this *NewMysql) Insert(user_audit common.UserAudit) (int64, error) {
	affected, err := this.conn.Insert(&user_audit)
	return affected, err

}

func NewConnect() (*NewMysql, error) {
	db := &NewMysql{}
	if err := db.Connect(); err != nil {
		return nil, err
	}
	return db, nil
}
