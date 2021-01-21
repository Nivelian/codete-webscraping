package api

import (
	"fmt"
	"github.com/Nivelian/codete-webscraping/helpers"
	"github.com/Nivelian/codete-webscraping/model"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
)

func substr(s string, start, end int) string {
	return string([]rune(s)[start:end])
}

func getHtmlVersion(content string) *model.Record {
	version := ""
	if strings.Contains(strings.ToLower(content), "<!doctype html>") {
		version = "HTML 5"
	} else {
		dtdIdx := strings.Index(strings.ToLower(content), "//dtd")
		if dtdIdx != -1 {
			rightPart := substr(content, dtdIdx+6, len(content))
			version = substr(rightPart, 0, strings.Index(rightPart, "/"))
		} else {
			version = "Can't detect html version: unknown doctype format"
		}
	}

	return &model.Record{
		Name:  "Html version",
		Value: version,
	}
}

func getPageTitle(document *goquery.Document) *model.Record {
	return &model.Record{
		Name:  "Page title",
		Value: document.Find("title").First().Text(),
	}
}

func getHeadingsNumber(document *goquery.Document) *model.Record {
	result := ""
	for i := 1; i <= 6; i++ {
		tag := fmt.Sprintf("h%v", i)
		tagsNum := document.Find(tag).Length()
		result += fmt.Sprintf("%v = %v", tag, tagsNum)
		if i < 6 {
			result += "\n"
		}
	}

	return &model.Record{
		Name:  "Headings number",
		Value: result,
	}
}

func getLinksInfo(document *goquery.Document) *model.Record {
	externalNum, internalNum := 0, 0
	var linkUrls []string
	document.Find("a").Each(func(i int, x *goquery.Selection) {
		href, ok := x.Attr("href")
		if ok {
			if strings.HasPrefix(href, "#") {
				internalNum++
			} else {
				externalNum++
				linkUrls = append(linkUrls, href)
			}
		}
	})

	// counting links in parallel in order to speed up the process
	var inaccessibleNum int64
	wg := new(sync.WaitGroup)
	wg.Add(len(linkUrls))
	for _, href := range linkUrls {
		currentHref := href
		go func() {
			defer wg.Done()
			resp, err := http.Head(currentHref)
			if err != nil || resp.StatusCode >= 400 {
				atomic.AddInt64(&inaccessibleNum, 1)
			}
		}()
	}
	wg.Wait()

	return &model.Record{
		Name:  "Links info",
		Value: fmt.Sprintf("Internal = %v\nExternal = %v\nInaccessible = %v", internalNum, externalNum, inaccessibleNum),
	}
}

func isLoginExist(url string) *model.Record {
	result := "Yes"
	resp, err := http.Head(url + "/login")
	if err != nil || resp.StatusCode >= 400 {
		result = "No"
	}

	return &model.Record{
		Name:  "Has login form?",
		Value: result,
	}
}

func GetWebsiteInfo(url string) ([]*model.Record, error) {
	var res []*model.Record

	resp, err := http.Get(url)
	if err != nil {
		return nil, helpers.LogErr(err, "Failed to get the page by url %v", url)
	}
	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, helpers.LogErr(err, "Failed to read body of response")
	}

	content := string(html)

	htmlVersion := getHtmlVersion(content)

	document, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		return nil, helpers.LogErr(err, "Failed to create document from html string")
	}

	pageTitle := getPageTitle(document)
	headingsNumber := getHeadingsNumber(document)
	linksInfo := getLinksInfo(document)
	loginInfo := isLoginExist(url)

	res = append(res, htmlVersion, pageTitle, headingsNumber, linksInfo, loginInfo)

	return res, nil
}
