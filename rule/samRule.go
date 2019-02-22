package rule

import (
	"strconv"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

type SamRule struct{}


type SamSku struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	SkuID    string        `bson:"SkuID"`
	Catagory string        `bson:"Catagory"`
	Name     string        `bson:"Name"`
	Price    string        `bson:"Price"`
	Desc     string        `bson:"Desc"`
}


func (p *SamRule) UrlRule() (url string) {
	//return "http://list.samsclub.cn/search/c141994-1/?advCode=SAMLBY_LM_LYGH&tp=2240.142101.0.0.0.M!IwQt4-10-FL5so&tps=x0.13349y0.04085&ti=M!IwQt4-10-FL5so_3XVY"
	return "http://list.samsclub.cn/searchPage/c141994-1/a/k-"
}

func (p *SamRule) PageRule(currentPage int) (page string) {
	//return "#page=" + strconv.Itoa(currentPage) + "&sort=1"
	return "p" + strconv.Itoa(currentPage) + "-s1-b0-pr-bm-f000"
}

func (p *SamRule) ImageRule(doc *goquery.Document, f func(image string)) {
	// doc.Find("a.view_img_link").Each(func(i int, s *goquery.Selection) {
	//     if img, exist := s.Attr("href"); exist {
	//         f(img)
	//     }
	// })
	//doc.Find("span.img-hash").Each(func(i int, s *goquery.Selection) {
	//    hash_value := s.Text()
	//    decoded, err := base64.StdEncoding.DecodeString(hash_value)
	//    if err == nil {
	//        img := "http:" + string(decoded)
	//        // fmt.Println(hash_value + "->" + img)
	//        f(img)
	//    }
	//})
	fmt.Println("--------")
	doc.Find(".mod_searchPro").Each(func(i int, s *goquery.Selection) {
		s.Find(".proName").Each(func(ii int, s2 *goquery.Selection) {
			t := s2.Find("a").Text()
			fmt.Println("test ", t)
		})
		//if img, exist := s.Attr("src"); exist {
		//	f(img)
		//}
	})
}

func (p *SamRule) DataRule(doc *goquery.Document, f func(data interface{}))  {
	//fmt.Println("--------")
	doc.Find(".mod_searchPro.jsModSearfhPro").Each(func(i int, s *goquery.Selection) {
		val, _ := s.Attr("prodid")
		//fmt.Println("id:", exist, val)
		//fmt.Println("name:", s.Find("p a").Text())
		//s.Find("p a").Text()
		//fmt.Println("price:", s.Find(".proPrice.clear").Find("em").Text())
		//fmt.Println("desc:", strings.TrimSpace(s.Find(".proName2").Text()))


		sam := SamSku {
			SkuID:val,
			Name:s.Find("p a").Text(),
			Price:s.Find(".proPrice.clear").Find("em").Text(),
			Desc:strings.TrimSpace(s.Find(".proName2").Text()),
		}
		f(sam)
		//fmt.Println("next sku-----")
	})
}