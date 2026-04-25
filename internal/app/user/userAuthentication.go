package user

func (s *Service) Authenticate(token string) (int, error) {

	claims, err := s.jwt.VerifyAccess(token)
	if err != nil {
		return 0, err
	}

	userId, ok := claims["user_id"].(float64)
	if !ok {
		return 0, ErrInvalidToken
	}

	return int(userId), nil
}
