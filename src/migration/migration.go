package migration

import (
	"gitlab.com/abhishek.k8/crud/src/database"
	"gitlab.com/abhishek.k8/crud/src/model"
)

//Migrate to migrate the models
func Migrate() {
	dbconn := database.GetSharedConnection()
	//DB migration
	dbconn.Debug().AutoMigrate(&model.Users{})
}
