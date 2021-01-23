package invoice

import (
	"github.com/doCodingWhileCoding/GO-DB/pkg/invoiceheader"
	"github.com/doCodingWhileCoding/GO-DB/pkg/invoiceitem"
)

//Model of invoice
type Model struct {
	Header *invoiceheader.Model
	Items  invoiceitem.Models
}
