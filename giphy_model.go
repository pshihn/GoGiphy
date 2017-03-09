package gogiphy

import (
	"strconv"
)

// Image represents a gif with a particular size
type Image struct {
	URL      string `json:"url"`
	Width    string `json:"width"`
	Height   string `json:"height"`
	Size     string `json:"size"`
	Frames   string `json:"frames"`
	Mp4      string `json:"mp4"`
	Mp4Size  string `json:"mp4_size"`
	Webp     string `json:"webp"`
	WebpSize string `json:"webp_size"`
}

// Gif represents a Gif object in Giphy
type Gif struct {
	Type             string           `json:"type"`
	ID               string           `json:"id"`
	URL              string           `json:"url"`
	Slug             string           `json:"slug"`
	BitlyGifURL      string           `json:"bitly_gif_url"`
	BitlyURL         string           `json:"bitly_url"`
	EmbedURL         string           `json:"embed_url"`
	Username         string           `json:"username"`
	Source           string           `json:"source"`
	Rating           string           `json:"rating"`
	Caption          string           `json:"caption"`
	ContentURL       string           `json:"content_url"`
	SourceTLD        string           `json:"source_tld"`
	SourcePostURL    string           `json:"source_post_url"`
	ImportDatetime   string           `json:"import_datetime"`
	TrendingDatetime string           `json:"trending_datetime"`
	Images           map[string]Image `json:"images"`
}

// Meta is the response to API call (HTTP Response data)
type Meta struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

// Pagination is the page information for the query result
type Pagination struct {
	TotalCount int `json:"total_count"`
	Count      int `json:"count"`
	Offset     int `json:"offset"`
}

// ListResponse is the response to the Giphy API that results in a lists of Gifs. e.g., search, trending
type ListResponse struct {
	Data       []Gif      `json:"data"`
	Meta       Meta       `json:"meta"`
	Pagination Pagination `json:"pagination"`
}

// GifResponse is the response to Giphy API that results in a single Gif e.g. random, translate, GifById
type GifResponse struct {
	Data Gif  `json:"data"`
	Meta Meta `json:"meta"`
}

// SearchParams - params use for search query. See https://github.com/Giphy/GiphyAPI#search-endpoint
type SearchParams struct {
	Limit  int
	Offset int64
	Rating string
	Lang   string
}

// GetWidth returns the width as int
func (image *Image) GetWidth() int64 {
	i, err := strconv.ParseInt(image.Width, 0, 64)
	if err != nil {
		return 0
	}
	return i
}

// GetHeight returns the height as int
func (image *Image) GetHeight() int64 {
	i, err := strconv.ParseInt(image.Height, 0, 64)
	if err != nil {
		return 0
	}
	return i
}

// GetSize returns the size as int64
func (image *Image) GetSize() int64 {
	i, err := strconv.ParseInt(image.Size, 0, 64)
	if err != nil {
		return 0
	}
	return i
}

// GetFrames returns the number of frames as int
func (image *Image) GetFrames() int64 {
	i, err := strconv.ParseInt(image.Frames, 0, 64)
	if err != nil {
		return 0
	}
	return i
}
