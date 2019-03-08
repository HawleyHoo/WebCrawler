package main

import (
	"./meizar"
	"./rule"
	"flag"
	"net/http"
	"runtime"
	"strings"

	"WebCrawler/spiderkit"
	"strconv"
	"sync"
)

var dir *string = flag.String("dir", "./images/", "")
var startPage *int = flag.Int("start", 1, "")
var userCookie *string = flag.String("cookie", "1092990552=903f5Z2FA411DxfPYORhRDmNqohZzkNsnLuvj76u; PHPSESSID=vq5be8aobr23gdt68ig0mmud31; 1092990552=0; _gat=1; jdna=596e6fb28c1bb47f949e65e1ae03f7f5#1465139210979; Hm_lvt_fd93b7fb546adcfbcf80c4fc2b54da2c=1465130953,1465137904; Hm_lpvt_fd93b7fb546adcfbcf80c4fc2b54da2c=1465139211; _ga=GA1.2.1645811469.1465130953", "")
var pageSort *int = flag.Int("pagesort", 1, "")
var client *http.Client = &http.Client{}

func init() {
	flag.Parse()
	if !strings.HasSuffix(*dir, "/") {
		*dir = *dir + "/"
	}
}

func main_meizar() {
	//ctrl := gomock.NewController()
	//defer ctrl.Finish()
	//mockRepo := mock_rule.NewMockRule(ctrl)
	//mockRepo.EXPECT().DataRule()

	runtime.GOMAXPROCS(runtime.NumCPU())
	meizar.New(*dir, *startPage, rule.RuleProvider(), *userCookie, client, *pageSort).Start()
}

var (
	chSem     = make(chan int, 10)
	chImgMaps = make(chan map[string]string, 1000)

	wg4Imginfo  sync.WaitGroup
	wg4Download sync.WaitGroup

	//图片存储地址
	imgDir = `D:\meizi3\`
)

func main() {
	baseUrl := "https://www.duotoo.com/zt/rbmn/index"
	for i := 1; i < 5; i++ {
		var url string
		if i != 1 {
			url = baseUrl + "_" + strconv.Itoa(i) + ".html"
		} else {
			url = baseUrl + ".html"
		}

		wg4Imginfo.Add(1)
		go func(theUrl string) {
			spiderkit.GetPageImginfos2Chan(theUrl, imgDir, chImgMaps)
			wg4Imginfo.Done()
		}(url)
	}

	go func() {
		wg4Imginfo.Wait()
		close(chImgMaps)
		//fmt.Println("chImgMaps closed!")
	}()

	for imgMap := range chImgMaps {
		//fmt.Println("imgMap got:",imgMap)
		wg4Download.Add(1)
		go func(im map[string]string) {
			chSem <- 123
			spiderkit.DownloadFileWithClient(im["url"], im["filename"])
			<-chSem
			wg4Download.Done()
		}(imgMap)
	}

	wg4Download.Wait()
}
