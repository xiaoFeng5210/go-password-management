package redisDB

import (
	"context"
	"errors"
	"fmt"

	"github.com/fatih/color"
)

func AddPassword(account string, password string, remark string) error {
	hmap := map[string]string{
		"account":  account,
		"password": password,
		"remark":   remark,
	}
	err := Client.HMSet(context.Background(), account, hmap).Err()
	if err != nil {
		fmt.Printf("add password failed: %v", err)
		return errors.New("add password failed")
	}

	fmt.Println("add password success")
	return nil
}

func UpdatePassword(account string, password string) error {
	hmap := make(map[string]interface{})
	for field, value := range Client.HGetAll(context.Background(), account).Val() {
		hmap[field] = value
	}

	if password != "" {
		hmap["password"] = password
		err := Client.HMSet(context.Background(), account, hmap).Err()
		if err != nil {
			color.Red("update password failed: %v", err)
			return errors.New("update password failed")
		}
		color.Green("update password success")
		return nil
	} else {
		color.Yellow("password is empty")
		return errors.New("password is empty")
	}
}
