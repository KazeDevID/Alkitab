package list

import "Alkitab/list/models"

// ListRepository is interface to implemented repository
type ListRepository interface {
	ListPassage() (*[]models.PassageList, error)
}
