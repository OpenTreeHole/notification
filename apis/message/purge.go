package message

import (
	"gorm.io/gorm"
	"notification/config"
	. "notification/models"
	"notification/utils"
	"time"
)

func purgeMessage() error {
	return DB.Transaction(func(tx *gorm.DB) error {
		// delete outdated messages
		result := tx.Exec(
			"DELETE FROM message WHERE created_at < ?",
			time.Now().Add(-time.Hour*24*time.Duration(config.Config.MessagePurgeDays)),
		)
		if result.Error != nil {
			return result.Error
		}

		// delete message
		result = tx.Exec(`
			DELETE message FROM message
			LEFT JOIN message_user ON id = message_id
			WHERE message_id IS NULL
		`)
		if result.Error != nil {
			return result.Error
		}

		// delete message_user
		result = tx.Exec(`
			DELETE message_user FROM message
			LEFT JOIN message_user ON id = message_id
			WHERE id IS NULL
		`)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})
}

func PurgeMessage() {
	ticker := time.NewTicker(time.Hour * 24)
	defer ticker.Stop()
	for range ticker.C {
		err := purgeMessage()
		if err != nil {
			utils.Logger.Error("error purge message: " + err.Error())
		}
	}
}
