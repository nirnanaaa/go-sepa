package sepa

const (
	// SepaPain00100203 credit transfer version
	SepaPain00100203 = 100203
	// SepaPain00100303 credit transfer version
	SepaPain00100303 = 100303
	// SepaPain00100103 credit transfer version
	SepaPain00100103 = 100103
	// SepaPain00100103Gbic credit transfer version
	SepaPain00100103Gbic = 1001031
	// SepaPain00800202 direct debit version
	SepaPain00800202 = 800202
	// SepaPain00800302 direct debit version
	SepaPain00800302 = 800302
	// SepaPain00800102 direct debit version
	SepaPain00800102 = 800102
	// SepaPain00800102Gbic direct debit version
	SepaPain00800102Gbic = 8001021
	// SepaPain00800102Austrian003 direct debit version
	SepaPain00800102Austrian003 = 8001022
)

const (
	// HTLMPatternIban for iban validation inside an html input field
	HTLMPatternIban = `([a-zA-Z]\s*){2}([0-9]\s?){2}\s*([a-zA-Z0-9]\s*){1,30}`

	// HTLMPatternBic for bic validation inside an html input field
	HTLMPatternBic = `([a-zA-Z]\s*){6}[a-zA-Z2-9]\s*[a-nA-Np-zP-Z0-9]\s*(([A-Z0-9]\s*){3}){0,1}`

	// PatternIban for programmatically matching ibans
	PatternIban = `[A-Z]{2}[0-9]{2}[A-Z0-9]{1,30}`

	// PatternBic for programmatically matching bics
	PatternBic = `[A-Z]{6}[A-Z2-9][A-NP-Z0-9]([A-Z0-9]{3}){0,1}`

	// PatternCreditorIdentifier validates creditor identifiers
	PatternCreditorIdentifier = `^[a-zA-Z]{2,2}[0-9]{2,2}([A-Za-z0-9]|[\+|\?|/|\-|:|\(|\)|\.|,|\']){3,3}([A-Za-z0-9]|[\+|\?|/|\-|:|\(|\)|\.|,|\']){1,28}$`
)
