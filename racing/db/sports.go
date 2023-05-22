package db

import (
	"database/sql"
	"strings"
	"sync"
	"time"

	"github.com/Kim-Hardie/entain-master/racing/proto/sports"
	"github.com/golang/protobuf/ptypes"
	_ "github.com/mattn/go-sqlite3"
)

// MatchesRepo provides repository access to matches.
type MatchesRepo interface {
	Init() error
	List(filter *sports.MatchFilter) ([]*sports.Match, error)
	GetByID(id int64) (*sports.Match, error)
}

type matchesRepo struct {
	db   *sql.DB
	init sync.Once
}

// NewMatchesRepo creates a new matches repository.
func NewMatchesRepo(db *sql.DB) MatchesRepo {
	return &matchesRepo{db: db}
}

// Init prepares the matches repository with dummy data.
func (r *matchesRepo) Init() error {
	var err error

	r.init.Do(func() {
		// For test/example purposes, we seed the DB with some dummy matches.
		err = r.seed()
	})

	return err
}

// List returns a list of matches based on the provided filter.
func (r *matchesRepo) List(filter *sports.MatchFilter) ([]*sports.Match, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	query = getMatchQueries()[matchesList]

	query, args = r.applyFilter(query, filter)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return r.scanMatches(rows)
}

// applyFilter applies the provided filter to the SQL query.
func (r *matchesRepo) applyFilter(query string, filter *sports.MatchFilter) (string, []interface{}) {
	var (
		clauses []string
		args    []interface{}
	)

	if filter == nil {
		return query, args
	}

	if filter.Stadium != "" {
		clauses = append(clauses, "stadium = ?")
		args = append(args, filter.Stadium)
	}

	if filter.Sport != "" {
		clauses = append(clauses, "sport = ?")
		args = append(args, filter.Sport)
	}

	if len(clauses) != 0 {
		query += " WHERE " + strings.Join(clauses, " AND ")
	}

	return query, args
}

// scanMatches scans the database rows and converts them to matches.
func (r *matchesRepo) scanMatches(rows *sql.Rows) ([]*sports.Match, error) {
	var matches []*sports.Match

	for rows.Next() {
		var match sports.Match
		var time time.Time

		if err := rows.Scan(&match.Id, &match.Name, &match.Stadium, &match.Sport, &match.Team1, &match.Team2, &time); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		ts, err := ptypes.TimestampProto(time)
		if err != nil {
			return nil, err
		}

		match.Time = ts

		matches = append(matches, &match)
	}

	return matches, nil
}

// GetByID retrieves a match by its ID.
func (r *matchesRepo) GetByID(matchID int64) (*sports.Match, error) {
	query := "SELECT id, name, stadium, sport, team1, team2, time FROM matches WHERE id = ?"
	row := r.db.QueryRow(query, matchID)

	match, err := r.scanMatch(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return match, nil
}

// scanMatch scans a single match from the database row.
func (r *matchesRepo) scanMatch(row *sql.Row) (*sports.Match, error) {
	var match sports.Match
	var time time.Time

	err := row.Scan(&match.Id, &match.Name, &match.Stadium, &match.Sport, &match.Team1, &match.Team2, &time)
	if err != nil {
		return nil, err
	}

	ts, err := ptypes.TimestampProto(time)
	if err != nil {
		return nil, err
	}

	match.Time = ts

	return &match, nil
}

// seed populates the matches table with dummy data.
func (r *matchesRepo) seed() error {
	// Implement the logic to seed the matches table with dummy data
	return nil
}

// getMatchQueries returns the SQL queries for matches.
func getMatchQueries() map[string]string {
	// Define the SQL queries for matches
	return map[string]string{
		matchesList: "SELECT id, name, stadium, sport, team1, team2, time FROM matches",
	}
}

const (
	matchesList = "matchesList"
)
