package model

import "fmt"

var (
	mappingAlreadyExist    = fmt.Errorf("mapping already exist")
	mappingNotFound        = fmt.Errorf("mapping doesn't exist")
	mappingContactNotFound = fmt.Errorf("contact doesn't exist")
)
