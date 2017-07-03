package cloudprt

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
	"unicode"
	"unicode/utf8"
)

//接口的好处在于，只规定标准，实现方式多种多样.
type Reciept interface {
	FormatToReciept() string
}

const (
	DefaultMaxshopNameLen = 8
	DefaultMaxWeight      = 12
	DefaultWeightHan      = 2
	DefaultWeightOther    = 1
	DefaultMaxLength      = 6
)

type xiaoPiao struct {
	MaxshopNameLen int
	MaxWeight      int
	WeightHan      int
	WeightOther    int
	MaxLength      int
	MenuPrt        *MenuList
}

type MenuList struct {
	ShopName   string
	ShopMess   []string //exp :{"桌号:","A14","订单号:","1232132131"}
	OtherMess  []string //exp :{"备注:","wu"}
	GoodList   []*Goods
	CreateTime time.Time
}

type Goods struct {
	Name  string
	Price float64
	Num   int
}

func NewDefaultxiaoPiao(m *MenuList) Reciept {
	return &xiaoPiao{DefaultMaxshopNameLen, DefaultMaxWeight, DefaultWeightHan, DefaultWeightOther, DefaultMaxLength, m}
}

func (x *xiaoPiao) NewxiaoPiaoWithParam(m *MenuList, param []int) Reciept {
	return &xiaoPiao{param[0], param[1], param[2], param[3], param[4], m}
}

func (x *xiaoPiao) FormatToReciept() string {
	m := x.MenuPrt
	shopName := m.ShopName
	var buf bytes.Buffer
	var total float64
	buf.WriteString("^N1^F1\n")
	shopnameLength := utf8.RuneCountInString(shopName)
	if shopnameLength <= x.MaxshopNameLen {
		if shopnameLength == 1 || shopnameLength == 2 {
			for i := 0; i < 3; i++ {
				shopName = "  " + shopName
			}
		} else if shopnameLength == 3 || shopnameLength == 4 {
			for i := 0; i < 2; i++ {
				shopName = "  " + shopName
			}
		} else if shopnameLength == 5 || shopnameLength == 6 {
			for i := 0; i < 1; i++ {
				shopName = "  " + shopName
			}
		}
		buf.WriteString(fmt.Sprintf("^B2%s\n\n", shopName))
	} else {

		if shopnameLength == 9 || shopnameLength == 10 {
			for i := 0; i < 3; i++ {
				shopName = "  " + shopName
			}
		} else if shopnameLength == 11 || shopnameLength == 12 {
			for i := 0; i < 2; i++ {
				shopName = "  " + shopName
			}
		} else if shopnameLength == 13 || shopnameLength == 14 {
			for i := 0; i < 1; i++ {
				shopName = "  " + shopName
			}
		}
		buf.WriteString(fmt.Sprintf("%s\n\n", shopName))
	}
	//end
	for i := 0; i < len(m.ShopMess); i = i + 2 {
		buf.WriteString(fmt.Sprintf("%s %s\n", m.ShopMess[i], m.ShopMess[i+1]))
	}

	buf.WriteString("名称\t\t\t\t\t单价\t\t数量\t 金额\n")
	buf.WriteString("^W2================================\n")

	for _, v := range m.GoodList {

		totalNow := float64(v.Num) * (v.Price)
		totalNowStr := strconv.FormatFloat(totalNow, 'f', 1, 64)
		total = total + totalNow
		priceStr := strconv.FormatFloat(v.Price, 'f', 1, 64)
		numStr := strconv.Itoa(v.Num)

		x.formatName(&buf, v.Name, priceStr, numStr, totalNowStr)

	}

	totalStr := strconv.FormatFloat(total, 'f', 1, 64)
	if m.OtherMess != nil {
		for i := 0; i < len(m.OtherMess); i = i + 2 {
			buf.WriteString(fmt.Sprintf("%s %s\n", m.OtherMess[i], m.OtherMess[i+1]))
		}
	}

	buf.WriteString("^W2================================\n")
	buf.WriteString(fmt.Sprintf("^H2合计: %s元\n", totalStr))
	buf.WriteString(fmt.Sprintf("^H2订餐时间："))
	buf.WriteString(fmt.Sprintf(m.CreateTime.Format("2006 15:04:05")))
	return buf.String()
}

func (x *xiaoPiao) formatName(buf *bytes.Buffer, name string, priceStr, numStr, totalNowStr string) {
	var weight int = 0
	var supWeight int = 0
	for _, r := range name {
		if unicode.Is(unicode.Scripts["Han"], r) {
			weight += x.WeightHan
		} else {
			weight += x.WeightOther
		}
	}

	if weight <= x.MaxWeight {
		supWeight = x.MaxWeight - weight
		if supWeight > 0 {
			for i := 0; i < supWeight; i++ {
				name = name + " "
			}
		}
		buf.WriteString(fmt.Sprintf("%s", name))
		buf.WriteString(fmt.Sprintf("%5s\t\t\t%2s\t%5s\n", priceStr, numStr, totalNowStr))

	} else {
		buf.WriteString(fmt.Sprintf("%s\n", name))
		buf.WriteString(fmt.Sprintf("%s%5s\t\t\t%2s\t%5s\n", createSpace(x.MaxWeight), priceStr, numStr, totalNowStr))
	}

}

func createSpace(len int) string {
	var space string
	for i := 0; i < len; i++ {
		space = space + " "
	}
	return space
}
