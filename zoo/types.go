package zoo

import "time"

type PostGateAuthBody struct {
	ID        string `json:"id"`
	NFCID     string `json:"nfc_id"`
	Direction string `json:"direction"`
}

type PostGateAuthResponse struct {
	Data struct {
		ID            int       `json:"id"`
		NFCSN         string    `json:"nfc_sn"`
		DeviceID      string    `json:"device_id"`
		DeviceType    string    `json:"device_type"`
		IPAddress     string    `json:"ip_address"`
		Action        string    `json:"action"`
		AccessGranted bool      `json:"access_granted"`
		CreatedAt     time.Time `json:"created_at"`
	} `json:"data"`

	Message string `json:"message"`
}
