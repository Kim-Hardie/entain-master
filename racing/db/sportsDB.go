package db

import (
	"database/sql"
	"time"
)

// SportsRepo provides repository access to sports matches.
type SportsRepo interface {
	Init() error
	CreateMatch(match *Match) error
	GetMatchByID(id int) (*Match, error)
}

type sportsRepo struct {
	db *sql.DB
}

// Match represents a sports match.
type Match struct {
	ID      int64
	Name    string
	Stadium string
	Sport   string
	Team1   string
	Team2   string
	Time    time.Time
}

// NewSportsRepo creates a new sports repository.
func NewSportsRepo(db *sql.DB) SportsRepo {
	return &sportsRepo{db: db}
}

func (r *sportsRepo) Init() error {
	// Create the matches table if it doesn't exist
	_, err := r.db.Exec(`CREATE TABLE IF NOT EXISTS matches (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		stadium TEXT,
		sport TEXT,
		team1 TEXT,
		team2 TEXT,
		time TIMESTAMP
	)`)

	return err
}

func (r *sportsRepo) CreateMatch(match *Match) error {
	// Insert the match into the database
	result, err := r.db.Exec(`INSERT INTO matches (name, stadium, sport, team1, team2, time)
		VALUES (?, ?, ?, ?, ?, ?)`,
		match.Name, match.Stadium, match.Sport, match.Team1, match.Team2, match.Time)
	if err != nil {
		return err
	}

	// Retrieve the auto-incremented ID of the inserted match
	matchID, _ := result.LastInsertId()
	match.ID = int64(matchID)

	return nil
}

func (r *sportsRepo) GetMatchByID(id int) (*Match, error) {
	// Retrieve the match from the database by ID
	row := r.db.QueryRow(`SELECT id, name, stadium, sport, team1, team2, time FROM matches WHERE id = ?`, id)

	match := &Match{}
	err := row.Scan(&match.ID, &match.Name, &match.Stadium, &match.Sport, &match.Team1, &match.Team2, &match.Time)
	if err != nil {
		return nil, err
	}

	return match, nil
}
