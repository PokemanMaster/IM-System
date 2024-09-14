package models

import (
	"IMProject/utils"
	"context"

	"time"
)

func SetUserOnlineInfo(key string, val []byte, timeTTL time.Duration) {
	ctx := context.Background()
	utils.Red.Set(ctx, key, val, timeTTL)
}
