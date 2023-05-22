package service

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"testing"
	"time"

	"github.com/Kim-Hardie/entain-master/racing/proto/racing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

//Testify Mock classes, used for its similarity with Java Mock classes im used to using.
// Brillianly useful https://github.com/stretchr/testify
type MockRacesRepo struct {
	mock.Mock
}

func (m *MockRacesRepo) Init() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockRacesRepo) List(filter *racing.ListRacesRequestFilter) ([]*racing.Race, error) {
	args := m.Called(filter)
	return args.Get(0).([]*racing.Race), args.Error(1)
}

func TestListRaces(t *testing.T) {
	raceTime, _ := ptypes.TimestampProto(time.Now())
	// Sample races to be used for testing
	races := []*racing.Race{
		{
			Id:                  1,
			MeetingId:           1,
			Name:                "Test Race 1",
			Number:              1,
			Visible:             true,
			AdvertisedStartTime: raceTime,
		},
		{
			Id:                  2,
			MeetingId:           2,
			Name:                "Test Race 2",
			Number:              2,
			Visible:             false,
			AdvertisedStartTime: raceTime,
		},
	}

	//Mock the repository and Create new Racing Service
	mockRepo := new(MockRacesRepo)
	s := NewRacingService(mockRepo)

	//Set expected outputs for List function Mock
	mockRepo.On("List", &racing.ListRacesRequestFilter{ShowOnlyVisible: &[]bool{true}[0]}).Return([]*racing.Race{races[0]}, nil)
	mockRepo.On("List", &racing.ListRacesRequestFilter{ShowOnlyVisible: &[]bool{false}[0]}).Return(races, nil)
	mockRepo.On("List", &racing.ListRacesRequestFilter{}).Return([]*racing.Race{races[0]}, nil)

	tests := []struct {
		name    string
		filter  *racing.ListRacesRequestFilter
		wantErr bool
		wantLen int
	}{
		//Test ShowOnlyVisible == true Returns 1 Result
		{
			name:    "ShowOnlyVisible is true",
			filter:  &racing.ListRacesRequestFilter{ShowOnlyVisible: &[]bool{true}[0]},
			wantErr: false,
			wantLen: 1,
		},
		//Test ShowOnlyVisible false Returns 2 Result
		{
			name:    "ShowOnlyVisible is false",
			filter:  &racing.ListRacesRequestFilter{ShowOnlyVisible: &[]bool{false}[0]},
			wantErr: false,
			wantLen: 2,
		},
		//Test ShowOnlyVisible nil returns 1 Result
		{
			name:    "ShowOnlyVisible is nil",
			filter:  &racing.ListRacesRequestFilter{},
			wantErr: false,
			wantLen: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := s.ListRaces(context.Background(), &racing.ListRacesRequest{Filter: tt.filter})
			if (err != nil) != tt.wantErr {
				t.Errorf("ListRaces() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.wantLen, len(resp.Races), "unexpected number of races")

			for _, race := range resp.Races {
				if tt.filter.ShowOnlyVisible != nil && *tt.filter.ShowOnlyVisible && !race.Visible {
					t.Errorf("expected all races to be visible")
				}
			}
		})
	}
}
