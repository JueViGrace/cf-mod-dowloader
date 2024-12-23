package main

type SearchReq struct {
	GameId       int    `json:"gameId"`
	GameVersion  string `json:"gameVersion"`
	SearchFilter string `json:"searchFilter"`
}

type APIResponse struct {
	Data []Mod
}

type Mod struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	LatestFiles []LatestFiles `json:"latestFiles"`
}

type LatestFiles struct {
	ID          int    `json:"id"`
	ModId       int    `json:"modId"`
	FileName    string `json:"fileName"`
	DownloadUrl string `json:"downloadUrl"`
}
