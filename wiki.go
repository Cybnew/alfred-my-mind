package main

import (
	"bufio"
	"io/ioutil"
	"regexp"
	"strings"
)

// Link holds a link with unique UID.
type Link struct {
	uid  string
	name string
	url  string
}

// parseSummaryFile parses GitBook Summary.md file and returns links.
func parseSummaryFile() []Link {
	bytes, _ := ioutil.ReadFile("summary.md")

	// regex to extract markdown links
	re := regexp.MustCompile(`\[([^\]]*)\]\(([^)]*)\)`)
	re1 := regexp.MustCompile(`(.md)`)

	var links []link

	// Read file line by line and extraxt links
	scanner := bufio.NewScanner(strings.NewReader(string(bytes)))
	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		links = append(links, link{
			uid:  matches[0][1],
			name: matches[0][1],
			url:  "https://knowledge.mrupp.dev" + re1.ReplaceAllString(matches[0][2], `.html`),
		})
	}
	return links
}

func searchWiki() {
	showUpdateStatus()
	links := parseSummaryFile()

	// TODO: implement caching
	for _, link := range links {
		wf.NewItem(link.name).UID(link.uid).Valid(true).Arg(link.url)
	}

	// TODO: message doesnt show
	wf.WarnEmpty("No matching items", "Try a different query?")

	wf.SendFeedback()
}
