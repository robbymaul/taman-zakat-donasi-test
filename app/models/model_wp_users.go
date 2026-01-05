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
	ID        uint64 `gorm:"column:id" json:"id"`
	UserLogin string `gorm:"column:user_login" json:"userLogin"`
	Email     string `gorm:"column:email" json:"email"`
	Password  string `gorm:"column:password" json:"password"`
	FullName  string `gorm:"column:full_name" json:"fullName"`
	Status    int    `gorm:"column:status" json:"status"`
}

type WpUserSession struct {
	ID       uint64 `gorm:"column:id" json:"id"`
	Email    string `gorm:"column:email" json:"email"`
	FullName string `gorm:"column:full_name" json:"fullName"`
}

type WpUserProfile struct {
	ID             uint64 `gorm:"column:id" json:"id"`
	CoverPicture   string `gorm:"column:cover_picture" json:"coverPicture"`
	ProfilePicture string `gorm:"column:profile_picture" json:"profilePicture"`
	FullName       string `gorm:"column:full_name" json:"fullName"`
	Greeting       string `gorm:"column:greeting" json:"greeting"`
	Biography      string `gorm:"column:biography" json:"biography"`
	Province       string `gorm:"column:province" json:"province"`
	City           string `gorm:"column:city" json:"city"`
	Address        string `gorm:"address" json:"address"`
	Username       string `gorm:"column:username" json:"username"`
	Email          string `gorm:"column:email" json:"email"`
	Whatsapp       string `gorm:"column:whatsapp" json:"whatsapp"`
	Bank           string `gorm:"column:bank" json:"bank"`
	Account        string `gorm:"column:account" json:"account"`
	AccountName    string `gorm:"column:account_name" json:"accountName"`
}
