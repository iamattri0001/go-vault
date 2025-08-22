package service

import customerrors "go-vault/custom_errors"

func (s *Service) GetSalts(username string) (*Salts, error) {
	user, err := s.userRepository.GetSaltsByUsername(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, customerrors.ErrUserNotFound
	}
	return &Salts{
		AuthSalt:       user.AuthSalt,
		EncryptionSalt: user.EncryptionSalt,
	}, nil
}
