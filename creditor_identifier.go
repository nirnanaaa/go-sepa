package sepa

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const emptyStringRegex = `\s+`

// SanitizeCreditorIdentifier checks if an creditor identifier (ci) is valid.
// Note that also if the ci is valid it does not have to exist
func SanitizeCreditorIdentifier(identifier string) (string, error) {
	re := regexp.MustCompile(emptyStringRegex)
	sanitized := re.ReplaceAllString(identifier, "")
	sanitized = strings.ToUpper(sanitized)
	if !isValidCreditorIdentifier(sanitized) {
		return "", fmt.Errorf("the creditor identifier isn't valid")
	}
	var checksum bytes.Buffer
	checksum.WriteString(getCreditorIdentifierChecksum(sanitized))
	checksum.WriteString(getCreditorIdentifierNationalID(sanitized))
	sanitizedChecksum := removeNonAlnumChars(checksum.String())
	convertedChecksum := ConvertStringToMatchingInts(sanitizedChecksum)
	valid := Iso7064Mod97m10ChecksumCheck(convertedChecksum)
	if !valid {
		return "", fmt.Errorf("the checksums didn't match")
	}
	return sanitized, nil
}

func isValidCreditorIdentifier(identifier string) bool {
	re := regexp.MustCompile(PatternCreditorIdentifier)
	return re.MatchString(identifier)
}

func removeNonAlnumChars(src string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	return re.ReplaceAllString(src, "")
}

func getCreditorIdentifierChecksum(identifier string) string {
	return identifier[7:]
}

func getCreditorIdentifierNationalID(identifier string) string {
	return identifier[:4]
}

var mod97Values = []int{
	1, 10, 3, 30, 9, 90, 27, 76, 81, 34, 49, 5, 50, 15, 53, 45, 62, 38,
	89, 17, 73, 51, 25, 56, 75, 71, 31, 19, 93, 57, 85, 74, 61, 28, 86,
	84, 64, 58, 95, 77, 91, 37, 79, 14, 43, 42, 32, 29, 96, 87, 94, 67,
	88, 7, 70, 21, 16, 63, 48, 92, 47, 82, 44, 52, 35, 59, 8, 80, 24,
}

var alphabetMappings = map[string]string{
	"A": "10",
	"B": "11",
	"C": "12",
	"D": "13",
	"E": "14",
	"F": "15",
	"G": "16",
	"H": "17",
	"I": "18",
	"J": "19",
	"K": "20",
	"L": "21",
	"M": "22",
	"N": "23",
	"O": "24",
	"P": "25",
	"Q": "26",
	"R": "27",
	"S": "28",
	"T": "29",
	"U": "30",
	"V": "31",
	"W": "32",
	"X": "33",
	"Y": "34",
	"Z": "35",
}

// Iso7064Mod97m10ChecksumCheck verifies iso 7064 defined checksums
func Iso7064Mod97m10ChecksumCheck(input string) bool {
	checksum := 0
	lenx := len(input)
	for i := 1; i <= lenx; i++ {
		modValue := mod97Values[i-1]
		value := input[lenx-i]
		ix, err := strconv.Atoi(string(value))
		if err != nil {
			return false
		}
		checksum = (checksum + modValue*ix) % 97
	}
	return checksum == 1
}

// ConvertStringToMatchingInts replaces alpha characters to registered integer
// values inside a string
func ConvertStringToMatchingInts(input string) string {
	var output []string
	for i := range input {
		r := string(input[i])
		if val, ok := alphabetMappings[r]; ok {
			output = append(output, val)
			continue
		}
		output = append(output, r)
	}
	return strings.Join(output, "")
}
