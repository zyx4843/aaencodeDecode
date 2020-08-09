package aaencodeDecode

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	b = []string{
		"(c^_^o)",
		"(ﾟΘﾟ)",
		"((o^_^o) - (ﾟΘﾟ))",
		"(o^_^o)",
		"(ﾟｰﾟ)",
		"((ﾟｰﾟ) + (ﾟΘﾟ))",
		"((o^_^o) +(o^_^o))",
		"((ﾟｰﾟ) + (o^_^o))",
		"((ﾟｰﾟ) + (ﾟｰﾟ))",
		"((ﾟｰﾟ) + (ﾟｰﾟ) + (ﾟΘﾟ))",
		"(ﾟДﾟ) .ﾟωﾟﾉ",
		"(ﾟДﾟ) .ﾟΘﾟﾉ",
		"(ﾟДﾟ) ['c']",
		"(ﾟДﾟ) .ﾟｰﾟﾉ",
		"(ﾟДﾟ) .ﾟДﾟﾉ",
		"(ﾟДﾟ) [ﾟΘﾟ]",
	}

	b2 = []string{
		"(c^_^o)",
		"(ﾟΘﾟ)",
		"((o^_^o)-(ﾟΘﾟ))",
		"(o^_^o)",
		"(ﾟｰﾟ)",
		"((ﾟｰﾟ)+(ﾟΘﾟ))",
		"((o^_^o)+(o^_^o))",
		"((ﾟｰﾟ)+(o^_^o))",
		"((ﾟｰﾟ)+(ﾟｰﾟ))",
		"((ﾟｰﾟ)+(ﾟｰﾟ)+(ﾟΘﾟ))",
		"(ﾟДﾟ).ﾟωﾟﾉ",
		"(ﾟДﾟ).ﾟΘﾟﾉ",
		"(ﾟДﾟ)['c']",
		"(ﾟДﾟ).ﾟｰﾟﾉ",
		"(ﾟДﾟ).ﾟДﾟﾉ",
		"(ﾟДﾟ)[ﾟΘﾟ]",
	}
)

func Aaencode(text string) string {

	var r = "ﾟωﾟﾉ= /｀ｍ´）ﾉ ~┻━┻   //*´∇｀*/ ['_']; o=(ﾟｰﾟ)  =_=3; c=(ﾟΘﾟ) =(ﾟｰﾟ)-(ﾟｰﾟ); "
	matched, err := regexp.MatchString(`ひだまりスケッチ×(365|３５６)\s*来週も見てくださいね[!！]`, text)
	if err != nil {
		fmt.Println("MatchString error :", err)
	}
	if matched {
		r += "X=_=3; "
		r += "\r\n\r\n    X / _ / X < \"来週も見てくださいね!\";\r\n\r\n"
	}
	r += "(ﾟДﾟ) =(ﾟΘﾟ)= (o^_^o)/ (o^_^o);" +
		"(ﾟДﾟ)={ﾟΘﾟ: '_' ,ﾟωﾟﾉ : ((ﾟωﾟﾉ==3) +'_') [ﾟΘﾟ] " +
		",ﾟｰﾟﾉ :(ﾟωﾟﾉ+ '_')[o^_^o -(ﾟΘﾟ)] " +
		",ﾟДﾟﾉ:((ﾟｰﾟ==3) +'_')[ﾟｰﾟ] }; (ﾟДﾟ) [ﾟΘﾟ] =((ﾟωﾟﾉ==3) +'_') [c^_^o];" +
		"(ﾟДﾟ) ['c'] = ((ﾟДﾟ)+'_') [ (ﾟｰﾟ)+(ﾟｰﾟ)-(ﾟΘﾟ) ];" +
		"(ﾟДﾟ) ['o'] = ((ﾟДﾟ)+'_') [ﾟΘﾟ];" +
		"(ﾟoﾟ)=(ﾟДﾟ) ['c']+(ﾟДﾟ) ['o']+(ﾟωﾟﾉ +'_')[ﾟΘﾟ]+ ((ﾟωﾟﾉ==3) +'_') [ﾟｰﾟ] + " +
		"((ﾟДﾟ) +'_') [(ﾟｰﾟ)+(ﾟｰﾟ)]+ ((ﾟｰﾟ==3) +'_') [ﾟΘﾟ]+" +
		"((ﾟｰﾟ==3) +'_') [(ﾟｰﾟ) - (ﾟΘﾟ)]+(ﾟДﾟ) ['c']+" +
		"((ﾟДﾟ)+'_') [(ﾟｰﾟ)+(ﾟｰﾟ)]+ (ﾟДﾟ) ['o']+" +
		"((ﾟｰﾟ==3) +'_') [ﾟΘﾟ];(ﾟДﾟ) ['_'] =(o^_^o) [ﾟoﾟ] [ﾟoﾟ];" +
		"(ﾟεﾟ)=((ﾟｰﾟ==3) +'_') [ﾟΘﾟ]+ (ﾟДﾟ) .ﾟДﾟﾉ+" +
		"((ﾟДﾟ)+'_') [(ﾟｰﾟ) + (ﾟｰﾟ)]+((ﾟｰﾟ==3) +'_') [o^_^o -ﾟΘﾟ]+" +
		"((ﾟｰﾟ==3) +'_') [ﾟΘﾟ]+ (ﾟωﾟﾉ +'_') [ﾟΘﾟ]; " +
		"(ﾟｰﾟ)+=(ﾟΘﾟ); (ﾟДﾟ)[ﾟεﾟ]='\\\\'; " +
		"(ﾟДﾟ).ﾟΘﾟﾉ=(ﾟДﾟ+ ﾟｰﾟ)[o^_^o -(ﾟΘﾟ)];" +
		"(oﾟｰﾟo)=(ﾟωﾟﾉ +'_')[c^_^o];" + //TODO
		"(ﾟДﾟ) [ﾟoﾟ]='\\\"';" +
		"(ﾟДﾟ) ['_'] ( (ﾟДﾟ) ['_'] (ﾟεﾟ+"
	r += "(ﾟДﾟ)[ﾟoﾟ]+ "
	untext := []rune(text)
	for i := 0; i < len(untext); i++ {
		n := untext[i]
		t := "(ﾟДﾟ)[ﾟεﾟ]+"
		if n <= 127 {
			re := regexp.MustCompile("[0-7]")
			m := fmt.Sprintf("%o", n)
			//fmt.Println(m + "*")
			//fmt.Println("#" + re.FindString(m))
			t += re.ReplaceAllStringFunc(m, func(c string) string {
				//fmt.Println(c)
				ci := c[0] - '0'
				return b[ci] + "+ "
			})
		} else {
			//fmt.Println(n)
			re1 := regexp.MustCompile("[0-9a-f]{4}$")
			m := re1.FindString("000" + fmt.Sprintf("%x", n))
			//fmt.Println(m)
			re := regexp.MustCompile("[0-9a-f]")
			t += "(oﾟｰﾟo)+ " + re.ReplaceAllStringFunc(m, func(c string) string {
				ci, _ := strconv.ParseInt(c, 16, 32)
				//fmt.Println(ci)
				return b[ci] + "+ "
			})
		}
		r += t
	}
	r += "(ﾟДﾟ)[ﾟoﾟ]) (ﾟΘﾟ)) ('_');"
	return r
}

func Aadecode(text string) string {

	result := make([]rune, 0)
	strs := strings.Split(text, "(ﾟДﾟ)[ﾟεﾟ]+")
	for i := 1; i < len(strs); i++ {
		charts := strings.Split(strs[i], "+")
		if strings.TrimSpace(charts[0]) == "(oﾟｰﾟo)" {
			//>127
			chart := rune(0)
			for j := 1; j < len(charts); j++ {
				tmp := strings.TrimSpace(charts[j])
				if len(tmp) > 0 {
					ch := cellDecode(tmp)
					if ch >= 0 {
						chart = chart*16 + rune(ch)
					} else if tmp[0] == '(' && tmp[1] == '(' && j+1 < len(charts) {
						//fmt.Println(charts[j])
						charts[j+1] = tmp + "+" + strings.TrimSpace(charts[j+1])
					} else {
						fmt.Println("cell err 1 ", charts[j])
					}
				}
			}
			result = append(result, chart)
		} else {
			//<= 127
			chart := rune(0)
			for j := 0; j < len(charts); j++ {
				tmp := strings.TrimSpace(charts[j])
				if len(tmp) > 0 {
					ch := cellDecode(tmp)
					if ch >= 0 {
						chart = chart*8 + rune(ch)
					} else if tmp[0] == '(' && tmp[1] == '(' && j+1 < len(charts) {
						//fmt.Println(charts[j])
						charts[j+1] = tmp + "+" + strings.TrimSpace(charts[j+1])
					} else {
						fmt.Println("cell err 2 ", charts[j])
					}
				}
			}
			result = append(result, chart)
		}
	}

	return string(result)
}

func cellDecode(chart string) int {
	for i := 0; i < len(b); i++ {
		if chart == b[i] {
			return i
		}
	}

	for i := 0; i < len(b2); i++ {
		if chart == b2[i] {
			return i
		}
	}

	return -1
}
