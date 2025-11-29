package utils

import (
	"errors"
	"strconv"

	"github.com/Nozomi9967/wmss-customer-api/internal/types"
	"github.com/Nozomi9967/wmss-customer-api/model"
)

func BankCardToBankCardInfo(req *model.CustomerBankCard) (*types.BankCardInfo, error) {
	if req == nil {
		return nil, errors.New("input CustomerBankCard model cannot be nil")
	}

	var resp *types.BankCardInfo
	moneyStr := strconv.FormatFloat(req.CardBalance, 'f', -1, 64)
	resp = &types.BankCardInfo{
		CardId:         req.CardId,
		CustomerId:     req.CustomerId,
		BankCardNumber: req.BankCardNumber,
		BankName:       req.BankName,
		CardBalance:    moneyStr,
		IsVirtual:      int(req.IsVirtual),
		BindStatus:     req.BindStatus,
		BindTime:       req.BindTime.Format("2006-01-02 15:04:05"),
		CreateTime:     req.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime:     req.UpdateTime.Format("2006-01-02 15:04:05"),
	}
	if !req.UnbindTime.Valid {
		resp.UnbindTime = ""
	}
	return resp, nil

}

func CustomerToCustomerInfo(req *model.CustomerInfo) (*types.CustomerInfo, error) {
	if req == nil {
		return nil, errors.New("input CustomerInfo model cannot be nil")
	}

	var resp *types.CustomerInfo

	// 执行数据映射和格式化操作
	resp = &types.CustomerInfo{
		CustomerId:               req.CustomerId,
		CustomerName:             req.CustomerName,
		CustomerType:             req.CustomerType,
		IdType:                   req.IdType,
		IdNumber:                 req.IdNumber,
		RiskLevel:                req.RiskLevel,
		RiskEvaluationTime:       req.RiskEvaluationTime.Format("2006-01-02 15:04:05"),
		RiskEvaluationExpireTime: req.RiskEvaluationExpireTime.Format("2006-01-02 15:04:05"),
		ContactPhone:             req.ContactPhone.String,
		Email:                    req.Email.String,
		CreateTime:               req.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime:               req.UpdateTime.Format("2006-01-02 15:04:05"),
	}

	return resp, nil
}
