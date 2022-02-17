package ch20_1_fedex

import (
	"fmt"
)

type FedexSender struct {
}

func (f *FedexSender) Send(parcel string) {
	fmt.Printf("Fedex sends %v parcel\n", parcel)
}
