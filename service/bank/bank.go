package bank

import (
	"context"
	"database/sql"
	"grpc-learn/protobuf"
	"log"
	"time"

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

type Body struct {
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type BankStruct struct {
	Id int `json:"id"`
	Name string `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
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

func (b *BankService) PostBanks(ctx context.Context, body *protobuf.Body) (*protobuf.Status, error) {
	var name string
	if body.GetName() != "" {
		name = body.GetName()
	}
	payload := Body{
		Name: name,
		CreatedAt: time.Now(),
	}

	err := b.DB.Table("banks").Create(&payload)
	if err.Error != nil {
		return &protobuf.Status{
			Status: false,
			Message: "Failed " + err.Error.Error(),
		}, status.Error(codes.Internal, "Failed internal " + err.Error.Error())
	}
	response := &protobuf.Status{
		Status: true,
		Message: "Success Add",
	}

	return response, nil
}

func (b *BankService) UpdateBanks(ctx context.Context, id *protobuf.Id) (*protobuf.Status, error){
	var id_data int32 
	var name string
	if id.GetId() != 0 {
		id_data = id.GetId()
	}
	if id.GetName() != ""{
		name = id.GetName()
	}

	var dataSelect BankStruct
	results := b.DB.Table("banks").Where(map[string]interface{}{"id": id_data}).First(&dataSelect);
	if  results.Error != nil {
		return &protobuf.Status{
			Status: false,
			Message: "Failed :: " + results.Error.Error(),
		}, nil
	}

	dataSelect.Name = name

	if errorS := results.Save(&dataSelect); errorS.Error != nil {
		return &protobuf.Status{
			Status: false,
			Message: "Failed :: " + errorS.Error.Error(),
		}, nil
	}
	

	response := &protobuf.Status{
		Status: true,
		Message: "Proto Update",
	}

	return response, nil
}

func (b *BankService) DeleteBanks(ctx context.Context, params *protobuf.Id) (*protobuf.Status, error){
	var bank BankStruct
	
	var id int32
	if params.GetId() != 0 {
		id = params.GetId()
	}

	results := b.DB.Table("banks").Where(map[string]interface{}{"id": id}).Delete(&bank)
	if results.Error != nil {
		return &protobuf.Status{
			Status: false,
			Message: "Failed :: " + results.Error.Error(),
		}, nil
	}
	response := &protobuf.Status{
		Status: true,
		Message: "Delete Success",
	}

	return response, nil
}
