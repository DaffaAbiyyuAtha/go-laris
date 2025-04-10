package lib

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func SetupMidtrans() {
	midtrans.ServerKey = "SB-Mid-server-hBkQqgnTCyEqJds6OMetyzvc" // pakai server key kamu
	midtrans.Environment = midtrans.Sandbox
}

func CreateSnapRequest(orderID string, amount int64, customerName, customerEmail string) (*snap.Response, error) {
	SetupMidtrans()

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: amount,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: customerName,
			Email: customerEmail,
		},
	}

	return snap.CreateTransaction(req)
}
