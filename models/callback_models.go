package models

import "time"

type ConsentUpdate struct {
	Ver       string    `json:"ver"`
	Timestamp time.Time `json:"timestamp"`
	Txnid     string    `json:"txnid"`
	Notifier  struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	} `json:"Notifier"`
	ConsentStatusNotification struct {
		ConsentID     string `json:"consentId"`
		ConsentStatus string `json:"ConsentStatus"`
	} `json:"ConsentStatusNotification"`
}

type DataFlowUpdate struct {
	Ver       string    `json:"ver"`
	Timestamp time.Time `json:"timestamp"`
	Txnid     string    `json:"txnid"`
	Notifier  struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	} `json:"Notifier"`
	FIStatusNotification struct {
		SessionID        string `json:"sessionId"`
		SessionStatus    string `json:"sessionStatus"`
		FIStatusResponse []struct {
			FipID    string `json:"fipID"`
			Accounts []struct {
				LinkRefNumber string `json:"linkRefNumber"`
				FIStatus      string `json:"FIStatus"`
				Description   string `json:"description"`
			} `json:"Accounts"`
		} `json:"FIStatusResponse"`
	} `json:"FIStatusNotification"`
}
