package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"

	"github.com/kjk/u"
)

func writeRobotsTxt() {
	sitemapURL := path.Join(siteBaseURL, "sitemap.txt")
	robotsTxt := fmt.Sprintf("Sitemap: %s\n", sitemapURL)
	robotsTxtPath := filepath.Join("www", "robots.txt")
	err := ioutil.WriteFile(robotsTxtPath, []byte(robotsTxt), 0644)
	u.PanicIfErr(err)
}
