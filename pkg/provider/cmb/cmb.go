package cmb

import (
	"encoding/csv"
	"fmt"
	"github.com/deb-sig/double-entry-generator/pkg/io/reader"
	"github.com/deb-sig/double-entry-generator/pkg/ir"
	"io"
	"log"
)

type Cmb struct {
	Statistics Statistics `json:"statistics,omitempty"`
	LineNum    int        `json:"line_num,omitempty"`
	Orders     []Order    `json:"orders,omitempty"`
}

func New() *Cmb {
	return &Cmb{
		Statistics: Statistics{},
		LineNum:    0,
		Orders:     make([]Order, 0),
	}
}

// Translate translates the alipay bill records to IR.
func (cmb *Cmb) Translate(filename string) (*ir.IR, error) {
	log.SetPrefix("[Provider-Cmb] ")

	billReader, err := reader.GetReader(filename)
	if err != nil {
		return nil, fmt.Errorf("can't get bill reader, err: %v", err)
	}

	csvReader := csv.NewReader(billReader)
	csvReader.LazyQuotes = true
	// If FieldsPerRecord is negative, no check is made and records
	// may have a variable number of fields.
	csvReader.FieldsPerRecord = -1

	for {
		line, err := csvReader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		cmb.LineNum++
		if cmb.LineNum <= 17 {
			// The first 17 lines are useless for us.
			continue
		}

		err = cmb.translateToOrders(line)
		if err != nil {
			return nil, fmt.Errorf("Failed to translate bill: line %d: %v",
				cmb.LineNum, err)
		}
	}
	log.Printf("Finished to parse the file %s", filename)
	return cmb.convertToIR(), nil
}
