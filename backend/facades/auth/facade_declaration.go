package auth

// Service manages authentication.
type Service interface {
	GenerateToken(userID uint64) (string, error)
	GetUserIDByToken(token string) (uint64, error)
}
