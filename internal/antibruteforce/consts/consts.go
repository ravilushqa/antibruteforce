package consts

const (
	// LoginPrefix is prefix for login key in bucket
	LoginPrefix = "login_"
	// PasswordPrefix is prefix for password key in bucket
	PasswordPrefix = "password_"
	// IPPrefix is prefix for ip key in bucket
	IPPrefix = "ip"
	// Whitelist is keyword for mysql query for checking that ip in whitelist
	Whitelist = "whitelist"
	// Blacklist is keyword for mysql query for checking that ip in blacklist
	Blacklist = "blacklist"
)
