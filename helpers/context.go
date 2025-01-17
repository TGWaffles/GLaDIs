package helpers

import "strings"

func RemoveContextIdFromString(customId string) string {
	sp := strings.Split(customId, ":")

	if len(sp) > 0 {
		newSlice := sp[:len(sp)-1]

		return strings.Join(newSlice, "")
	} else {
		return ""
	}
}

func GetContextFromId(customId string) *string {
	sp := strings.Split(customId, ":")

	if len(sp) > 1 {
		return &sp[len(sp)-1]
	} else {
		return nil
	}
}
