package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	// BingURL is the base url for bing
	BingURL = "http://www.bing.com"
	// ImageAPIEndpoint is the API endpoint to get today's wallpaper
	ImageAPIEndpoint = "/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=en-IN"
)

// ImageResponse is the response obtained from ImageApiEndpoint
type ImageResponse struct {
	Images []Image `json:"images"`
}

// Image is the properties of the image in the ImageResponse
type Image struct {
	URL string `json:"url"`
}

func main() {
	url := GetWallpaperURL()
	log.Println("Downloading wallpaper")
	wallpaperPath := GetWallpaperPath()
	DownloadWallpaper(url, wallpaperPath)
	log.Println("Wallpaper downloaded")
	log.Printf("Setting wallpaper to %s\n", wallpaperPath)
	SetWallpaper(wallpaperPath)
}

// GetWallpaperURL returns the url for today's wallpaper from bing
func GetWallpaperURL() string {
	res, err := http.Get(BingURL + ImageAPIEndpoint)
	if err != nil {
		log.Fatalf("Failed to get response.\nError is: %v\n", err)
	}
	defer res.Body.Close()
	// Decode json
	decoder := json.NewDecoder(res.Body)
	var imgResponse ImageResponse
	err = decoder.Decode(&imgResponse)
	if err != nil {
		log.Fatalf("Failed to decode json.\nError is: %v\n", err)
	}
	return imgResponse.Images[0].URL
}

// DownloadWallpaper downloads the wallpaper from the provided url
// It stores the wallpaper in the path provided
func DownloadWallpaper(url string, path string) {
	res, err := http.Get(BingURL + url)
	if err != nil {
		log.Fatalf("Failed to download image.\nError is: %v\n", err)
	}
	wallpaper, err := os.Create(path)
	if err != nil {
		log.Fatalf("Unable to create file.\nError is: %v\n", err)
	}
	defer res.Body.Close()
	defer wallpaper.Close()
	io.Copy(wallpaper, res.Body)
}
