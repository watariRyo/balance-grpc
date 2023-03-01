package service

// ListIncomeAndExpenditure(context.Context, *IncomeAndExpenditureListRequest) (*IncomeAndExpenditureListResponse, error)
// GetIncomeAndExpenditure(context.Context, *IncomeAndExpenditureID) (*IncomeAndExpenditureResponse, error)
// RegisterIncomeAndExpenditure(context.Context, *IncomeAndExpenditureRequest) (*IncomeAndExpenditureResponse, error)
// UpdateIncomeAndExpenditure(context.Context, *IncomeAndExpenditureRequest) (*IncomeAndExpenditureResponse, error)
// DeleteIncomeAndExpenditure(context.Context, *IncomeAndExpenditureID) (*IncomeAndExpenditureID, error)

import (
	"context"
	"log"

	"github.com/watariRyo/balance/server/domain/repository"
	pb "github.com/watariRyo/balance/server/proto"
)

func NewIncomAndExpenditureService(r repository.AllRepository) *incomeAndExpenditureService {
	return &incomeAndExpenditureService{
		AllRepository: r,
	}
}

func (s *incomeAndExpenditureService) ListIncomeAndExpenditure(ctx context.Context, request *pb.IncomeAndExpenditureListRequest) (*pb.IncomeAndExpenditureListResponse, error) {
	log.Println("ListIncomeAndExpenditure was invoked.")

	return &pb.IncomeAndExpenditureListResponse{
		IncomeAndExpenditureList: []*pb.IncomeAndExpenditureResponse{},
	}, nil
}

func (s *incomeAndExpenditureService) GetIncomeAndExpenditure(ctx context.Context, incomeAndExpenditureID *pb.IncomeAndExpenditureID) (*pb.IncomeAndExpenditureResponse, error) {
	log.Println("GetIncomeAndExpenditure was invoked.")

	return &pb.IncomeAndExpenditureResponse{
		Id: 1,
		UserId: "userId",
		Amount: 100,
		OccurrenceDate: "2023-01-01",
		UserTagId: 1,
		Classification: "INCOME",
	}, nil
}

func (s *incomeAndExpenditureService) RegisterIncomeAndExpenditure(ctx context.Context, request *pb.IncomeAndExpenditureRequest) (*pb.IncomeAndExpenditureResponse, error) {
	log.Println("RegisterIncomeAndExpenditure was invoked.")

	return &pb.IncomeAndExpenditureResponse{
		Id: 1,
		UserId: "userId",
		Amount: 100,
		OccurrenceDate: "2023-01-01",
		UserTagId: 1,
		Classification: "INCOME",
	}, nil
}

func (s *incomeAndExpenditureService) UpdateIncomeAndExpenditure(ctx context.Context, request *pb.IncomeAndExpenditureRequest) (*pb.IncomeAndExpenditureResponse, error) {
	log.Println("UpdateIncomeAndExpenditure was invoked.")

	return &pb.IncomeAndExpenditureResponse{
		Id: 1,
		UserId: "userId",
		Amount: 100,
		OccurrenceDate: "2023-01-01",
		UserTagId: 1,
		Classification: "INCOME",
	}, nil
}

func (s *incomeAndExpenditureService) DeleteIncomeAndExpenditure(ctx context.Context, incomeAndExpenditureID *pb.IncomeAndExpenditureID) (*pb.IncomeAndExpenditureID, error) {
	log.Println("DeleteIncomeAndExpenditure was invoked.")

	return &pb.IncomeAndExpenditureID{
		Id: 1,
		UserId: "userId",
	}, nil
}