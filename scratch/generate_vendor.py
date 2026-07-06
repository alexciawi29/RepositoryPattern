import os

vendor_go_content = '''package Models

import (
	"gorm.io/gorm"
)

type Vendor struct {
	gorm.Model

	// Step 1
	RegistrationPurposes   []string `json:"RegistrationPurposes" gorm:"serializer:json"`
	VendorType             string   `json:"VendorType"`
	Details                string   `json:"Details"`
	OrderType              string   `json:"OrderType"`
	ShipmentMethod         string   `json:"ShipmentMethod"`
	ReimburseTo            string   `json:"ReimburseTo"`
	DepartmentDivision     string   `json:"DepartmentDivision"`
	DeliveryType           string   `json:"DeliveryType"`
	WorkOrderRemarks       string   `json:"WorkOrderRemarks"`
	DiverWork              string   `json:"DiverWork"`
	ExistingVendorName     string   `json:"ExistingVendorName"`
	DisableExistingVendor  string   `json:"DisableExistingVendor"`
	ChangeReason           string   `json:"ChangeReason"`

	// Step 2
	PaymentTermType             string   `json:"PaymentTermType"`
	CreditTermDays              string   `json:"CreditTermDays"`
	AdvanceTermType             *string  `json:"AdvanceTermType"`
	CustomPaymentTerm           *string  `json:"CustomPaymentTerm"`
	TaxType                     string   `json:"TaxType"`
	PpnType                     string   `json:"PpnType"`
	DefaultDiscount             string   `json:"DefaultDiscount"`
	Step2DeliveryType           string   `json:"Step2DeliveryType"`
	DeliveryMethod              string   `json:"DeliveryMethod"`
	DeliveryCityAreas           []string `json:"DeliveryCityAreas" gorm:"serializer:json"`
	Country                     string   `json:"Country"`
	StateCity                   string   `json:"StateCity"`
	InformationSource           string   `json:"InformationSource"`
	InternalPicFullName         string   `json:"InternalPicFullName"`
	InternalPicJobTitle         string   `json:"InternalPicJobTitle"`
	InternalPicWhatsappCode     string   `json:"InternalPicWhatsappCode"`
	InternalPicWhatsappNumber   string   `json:"InternalPicWhatsappNumber"`
	RecommenderVendorName       *string  `json:"RecommenderVendorName"`
	RecommenderPicFullName      *string  `json:"RecommenderPicFullName"`
	RecommenderPicJobTitle      *string  `json:"RecommenderPicJobTitle"`
	RecommenderPicWhatsappCode  *string  `json:"RecommenderPicWhatsappCode"`
	RecommenderPicWhatsappNumber *string  `json:"RecommenderPicWhatsappNumber"`
	ContactPersonFullName       string   `json:"ContactPersonFullName"`
	ContactPersonJobTitle       string   `json:"ContactPersonJobTitle"`
	ContactPersonWhatsappCode   string   `json:"ContactPersonWhatsappCode"`
	ContactPersonWhatsappNumber string   `json:"ContactPersonWhatsappNumber"`
	SubmissionType              string   `json:"SubmissionType"`

	// Step 3
	BusinessType       string `json:"BusinessType"`
	CompanyName        string `json:"CompanyName"`
	YearEstablished    string `json:"YearEstablished"`
	Website            string `json:"Website"`

	MainContactName      string `json:"MainContactName"`
	MainContactTitle     string `json:"MainContactTitle"`
	MainContactEmail     string `json:"MainContactEmail"`
	MainContactPhoneCode string `json:"MainContactPhoneCode"`
	MainContactPhone     string `json:"MainContactPhone"`
	
	InquirySameAsMain  bool `json:"InquirySameAsMain"`
	
	InquiryContactName      *string `json:"InquiryContactName"`
	InquiryContactTitle     *string `json:"InquiryContactTitle"`
	InquiryContactEmail     *string `json:"InquiryContactEmail"`
	InquiryContactPhoneCode *string `json:"InquiryContactPhoneCode"`
	InquiryContactPhone     *string `json:"InquiryContactPhone"`

	PoSameAsMain       bool `json:"PoSameAsMain"`
	
	PoContactName      *string `json:"PoContactName"`
	PoContactTitle     *string `json:"PoContactTitle"`
	PoContactEmail     *string `json:"PoContactEmail"`
	PoContactPhoneCode *string `json:"PoContactPhoneCode"`
	PoContactPhone     *string `json:"PoContactPhone"`

	PaymentSameAsMain  bool `json:"PaymentSameAsMain"`
	
	PaymentContactName      *string `json:"PaymentContactName"`
	PaymentContactTitle     *string `json:"PaymentContactTitle"`
	PaymentContactEmail     *string `json:"PaymentContactEmail"`
	PaymentContactPhoneCode *string `json:"PaymentContactPhoneCode"`
	PaymentContactPhone     *string `json:"PaymentContactPhone"`

	StreetAddress   string `json:"StreetAddress"`
	Province        string `json:"Province"`
	CityArea        string `json:"CityArea"`
	PostalCode      string `json:"PostalCode"`
	FullAddress     string `json:"FullAddress"`
	MapsUrl         string `json:"MapsUrl"`
	OfficePhoneCode string `json:"OfficePhoneCode"`
	OfficePhone     string `json:"OfficePhone"`

	// Step 4
	IndustryType       string   `json:"IndustryType"`
	CustomIndustryType *string  `json:"CustomIndustryType"`
	CompanySize        string   `json:"CompanySize"`
	VendorBackground   string   `json:"VendorBackground"`
	Certifications     []string `json:"Certifications" gorm:"serializer:json"`

	// Step 5
	HasCompanyBankAccount bool    `json:"HasCompanyBankAccount"`
	BankName              string  `json:"BankName"`
	CustomBankName        *string `json:"CustomBankName"`
	BankAccountNumber     string  `json:"BankAccountNumber"`
	BankAccountHolder     string  `json:"BankAccountHolder"`
	BankCurrency          string  `json:"BankCurrency"`

	// Step 6
	IsIntegrityPactAccepted bool   `json:"IsIntegrityPactAccepted"`
	IntegrityPactSignerName string `json:"IntegrityPactSignerName"`
	IntegrityPactSignerRole string `json:"IntegrityPactSignerRole"`

	// Attachments
	CustomPaymentTermApproval *string `json:"CustomPaymentTermApproval"`
	InformationScreenshot     *string `json:"InformationScreenshot"`
	CompanyStampInvoiceHeader *string `json:"CompanyStampInvoiceHeader"`
	FrontOfficePhoto          *string `json:"FrontOfficePhoto"`
	InsideOfficePhoto         *string `json:"InsideOfficePhoto"`
	OfficeVideo               *string `json:"OfficeVideo"`
	BankAccountProof          *string `json:"BankAccountProof"`
	OwnerIdentityCard         *string `json:"OwnerIdentityCard"`
}
'''

with open(r'D:\AlexCiawi\mobile\ReposityPattern\Models\Vendor.go', 'w', encoding='utf-8') as f:
    f.write(vendor_go_content.replace('`', ''))
