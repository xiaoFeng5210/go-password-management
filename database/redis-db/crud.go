package redisDB

import (
	"context"
	"errors"
	"fmt"

	myCommand "github.com/xiaoFeng5210/go-password-management/command"

	"github.com/fatih/color"
	"github.com/redis/go-redis/v9"
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
	hmap := make(map[string]string)
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

func DeletePassword(account string) error {
	err := Client.Del(context.Background(), account).Err()
	if err != nil {
		color.Red("delete password failed: %v", err)
		return errors.New("delete password failed")
	}
	color.Green("delete password success")
	return nil
}

func GetPassword(account string) error {
	var err error
	var password string
	var remark string
	password, err = Client.HGet(context.Background(), account, "password").Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			color.Yellow("key %s 不存在", account)
			return errors.New("key not found")
		}
	}
	remark, err = Client.HGet(context.Background(), account, "remark").Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			color.Yellow("key %s 不存在", account)
			return errors.New("key not found")
		}
	}
	color.Cyan("password: %s, remark: %s", password, remark)

	myCommand.CopyToClipboard(password)
	if err != nil {
		color.Yellow("copy password failed: %v", err)
		return errors.New("copy password failed")
	}
	color.Green("copy password success")
	return nil
}
