package Helper

import "regexp"

const FirstCharacterRussianPhone = string("7")

func FormattingPhone(phone string) string {
	regex := regexp.MustCompile("/[^0-9]/")
	phone = regex.ReplaceAllString(phone, "")

	if len(phone) <= 11 {
		phone = FirstCharacterRussianPhone + phone[1:]
	}

	return phone
}
