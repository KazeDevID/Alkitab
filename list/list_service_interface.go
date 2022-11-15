package list

import "Alkitab/list/models"

// ListService is interface to implemented repository
type ListService interface {
	ListPassage() (*[]models.PassageList, error)
}
