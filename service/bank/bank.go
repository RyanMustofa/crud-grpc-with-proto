package bank

import (
	"context"
	"database/sql"
	"grpc-learn/protobuf"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type BankService struct {
	protobuf.UnimplementedBankServiceServer
	DB *gorm.DB
}

type DateValue struct {
	UpdatedAt sql.NullTime `json:"updated_at,omitempty"`
	CreatedAt sql.NullTime `json:"created_at,omitempty"`
}

func (b *BankService) GetBanks(ctx context.Context, params *protobuf.Params) (*protobuf.Banks, error){
	var page int = 1
	var limit int = 10

	if(params.GetPage() != 0){
		page = int(params.GetPage())
	}
	if(params.GetLimit() != 0){
		limit = int(params.GetLimit())
	}

	offset := (page - 1) * limit

	var banks []*protobuf.Bank

	rows, err := b.DB.Table("banks as b").Limit(limit).Offset(offset).Select("b.id", "b.name", "b.created_at", "b.updated_at").Rows()

	if err != nil {
		return nil, status.Error(codes.Internal, "internal server")
	}

	defer rows.Close()

	for rows.Next() {
		var bank protobuf.Bank
		var dates DateValue
		if err := rows.Scan(&bank.Id, &bank.Name, &dates.CreatedAt, &dates.UpdatedAt); err != nil {
			log.Fatal("Failed" + err.Error())
		}


		if dates.CreatedAt.Valid {
			bank.CreatedAt = timestamppb.New(dates.CreatedAt.Time)
		}
		if dates.UpdatedAt.Valid {
			bank.UpdatedAt = timestamppb.New(dates.UpdatedAt.Time)
		}

		banks = append(banks, &bank)
	}


	response := &protobuf.Banks{
		Meta: &protobuf.Pagination{
			Page: 1,
			Limit: 10,
		},
		Data: banks,
	}
	return response, nil
}
