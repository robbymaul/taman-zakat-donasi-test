package repositories

import (
	"context"
	"donasitamanzakattest/app/models"
	"donasitamanzakattest/app/web"

	"gorm.io/gorm"
)

func (rc *RepositoryContext) RegistrationRepository(ctx context.Context, register *models.WpDjaRegister, wpUser *models.WpUsers, wpDjaUser *models.WpDjaUsers, wpDjaVerificationDetail *models.WpDjaVerificationDetails) error {
	var err error

	err = rc.WithTransaction(func(tx *gorm.DB) error {
		// wp_dja_register
		queryRegister := "INSERT INTO " +
			"wp_dja_register (r_nama_lengkap, r_email, r_whatsapp, r_password, r_code, r_status, created_at) " +
			"VALUES (?, ?, ?, ?, ?, ?, ?)"
		tx = tx.WithContext(ctx).Exec(queryRegister, register.RNamaLengkap, register.REmail, register.RWhatsapp, register.RPassword, register.RCode, register.RStatus, register.CreatedAt)
		err = tx.Error
		if err != nil {
			return err
		}

		// wp_users
		queryWpUser := "INSERT INTO " +
			"wp_users (user_login, user_pass, user_nicename, user_email, user_url, user_registered, user_activation_key, user_status, display_name) " +
			"values (?, ?, ?, ?, ?, ?, ?, ?, ?)"
		tx = tx.WithContext(ctx).Exec(queryWpUser, wpUser.UserLogin, wpUser.UserPass, wpUser.UserNicename, wpUser.UserEmail, wpUser.UserURL, wpUser.UserRegistered, wpUser.UserActivationKey, wpUser.UserStatus, wpUser.DisplayName)
		err = tx.Error
		if err != nil {
			return err
		}

		// get last id wp_users
		var wpUserId int64
		tx = tx.Raw("SELECT LAST_INSERT_ID()").Scan(&wpUserId)
		err = tx.Error
		if err != nil {
			return err
		}

		// wp_dja_users
		queryWpDjaUser := "INSERT INTO " +
			"wp_dja_users (user_id, user_randid, user_type, user_verification, user_bio, user_wa, user_provinsi, user_kabkota, user_kecamatan, user_provinsi_id, user_kabkota_id, user_kecamatan_id, user_alamat, user_bank_name, user_bank_no, user_bank_an, user_pp_img, user_cover_img, created_at, user_anonim_f, user_commission_f, user_sapaan) " +
			"values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

		wpDjaUser.UserID = &wpUserId

		tx = tx.WithContext(ctx).Exec(queryWpDjaUser, wpDjaUser.UserID, wpDjaUser.UserRandid, wpDjaUser.UserType, wpDjaUser.UserVerification, wpDjaUser.UserBio, wpDjaUser.UserWa, wpDjaUser.UserProvinsi, wpDjaUser.UserKabkota, wpDjaUser.UserKecamatan, wpDjaUser.UserProvinsiID, wpDjaUser.UserKabkotaID, wpDjaUser.UserKecamatanID, wpDjaUser.UserAlamat, wpDjaUser.UserBankName, wpDjaUser.UserBankNo, wpDjaUser.UserBankAn, wpDjaUser.UserPpImg, wpDjaUser.UserCoverImg, wpDjaUser.CreatedAt, wpDjaUser.UserAnonimF, wpDjaUser.UserCommissionF, wpDjaUser.UserSapaan)
		err = tx.Error
		if err != nil {
			return err
		}

		// wp_dja_verification_details
		queryWpDjaVerificationDetail := "INSERT INTO " +
			"wp_dja_verification_details (u_id, u_nama_lengkap, u_email, u_whatsapp, u_ktp, u_ktp_selfie, u_jabatan, u_nama_ketua, u_alamat_lengkap, u_program_unggulan, u_profile, u_legalitas, created_at) " +
			"values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

		UID := int(wpUserId)

		wpDjaVerificationDetail.UID = &UID

		tx = tx.WithContext(ctx).Exec(queryWpDjaVerificationDetail, wpDjaVerificationDetail.UID, wpDjaVerificationDetail.UNamaLengkap, wpDjaVerificationDetail.UEmail, wpDjaVerificationDetail.UWhatsapp, wpDjaVerificationDetail.UKtp, wpDjaVerificationDetail.UKtpSelfie, wpDjaVerificationDetail.UJabatan, wpDjaVerificationDetail.UNamaKetua, wpDjaVerificationDetail.UAlamatLengkap, wpDjaVerificationDetail.UProgramUnggulan, wpDjaVerificationDetail.UProfile, wpDjaVerificationDetail.ULegalitas, wpDjaVerificationDetail.CreatedAt)
		err = tx.Error
		if err != nil {
			return err
		}

		return err
	})

	return err
}

func (rc *RepositoryContext) LoginRepository(ctx context.Context, req *web.LoginRequest) (*models.WpUsersLogin, error) {
	var wpUser models.WpUsersLogin
	var err error
	var db = rc.db

	query := "SELECT wp_users.id as id, wp_users.user_email as email, wp_users.user_pass as password, wp_users.display_name as full_name, wp_dja_users.user_verification as status " +
		"FROM wp_users " +
		"JOIN wp_dja_users ON wp_users.id = wp_dja_users.user_id " +
		"WHERE user_email = ?"

	db = db.WithContext(ctx).Raw(query, req.Email)

	db = db.Scan(&wpUser)

	err = db.Error
	if err != nil {
		return nil, err
	}

	return &wpUser, err
}
