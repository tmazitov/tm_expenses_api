package user

import "context"

type ProfileOutput struct {
	Id        int
	FirstName string
	LastName  string
}

func (s *Service) Profile(ctx context.Context, userId int) (*ProfileOutput, error) {

	user, err := s.repo.GetById(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &ProfileOutput{
		Id:        user.Id(),
		FirstName: user.FirstName(),
		LastName:  user.LastName(),
	}, nil
}
