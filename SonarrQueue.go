package main

import (
	"encoding/json"
	"time"
)

type SonarrQueue struct {
	Page          int              `json:"page"`
	PageSize      int              `json:"pageSize"`
	SortKey       string           `json:"sortKey"`
	SortDirection string           `json:"sortDirection"`
	TotalRecords  int              `json:"totalRecords"`
	Records       []SonarrQueueItem `json:"records"`
}

type SonarrQueueItem struct {
	SeriesID                   int           `json:"seriesId"`
	EpisodeID                  int           `json:"episodeId"`
	SeasonNumber               int           `json:"seasonNumber"`
	Languages                  []Language    `json:"languages"`
	Quality                    Quality       `json:"quality"`
	CustomFormats              []interface{} `json:"customFormats"`
	CustomFormatScore          int           `json:"customFormatScore"`
	Size                       float64       `json:"size"`
	Title                      string        `json:"title"`
	Sizeleft                   float64       `json:"sizeleft"`
	Added                      time.Time     `json:"added"`
	Status                     string        `json:"status"`
	TrackedDownloadStatus      string        `json:"trackedDownloadStatus"`
	TrackedDownloadState       string        `json:"trackedDownloadState"`
	StatusMessages             []interface{} `json:"statusMessages"`
	ErrorMessage               string        `json:"errorMessage"`
	DownloadID                 string        `json:"downloadId"`
	Protocol                   string        `json:"protocol"`
	DownloadClient             string        `json:"downloadClient"`
	DownloadClientHasPostImportCategory bool `json:"downloadClientHasPostImportCategory"`
	Indexer                    string        `json:"indexer"`
	EpisodeHasFile             bool          `json:"episodeHasFile"`
	ID                         int           `json:"id"`
}

type Language struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Quality struct {
	Quality struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Source     string `json:"source"`
		Resolution int    `json:"resolution"`
	} `json:"quality"`
	Revision struct {
		Version  int  `json:"version"`
		Real     int  `json:"real"`
		IsRepack bool `json:"isRepack"`
	} `json:"revision"`
}
