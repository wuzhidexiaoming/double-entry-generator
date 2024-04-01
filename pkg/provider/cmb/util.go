package cmb

import "strings"

func getOrderType(ot string) OrderType {
	switch ot {
	case string(OrderTypeRecv):
		return OrderTypeRecv
	case string(OrderTypeSend):
		return OrderTypeSend
	default:
		return OrderTypeUnknown
	}
}

func getTxType(tt string) TxType {
	if strings.Contains(tt, string(TxTypeSalary)) {
		return TxTypeSalary
	} else if strings.Contains(tt, string(TxTypeConsume)) {
		return TxTypeConsume
	} else if strings.Contains(tt, string(TxTypeTransfer)) {
		return TxTypeTransfer
	} else if strings.Contains(tt, string(TxTypeUnionPay)) {
		return TxTypeUnionPay
	} else if strings.Contains(tt, string(TxTypeUnionPayOnline)) {
		return TxTypeUnionPayOnline
	} else if strings.Contains(tt, string(TxTypeInterest)) {
		return TxTypeInterest
	} else if strings.Contains(tt, string(TxTypeRefund)) {
		return TxTypeRefund
	} else if strings.Contains(tt, string(TxTypeOnlinePayment)) {
		return TxTypeOnlinePayment
	} else if strings.Contains(tt, string(TxTypeOnlineRefund)) {
		return TxTypeOnlineRefund
	} else if strings.Contains(tt, string(TxTypeOnlinePaymentTransaction)) {
		return TxTypeOnlinePaymentTransaction
	} else if strings.Contains(tt, string(TxTypeNextOnlinePayment)) {
		return TxTypeNextOnlinePayment
	} else if strings.Contains(tt, string(TxTypeCashAdvance)) {
		return TxTypeCashAdvance
	} else if strings.Contains(tt, string(TxTypeRemittance)) {
		return TxTypeRemittance
	} else if strings.Contains(tt, string(TxTypeRedemption)) {
		return TxTypeRedemption
	} else if strings.Contains(tt, string(TxTypePurchase)) {
		return TxTypePurchase
	} else if strings.Contains(tt, string(TxTypeUnionPayQuickPayment)) {
		return TxTypeUnionPayQuickPayment
	} else if strings.Contains(tt, string(TxTypeCreditCardRepayment)) {
		return TxTypeCreditCardRepayment
	} else {
		return TxTypeUnknown
	}
}
