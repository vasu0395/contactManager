package service

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"

	"UrlShortner/model"
)

type ContactService struct {
	store model.ContactStore
}

func createContactService(s model.ContactStore) ContactService {
	cs := ContactService{
		store: s,
	}

	return cs
}

func (s *ContactService) GetContactNumber(ctx context.Context, contactName string) (string, error) {
	resp, err := s.store.GetContactNumber(ctx, contactName)
	logger := zerolog.Ctx(ctx)
	if err != nil {
		logger.Err(fmt.Errorf("error while fetching contactNumber %s", err.Error()))
		return "", err
	}
	return resp, nil
}

func (s *ContactService) SetContactNumber(ctx context.Context, contactName, phoneNumber string) error {
	err := s.store.SetContactMapping(ctx, contactName, phoneNumber)
	logger := zerolog.Ctx(ctx)
	if err != nil {
		logger.Err(fmt.Errorf("error while saving contactName and phoneNumber %s", err.Error()))
		return err
	}
	return nil
}

func (s *ContactService) DeleteContactNumber(ctx context.Context, contactName string) error {
	err := s.store.DeleteContactNumber(ctx, contactName)
	logger := zerolog.Ctx(ctx)
	if err != nil {
		logger.Err(fmt.Errorf("error while deleting contactName %s", err.Error()))
		return err
	}
	return nil
}

func (s *ContactService) UpdateContactNumber(ctx context.Context, contactName string, updatedNumber string) error {
	err := s.store.UpdateContactNumber(ctx, contactName, updatedNumber)
	logger := zerolog.Ctx(ctx)
	if err != nil {
		logger.Err(fmt.Errorf("error while updating contactNumber %s", err.Error()))
		return err
	}
	return nil
}
