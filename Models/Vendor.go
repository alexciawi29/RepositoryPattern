package Models

import (
	"gorm.io/gorm"
)

type Vendor struct {
	gorm.Model

	ReferenceNumber string `json:"ReferenceNumber"`
	Deleted string `json:"Deleted" gorm:"default:'F'"`

	// Step 1
	RegistrationPurpose   []string `json:"RegistrationPurpose" gorm:"serializer:json"`
	VendorType             string   `json:"VendorType"`
	Detail                 string   `json:"Detail"`
	OrderType              string   `json:"OrderType"`
	ShipmentMethod         string   `json:"ShipmentMethod"`
	ReimburseTo            string   `json:"ReimburseTo"`
	DepartmentDivision     string   `json:"DepartmentDivision"`
	DeliveryType           string   `json:"DeliveryType"`
	WorkOrderRemark       string   `json:"WorkOrderRemark"`
	DiverWork              string   `json:"DiverWork"`
	ExistingVendorName     string   `json:"ExistingVendorName"`
	DisableExistingVendor  string   `json:"DisableExistingVendor"`
	ChangeReason           string   `json:"ChangeReason"`

	// Step 2
	PaymentTermType             string   `json:"PaymentTermType"`
	CreditTermDay              string   `json:"CreditTermDay"`
	AdvanceTermType             *string  `json:"AdvanceTermType"`
	CustomPaymentTerm           *string  `json:"CustomPaymentTerm"`
	TaxType                     string   `json:"TaxType"`
	PpnType                     string   `json:"PpnType"`
	DefaultDiscount             string   `json:"DefaultDiscount"`
	Step2DeliveryType           string   `json:"Step2DeliveryType"`
	DeliveryMethod              string   `json:"DeliveryMethod"`
	DeliveryCityArea           []string `json:"DeliveryCityArea" gorm:"serializer:json"`
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
	MapUrl         string `json:"MapUrl"`
	OfficePhoneCode string `json:"OfficePhoneCode"`
	OfficePhone     string `json:"OfficePhone"`

	// Step 4
	IndustryType       string   `json:"IndustryType"`
	CustomIndustryType *string  `json:"CustomIndustryType"`
	CompanySize        string   `json:"CompanySize"`
	VendorBackground   string   `json:"VendorBackground"`
	Certification     []string `json:"Certification" gorm:"serializer:json"`

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
	CompanyBankAccountStatement *string `json:"CompanyBankAccountStatement"`
	OwnerIdentityCard         *string `json:"OwnerIdentityCard"`
	SppkpDocument             *string `json:"SppkpDocument"`
	SkptDocument              *string `json:"SkptDocument"`
	NibDocument               *string `json:"NibDocument"`
	NpwpDocument              *string `json:"NpwpDocument"`
	DeedOfEstablishmentDocument *string `json:"DeedOfEstablishmentDocument"`
	DeedOfAmendmentDocument   *string `json:"DeedOfAmendmentDocument"`
	DirectorIdCardDocument    *string `json:"DirectorIdCardDocument"`
	CompanyProfileDocument    *string `json:"CompanyProfileDocument"`

	// NPWP Extended Details
	NpwpNumber         *string `json:"NpwpNumber"`
	NpwpTaxOffice      *string `json:"NpwpTaxOffice"`
	NpwpRegisteredDate *string `json:"NpwpRegisteredDate"`

	// NIB Extended Details
	NibNumber     *string `json:"NibNumber"`
	NibAddress    *string `json:"NibAddress"`
	NibPhone      *string `json:"NibPhone"`
	NibEmail      *string `json:"NibEmail"`
	NibIssuedDate *string `json:"NibIssuedDate"`

	// Internal Review Checklist States
	ChecklistSummary      *string `json:"ChecklistSummary"`
	ReviewIsLocalBusiness *bool   `json:"ReviewIsLocalBusiness"`
	ReviewIsShipment      *bool   `json:"ReviewIsShipment"`
	ReviewIsInternal      *bool   `json:"ReviewIsInternal"`
	ReviewIsThirdParty    *bool   `json:"ReviewIsThirdParty"`
	ReviewIsForeignVendor *bool   `json:"ReviewIsForeignVendor"`
	ReviewIsBunker        *bool   `json:"ReviewIsBunker"`
	ReviewIsGovernment    *bool   `json:"ReviewIsGovernment"`
	ReviewIsReimburse     *bool   `json:"ReviewIsReimburse"`
	ReviewIsViaWeb        *bool   `json:"ReviewIsViaWeb"`
	ReviewIsViaHardcopy   *bool   `json:"ReviewIsViaHardcopy"`
	ReviewIsIdFilled      *bool   `json:"ReviewIsIdFilled"`
	ReviewIsVendorAcc     *bool   `json:"ReviewIsVendorAcc"`
	ReviewIsDataMatch     *bool   `json:"ReviewIsDataMatch"`
}
