package goeditorjs_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/davidscottmills/goeditorjs"
	"github.com/stretchr/testify/require"
)

func Test_HeaderHandler_Type(t *testing.T) {
	h := &goeditorjs.HeaderHandler{}
	require.Equal(t, "header", h.Type())
}

func Test_HeaderHandler_GenerateHTML_Returns_Parse_Err(t *testing.T) {
	h := &goeditorjs.HeaderHandler{}
	_, err := h.GenerateHTML(goeditorjs.EditorJSBlock{Type: "header", Data: []byte{}})
	require.Error(t, err)
}

func Test_HeaderHandler_GenerateHTML(t *testing.T) {
	bhh := &goeditorjs.HeaderHandler{}
	testData := []struct {
		data           string
		expectedResult string
	}{
		{data: `{"text": "Heading","level": 1}`, expectedResult: "<h1>Heading</h1>"},
		{data: `{"text": "Heading","level": 2}`, expectedResult: "<h2>Heading</h2>"},
		{data: `{"text": "Heading","level": 3}`, expectedResult: "<h3>Heading</h3>"},
		{data: `{"text": "Heading","level": 4}`, expectedResult: "<h4>Heading</h4>"},
		{data: `{"text": "Heading","level": 5}`, expectedResult: "<h5>Heading</h5>"},
		{data: `{"text": "Heading","level": 6}`, expectedResult: "<h6>Heading</h6>"},
	}

	for _, td := range testData {
		jsonData := []byte(td.data)
		ejsBlock := goeditorjs.EditorJSBlock{Type: "header", Data: jsonData}
		html, _ := bhh.GenerateHTML(ejsBlock)
		require.Equal(t, td.expectedResult, html)
	}
}

func Test_HeaderHandler_GenerateMarkdown_Returns_Parse_Err(t *testing.T) {
	h := &goeditorjs.HeaderHandler{}
	_, err := h.GenerateMarkdown(goeditorjs.EditorJSBlock{Type: "header", Data: []byte{}})
	require.Error(t, err)
}

func Test_HeaderHandler_GenerateMarkdown(t *testing.T) {
	bhh := &goeditorjs.HeaderHandler{}
	testData := []struct {
		data           string
		expectedResult string
	}{
		{data: `{"text": "Heading","level": 1}`, expectedResult: "# Heading"},
		{data: `{"text": "Heading","level": 2}`, expectedResult: "## Heading"},
		{data: `{"text": "Heading","level": 3}`, expectedResult: "### Heading"},
		{data: `{"text": "Heading","level": 4}`, expectedResult: "#### Heading"},
		{data: `{"text": "Heading","level": 5}`, expectedResult: "##### Heading"},
		{data: `{"text": "Heading","level": 6}`, expectedResult: "###### Heading"},
	}

	for _, td := range testData {
		jsonData := []byte(td.data)
		ejsBlock := goeditorjs.EditorJSBlock{Type: "header", Data: jsonData}
		md, _ := bhh.GenerateMarkdown(ejsBlock)
		require.Equal(t, td.expectedResult, md)
	}
}

func Test_ParagraphHandler_Type(t *testing.T) {
	h := &goeditorjs.ParagraphHandler{}
	require.Equal(t, "paragraph", h.Type())
}

func Test_ParagraphHandler_GenerateHTML_Returns_Parse_Err(t *testing.T) {
	h := &goeditorjs.ParagraphHandler{}
	_, err := h.GenerateHTML(goeditorjs.EditorJSBlock{Type: "paragraph", Data: []byte{}})
	require.Error(t, err)
}

func Test_ParagraphHandler_GenerateHTML_Left(t *testing.T) {
	bph := &goeditorjs.ParagraphHandler{}
	jsonData := []byte(`{"text": "paragraph","alignment": "left"}`)
	ejsBlock := goeditorjs.EditorJSBlock{Type: "paragraph", Data: jsonData}
	html, _ := bph.GenerateHTML(ejsBlock)
	require.Equal(t, "<p>paragraph</p>", html)
}

func Test_ParagraphHandler_GenerateHTML_Center_Right(t *testing.T) {
	bph := &goeditorjs.ParagraphHandler{}
	testData := []struct {
		alignment string
		data      string
	}{
		{alignment: "center", data: `{"text": "paragraph","alignment": "center"}`},
		{alignment: "right", data: `{"text": "paragraph","alignment": "right"}`},
	}

	for _, td := range testData {
		jsonData := []byte(td.data)
		ejsBlock := goeditorjs.EditorJSBlock{Type: "paragraph", Data: jsonData}
		html, _ := bph.GenerateHTML(ejsBlock)
		require.Equal(t, fmt.Sprintf(`<p style="text-align:%s">paragraph</p>`, td.alignment), html)
	}
}

func Test_ParagraphHandler_GenerateMarkdown_Returns_Parse_Err(t *testing.T) {
	h := &goeditorjs.ParagraphHandler{}
	_, err := h.GenerateMarkdown(goeditorjs.EditorJSBlock{Type: "paragraph", Data: []byte{}})
	require.Error(t, err)
}

func Test_ParagraphHandler_GenerateMarkdown_Left(t *testing.T) {
	bph := &goeditorjs.ParagraphHandler{}
	jsonData := []byte(`{"text": "paragraph","alignment": "left"}`)
	ejsBlock := goeditorjs.EditorJSBlock{Type: "paragraph", Data: jsonData}
	md, _ := bph.GenerateMarkdown(ejsBlock)
	require.Equal(t, "paragraph", md)
}

func Test_ParagraphHandler_GenerateMarkdown_WithATag(t *testing.T) {
	bph := &goeditorjs.ParagraphHandler{}
	jsonData := []byte(`{"text": "paragraph<a href=\"123\">456</a>","alignment": "left"}`)
	ejsBlock := goeditorjs.EditorJSBlock{Type: "paragraph", Data: jsonData}
	md, err := bph.GenerateMarkdown(ejsBlock)
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "paragraph[456](123)", md)
}

func Test_ParagraphHandler_GenerateMarkdown_WithCodeTag(t *testing.T) {
	bph := &goeditorjs.ParagraphHandler{}
	jsonData := []byte(`{"text": "paragraph<code>456</code>","alignment": "left"}`)
	ejsBlock := goeditorjs.EditorJSBlock{Type: "paragraph", Data: jsonData}
	md, _ := bph.GenerateMarkdown(ejsBlock)
	require.Equal(t, "paragraph`456`", md)
}

func Test_ParagraphHandler_GenerateMarkdown_Center_Right(t *testing.T) {
	bph := &goeditorjs.ParagraphHandler{}
	testData := []struct {
		alignment string
		data      string
	}{
		{alignment: "center", data: `{"text": "paragraph","alignment": "center"}`},
		{alignment: "right", data: `{"text": "paragraph","alignment": "right"}`},
	}

	for _, td := range testData {
		jsonData := []byte(td.data)
		ejsBlock := goeditorjs.EditorJSBlock{Type: "paragraph", Data: jsonData}
		md, _ := bph.GenerateMarkdown(ejsBlock)
		require.Equal(t, fmt.Sprintf(`<p style="text-align:%s">paragraph</p>`, td.alignment), md)
	}
}

func Test_ListHandler_Type(t *testing.T) {
	h := &goeditorjs.ListHandler{}
	require.Equal(t, "list", h.Type())
}

func Test_ListHandler_GenerateHTML_Returns_Parse_Err(t *testing.T) {
	h := &goeditorjs.ListHandler{}
	_, err := h.GenerateHTML(goeditorjs.EditorJSBlock{Type: "list", Data: []byte{}})
	require.Error(t, err)
}

func Test_ListHandler_GenerateHTML(t *testing.T) {
	blh := &goeditorjs.ListHandler{}
	testData := []struct {
		data           string
		expectedResult string
	}{
		{data: `{"style": "ordered", "items": ["one", "two", "three"]}`,
			expectedResult: "<ol><li>one</li><li>two</li><li>three</li></ol>"},
		{data: `{"style": "unordered", "items": ["one", "two", "three"]}`,
			expectedResult: "<ul><li>one</li><li>two</li><li>three</li></ul>"},
	}

	for _, td := range testData {
		jsonData := []byte(td.data)
		ejsBlock := goeditorjs.EditorJSBlock{Type: "list", Data: jsonData}
		html, _ := blh.GenerateHTML(ejsBlock)
		require.Equal(t, td.expectedResult, html)
	}
}

func Test_ListHandler_GenerateMarkdown_Returns_Parse_Err(t *testing.T) {
	h := &goeditorjs.ListHandler{}
	_, err := h.GenerateMarkdown(goeditorjs.EditorJSBlock{Type: "list", Data: []byte{}})
	require.Error(t, err)
}

func Test_ListHandler_GenerateMarkdown(t *testing.T) {
	blh := &goeditorjs.ListHandler{}
	testData := []struct {
		data           string
		expectedResult string
	}{
		{data: `{"style": "ordered", "items": ["one", "two", "three"]}`,
			expectedResult: "1. one\n1. two\n1. three"},
		{data: `{"style": "unordered", "items": ["one", "two", "three"]}`,
			expectedResult: "- one\n- two\n- three"},
	}

	for _, td := range testData {
		jsonData := []byte(td.data)
		ejsBlock := goeditorjs.EditorJSBlock{Type: "list", Data: jsonData}
		md, _ := blh.GenerateMarkdown(ejsBlock)
		require.Equal(t, td.expectedResult, md)
	}
}

func Test_CodeBoxHandler_Type(t *testing.T) {
	h := &goeditorjs.CodeBoxHandler{}
	require.Equal(t, "codeBox", h.Type())
}

func Test_CodeBoxHandler_GenerateHTML_Returns_Parse_Err(t *testing.T) {
	h := &goeditorjs.CodeBoxHandler{}
	_, err := h.GenerateHTML(goeditorjs.EditorJSBlock{Type: "codeBox", Data: []byte{}})
	require.Error(t, err)
}

func Test_CodeBoxHandler_GenerateHTML(t *testing.T) {
	bcbh := &goeditorjs.CodeBoxHandler{}
	jsonData := []byte(`{"language": "go", "code": "func main(){fmt.Println(\"HelloWorld\")}"}`)
	ejsBlock := goeditorjs.EditorJSBlock{Type: "codeBox", Data: jsonData}
	expectedResult := `<pre><code class="go">func main(){fmt.Println("HelloWorld")}</code></pre>`
	html, _ := bcbh.GenerateHTML(ejsBlock)
	require.Equal(t, expectedResult, html)
}

func Test_CodeBoxHandler_GenerateMarkdown_Returns_Parse_Err(t *testing.T) {
	h := &goeditorjs.CodeBoxHandler{}
	_, err := h.GenerateMarkdown(goeditorjs.EditorJSBlock{Type: "codeBox", Data: []byte{}})
	require.Error(t, err)
}

func Test_CodeBoxHandler_GenerateMarkdown(t *testing.T) {
	bcbh := &goeditorjs.CodeBoxHandler{}
	jsonData := []byte(`{"language": "go", "code": "func main(){fmt.Println(\"HelloWorld\")}"}`)
	ejsBlock := goeditorjs.EditorJSBlock{Type: "codeBox", Data: jsonData}
	expectedResult := "```go\nfunc main(){fmt.Println(\"HelloWorld\")}\n```"
	md, _ := bcbh.GenerateMarkdown(ejsBlock)
	require.Equal(t, expectedResult, md)
}

func Test_CodeBoxHandler_GenerateMarkdown_Clean(t *testing.T) {
	bcbh := &goeditorjs.CodeBoxHandler{}
	jsonData := []byte(`{"language": "go", "code": "<span class=\"hljs-keyword\"><span class=\"hljs-keyword\">package</span></span> main<div><br></div><div>import <span class=\"hljs-string\"><span class=\"hljs-string\">\"fmt\"</span></span></div><div><br></div><div><span class=\"hljs-function\"><span class=\"hljs-keyword\"><span class=\"hljs-function\"><span class=\"hljs-keyword\">func</span></span></span><span class=\"hljs-function\"> </span><span class=\"hljs-title\"><span class=\"hljs-function\"><span class=\"hljs-title\">main</span></span></span><span class=\"hljs-params\"><span class=\"hljs-function\"><span class=\"hljs-params\">()</span></span></span></span> {</div><div>  fmt.Println(<span class=\"hljs-string\"><span class=\"hljs-string\">\"Hello World\"</span></span>)</div><div>}</div>"}`)
	ejsBlock := goeditorjs.EditorJSBlock{Type: "codeBox", Data: jsonData}
	expectedResult := "```go\npackage main\n\nimport \"fmt\"\n\nfunc main() {\n  fmt.Println(\"Hello World\")\n}\n```"
	md, _ := bcbh.GenerateMarkdown(ejsBlock)
	require.Equal(t, expectedResult, md)
}

func Test_RawHTMLHandler_Type(t *testing.T) {
	h := &goeditorjs.RawHTMLHandler{}
	require.Equal(t, "raw", h.Type())
}

func Test_RawHTMLHandler_GenerateHTML_Returns_Parse_Err(t *testing.T) {
	h := &goeditorjs.RawHTMLHandler{}
	_, err := h.GenerateHTML(goeditorjs.EditorJSBlock{Type: "raw", Data: []byte{}})
	require.Error(t, err)
}

func Test_RawHTMLHandler_GenerateHTML(t *testing.T) {
	h := &goeditorjs.RawHTMLHandler{}
	jsonData := []byte(`{"html": "<div style=\"background: #000; color: #fff; font-size: 30px; padding: 50px;\">Any HTML code</div>"}`)
	ejsBlock := goeditorjs.EditorJSBlock{Type: "raw", Data: jsonData}
	expectedResult := `<div style="background: #000; color: #fff; font-size: 30px; padding: 50px;">Any HTML code</div>`
	html, _ := h.GenerateHTML(ejsBlock)
	require.Equal(t, expectedResult, html)
}

func Test_RawHTMLHandler_GenerateMarkdown_Returns_Parse_Err(t *testing.T) {
	h := &goeditorjs.RawHTMLHandler{}
	_, err := h.GenerateMarkdown(goeditorjs.EditorJSBlock{Type: "raw", Data: []byte{}})
	require.Error(t, err)
}

func Test_RawHTMLHandler_GenerateMarkdown(t *testing.T) {
	h := &goeditorjs.RawHTMLHandler{}
	jsonData := []byte(`{"html": "<div style=\"background: #000; color: #fff; font-size: 30px; padding: 50px;\">Any HTML code</div>"}`)
	ejsBlock := goeditorjs.EditorJSBlock{Type: "raw", Data: jsonData}
	expectedResult := `<div style="background: #000; color: #fff; font-size: 30px; padding: 50px;">Any HTML code</div>`
	md, _ := h.GenerateMarkdown(ejsBlock)
	require.Equal(t, expectedResult, md)
}

func Test_ImageHandler_Type(t *testing.T) {
	h := &goeditorjs.ImageHandler{}
	require.Equal(t, "image", h.Type())
}

func Test_ImageHandler_GenerateHTML_Returns_Parse_Err(t *testing.T) {
	h := &goeditorjs.ImageHandler{}
	_, err := h.GenerateHTML(goeditorjs.EditorJSBlock{Type: "image", Data: []byte{}})
	require.Error(t, err)
}

func Test_ImageHandler_GenerateHTML(t *testing.T) {
	h := &goeditorjs.ImageHandler{}

	testData := []struct {
		data           string
		expectedResult string
	}{
		// Full test with defaults
		{data: `{"file":{"url": "https://www.w3schools.com/html/pic_trulli.jpg"},"caption": "Example Captions","withBorder": true,"stretched": true,"withBackground": true}`,
			expectedResult: `<img src="https://www.w3schools.com/html/pic_trulli.jpg" alt="Example Captions" class="image-tool--stretched image-tool--withBorder image-tool--withBackground"/>`},
		// No captions
		{data: `{"file":{"url": "https://www.w3schools.com/html/pic_trulli.jpg"},"caption": "","withBorder": true,"stretched": true,"withBackground": true}`,
			expectedResult: `<img src="https://www.w3schools.com/html/pic_trulli.jpg" alt="" class="image-tool--stretched image-tool--withBorder image-tool--withBackground"/>`},
		// Border
		{data: `{"file":{"url": "https://www.w3schools.com/html/pic_trulli.jpg"},"caption": "","withBorder": true,"stretched": false,"withBackground": false}`,
			expectedResult: `<img src="https://www.w3schools.com/html/pic_trulli.jpg" alt="" class="image-tool--withBorder"/>`},
		// Stretch
		{data: `{"file":{"url": "https://www.w3schools.com/html/pic_trulli.jpg"},"caption": "","withBorder": false,"stretched": true,"withBackground": false}`,
			expectedResult: `<img src="https://www.w3schools.com/html/pic_trulli.jpg" alt="" class="image-tool--stretched"/>`},
		// Background
		{data: `{"file":{"url": "https://www.w3schools.com/html/pic_trulli.jpg"},"caption": "","withBorder": false,"stretched": false,"withBackground": true}`,
			expectedResult: `<img src="https://www.w3schools.com/html/pic_trulli.jpg" alt="" class="image-tool--withBackground"/>`},
		// No classes
		{data: `{"file":{"url": "https://www.w3schools.com/html/pic_trulli.jpg"},"caption": "","withBorder": false,"stretched": false,"withBackground": false}`,
			expectedResult: `<img src="https://www.w3schools.com/html/pic_trulli.jpg" alt="" />`},
	}

	for _, td := range testData {
		jsonData := []byte(td.data)
		ejsBlock := goeditorjs.EditorJSBlock{Type: "image", Data: jsonData}
		result, _ := h.GenerateHTML(ejsBlock)
		require.Equal(t, td.expectedResult, result)
	}
}

func Test_ImageHandler_GenerateMarkdown_Returns_Parse_Err(t *testing.T) {
	h := &goeditorjs.ImageHandler{}
	_, err := h.GenerateMarkdown(goeditorjs.EditorJSBlock{Type: "image", Data: []byte{}})
	require.Error(t, err)
}

func Test_ImageHandler_GenerateMarkdown(t *testing.T) {
	h := &goeditorjs.ImageHandler{}

	testData := []struct {
		data           string
		expectedResult string
	}{
		// Full test with defaults
		{data: `{"file":{"url": "https://www.w3schools.com/html/pic_trulli.jpg"},"caption": "Example Captions","withBorder": true,"stretched": true,"withBackground": true}`,
			expectedResult: `<img src="https://www.w3schools.com/html/pic_trulli.jpg" alt="Example Captions" class="image-tool--stretched image-tool--withBorder image-tool--withBackground"/>`},
		// No captions
		{data: `{"file":{"url": "https://www.w3schools.com/html/pic_trulli.jpg"},"caption": "","withBorder": true,"stretched": true,"withBackground": true}`,
			expectedResult: `<img src="https://www.w3schools.com/html/pic_trulli.jpg" alt="" class="image-tool--stretched image-tool--withBorder image-tool--withBackground"/>`},
		// Border
		{data: `{"file":{"url": "https://www.w3schools.com/html/pic_trulli.jpg"},"caption": "","withBorder": true,"stretched": false,"withBackground": false}`,
			expectedResult: `<img src="https://www.w3schools.com/html/pic_trulli.jpg" alt="" class="image-tool--withBorder"/>`},
		// Stretch
		{data: `{"file":{"url": "https://www.w3schools.com/html/pic_trulli.jpg"},"caption": "","withBorder": false,"stretched": true,"withBackground": false}`,
			expectedResult: `<img src="https://www.w3schools.com/html/pic_trulli.jpg" alt="" class="image-tool--stretched"/>`},
		// Background
		{data: `{"file":{"url": "https://www.w3schools.com/html/pic_trulli.jpg"},"caption": "","withBorder": false,"stretched": false,"withBackground": true}`,
			expectedResult: `<img src="https://www.w3schools.com/html/pic_trulli.jpg" alt="" class="image-tool--withBackground"/>`},
		// No classes or caption
		{data: `{"file":{"url": "https://www.w3schools.com/html/pic_trulli.jpg"},"caption": "","withBorder": false,"stretched": false,"withBackground": false}`,
			expectedResult: `![alt text](https://www.w3schools.com/html/pic_trulli.jpg "")`},
		// No classes
		{data: `{"file":{"url": "https://www.w3schools.com/html/pic_trulli.jpg"},"caption": "Some caption","withBorder": false,"stretched": false,"withBackground": false}`,
			expectedResult: `![alt text](https://www.w3schools.com/html/pic_trulli.jpg "Some caption")`},
	}

	for _, td := range testData {
		jsonData := []byte(td.data)
		ejsBlock := goeditorjs.EditorJSBlock{Type: "image", Data: jsonData}
		result, _ := h.GenerateMarkdown(ejsBlock)
		require.Equal(t, td.expectedResult, result)
	}
}

func Test_TableHandler_GenerateMarkdown(t *testing.T) {
	h := &goeditorjs.TableHandler{}

	result, err := h.GenerateMarkdown(goeditorjs.EditorJSBlock{Type: "table", Data: json.RawMessage(`{"withHeadings":false,"stretched":false,"content":[["title","subtitle",""],["123","111",""],["333","2222",""]]}`)})
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, result, "| title | subtitle |  |\n| --- | --- | --- |\n| 123 | 111 |  |\n| 333 | 2222 |  |\n")
}
