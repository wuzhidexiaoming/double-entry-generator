package jd

import "time"

type Config struct {
	Rules []Rule `mapstructure:"rules,omitempty"`
}

// Rule is the type for match rules.
type Rule struct {
	Peer          *string   `mapstructure:"peer,omitempty"`
	Item          *string   `mapstructure:"item,omitempty"`
	Type          *string   `mapstructure:"type,omitempty"`
	Method        *string   `mapstructure:"method,omitempty"`
	PayTime       time.Time `json:"payTime,omitempty"` // 交易时间
	MethodAccount *string   `mapstructure:"methodAccount,omitempty"`
	TargetAccount *string   `mapstructure:"targetAccount,omitempty"`
	FullMatch     bool      `mapstructure:"fullMatch,omitempty"` // default: false
	Ignore        bool      `mapstructure:"ignore,omitempty"`    // default: false
	DealNo        string    `json:"dealNo,omitempty"`            // 交易订单号
	MerchantId    string    `json:"merchantId,omitempty"`        // 商家订单号
	Status        string    `json:"status,omitempty"`            // 交易状态
	Category      string    `json:"category,omitempty"`          // 交易分类
}
