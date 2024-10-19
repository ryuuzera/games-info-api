package controllers

import (
	irepository "games-info-api/src/models/interface"
	"games-info-api/src/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGames(ctx *gin.Context){
	var gameRepo irepository.GamesRepository = repository.NewSQLiteGamesRepository() 

	games, err := gameRepo.ListAllGames();

	if (err != nil) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()});
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"games": games})
}