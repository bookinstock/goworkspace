package str

import (
	"encoding/base64"
	"fmt"
)

func RunBase64() {
	p := fmt.Println

	d := "base64!@#$%^&*()'~"

	enStd := base64.StdEncoding.EncodeToString([]byte(d))
	deStd, _ := base64.StdEncoding.DecodeString(enStd)
	p("enStd=", enStd)
	p("deStd=", string(deStd))

	enUrl := base64.URLEncoding.EncodeToString([]byte(d))
	deUrl, _ := base64.URLEncoding.DecodeString(enUrl)
	p("enUrl=", enUrl)
	p("deUrl=", string(deUrl))
}
