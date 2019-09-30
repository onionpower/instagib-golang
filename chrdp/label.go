package chrdp

type Label struct {
	PackDate     string
	ReasonId     int64
	Seller       string
	OrderNum     string
	OrderPkgL    string
	OrderPkgR    string
	Client       string
	ClientPhone  string
	SrtRoute     string
	Destination  string
	Contractor   string
	DestAddr     string
	Width        int64
	Height       int64
	Length       int64
	Count        int64
	Weight       int64
	LowzoneId    int64
	DeliveryDate string
}

func NewLabel() Label {
	return Label{
		Width: 999,
	}
}
