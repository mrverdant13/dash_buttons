package auth

// Service manages authentication.
type Service interface {
	GenerateToken(userID string) (string, error)
	GetUserIDByToken(token string) (string, error)
}
