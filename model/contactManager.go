package model

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
)

type ContactDB struct {
	contactStore map[string]string
}

func (s *ContactDB) GetContactNumber(ctx context.Context, contactName string) (string, error) {
	logger := zerolog.Ctx(ctx)
	contactNumber, exists := s.contactStore[contactName]
	if exists {
		return contactNumber, nil
	} else {
		logger.Err(fmt.Errorf("contact doesn't exist : %s", contactName))
		return "", mappingContactNotFound
	}
}

func (s *ContactDB) SetContactMapping(ctx context.Context, contactName, phoneNumber string) error {
	logger := zerolog.Ctx(ctx)
	_, exists := s.contactStore[contactName]
	if exists {
		logger.Err(fmt.Errorf("contactName already exist : %s", contactName))
		return mappingAlreadyExist
	}

	s.contactStore[contactName] = phoneNumber
	return nil
}

func (s *ContactDB) DeleteContactNumber(ctx context.Context, contactName string) error {
	logger := zerolog.Ctx(ctx)
	_, exists := s.contactStore[contactName]
	if !exists {
		logger.Err(fmt.Errorf("contactName mapping doesn't exist : %s", contactName))
		return mappingNotFound
	}

	delete(s.contactStore, contactName)
	return nil
}

func (s *ContactDB) UpdateContactNumber(ctx context.Context, contactName string, updatedNumber string) error {
	logger := zerolog.Ctx(ctx)
	_, exists := s.contactStore[contactName]
	if !exists {
		logger.Err(fmt.Errorf("contactName doesn't exist : %s", contactName))
	}

	s.contactStore[contactName] = updatedNumber
	return nil
}
