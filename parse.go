package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	anodes := linkNodes(doc)
	var links []Link
	for _, n := range anodes {
		links = append(links, buildLink(n))
	}
	//fmt.Printf("%+v", links)
	return links, nil
}

func buildLink(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			ret.Text = nodeText(n)
			break
		}
	}
	return ret
}

func nodeText(n *html.Node) string {
	var ret string
	//base cases
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += nodeText(c) + " "
	}
	return strings.Join(strings.Fields(ret), " ")
}
func linkNodes(n *html.Node) []*html.Node {
	var ret []*html.Node
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}
