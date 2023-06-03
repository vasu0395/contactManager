package service

import "context"

type Service interface {
	GetContactNumber(ctx context.Context, contactName string) (string, error)
	SetContactNumber(ctx context.Context, contactName, phoneNumber string) error
	DeleteContactNumber(ctx context.Context, contactName string) error
	UpdateContactNumber(ctx context.Context, contactName string, updatedNumber string) error
}
