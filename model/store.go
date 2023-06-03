package model

import "context"

type ContactStore interface {
	GetContactNumber(ctx context.Context, contactName string) (string, error)
	SetContactMapping(ctx context.Context, contactName, phoneNumber string) error
	DeleteContactNumber(ctx context.Context, contactName string) error
	UpdateContactNumber(ctx context.Context, contactName string, updatedNumber string) error
}
