package audit

import (
	"errors"
	"os"

	"github.com/467754239/GolangNote/lh/config"
	"github.com/467754239/GolangNote/lh/utils"
)

func Authentication() error {
	// 1. 验证主机名
	name, err := os.Hostname()
	if err != nil {
		return err
	}
	if name != config.TIGER_SERVER_HOSTNAME {
		return errors.New("hostname verify failed.")
	}

	// 2. 验证IP地址
	addrpool, err := utils.LocalAddrpool()
	if err != nil {
		return err
	}
	ok := utils.IsValueInList(config.TIGER_SERVER_IPADDR, addrpool)
	if !ok {
		return errors.New("addr verify failed.")
	}
	return nil

}
