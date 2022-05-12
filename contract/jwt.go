package contract

type (
	GenerateTokenPair func(secret []byte, userID string) (map[string]string, error)
	ParseToken        func(secret []byte, token string) (string, error)
)
