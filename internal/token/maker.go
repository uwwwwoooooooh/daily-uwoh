package token

import "time"

type TokenMaker interface {
	CreateToken(userID uint, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
