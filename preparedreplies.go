package main

import (
	"strings"
)

var preparedReplies = map[string]string{
	"rick and morty": "شتتتتتتتتتتت 😍😍😍",
	"mr robot":       "فاااااااااااااااااک!!!! 😍",
	"اوکیه":          "از نظر من اوکی نیست 😐",
	"chakerim":       "9karam",
	"4krim":          "9karam",
	"nokaram":        "4krim",
	"9karam":         "4krim",
	"mokhlesam":      "sultani",
	"4820":           "🚨⚠️🚨⚠️🚨⚠️🚨⚠️🚨⚠️",
	"غذات":           "برو غذات رو بگیر",
}

func PrepareReply(str string) string {
	str = strings.ToLower(str)
	for msg, reply := range preparedReplies {
		if strings.Contains(str, msg) {
			return reply
		}
	}

	return ""
}
