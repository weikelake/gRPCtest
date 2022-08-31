package src

import (
	"context"
	"gRPCtest/pkg/proto"
)

type Server struct {
	api.UnimplementedRusProfileServer
}

func (s *Server) GetRusProfileData(ctx context.Context, req *api.RpRequest) (*api.RpResponse, error) {
	Data, err := ParseRusProfile(req.GetInn())

	return &api.RpResponse{
		CompanyName: Data.CompanyName,
		DirectorFio: Data.DirectorFIO,
		Inn:         Data.Inn,
		Kpp:         Data.Kpp,
	}, err
}

//func (s *Server) mustEmbedUnimplementedRusProfileServer() {}
