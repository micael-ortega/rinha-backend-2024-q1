package service

import (
	"github.com/micael-ortega/crebito/internal/dto"
	"github.com/micael-ortega/crebito/internal/dto/response"
	"github.com/micael-ortega/crebito/internal/repository"
)

type Service struct {
	repo *repository.Repo
}

func NewService(repo *repository.Repo) *Service {
	return &Service{
		repo: repo,
	}
}


func (s *Service) CreditTransaction(value int, description string, client_id int) (*dto.TransactionDTO, error) {
	response, err := s.repo.CreditTransaction(value, description, client_id)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *Service) DebitTransaction(value int, description string, client_id int) (*dto.TransactionDTO, error) {
	response, err := s.repo.DebitTransaction(value, description, client_id)
	if err != nil {
		return nil, err
	}

	return response, nil
}


func (s *Service) GetLastTransactions(id int) (*response.BankStatement, error) {
	res, err := s.repo.GetLastTransactions(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

