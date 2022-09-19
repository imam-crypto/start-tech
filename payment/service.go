package payment

import (
	"log"

	"pustaka-api/user"
	"strconv"

	"github.com/veritrans/go-midtrans"
)

type service struct {
}
type Service interface {
	GetPaymentURL(tr TransactionPayment, user user.User) (string, error)
}

func NewServicePayment() *service {
	return &service{}
}
func (s *service) GetPaymentURL(tr TransactionPayment, user user.User) (string, error) {
	// midClient := midtrans.midClient()
	// midclient.ServerKey = "YOUR-VT-SERVER-KEY"
	// midclient.ClientKey = "YOUR-VT-CLIENT-KEY"
	// midclient.APIEnvType = midtrans.Sandbox

	// snapGateway := midtrans.SnapGateway{
	// 	Client: midClient,
	// }
	// snapReq := &midtrans.SnapReq{
	// 	CustomerDetail: &midtrans.CustDetail{
	// 		Fname: user.First_name,
	// 		Email: user.Email,
	// 	},
	// 	TransactionDetails: midtrans.TransactionDetails{
	// 		OrderID:  strconv.itoa(tr.ID),
	// 		GrossAmt: int64(tr.Amount),
	// 	},
	// }

	// response, err := snapGateway.GetToken(snapReq)
	// if err != nil {
	// 	return "", err
	// }
	// return response.RedirectURL, nil

	midclient := midtrans.NewClient()
	midclient.ServerKey = "YOUR-VT-SERVER-KEY"
	midclient.ClientKey = "YOUR-VT-CLIENT-KEY"
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}
	snapReq := &midtrans.SnapReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(tr.ID),
			GrossAmt: int64(tr.Amount),
		},
		CustomerDetail: &midtrans.CustDetail{
			FName: user.First_name,

			Email: user.Email,
		},
	}

	log.Println("GetToken:")
	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}
	return snapTokenResp.RedirectURL, nil
}
