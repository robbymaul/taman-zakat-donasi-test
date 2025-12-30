package models

import "time"

type WpUsers struct {
	ID                uint64    `gorm:"column:ID;primaryKey" json:"id"`
	UserLogin         string    `gorm:"column:user_login" json:"userLogin"`
	UserPass          string    `gorm:"column:user_pass" json:"userPass"`
	UserNicename      string    `gorm:"column:user_nicename" json:"userNicename"`
	UserEmail         string    `gorm:"column:user_email" json:"userEmail"`
	UserURL           string    `gorm:"column:user_url" json:"userUrl"`
	UserRegistered    time.Time `gorm:"column:user_registered" json:"userRegistered"`
	UserActivationKey string    `gorm:"column:user_activation_key" json:"userActivationKey"`
	UserStatus        int       `gorm:"column:user_status" json:"userStatus"`
	DisplayName       string    `gorm:"column:display_name" json:"displayName"`
}

func (WpUsers) TableName() string {
	return "wp_users"
}

type WpUsersLogin struct {
	ID       uint64 `gorm:"column:id" json:"id"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
	FullName string `gorm:"column:full_name" json:"fullName"`
	Status   int    `gorm:"column:status" json:"status"`
}
