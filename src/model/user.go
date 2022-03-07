package model

import (
	"gitlab.com/abhishek.k8/crud/src/database"
)

type (
	//Users Model for user table
	Users struct {
		ID             uint   `gorm:"primary_key" json:"id,omitempty" `
		FirstName      string `gorm:"type:varchar(100);" json:"first_name" valid:"required,length(3|100)"`
		LastName       string `gorm:"type:varchar(100);" json:"last_name" valid:"required,length(1|100)"`
		Email          string `gorm:"type:varchar(100);unique_index; not null" json:"email" valid:"email,required"`
		Phone          string `gorm:"type:varchar(100);unique_index; not null" json:"phone" valid:"required,length(5|15)"`
		ProfilePicture string `gorm:"type:varchar(100);default:''" json:"profile_picture"`
		Password       string `json:"password,omitempty"`
		Country        string `gorm:"type:varchar(100);default:''" json:"country" valid:"required,length(1|100)"`
		State          string `gorm:"type:varchar(100);default:''" json:"state" valid:"required,length(2|100)"`
		Address        string `gorm:"type:varchar(255);default:''" json:"address" valid:"required,length(3|255)"`
		ZipCode        string `gorm:"type:varchar(12);default:''" json:"zip_code" valid:"required,length(3|10)"`
	}
	//ReferralInfo struct
	Update struct {
		ID             uint   `gorm:"primary_key" json:"id,omitempty" `
		FirstName      string `gorm:"type:varchar(100);" json:"first_name"`
		LastName       string `gorm:"type:varchar(100);" json:"last_name"`
		ProfilePicture string `gorm:"type:varchar(100);default:''" json:"profile_picture"`
		Password       string `json:"password,omitempty"`
		Country        string `gorm:"type:varchar(100);default:''" json:"country"`
		State          string `gorm:"type:varchar(100);default:''" json:"state" `
		Address        string `gorm:"type:varchar(255);default:''" json:"address"`
		ZipCode        string `gorm:"type:varchar(12);default:''" json:"zip_code"`
	}
)

//Register register a user
func (u *Users) Register() error {
	//database connection
	dbconn := database.GetSharedConnection()
	tx := dbconn.Begin()
	if err := tx.Create(u).Error; err != nil {

		//log error and return
		// log.Error(u)
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

//GetUser update user status
func (u *Users) GetUser(userID uint) error {
	//database connection
	dbconn := database.GetSharedConnection()
	if err := dbconn.Debug().Select(`id,first_name,last_name,email,phone,profile_picture,country,state,address,zip_code`).
		Where(`id=?`, userID).First(&u).Error; err != nil {
		return err
	}
	return nil
}

func (u *Users) GetAllUsers() []Users {
	var (
		users []Users
	)
	db := database.GetSharedConnection()
	query := db.Debug().Select(`
			first_name,
			last_name,
			email,
			phone,
			profile_picture,
			password,
			country,
			state,
			address,
			zip_code
	`).Table("users")
	query.Scan(&users)
	return users
}

func (u *Users) Updateuser(user Update) (*Users, error) {
	dbconn := database.GetSharedConnection()
	if err := dbconn.Debug().Where(`id=?`, user.ID).Find(&u).Updates(map[string]interface{}{
		"first_name":      user.FirstName,
		"last_name":       user.LastName,
		"profile_picture": user.ProfilePicture,
		"password":        user.Password,
		"country":         user.Country,
		"state":           user.State,
		"address":         user.Address,
		"zip_code":        user.ZipCode,
	}).Error; err != nil {
		return nil, err
	}
	return u, nil
}

//delete user
func (u *Users) DeleteUser(userID uint) error {
	//database connection
	dbconn := database.GetSharedConnection()
	if err := dbconn.Debug().Where(`id=?`, userID).First(&u).Delete(&u).Error; err != nil {
		return err
	}
	return nil
}
