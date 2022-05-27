package types

type SuccessfulPayment struct {
	Currency string `json:"currency"`
	InvoicePayload string `json:"invoice_payload"`
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
	ProviderPaymentChargeId string `json:"provider_payment_charge_id"`
	ShippingOptionId string `json:"shipping_option_id,omitempty"`
	TelegramPaymentChargeId string `json:"telegram_payment_charge_id"`
	TotalAmount int `json:"total_amount"`
}
