package billplz

// Bank SWIFT codes supported by the Billplz API.
//
// Extracted from https://www.billplz.com/api#create-a-bank-account.
const (
	BankCodeAffinBank             = "PHBMMYKL"
	BankCodeAgrobank              = "BPMBMYKL"
	BankCodeAllianceBank          = "MFBBMYKL"
	BankCodeAlRajhiBank           = "RJHIMYKL"
	BankCodeAmBank                = "ARBKMYKL"
	BankCodeBankIslam             = "BIMBMYKL"
	BankCodeBankKerjasamaRakyat   = "BKRMMYKL"
	BankCodeBankMuamalat          = "BMMBMYKL"
	BankCodeBankSimpananNasional  = "BSNAMYK1"
	BankCodeCIMBBank              = "CIBBMYKL"
	BankCodeCitibank              = "CITIMYKL"
	BankCodeHongLeongBank         = "HLBBMYKL"
	BankCodeHSBCBank              = "HBMBMYKL"
	BankCodeMaybank               = "MBBEMYKL"
	BankCodeOCBCBank              = "OCBCMYKL"
	BankCodePublicBank            = "PBBEMYKL"
	BankCodeRHBBank               = "RHBBMYKL"
	BankCodeStandardCharteredBank = "SCBLMYKX"
	BankCodeUnitedOverseasBank    = "UOVBMYKL"
)

// Base URLs for Billplz API endpoints supported by the package.
const (
	endpointStaging = "https://billplz-staging.herokuapp.com/api/v3"
	endpointProdV3  = "https://www.billplz.com/api/v3"
)
