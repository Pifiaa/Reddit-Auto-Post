package credentials

type credentials struct {
	id            int    "`json:id`"
	usename       int    "`json:usename`"
	password      int    "`json:password`"
	client_secret string "`json:client_secret `"
	client_id     string "`json:client_id `"
}
