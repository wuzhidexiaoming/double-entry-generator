package cmb

import "time"

const (
	// localTimeFmt set time format to utc+8
	localTimeFmt = "2006-01-02 15:04:05 -0700 CST"
)

// Statistics is the Statistics of the bill file.
type Statistics struct {
	UserID          string    `json:"user_id,omitempty"`
	Username        string    `json:"username,omitempty"`
	ParsedItems     int       `json:"parsed_items,omitempty"`
	Start           time.Time `json:"start,omitempty"`
	End             time.Time `json:"end,omitempty"`
	TotalInRecords  int       `json:"total_in_records,omitempty"`
	TotalInMoney    float64   `json:"total_in_money,omitempty"`
	TotalOutRecords int       `json:"total_out_records,omitempty"`
	TotalOutMoney   float64   `json:"total_out_money,omitempty"`
}

// Order is the single order.
type Order struct {
	PayDate           time.Time // 交易日期
	PayTime           time.Time // 交易时间
	Income            float64   // 收入
	Expenditure       float64   // 支出
	Balance           float64   // 余额
	TxType            TxType    // 交易类型
	TransactionRemark string    // 交易备注
	Type              OrderType // 收/支 (数据中无该列，推测而来)
}

// OrderType is the type of the order.
type OrderType string

const (
	OrderTypeSend    OrderType = "支出"
	OrderTypeRecv    OrderType = "收入"
	OrderTypeUnknown OrderType = "Unknown"
)

type TxType string

const (
	TxTypeSalary                   TxType = "薪资"
	TxTypeUnionPay                 TxType = "银联代付"
	TxTypeUnionPayOnline           TxType = "银联在线支付"
	TxTypeTransfer                 TxType = "转账汇款"
	TxTypeInterest                 TxType = "账户结息"
	TxTypeOnlinePayment            TxType = "网联协议支付"
	TxTypeOnlineRefund             TxType = "网联退款"
	TxTypeOnlinePaymentTransaction TxType = "网联付款交易"
	TxTypeNextOnlinePayment        TxType = "一网通支付"
	TxTypeCashAdvance              TxType = "预借现金"
	TxTypeConsume                  TxType = "消费"
	TxTypeRefund                   TxType = "退款"
	TxTypeRemittance               TxType = "汇入汇款"
	TxTypeRedemption               TxType = "朝朝宝赎回"
	TxTypePurchase                 TxType = "朝朝宝购买"
	TxTypeUnionPayQuickPayment     TxType = "银联快捷支付"
	TxTypeCreditCardRepayment      TxType = "信用卡还款"
	TxTypeUnknown                  TxType = "未知"
)
