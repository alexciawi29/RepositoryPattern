package Models

import (
	"gorm.io/gorm"
)

type InternalApproval struct {
	gorm.Model

	VendorId uint `json:"VendorId"`

	// Section 1: Internal Procurement Data
	TeamName              string  `json:"TeamName"`
	PicName               string  `json:"PicName"`
	VendorType            string  `json:"VendorType"` // PKP, Non-PKP, Overseas
	ProcuredGoodsServices string  `json:"ProcuredGoodsServices"`
	AgreedPaymentTerm     string  `json:"AgreedPaymentTerm"`
	AgreedPaymentTermCheck bool   `json:"AgreedPaymentTermCheck"`
	NegotiationHistory    string  `json:"NegotiationHistory"`
	EstimatedTotalPoValue float64 `json:"EstimatedTotalPoValue"`
	Currency              string  `json:"Currency"`

	// Section 2: Bank Account & Sourcing
	BankAccountHolder      string `json:"BankAccountHolder"`
	BankAccountHolderCheck bool   `json:"BankAccountHolderCheck"`
	FraudCheckStatus       string `json:"FraudCheckStatus"`
	FraudCheckStatusCheck  bool   `json:"FraudCheckStatusCheck"`
	DiscoverySource        string `json:"DiscoverySource"`
	VendorHistory          string `json:"VendorHistory"`
	NumberOfEmployees      string `json:"NumberOfEmployees"`
	ContactPerson          string `json:"ContactPerson"`

	// Section 3: Verification & Legal Status
	CompanyWebsite           string `json:"CompanyWebsite"`
	IsCvPersonalAccount      *bool  `json:"IsCvPersonalAccount"`
	NpwpDocument             string `json:"NpwpDocument"`
	NpwpStatus               string `json:"NpwpStatus"`
	NibDocument              string `json:"NibDocument"`
	NibStatus                string `json:"NibStatus"`
	SppkpDocument            string `json:"SppkpDocument"`
	SppkpStatus              string `json:"SppkpStatus"`
	UsccDocument             string `json:"UsccDocument"`
	UsccStatus               string `json:"UsccStatus"`
	ThirdPartyValidation     string `json:"ThirdPartyValidation"`
	AlibabaStatus            string `json:"AlibabaStatus"`
	PersonalAccountStatement string `json:"PersonalAccountStatement"`
	PersonalAccountStatus    string `json:"PersonalAccountStatus"`

	// Section 4: Evidence Attachments
	GoogleStreetView          string `json:"GoogleStreetView"`
	GoogleStreetViewStatus    string `json:"GoogleStreetViewStatus"`
	GoogleStreetViewNotes     string `json:"GoogleStreetViewNotes"`
	PhysicalLocationPhoto     string `json:"PhysicalLocationPhoto"`
	AccreditationCertificates string `json:"AccreditationCertificates"`
	AdditionalAttachmentsJson string `json:"AdditionalAttachmentsJson"`

	// Highlights
	HighlightedFields string `json:"HighlightedFields"`
}
