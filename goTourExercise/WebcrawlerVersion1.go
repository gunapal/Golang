package main

import (
		"fmt"
		"log"
		"github.com/PuerkitoBio/goquery"
		"regexp"
)

type Page struct {
	Url string
	Links map[string]string
	StaticAssets []string
	IsCrawled bool
}

type SiteMap struct {
	SeedURL string
	UrlSet map[string]Page
}

func main() {
	//data structure to hold the site map 
	siteMap := SiteMap{SeedURL: "http://digitalocean.com"}
	siteMap.UrlSet = make(map[string]Page)
	//create the seed page
	digitalOceanPage := Page{Url: "http://digitalocean.com"}
	
	//start crawling
	digitalOceanPage.Crawl(&siteMap)

	siteMap.Print()
}

func (siteMap *SiteMap) IsPagedVisited(pgUrl string) bool {
	if pgUrl == "" {
		fmt.Println("url string is empty")
	}
	_, isVisited := siteMap.UrlSet[pgUrl]

	return isVisited
}

func (siteMap *SiteMap) Print(){
	if siteMap == nil || siteMap.UrlSet == nil || len(siteMap.UrlSet) == 0 {
		fmt.Println("site map is empty !!! hmmmm")
		return
	}

	fmt.Println("Total number of sites crawled : ", len(siteMap.UrlSet))
	fmt.Println("======================================================")

	for _, pg := range(siteMap.UrlSet){
		fmt.Println(pg.Url)
		fmt.Println("------------------")
		for _, val := range(pg.Links){
			fmt.Printf("\t %s \n", val)
		}
	}
}

/*
 * crawls the page specified
 */
func (pg *Page) Crawl(siteMap *SiteMap) {

	if pg.Url == ""	{
		fmt.Println("url is empty")
		return
	}

	if pg.IsCrawled {
		return
	}

	doc, err := goquery.NewDocument(pg.Url)

	if err != nil 	{
		log.Fatal(err)
	}
	var userURL = regexp.MustCompile(`.*/community/`)

	pg.Parse(doc, siteMap)
	pg.IsCrawled = true
	siteMap.UrlSet[pg.Url] = *pg


	//can i put this in a seperate logic??
	for _, url2crawl := range(pg.Links)	{
		
		_, isVisited := siteMap.UrlSet[url2crawl]

		childPage := NewPage(url2crawl)

		//crawl only if the page is not already visited.
		if ! isVisited && ! userURL.MatchString(url2crawl){
			//Crawl the child page

			childPage.Crawl(siteMap)
		}
	}
}

func (pg *Page) Parse(doc *goquery.Document, siteMap *SiteMap) {

	if len(pg.Links) == 0	{
		pg.Links = make(map[string]string)
	}

	AddChildURL := func(i int, s *goquery.Selection){
		url, hrefAttrExists := s.Attr("href")
		if hrefAttrExists {
			url, ok := GetAbsoluteURL(url)
			if ok {
				pg.Links[url] = url
			}
		}
	}
	
	doc.Find("a").Each(AddChildURL)
}

func (pg *Page) AddStaticAssets(assetUrl string) {

}

func NewPage(pageUrl string) (newPage *Page) {
	pg := Page{Url: pageUrl}
	return &pg
}

func GetAbsoluteURL(url string) (string, bool){
	var relativeURL = regexp.MustCompile(`^\/`)
	var absoluteURL = regexp.MustCompile(`^http://digitalocean.com`)
	var userURL = regexp.MustCompile(`.*/community/users/`)

	if relativeURL.MatchString(url){
		return "http://digitalocean.com"+url, true
	}

	if absoluteURL.MatchString(url) && ! userURL.MatchString(url) {
		return url, true
	}

	return "", false
}