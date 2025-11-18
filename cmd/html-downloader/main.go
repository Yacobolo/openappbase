package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/JohannesKaufmann/dom"
	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/JohannesKaufmann/html-to-markdown/v2/converter"
	"golang.org/x/net/html"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failure", "error", err)
		os.Exit(1)
	}
}

func run() error {
	// Configuration
	url := "https://data-star.dev/docs"
	outputDir := "context/datastar"
	outputFile := "datastar.md"

	slog.Info("Starting HTML to Markdown conversion", "url", url)

	// Download HTML content
	html, err := downloadHTML(url)
	if err != nil {
		return fmt.Errorf("failed to download HTML: %w", err)
	}

	// Convert HTML to Markdown
	markdown, err := convertHTMLToMarkdown(html, url)
	if err != nil {
		return fmt.Errorf("failed to convert HTML to markdown: %w", err)
	}

	// Ensure output directory exists
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Write markdown to file
	outputPath := filepath.Join(outputDir, outputFile)
	if err := os.WriteFile(outputPath, []byte(markdown), 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	slog.Info("Successfully converted HTML to Markdown", "output", outputPath)
	return nil
}

func downloadHTML(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP request returned non-OK status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(body), nil
}

func convertHTMLToMarkdown(htmlContent string, baseURL string) (string, error) {
	// Parse the HTML document
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return "", fmt.Errorf("failed to parse HTML: %w", err)
	}

	// Find the article element
	articleNode := dom.FindFirstNode(doc, func(node *html.Node) bool {
		return node.Type == html.ElementNode && node.Data == "article"
	})

	if articleNode == nil {
		return "", fmt.Errorf("no article element found in HTML")
	}

	// Preprocess HTML: remove unwanted elements before markdown conversion
	removeLineNumberSpans(articleNode)
	removeAnchors(articleNode)

	// Convert the article node to a string
	var articleHTML strings.Builder
	if err := html.Render(&articleHTML, articleNode); err != nil {
		return "", fmt.Errorf("failed to render article HTML: %w", err)
	}

	// Convert the article HTML to markdown
	markdown, err := htmltomarkdown.ConvertString(
		articleHTML.String(),
		converter.WithDomain(baseURL),
	)
	if err != nil {
		return "", fmt.Errorf("conversion failed: %w", err)
	}

	return markdown, nil
}

// removeLineNumberSpans removes line number anchors from code blocks in the HTML DOM
// It targets <a> tags inside <code> elements that contain only digits (line numbers)
// This approach is more precise and robust than relying on CSS styling attributes because:
// - It targets the semantic pattern (digit-only anchors) rather than presentation (CSS)
// - It's resilient to changes in the site's styling classes or attributes
// - It has fewer false positives since digit-only anchors in code are almost always line numbers
func removeLineNumberSpans(node *html.Node) {
	// Track nodes to remove (can't modify during traversal)
	var nodesToRemove []*html.Node

	// Recursively find all code elements
	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		// If this is a <code> element, look for line number anchors inside it
		if n.Type == html.ElementNode && n.Data == "code" {
			// Find all <a> tags that contain only digits
			var findLineNumberAnchors func(*html.Node)
			findLineNumberAnchors = func(child *html.Node) {
				if child.Type == html.ElementNode && child.Data == "a" {
					// Check if this anchor contains only digits (it's a line number)
					if isLineNumberAnchor(child) {
						// Remove the parent span if it exists and wraps only this anchor
						if child.Parent != nil && child.Parent.Data == "span" && isLineNumberSpanWrapper(child.Parent) {
							nodesToRemove = append(nodesToRemove, child.Parent)
						} else {
							nodesToRemove = append(nodesToRemove, child)
						}
					}
				}
				// Continue traversing children
				for c := child.FirstChild; c != nil; c = c.NextSibling {
					findLineNumberAnchors(c)
				}
			}

			// Start searching within this code block
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				findLineNumberAnchors(c)
			}
		}

		// Continue traversing the tree
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}

	traverse(node)

	// Remove collected nodes
	for _, n := range nodesToRemove {
		if n.Parent != nil {
			n.Parent.RemoveChild(n)
		}
	}
}

// isLineNumberAnchor checks if an anchor element contains only digits (and whitespace)
// These anchors are line numbers in syntax-highlighted code blocks
func isLineNumberAnchor(node *html.Node) bool {
	if node.Type != html.ElementNode || node.Data != "a" {
		return false
	}

	// Get the text content of the anchor
	text := getTextContent(node)
	text = strings.TrimSpace(text)

	// Check if it's only digits
	if text == "" {
		return false
	}

	for _, ch := range text {
		if ch < '0' || ch > '9' {
			return false
		}
	}

	return true
}

// isLineNumberSpanWrapper checks if a span only wraps a line number anchor
// This helps us remove the entire wrapper span, not just the anchor
func isLineNumberSpanWrapper(node *html.Node) bool {
	if node.Type != html.ElementNode || node.Data != "span" {
		return false
	}

	// Count meaningful children (excluding text nodes with only whitespace)
	meaningfulChildren := 0
	hasAnchorChild := false

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode {
			meaningfulChildren++
			if c.Data == "a" {
				hasAnchorChild = true
			}
		} else if c.Type == html.TextNode {
			if strings.TrimSpace(c.Data) != "" {
				meaningfulChildren++
			}
		}
	}

	// It's a wrapper if it only has one meaningful child (the anchor)
	return meaningfulChildren == 1 && hasAnchorChild
}

// getTextContent recursively extracts all text from a node and its children
func getTextContent(node *html.Node) string {
	if node.Type == html.TextNode {
		return node.Data
	}

	var text strings.Builder
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		text.WriteString(getTextContent(c))
	}
	return text.String()
}

// removeAnchors removes all <a> anchor elements from the HTML DOM
// This prevents links from being converted to markdown links [text](url)
func removeAnchors(node *html.Node) {
	var nodesToRemove []*html.Node

	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		// If this is an anchor element, mark it for removal
		if n.Type == html.ElementNode && n.Data == "a" {
			nodesToRemove = append(nodesToRemove, n)
		}

		// Continue traversing children
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}

	traverse(node)

	// Remove anchor elements and preserve their text content (except for class="anchor" links)
	for _, anchor := range nodesToRemove {
		if anchor.Parent == nil {
			continue
		}

		// Check if this is a heading anchor (class="anchor") - remove completely
		isHeadingAnchor := false
		for _, attr := range anchor.Attr {
			if attr.Key == "class" && strings.Contains(attr.Val, "anchor") {
				isHeadingAnchor = true
				break
			}
		}

		if isHeadingAnchor {
			// Just remove the anchor completely (don't preserve content)
			anchor.Parent.RemoveChild(anchor)
		} else {
			// For regular links, preserve the text content
			for child := anchor.FirstChild; child != nil; {
				next := child.NextSibling
				anchor.RemoveChild(child)
				anchor.Parent.InsertBefore(child, anchor)
				child = next
			}
			anchor.Parent.RemoveChild(anchor)
		}
	}
}
