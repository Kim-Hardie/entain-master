package db

import (
	"database/sql"
	"github.com/golang/protobuf/ptypes"
	_ "github.com/mattn/go-sqlite3"
	"strings"
	"sync"
	"time"

	"github.com/Kim-Hardie/entain-master/racing/proto/racing"
)

// RacesRepo provides repository access to races.
type RacesRepo interface {
	// Init will initialise our races repository.
	Init() error

	// List will return a list of races.
	List(filter *racing.ListRacesRequestFilter) ([]*racing.Race, error)

	// GetByID will return a race by its ID.
	GetByID(id int64) (*racing.Race, error)
}

type racesRepo struct {
	db   *sql.DB
	init sync.Once
}

// NewRacesRepo creates a new races repository.
func NewRacesRepo(db *sql.DB) RacesRepo {
	return &racesRepo{db: db}
}

// Init prepares the race repository dummy data.
func (r *racesRepo) Init() error {
	var err error

	r.init.Do(func() {
		// For test/example purposes, we seed the DB with some dummy races.
		err = r.seed()
	})

	return err
}

func (r *racesRepo) List(filter *racing.ListRacesRequestFilter) ([]*racing.Race, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	query = getRaceQueries()[racesList]

	query, args = r.applyFilter(query, filter)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return r.scanRaces(rows)
}

func (r *racesRepo) applyFilter(query string, filter *racing.ListRacesRequestFilter) (string, []interface{}) {
	var (
		clauses []string
		args    []interface{}
		order   string
	)

	if filter == nil {
		return query, args
	}

	if len(filter.MeetingIds) > 0 {
		clauses = append(clauses, "meeting_id IN ("+strings.Repeat("?,", len(filter.MeetingIds)-1)+"?)")

		for _, meetingID := range filter.MeetingIds {
			args = append(args, meetingID)
		}
	}

	//if no filter is set defaults to only show visible races
	if filter.ShowOnlyVisible != nil {
		clauses = append(clauses, "visible = ?")
		args = append(args, *filter.ShowOnlyVisible)
	} else {
		clauses = append(clauses, "visible = ?")
		args = append(args, true)
	}
	//if no filter is set Default to Ascending order by DateTime
	if filter.OrderAscending != nil {
		if *filter.OrderAscending {
			order = " ORDER BY advertised_start_time ASC"
		} else {
			order = " ORDER BY advertised_start_time DESC"
		}
	} else {
		order = " ORDER BY advertised_start_time ASC"
	}

	if len(clauses) != 0 {
		query += " WHERE " + strings.Join(clauses, " AND ")
	}
	query += order
	return query, args
}

func (r *racesRepo) scanRaces(rows *sql.Rows) ([]*racing.Race, error) {
	var races []*racing.Race

	for rows.Next() {
		var race racing.Race
		var advertisedStart time.Time
		var status string

		if err := rows.Scan(&race.Id, &race.MeetingId, &race.Name, &race.Number, &race.Visible, &advertisedStart, &status); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		ts, err := ptypes.TimestampProto(advertisedStart)
		if err != nil {
			return nil, err
		}

		race.AdvertisedStartTime = ts
		race.Status = status

		races = append(races, &race)
	}

	return races, nil
}
func (r *racesRepo) GetByID(raceID int64) (*racing.Race, error) {
	query := "SELECT id, meeting_id, name, number, visible, advertised_start_time, status FROM races WHERE id = ?"
	row := r.db.QueryRow(query, raceID)

	race, err := r.scanRace(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return race, nil
}

func (r *racesRepo) scanRace(row *sql.Row) (*racing.Race, error) {
	var race racing.Race
	var advertisedStart time.Time
	var status string

	err := row.Scan(&race.Id, &race.MeetingId, &race.Name, &race.Number, &race.Visible, &advertisedStart, &status)
	if err != nil {
		return nil, err
	}

	ts, err := ptypes.TimestampProto(advertisedStart)
	if err != nil {
		return nil, err
	}

	race.AdvertisedStartTime = ts
	race.Status = status

	return &race, nil
}
