package models

type WpWpfmBackup struct {
	ID         int     `gorm:"column:id;primaryKey" json:"id"`
	BackupName *string `gorm:"column:backup_name" json:"backupName"`
	BackupDate *string `gorm:"column:backup_date" json:"backupDate"`
}

func (WpWpfmBackup) TableName() string {
	return "wp_wpfm_backup"
}
