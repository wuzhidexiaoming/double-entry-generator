package jd

import "github.com/deb-sig/double-entry-generator/pkg/ir"

func (h *Jd) convertToIR() *ir.IR {
	i := ir.New()
	for _, o := range h.Orders {
		irO := ir.Order{
			Peer:           "Jd",
			PayTime:        o.PayTime,
			TxTypeOriginal: o.TxTypeOriginal,
			Type:           convertType(o.Type),
			TypeOriginal:   string(o.Type),
			Money:          o.Money,
		}
		i.Orders = append(i.Orders, irO)
	}
	return i
}

func convertType(t OrderType) ir.Type {
	switch t {
	case OrderTypeSend:
		return ir.TypeSend
	case OrderTypeRecv:
		return ir.TypeRecv
	default:
		return ir.TypeUnknown
	}
}
