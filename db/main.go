package db

import (
	"database/sql"
	"github.com/joram/game-server/utils"
	_ "github.com/mattn/go-sqlite3"
)

func dbConn() *sql.DB {
	database, err := sql.Open("sqlite3", "./db/game.db")
	if err != nil {
		panic(err)
	}
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS players (" +
		"id INTEGER PRIMARY KEY," +
		"firstname TEXT," +
		"lastname TEXT," +
		"email TEXT UNIQUE," +
		"x INTEGER," +
		"y INTEGER" +
	")")
	if err != nil {
		panic(err)
	}
	statement.Exec()
	return database
}

type Player struct {
	Id        int
	X         int
	Y         int
	firstName string
	lastName  string
	email     string
}

func GetPlayerByEmail(email string) *Player {
	database := dbConn()
	row := database.QueryRow("SELECT id, x, y, firstname, lastname, email FROM players WHERE email=?", email)
	if row == nil {
		return nil
	}
	p := Player{}
	row.Scan(&p.Id, &p.X, &p.Y, &p.firstName, &p.lastName, &p.email)
	if p.email == "" {
		return nil
	}
	return &p
}

func GetOrCreatePlayer(email,firstName,lastName string) Player {
	player := GetPlayerByEmail(email)
	if player != nil {
		return *player
	}

	p := CreatePlayer(-utils.NextID(),0,0,firstName,lastName,email)
	return p
}

func GetPlayerById(id int) Player {
	database := dbConn()
	row := database.QueryRow("SELECT id, x, y, firstname, lastname, email FROM players WHERE id=?", id)
	p := Player{}
	row.Scan(&p.Id, &p.X, &p.Y, &p.firstName, &p.lastName, &p.email)
	return p
}

func CreatePlayer(id,x,y int, firstName, lastName, email string) Player {
	database := dbConn()
	statement, _ := database.Prepare("INSERT INTO players (id, x, y, firstName, lastName, email) VALUES (?,?,?,?,?,?)")
	statement.Exec(id,x,y,firstName,lastName,email)
	return Player{id,x,y,firstName,lastName,email}
}

func UpdatePlayer(id, x, y int) Player {
	database := dbConn()
	statement, err := database.Prepare("UPDATE players SET x=?, y=? WHERE id=?")
	if err != nil {
		panic(err)
	}
	_, err = statement.Exec(x,y, id)
	if err != nil {
		panic(err)
	}
	return GetPlayerById(id)
}

