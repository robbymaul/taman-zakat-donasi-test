package enums

const (
	USER_VERIFICATION_NO_STATUS int = 0
	USER_VERIFICATION_VERIFIED  int = 1
	USER_VERIFICATION_ON_REVIEW int = 2
	USER_VERIFICATION_REJECTED  int = 3
	USER_VERIFICATION_BANNED    int = 4
)

const (
	NO_STATUS string = "No Status"
	VERIFIED  string = "Verified"
	ON_REVIEW string = "On Review"
	REJECTED  string = "Rejected"
	BANNED    string = "Banned"
)

func GetUserVerification(status int) string {
	var result string

	switch status {
	case USER_VERIFICATION_NO_STATUS:
		result = NO_STATUS
	case USER_VERIFICATION_VERIFIED:
		result = VERIFIED
	case USER_VERIFICATION_ON_REVIEW:
		result = ON_REVIEW
	case USER_VERIFICATION_REJECTED:
		result = REJECTED
	case USER_VERIFICATION_BANNED:
		result = BANNED
	default:
		result = NO_STATUS
	}

	return result
}
