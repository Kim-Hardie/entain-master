package service

import (
	"context"

	"github.com/Kim-Hardie/entain-master/racing/db"
	pb "github.com/Kim-Hardie/entain-master/racing/proto/racing"
)

// RacingService implements the RacingServer interface.
type RacingService struct {
	pb.UnimplementedRacingServer // embed this
	racesRepo                    db.RacesRepo
}

// NewRacingService instantiates and returns a new RacingService.
func NewRacingService(racesRepo db.RacesRepo) *RacingService {
	return &RacingService{racesRepo: racesRepo}
}

func (s *RacingService) ListRaces(ctx context.Context, in *pb.ListRacesRequest) (*pb.ListRacesResponse, error) {
	races, err := s.racesRepo.List(in.Filter)
	if err != nil {
		return nil, err
	}

	return &pb.ListRacesResponse{Races: races}, nil
}
