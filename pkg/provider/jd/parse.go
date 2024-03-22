package jd

import (
	"strings"
)

func (h *Jd) translateToOrders(arr []string) error {
	// trim strings
	for idx, a := range arr {
		a = strings.Trim(a, " ")
		a = strings.Trim(a, "\t")
		arr[idx] = a
	}
	var bill Order

	h.Orders = append(h.Orders, bill)
	return nil
}
