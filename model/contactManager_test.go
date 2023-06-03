package model

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetContactNumber(t *testing.T) {
	mp := map[string]string{"vasu": "9988776655"}
	s := ContactDB{contactStore: mp}

	t.Run("Should return nil error, when error not exist", func(t *testing.T) {
		ctx := context.Background()
		defer ctx.Done()
		resp, err := s.GetContactNumber(ctx, "vasu")
		assert.Equal(t, resp, "9988776655")
		assert.Equal(t, err, nil)
	})

	t.Run("Should return error, when contact does not exist", func(t *testing.T) {
		ctx := context.Background()
		defer ctx.Done()
		resp, err := s.GetContactNumber(ctx, "name")
		assert.Equal(t, resp, "")
		assert.Equal(t, err, mappingContactNotFound)
	})
}

func TestSetContactMapping(t *testing.T) {
	mp := map[string]string{"vasu": "9988776655"}
	s := ContactDB{contactStore: mp}

	t.Run("Should return nil error, when error not exist", func(t *testing.T) {
		ctx := context.Background()
		defer ctx.Done()
		err := s.SetContactMapping(ctx, "name", "9988999889")
		assert.Equal(t, err, nil)
	})

	t.Run("Should return error, when contact already exist", func(t *testing.T) {
		ctx := context.Background()
		defer ctx.Done()
		err := s.SetContactMapping(ctx, "vasu", "9999889989")
		assert.Equal(t, err, mappingAlreadyExist)
	})
}

func TestDeleteContactNumber(t *testing.T) {
	mp := map[string]string{"vasu": "9988776655"}
	s := ContactDB{contactStore: mp}

	t.Run("Should return nil error, when successful deletion occurs", func(t *testing.T) {
		ctx := context.Background()
		defer ctx.Done()
		err := s.DeleteContactNumber(ctx, "vasu")
		assert.Equal(t, err, nil)
	})

	t.Run("Should return error, when contact doesn't exist exist", func(t *testing.T) {
		ctx := context.Background()
		defer ctx.Done()
		err := s.DeleteContactNumber(ctx, "vasu")
		assert.Equal(t, err, mappingNotFound)
	})
}

func TestUpdateContactNumber(t *testing.T) {
	mp := map[string]string{"vasu": "9988776655"}
	s := ContactDB{contactStore: mp}

	t.Run("Should return nil error, when update successful", func(t *testing.T) {
		ctx := context.Background()
		defer ctx.Done()
		err := s.UpdateContactNumber(ctx, "vasu", "9897797879")
		assert.Equal(t, err, nil)
		assert.Equal(t, "9897797879", s.contactStore["vasu"])
	})

	t.Run("Should return nil error, when contact doesn't exist exist", func(t *testing.T) {
		ctx := context.Background()
		defer ctx.Done()
		err := s.UpdateContactNumber(ctx, "name", "9897797879")
		assert.Equal(t, err, nil)
		assert.Equal(t, "9897797879", s.contactStore["vasu"])
	})
}
