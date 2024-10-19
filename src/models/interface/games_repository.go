package repositoryinterface

import "games-info-api/src/models"

type GamesRepository interface {
	ListAllGames() (*models.Games, error)
	GetGameByFilter(filter string, value interface{}) (*models.Games, error)
}