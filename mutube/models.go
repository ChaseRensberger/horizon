package mutube

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type Token struct {
	Token     string `json:"token" dynamodbav:"token"`
	ExpiresAt int64  `json:"expires_at" dynamodbav:"expires_at"`
	UserID    string `json:"user_id" dynamodbav:"user_id"`
}
