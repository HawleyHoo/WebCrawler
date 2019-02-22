package rule

import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"fmt"
)

type MeiRule struct{}

func (p *MeiRule) UrlRule() string {
	return "https://meizi.us/"
}

func (p *MeiRule) PageRule(currentPage int) (page string) {
	return "?page=" + strconv.Itoa(currentPage)
}

func (p *MeiRule) ImageRule(doc *goquery.Document, f func(image string)) {
	doc.Find(".mod_searchPro").Each(func(i int, s *goquery.Selection) {
		s.Find(".proName i-size-tit is-tags-mark").Each(func(ii int, s2 *goquery.Selection) {
			t := s2.Find("a").Text()
			fmt.Println("test ", t)
		})
		//if img, exist := s.Attr("src"); exist {
		//	f(img)
		//}
	})
}
