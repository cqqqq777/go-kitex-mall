package pkg

import (
	"fmt"
	"github.com/cqqqq777/go-kitex-mall/shared/consts"
	"github.com/smartwalle/alipay/v3"
	"os"
	"strconv"
)

func readPrivateKey() (string, error) {
	content, err := os.ReadFile("secret")
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func newClient() (*alipay.Client, error) {
	privateKey, err := readPrivateKey()
	if err != nil {
		return nil, err
	}
	client, err := alipay.New(consts.AppId, privateKey, false)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func Pay(payID, orderID, amount int64) (string, error) {
	client, err := newClient()
	if err != nil {
		return "", err
	}
	if err = client.LoadAppPublicCertFromFile("appPublicCert.crt"); err != nil {
		return "", err
	}
	if err = client.LoadAliPayRootCertFromFile("alipayRootCert.crt"); err != nil {
		return "", err
	}
	if err = client.LoadAliPayPublicCertFromFile("alipayPublicCert.crt"); err != nil {
		return "", err
	}
	pay := alipay.TradePagePay{}
	pay.Subject = fmt.Sprintf("%d", payID)
	pay.OutTradeNo = strconv.FormatInt(orderID, 10)
	pay.ProductCode = "FAST_INSTANT_TRADE_PAY"
	pay.TotalAmount = fmt.Sprintf("%.2f", float64(amount)/100)
	pay.TimeoutExpress = "15m"
	pay.NotifyURL = ""
	url, err := client.TradePagePay(pay)
	if err != nil {
		return "", nil
	}
	payURL := url.String()
	return payURL, nil
}
