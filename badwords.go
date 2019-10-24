package main

import (
	"strings"
)

var badWords = []string{
	"آب کیر",
	"ارگاسم",
	"برهنه",
	"پورن",
	"پورنو",
	"تجاوز",
	"تخمی",
	"جق",
	"جقی",
	"جلق",
	"جنده",
	"چوچول",
	"حشر",
	"حشری",
	"داف",
	"دودول",
	"ساک زدن",
	"سکس",
	"سکس کردن",
	"سکسی",
	"شق کردن",
	"شهوت",
	"شهوتی",
	"شونبول",
	"فیلم سوپر",
	"کس دادن",
	"کس کردن",
	"کسکش",
	"کوس",
	"کون",
	"کون دادن",
	"کون کردن",
	"کونکش",
	"کونی",
	"کیر",
	"کیری",
	"لاپا",
	"لاپایی",
	"لاشی",
	"لخت",
	"لش",
	"منی",
	"هرزه",
	"کص",
	"kir",
	"kiri",
	"jende",
	"kososher",
	"kossher",
	"kosher",
	"koon",
	"kooni",
	"koskesh",
	"jakesh",
	"koonkesh",
	"konkesh",
	"obi",
	"jagh",
	"jaghi",
	"porn",
	"kos",
	"koos",
	"pofuz",
	"pofiooz",
	"daus",
	"dayoos",
	"dayyus",
	"lashi",
	"ankoloft",
	"anagha",
}

func HasBadwords(str string) bool {
	str = strings.ToLower(str)
	for _, badWord := range badWords {
		if strings.Contains(str, badWord) {
			return true
		}
	}
	return false
}
