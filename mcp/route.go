package mcp

import (
	"fmt"
	"mangadex_mcp/api"
	"mangadex_mcp/repository"

	"github.com/mark3labs/mcp-go/server"
)

type MCPRoute struct {
	s *server.MCPServer
}

type Handler struct {
	MangaAPI *api.MangaAPI
	mapRepo  *repository.MapNameWithIDRepo
}

func NewMCPRoute(mangaApiPackage *api.MangaAPI, mapRepo *repository.MapNameWithIDRepo) *MCPRoute {
	// Add tool

	// Create MCP server
	s := server.NewMCPServer(
		"Mangadex MCP 🚀",
		"1.0.0",
		// server.WithResourceCapabilities(true, true),
		// server.WithLogging(),
		server.WithInstructions("The is a MCP server for MangaDex that's contain only api for manga"),
	)

	handler := &Handler{
		MangaAPI: mangaApiPackage,
		mapRepo:  mapRepo,
	}

	s.AddTool(handler.mangaTool(), handler.mangaHandler)

	return &MCPRoute{s: s}
}

func (r *MCPRoute) Run() {
	// Add tool handler

	// Start the stdio server
	if err := server.ServeStdio(r.s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
