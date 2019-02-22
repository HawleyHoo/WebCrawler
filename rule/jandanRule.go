package rule

import (
    // "fmt"
    "github.com/PuerkitoBio/goquery"

    "strconv"
    "fmt"
)

type JandanRule struct{}

func (p *JandanRule) UrlRule() (url string) {
    //return "http://jandan.net/ooxx/"
    //return "http://www.samsclub.cn/1/"
    return "http://list.samsclub.cn/search/c166041-1/?tp=2252.1.0.0.0.M!EcQTx-10-FL5so&tps=x0.07817y0.03106&ti=M!EcQTx-10-FL5so_FRBB"
}

func (p *JandanRule) PageRule(currentPage int) (page string) {
    return "page-" + strconv.Itoa(currentPage)
}

func (p *JandanRule) ImageRule(doc *goquery.Document, f func(image string)) {
    // doc.Find("a.view_img_link").Each(func(i int, s *goquery.Selection) {
    //     if img, exist := s.Attr("href"); exist {
    //         f(img)
    //     }
    // })
    fmt.Println("--------")
    //doc.Find("span.img-hash").Each(func(i int, s *goquery.Selection) {
    //    hash_value := s.Text()
    //    decoded, err := base64.StdEncoding.DecodeString(hash_value)
    //    if err == nil {
    //        img := "http:" + string(decoded)
    //        // fmt.Println(hash_value + "->" + img)
    //        f(img)
    //    }
    //})
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
