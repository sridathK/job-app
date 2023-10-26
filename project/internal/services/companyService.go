package services

import (
	"errors"
	"project/internal/model"

	"github.com/rs/zerolog/log"
)

func (s *Service) CompanyCreate(nc model.CreateCompany) (model.Company, error) {
	company := model.Company{CompanyName: nc.CompanyName, Adress: nc.Adress, Domain: nc.Domain}
	cu, err := s.r.CreateCompany(company)
	if err != nil {
		log.Error().Err(err).Msg("couldnot create user")
		return model.Company{}, errors.New("user creation failed")
	}

	return cu, nil
}

func (s *Service) GetAllCompanies() ([]model.Company, error) {

	AllCompanies, err := s.r.GetAllCompany()
	if err != nil {
		return nil, err
	}
	return AllCompanies, nil

}

func (s *Service) GetCompany(id int) (model.Company, error) {

	AllCompanies, err := s.r.GetCompany(id)
	if err != nil {
		return model.Company{}, err
	}
	return AllCompanies, nil

}
