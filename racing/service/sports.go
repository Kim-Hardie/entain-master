package service

import (
	"context"

	"github.com/Kim-Hardie/entain-master/racing/db"
	pb "github.com/Kim-Hardie/entain-master/racing/proto/sports"
)

// SportsService implements the SportsServer interface.
type SportsService struct {
	pb.UnimplementedSportsServer
	matchesRepo db.MatchesRepo
}

// NewSportsService instantiates and returns a new SportsService.
func NewSportsService(matchesRepo db.MatchesRepo) *SportsService {
	return &SportsService{matchesRepo: matchesRepo}
}

func (s *SportsService) ListMatches(ctx context.Context, in *pb.ListMatchesRequest) (*pb.ListMatchesResponse, error) {
	// Call the repository to get the list of matches based on the filter
	matches, err := s.matchesRepo.List(in.Filter)
	if err != nil {
		return nil, err
	}

	return &pb.ListMatchesResponse{Matches: matches}, nil
}

func (s *SportsService) GetMatchByID(ctx context.Context, req *pb.GetMatchByIDRequest) (*pb.GetMatchByIDResponse, error) {
	match, err := s.matchesRepo.GetByID(req.MatchId)
	if err != nil {
		return nil, err
	}

	return &pb.GetMatchByIDResponse{
		Match: match,
	}, nil
}
