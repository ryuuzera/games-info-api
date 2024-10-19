package models

import (
	"database/sql"
)

type GameDto struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	System          string `json:"system"`
	ComercialSystem string `json:"comercialSystem"`
	RomName         string `json:"romName"`
	Developer       string `json:"developer"`
	CRC32           string `json:"crc32"`
	Serial          string `json:"serial"`
}
	
type Game struct {
	ID              int    
	Name            sql.NullString 
	System          sql.NullString 
	ComercialSystem sql.NullString 
	RomName         sql.NullString 
	Developer       sql.NullString 
	CRC32           sql.NullString 
	Serial          sql.NullString 
}

type Games struct {
	Games []GameDto `json:"games"`
}

