package htmlparser

import (
	"strings"

	"golang.org/x/net/html"
)

type HtmlParser struct {
	Content string
}

func New(content string) *HtmlParser {
	return &HtmlParser{Content: content}
}

func (p *HtmlParser) GetNumLinksAndImages() (int, int, error) {
	numLinks := 0
	numImages := 0

	htmlParser := html.NewTokenizer(strings.NewReader(p.Content))
	for {
		tokenType := htmlParser.Next()
		if tokenType == html.StartTagToken || tokenType == html.SelfClosingTagToken {
			token := htmlParser.Token()
			if token.Data == "a" {
				numLinks++
			}
			if token.Data == "img" {
				numImages++
			}
		} else if tokenType == html.ErrorToken {
			break
		}
	}

	return numLinks, numImages, nil
}
