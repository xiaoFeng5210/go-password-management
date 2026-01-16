package redisDB

import (
	"context"
	"errors"
	"fmt"
)

func AddPassword(account string, password string, remark string) error {
	hmap := map[string]string{
		"account":  account,
		"password": password,
		"remark":   remark,
	}
	err := Client.HMSet(context.Background(), account, hmap)
	if err != nil {
		fmt.Println("add password failed", err)
		return errors.New("add password failed")
	}

	fmt.Println("add password success")
	return nil
}
