package dto

import (
	"database/sql"
	"games-info-api/src/models"
	"reflect"
)

func ParseGameDto(game models.Game) *models.GameDto {
	gameDto := &models.GameDto{
			ID: game.ID,
	}

	gameValue := reflect.ValueOf(game)
	gameDtoValue := reflect.ValueOf(gameDto).Elem()
	gameType := reflect.TypeOf(game)

	for i := 0; i < gameValue.NumField(); i++ {
			field := gameValue.Field(i)
			fieldType := gameType.Field(i)

			if fieldType.Name == "ID" {
					continue
			}

			if field.Type() == reflect.TypeOf(sql.NullString{}) {
					nullString := field.Interface().(sql.NullString)
					if nullString.Valid {
							gameDtoValue.FieldByName(fieldType.Name).SetString(nullString.String)
					} else {
							gameDtoValue.FieldByName(fieldType.Name).SetString("")
					}
			}
	}

	return gameDto
}
// var consoleNames = map[string]string{
// 	"cpc":                 "Amstrad CPC",
// 	"gx4000":              "Amstrad GX4000",
// 	"arduboy":             "Arduboy",
// 	"atari2600":           "Atari 2600",
// 	"atari5200":           "Atari 5200",
// 	"atari7800":           "Atari 7800",
// 	"atari8bit":           "Atari 8-bit",
// 	"jaguar":              "Atari Jaguar",
// 	"lynx":                "Atari Lynx",
// 	"atarist":             "Atari ST",
// 	"atomiswave":          "Sega Atomiswave",
// 	"wsc":                 "WonderSwan Color",
// 	"ws":                  "WonderSwan",
// 	"cannonball":          "OutRun Engine Cannonball",
// 	"loopy":               "Casio Loopy",
// 	"pv1000":              "Casio PV-1000",
// 	"cavestory":           "Cave Story",
// 	"chailove":            "ChaiLove",
// 	"chip8":               "CHIP-8",
// 	"coleco":              "ColecoVision",
// 	"c64":                 "Commodore 64",
// 	"amiga":               "Commodore Amiga",
// 	"cd32":                "Commodore CD32",
// 	"cdtv":                "Commodore CDTV",
// 	"pet":                 "Commodore PET",
// 	"plus4":               "Commodore Plus/4",
// 	"vic20":               "Commodore VIC-20",
// 	"dinothawr":           "DinoThawr",
// 	"doom":                "DOOM",
// 	"dos":                 "MS-DOS",
// 	"arcadia2001":         "Emerson Arcadia 2001",
// 	"enterprise128":       "Enterprise 128",
// 	"adventurevision":     "Entex Adventure Vision",
// 	"supercassettevision": "Epoch Super Cassette Vision",
// 	"channelf":            "Fairchild Channel F",
// 	"fbneo":               "Final Burn Neo",
// 	"flashback":           "Atari Flashback",
// 	"superacan":           "Funtech Super A'Can",
// 	"gp32":                "GamePark GP32",
// 	"vectrex":             "GCE Vectrex",
// 	"heg":                 "Handheld Electronic Game",
// 	"gamemaster":          "Hartung Game Master",
// 	"hbmame":              "HBMAME",
// 	"zmachine":            "Interactive Fiction (Z-Machine)",
// 	"jumpnbump":           "Jump 'n Bump",
// 	"leapster":            "Leapster",
// 	"lowresnx":            "LowRes NX",
// 	"lutro":               "Lutro",
// 	"odyssey2":            "Magnavox Odyssey 2",
// 	"mame2000":            "MAME 2000",
// 	"mame2003plus":        "MAME 2003 Plus",
// 	"mame2003":            "MAME 2003",
// 	"mame2010":            "MAME 2010",
// 	"mame2015":            "MAME 2015",
// 	"mame2016":            "MAME 2016",
// 	"mame":                "MAME",
// 	"intellivision":       "Mattel Intellivision",
// 	"msx":                 "MSX",
// 	"msx2":                "MSX2",
// 	"xbox":                "Microsoft Xbox",
// 	"microw8":             "MicroW8",
// 	"mrboom":              "Mr. Boom",
// 	"pce":                 "NEC PC Engine",
// 	"tgcd":                "NEC TurboGrafx-CD",
// 	"sgx":                 "NEC SuperGrafx",
// 	"pc8801":              "NEC PC-8801",
// 	"pc98":                "NEC PC-9801",
// 	"pcfx":                "NEC PC-FX",
// 	"ereader":             "Nintendo e-Reader",
// 	"fds":                 "Nintendo Famicom Disk System",
// 	"gba":                 "Nintendo Game Boy Advance",
// 	"gbc":                 "Nintendo Game Boy Color",
// 	"gb":                  "Nintendo Game Boy",
// 	"gc":                  "Nintendo GameCube",
// 	"3ds":                 "Nintendo 3DS",
// 	"n64":                 "Nintendo 64",
// 	"n64dd":               "Nintendo 64DD",
// 	"nds":                 "Nintendo DS",
// 	"dsi":                 "Nintendo DSi",
// 	"nes":                 "Nintendo Entertainment System (NES)",
// 	"pokemonmini":         "Nintendo Pokémon Mini",
// 	"satellaview":         "Nintendo Satellaview",
// 	"sufami":              "Nintendo Sufami Turbo",
// 	"snes":                "Super Nintendo Entertainment System (SNES)",
// 	"vb":                  "Nintendo Virtual Boy",
// 	"wii":                 "Nintendo Wii",
// 	"cdi":                 "Philips CD-i",
// 	"videopac":            "Philips Videopac",
// 	"pico8":               "PICO-8",
// 	"puzzlescript":        "PuzzleScript",
// 	"quake2":              "Quake II",
// 	"quake3":              "Quake III Arena",
// 	"quake":               "Quake",
// 	"studio2":             "RCA Studio II",
// 	"rickdangerous":       "Rick Dangerous",
// 	"rpgmaker":            "RPG Maker",
// 	"scummvm":             "ScummVM",
// 	"32x":                 "Sega 32X",
// 	"dc":                  "Sega Dreamcast",
// 	"gg":                  "Sega Game Gear",
// 	"sms":                 "Sega Master System",
// 	"md":                  "Sega Mega Drive/Genesis",
// 	"scd":                 "Sega CD",
// 	"naomi2":              "Sega NAOMI 2",
// 	"naomi":               "Sega NAOMI",
// 	"pico":                "Sega Pico",
// 	"saturn":              "Sega Saturn",
// 	"sg1000":              "Sega SG-1000",
// 	"x1":                  "Sharp X1",
// 	"x68000":              "Sharp X68000",
// 	"zx81":                "Sinclair ZX81",
// 	"zxsp3":               "Sinclair ZX Spectrum +3",
// 	"zxspectrum":          "Sinclair ZX Spectrum",
// 	"ngcd":                "SNK Neo Geo CD",
// 	"ngc":                 "SNK Neo Geo Pocket Color",
// 	"ngp":                 "SNK Neo Geo Pocket",
// 	"neogeo":              "SNK Neo Geo",
// 	"ps2":                 "Sony PlayStation 2",
// 	"ps3psn":              "Sony PlayStation 3 (PSN)",
// 	"ps3":                 "Sony PlayStation 3",
// 	"psp_psn":             "Sony PSP (PSN)",
// 	"psp":                 "Sony PlayStation Portable (PSP)",
// 	"psvita":              "Sony PlayStation Vita",
// 	"psx":                 "Sony PlayStation",
// 	"svi":                 "Spectravideo SVI",
// 	"3do":                 "Panasonic 3DO",
// 	"moto":                "Moto Racer",
// 	"tic80":               "TIC-80",
// 	"gamecom":             "Tiger Game.com",
// 	"tombraider":          "Tomb Raider",
// 	"uzebox":              "Uzebox",
// 	"vircon32":            "Vircon32",
// 	"creativision":        "VTech CreatiVision",
// 	"vsmile":              "VTech V.Smile",
// 	"wasm4":               "WASM-4",
// 	"supervision":         "Watara Supervision",
// 	"wolf3d":              "Wolfenstein 3D",
// }
