package repository

import (
	"database/sql"
	"games-info-api/src/models"
	"games-info-api/src/models/dto"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteGamesRepository struct {
    db *sql.DB
}

func NewSQLiteGamesRepository() *SQLiteGamesRepository {
	basePath, err := os.Getwd()

	dbPath := filepath.Join(basePath, "db", "data.games")

	db, err := sql.Open("sqlite3", dbPath)

	if (err != nil) {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
  return &SQLiteGamesRepository{db: db}
}

func (r *SQLiteGamesRepository) ListAllGames() (*models.Games, error) {
    query := "SELECT id, name, system, romName, developer, crc32, serial FROM games limit 100"
    rows, err := r.db.Query(query)

    if err != nil {
        return nil, err
    }
		
    defer rows.Close()

    var games models.Games

    for rows.Next() {
        var game models.Game
        err := rows.Scan(&game.ID, &game.Name, &game.System, &game.RomName, &game.Developer, &game.CRC32, &game.Serial)
        if err != nil {
            return nil, err
        }
				gameDto := dto.ParseGameDto(game);
        games.Games = append(games.Games, *gameDto)
    }

    return &games, nil
}

func (r *SQLiteGamesRepository) GetGameByFilter(filter string, value interface{}) (*models.Games, error) {
    query := "SELECT id, name, system, comercial_system, rom_name, developer, crc32, serial FROM games WHERE " + filter + " = ?"
    rows, err := r.db.Query(query, value)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var games models.Games
    for rows.Next() {
        var game models.Game
        err := rows.Scan(&game.ID, &game.Name, &game.System, &game.ComercialSystem, &game.RomName, &game.Developer, &game.CRC32, &game.Serial)
        if err != nil {
            return nil, err
        }
				gameDto := dto.ParseGameDto(game);
        games.Games = append(games.Games, *gameDto)
    }

    return &games, nil
}