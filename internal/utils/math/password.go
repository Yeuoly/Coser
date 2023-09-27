package math

func HashPassword(password string) string {
	return Md5(Sha512(Sha256(Md5(password)) + "billboards" + Sha256(password) + Md5(password)))
}
