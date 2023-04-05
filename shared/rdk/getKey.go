package rdk

func GetVerificationKey(email string) string {
	return Prefix + email + RKVerification
}
