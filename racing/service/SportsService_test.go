package service

import (
	"context"
	"errors"
	"testing"

	pb "github.com/Kim-Hardie/entain-master/racing/proto/sports"
	"github.com/stretchr/testify/assert"
)

type MockMatchesRepo struct {
	Matches []*pb.Match
}

func (m *MockMatchesRepo) Init() error {
	return nil
}

func (m *MockMatchesRepo) List(filter *pb.MatchFilter) ([]*pb.Match, error) {
	var matches []*pb.Match
	for _, match := range m.Matches {
		if (filter.Sport == "" || match.Sport == filter.Sport) &&
			(filter.Stadium == "" || match.Stadium == filter.Stadium) {
			matches = append(matches, match)
		}
	}
	return matches, nil
}

func (m *MockMatchesRepo) GetByID(id int64) (*pb.Match, error) {
	for _, match := range m.Matches {
		if match.Id == id {
			return match, nil
		}
	}
	return nil, errors.New("Match not found")
}

func TestSportsService_ListMatches(t *testing.T) {
	// Create a mock MatchesRepo
	mockRepo := &MockMatchesRepo{
		Matches: []*pb.Match{
			{
				Id:      1,
				Name:    "Match 1",
				Stadium: "Stadium A",
				Sport:   "Sport A",
			},
			{
				Id:      2,
				Name:    "Match 2",
				Stadium: "Stadium B",
				Sport:   "Sport A",
			},
			{
				Id:      3,
				Name:    "Match 3",
				Stadium: "Stadium C",
				Sport:   "Sport B",
			},
		},
	}

	// Create a new SportsService instance with the mock repository
	service := NewSportsService(mockRepo)

	// Create a ListMatchesRequest with the filter by sport
	sportFilter := &pb.MatchFilter{
		Sport: "Sport A",
	}
	request := &pb.ListMatchesRequest{
		Filter: sportFilter,
	}

	// Call the ListMatches method
	response, err := service.ListMatches(context.Background(), request)
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Len(t, response.Matches, 2)

	// Create a ListMatchesRequest with the filter by stadium
	stadiumFilter := &pb.MatchFilter{
		Stadium: "Stadium B",
	}
	request = &pb.ListMatchesRequest{
		Filter: stadiumFilter,
	}

	// Call the ListMatches method
	response, err = service.ListMatches(context.Background(), request)
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Len(t, response.Matches, 1)
	assert.Equal(t, int64(2), response.Matches[0].Id)
}

func TestSportsService_GetMatchByID(t *testing.T) {
	// Create a mock MatchesRepo
	mockRepo := &MockMatchesRepo{
		Matches: []*pb.Match{
			{
				Id:      1,
				Name:    "Match 1",
				Stadium: "Stadium A",
				Sport:   "Sport A",
			},
			{
				Id:      2,
				Name:    "Match 2",
				Stadium: "Stadium B",
				Sport:   "Sport A",
			},
			{
				Id:      3,
				Name:    "Match 3",
				Stadium: "Stadium C",
				Sport:   "Sport B",
			},
		},
	}

	// Create a new SportsService instance with the mock repository
	service := NewSportsService(mockRepo)

	// Create a GetMatchByIDRequest for an existing match
	request := &pb.GetMatchByIDRequest{
		MatchId: 2,
	}

	// Call the GetMatchByID method
	response, err := service.GetMatchByID(context.Background(), request)
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, int64(2), response.Match.Id)
	assert.Equal(t, "Match 2", response.Match.Name)

}
