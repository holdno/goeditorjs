package goeditorjs

import (
	"encoding/json"
	"fmt"
	"strings"
)

// MarkdownEngine is the engine that creates the HTML from EditorJS blocks
type MarkdownEngine struct {
	StaticDomain  string
	BlockHandlers map[string]MarkdownBlockHandler
}

// MarkdownBlockHandler is an interface for a plugable EditorJS HTML generator
type MarkdownBlockHandler interface {
	Type() string // Type returns the type the block handler supports as a string
	GenerateMarkdown(editorJSBlock EditorJSBlock) (string, error)
}

type MarkdownEngineOptions func(m *MarkdownEngine)

func WithStaticDomain(domain string) MarkdownEngineOptions {
	return func(m *MarkdownEngine) {
		m.StaticDomain = domain
	}
}

// NewMarkdownEngine creates a new MarkdownEngine
func NewMarkdownEngine(opts ...MarkdownEngineOptions) *MarkdownEngine {
	bhs := make(map[string]MarkdownBlockHandler)
	m := &MarkdownEngine{BlockHandlers: bhs}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// RegisterBlockHandlers registers or overrides a block handlers for blockType given by MarkdownBlockHandler.Type()
func (markdownEngine *MarkdownEngine) RegisterBlockHandlers(handlers ...MarkdownBlockHandler) {
	for _, bh := range handlers {
		markdownEngine.BlockHandlers[bh.Type()] = bh
	}
}

// GenerateMarkdown generates markdown from the editorJS using configured set of markdown handlers
func (markdownEngine *MarkdownEngine) GenerateMarkdown(editorJSData string) (string, error) {
	results := []string{}
	ejs, err := parseEditorJSON(editorJSData)
	if err != nil {
		return "", err
	}
	for _, block := range ejs.Blocks {
		if generator, ok := markdownEngine.BlockHandlers[block.Type]; ok {
			md, err := generator.GenerateMarkdown(block)
			if err != nil {
				return "", err
			}
			results = append(results, md)
		} else {
			return "", fmt.Errorf("%w, Block Type: %s", ErrBlockHandlerNotFound, block.Type)
		}
	}

	return strings.Join(results, "\n\n"), nil
}

func unknownMarkdownBlockHandler(data EditorJSBlock) string {
	raw, _ := json.MarshalIndent(data.Data, "", "  ")
	return fmt.Sprintf("```json\n// type: %s\n%s\n```", data.Type, string(raw))
}

// GenerateMarkdown generates markdown from the editorJS using configured set of markdown handlers
func (markdownEngine *MarkdownEngine) GenerateMarkdownWithUnknownBlock(editorJSData string) (string, error) {
	results := []string{}
	ejs, err := parseEditorJSON(editorJSData)
	if err != nil {
		return "", err
	}
	for _, block := range ejs.Blocks {
		if generator, ok := markdownEngine.BlockHandlers[block.Type]; ok {
			md, err := generator.GenerateMarkdown(block)
			if err != nil {
				results = append(results, unknownMarkdownBlockHandler(block))
				continue
			}
			results = append(results, md)
		} else {
			results = append(results, unknownMarkdownBlockHandler(block))
		}
	}

	return strings.Join(results, "\n\n"), nil
}
