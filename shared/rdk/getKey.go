package rdk

import "strconv"

func GetVerificationKey(email string) string {
	return Prefix + email + RKVerification
}

func GetCacheUserInfoKey(id int64) string {
	return Prefix + strconv.FormatInt(id, 10) + RKCacheUserInfo
}
