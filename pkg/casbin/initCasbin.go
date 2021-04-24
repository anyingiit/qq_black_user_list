package casbin

import (
	"github.com/casbin/casbin/v2"
)

var Casbin *casbin.Enforcer

func InitCasbin() (err error) {
	if Casbin, err = casbin.NewEnforcer("conf/casbin/casbin_model.conf", "conf/casbin/casbin_rabc.csv"); err != nil {
		return err
	}
	if err = Casbin.LoadModel(); err != nil {
		return err
	}
	if err = Casbin.LoadPolicy(); err != nil {
		return err
	}
	return nil
}
