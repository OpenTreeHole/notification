package models

func init() {
	err := DB.AutoMigrate(&PushToken{})
	if err != nil {
		panic(err)
	}
}
