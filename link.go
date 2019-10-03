package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents a link (<a href="..."> tag) in an HTML document.
type Link struct {
	Href string
	Text string
}

// ParseLinks will take an HTML document and parse all the links it contains
// into a slice of Link.
func ParseLinks(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	nodes := findLinkNodes(doc)
	var links []Link
	for _, n := range nodes {
		links = append(links, buildLink(n))
	}

	return links, nil
}

func findLinkNodes(n *html.Node) []*html.Node {
	// If the element is a there is no need to go further
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var ret []*html.Node

	// Calls itself for every child and sibling until an a is returned
	// in which case it's added to the ret variable over and over
	// until there is a single ret returned
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, findLinkNodes(c)...)
	}

	return ret
}

func buildLink(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	ret.Text = findText(n)

	return ret
}

func findText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}

	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += findText(c)
	}

	return strings.Join(strings.Fields(ret), " ")
}
