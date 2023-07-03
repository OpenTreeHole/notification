package token

import (
	"context"
	. "notification/models"
	"time"
)

func deleteExpiredTokens(context context.Context) {
	ticker := time.NewTicker(time.Hour)
	for {
		select {
		case <-ticker.C:
			// delete tokens created_at more than a month ago
			DB.Where("created_at < ?", time.Now().AddDate(0, -1, 0)).Delete(PushToken{})
		case <-context.Done():
			return
		}
	}
}
