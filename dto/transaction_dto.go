package dto

import "git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/entity"

type TopUpRequest struct {
	Amount         float64 `json:"amount" validate:"required,numeric,gte=10000,lte=100000000"`
	SourceOfFundId uint    `json:"source_of_fund_id" validate:"numeric" default:"1"`
}

func (r *TopUpRequest) ReqToTransaction(userID int) entity.Transaction {
	return entity.Transaction{
		UserID:          uint(userID),
		Balance:         r.Amount,
		SourceOfFundsId: r.SourceOfFundId,
	}
}

type TopUpResponse struct {
	UserId          uint    `json:"user_id"`
	Balance         float64 `json:"balance"`
	SourceOfFundsId uint    `json:"source_of_funds_id"`
	Message         string  `json:"message,omitempty"`
}

func (r *TopUpResponse) TransactionToRes(transaction entity.Transaction) TopUpResponse {
	return TopUpResponse{
		UserId:          transaction.UserID,
		Balance:         transaction.Balance,
		SourceOfFundsId: transaction.SourceOfFundsId,
	}
}
