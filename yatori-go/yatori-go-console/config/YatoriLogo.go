package config

import (
	"encoding/base64"
	"log"
)

// 读取logo
func YaotirLogo() string {
	encoded := "CgogICAgICAgICAgICAgICAgICAgICAgICAgICAgIF9fXwogICAgICAgICwtLS0sICAgICAgICAgICAgICAsLS0uJ3xfICAgICAgICAgICAgICAgICAgICAgICwtLSwKICAgICAgIC9fIC4vfCAgICAgICAgICAgICAgfCAgfCA6LCcgICAsLS0tLiAgICBfXyAgLC0uLC0tLid8CiAsLS0tLCB8ICAnIDogICAgICAgICAgICAgIDogIDogJyA6ICAnICAgLCdcICwnICwnLyAvfHwgIHwsCi9fX18vIFwuICA6IHwgICAsLS0uLS0uICAuO19fLCcgIC8gIC8gICAvICAgfCcgIHwgfCcgfGAtLSdfCiAuICBcICBcICwnICcgIC8gICAgICAgXCB8ICB8ICAgfCAgLiAgIDsgLC4gOnwgIHwgICAsJywnICwnfAogIFwgIDsgIGAgICwnIC4tLS4gIC4tLiB8Ol9fLCd8IDogICcgICB8IHw6IDonICA6ICAvICAnICB8IHwKICAgXCAgXCAgICAnICAgXF9fXC86IC4gLiAgJyAgOiB8X18nICAgfCAuOyA6fCAgfCAnICAgfCAgfCA6CiAgICAnICBcICAgfCAgICwiIC4tLS47IHwgIHwgIHwgJy4nfCAgIDogICAgfDsgIDogfCAgICcgIDogfF9fCiAgICAgXCAgOyAgOyAgLyAgLyAgLC4gIHwgIDsgIDogICAgO1wgICBcICAvIHwgICwgOyAgIHwgIHwgJy4nfAogICAgICA6ICBcICBcOyAgOiAgIC4nICAgXCB8ICAsICAgLyAgYC0tLS0nICAgLS0tJyAgICA7ICA6ICAgIDsKICAgICAgIFwgICcgO3wgICwgICAgIC4tLi8gIC0tLWAtJyAgICAgICAgICAgICAgICAgICAgfCAgLCAgIC8KICAgICAgICBgLS1gICBgLS1gLS0tJyAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIC0tLWAtJwogICAgICAgICAgICAgICAgICAgICAgWWF0b3JpLWdvLWNvbnNvbGUgdjAuMC4xLUJldGEuMwogICAgICAgICAgICAgICAgICAgICDku4XnlKjkuo7lrabkuaDkuqTmtYHvvIzor7fli7/nlKjkuo7ov53ms5XlkozllYbkuJrnlKjpgJTvvIHvvIHvvIEKICAgICAgICAgICAgICBHaXRIdWLlvIDmupDlnLDlnYDvvJpodHRwczovL2dpdGh1Yi5jb20vQ2hhbmdiYWlxaS95YXRvcmkKICAgICAgICAgICAgICDkuKrkurrljZrlrqLvvJpodHRwczovL2Jsb2dzLmNoYW5nYmFpcWkudG9w"
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		log.Fatal(err)
	}
	return string(decoded)
}
