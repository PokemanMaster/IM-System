package models

import (
	"IMProject/utils"
	"fmt"
)

func Migration() {
	// 自动迁移模式
	err := utils.DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&UserBasic{},
			&Message{},
			&GroupBasic{},
			&Contact{},
			&Community{},
		)
	if err != nil {
		fmt.Println("err", err)
	}
}
