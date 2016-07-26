package token

// Token ...
type Token struct {
	Key         string
	UserID      string
	LastLoginAt int64
	ExpireAt    int64
}
