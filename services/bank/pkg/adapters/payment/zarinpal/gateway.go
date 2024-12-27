package zarinpal

import (
	moneyDomain "bank_service/internal/common/types"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	paymentRequestURL = "https://sandbox.zarinpal.com/pg/v4/payment/request.json"
	paymentVerifyURL  = "https://sandbox.zarinpal.com/pg/v4/payment/verify.json"
	paymentStartURL   = "https://sandbox.zarinpal.com/pg/StartPay/"
)

type ZarinpalGateway struct {
	merchantID  string
	callbackURL string
	client      *http.Client
}

type PaymentRequestData struct {
	MerchantID  string                 `json:"merchant_id"`
	Amount      int64                  `json:"amount"`
	CallbackURL string                 `json:"callback_url"`
	Description string                 `json:"description"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

type PaymentRequestResponse struct {
	Data struct {
		Authority string `json:"authority"`
		Fee       int    `json:"fee"`
		FeeType   string `json:"fee_type"`
		Code      int    `json:"code"`
		Message   string `json:"message"`
	} `json:"data"`
	Errors []interface{} `json:"errors"`
}

type VerifyRequest struct {
	MerchantID string `json:"merchant_id"`
	Authority  string `json:"authority"`
	Amount     int64  `json:"amount"`
}

type VerifyResponse struct {
	Data struct {
		Code        int         `json:"code"`
		RefID       interface{} `json:"ref_id"`
		CardPan     interface{} `json:"card_pan"`
		Wages       interface{} `json:"wages"`
		Message     interface{} `json:"message"`
		CardHash    interface{} `json:"card_hash"`
		FeeType     interface{} `json:"fee_type"`
		Fee         interface{} `json:"fee"`
		ShaparakFee interface{} `json:"shaparak_fee"`
		OrderId     interface{} `json:"order_id"`
	} `json:"data"`
	Errors *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"errors,omitempty"`
}
type T struct {
	Data struct {
		Wages       interface{} `json:"wages"`
		Code        int         `json:"code"`
		Message     string      `json:"message"`
		CardHash    string      `json:"card_hash"`
		CardPan     string      `json:"card_pan"`
		RefId       int         `json:"ref_id"`
		FeeType     string      `json:"fee_type"`
		Fee         int         `json:"fee"`
		ShaparakFee int         `json:"shaparak_fee"`
		OrderId     interface{} `json:"order_id"`
	} `json:"data"`
	Errors []interface{} `json:"errors"`
}

func NewZarinpalGateway(merchantID, callbackURL string) *ZarinpalGateway {
	return &ZarinpalGateway{
		merchantID:  merchantID,
		callbackURL: callbackURL,
		client:      &http.Client{},
	}
}

func (g *ZarinpalGateway) InitiatePayment(ctx context.Context, amount *moneyDomain.Money, metadata map[string]interface{}) (string, error) {
	reqData := PaymentRequestData{
		MerchantID:  g.merchantID,
		Amount:      amount.Amount,
		CallbackURL: g.callbackURL,
		Description: "Payment Description",
		Metadata:    metadata,
	}

	jsonData, err := json.Marshal(reqData)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request data: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", paymentRequestURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "ZarinPal Rest Api v4")

	resp, err := g.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var result PaymentRequestResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("response error: %v", string(body))
	}

	if result.Errors != nil {
		if len(result.Errors) != 0 {
			return "", fmt.Errorf("payment request error: %v", result.Errors)

		}
	}

	if result.Data.Code != 100 {
		return "", fmt.Errorf("unexpected response code: %d", result.Data.Code)
	}

	return paymentStartURL + result.Data.Authority, nil
}

func (g *ZarinpalGateway) VerifyPayment(ctx context.Context, referenceID string, amount int64) (bool, error) {

	reqData := VerifyRequest{
		MerchantID: g.merchantID,
		Authority:  referenceID,
		Amount:     amount,
	}

	jsonData, err := json.Marshal(reqData)
	if err != nil {
		return false, fmt.Errorf("failed to marshal verify request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", paymentVerifyURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return false, fmt.Errorf("failed to create verify request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "ZarinPal Rest Api v4")

	resp, err := g.client.Do(req)
	if err != nil {
		return false, fmt.Errorf("verify request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("failed to read verify response: %w", err)
	}

	var result VerifyResponse
	_ = json.Unmarshal(body, &result)

	if result.Errors != nil && result.Errors.Code != 0 {

		return false, fmt.Errorf("verify error: %d - %s", result.Errors.Code, result.Errors.Message)
	}

	return result.Data.Code == 100, nil
}

func (g *ZarinpalGateway) GetPaymentStatus(ctx context.Context, referenceID string) (string, error) {

	return "failed", fmt.Errorf("this service not implemented")
}

func (g *ZarinpalGateway) RefundPayment(ctx context.Context, referenceID string, amount *moneyDomain.Money) error {

	return fmt.Errorf("automatic refunds not supported by Zarinpal")
}
