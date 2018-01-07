package config

import (
	"fmt"
)

const (
	VKLH_ASSETS_PROD_URL string = "xxxxxxxxxxxxxxxxxxxxxxxxxx"
)

const (
	VKLH_CLIENT_VER_MAJOR uint8 = 1
	VKLH_CLIENT_VER_MINOR uint8 = 0
)

const (
	TIGER_SERVER_IPADDR   string = "xxxxxx"
	TIGER_SERVER_HOSTNAME string = "xxxxx"
)

const (
	VKLH_AUDIT_PROD_ADDR string = "127.0.0.1"
	VKLH_AUDIT_PROD_PORT int    = 9999
)

const (
	VKLH_DB_TYPE            string = "mysql"
	VKLH_DB_SERVER_HOST     string = "127.0.0.1"
	VKLH_DB_SERVER_PORT     int    = 3306
	VKLH_DB_SERVER_USER     string = "xxx"
	VKLH_DB_SERVER_PASS     string = "xxx"
	VKLH_DB_SERVER_DATABASE string = "go_opsweb"
)

func DatabaseDialString() string {
	return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8",
		VKLH_DB_SERVER_USER,
		VKLH_DB_SERVER_PASS,
		VKLH_DB_SERVER_HOST,
		VKLH_DB_SERVER_PORT,
		VKLH_DB_SERVER_DATABASE)
}
