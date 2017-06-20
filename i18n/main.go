package main

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {

	userLang := language.SimplifiedChinese
	p := message.NewPrinter(userLang)

	demo := `template1:%s`
	zh_Hans := `你好，%s:
  邮件内容什么的。
`
	zh_Hant := `你好，%s:
  郵件內容什麼的。
`
	en := `Hi,%s:
  email content.
`
	message.SetString(language.English, demo, en)                 //英语
	message.SetString(language.SimplifiedChinese, demo, zh_Hans)  //简体中文
	message.SetString(language.TraditionalChinese, demo, zh_Hant) //繁体中文

	userLang = language.MustParse("en")
	p = message.NewPrinter(userLang)
	fmt.Println(p.Sprintf(demo, "Joe"))

	userLang = language.MustParse("zh-Hans")
	p = message.NewPrinter(userLang)
	fmt.Println(p.Sprintf(demo, "李四"))

	userLang = language.MustParse("zh-Hant")
	p = message.NewPrinter(userLang)
	fmt.Println(p.Sprintf(demo, "張三"))
}
