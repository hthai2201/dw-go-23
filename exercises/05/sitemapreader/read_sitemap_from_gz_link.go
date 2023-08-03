package sitemapreader

import (
	"compress/gzip"
	"encoding/xml"
	"io"
	"net/http"
	"os"
	"strings"
)

type Urlset struct {
	XMLUrlSet xml.Name `xml:"urlset"`
	Urls   []Url   `xml:"url"`	
}

type Url struct {
	Loc      string `xml:"loc"`
	Lastmod  string `xml:"lastmod"`
}

func ReadSitemapFromGZLink(url string) (*Urlset, error) {
	gzFilePath := "temp_sitemap.xml.gz"

	err := downloadFile(url, gzFilePath)
	if err != nil {
		return nil, err
	}
	defer os.Remove(gzFilePath)

	xmlFilePath, err := decompressGzFile(gzFilePath)
	if err != nil {
		return nil, err
	}
	defer os.Remove(xmlFilePath)

	xmlContent, err := readXMLFile(xmlFilePath)
	if err != nil {
		return nil, err
	}

	var sitemap Urlset
	err = xml.Unmarshal([]byte(xmlContent), &sitemap)
	if err != nil {
		return nil, err
	}

	return &sitemap, nil
}

func downloadFile(url, filePath string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, response.Body)
	return err
}

func decompressGzFile(gzFilePath string) (string, error) {
	xmlFilePath := strings.TrimSuffix(gzFilePath, ".gz")
	gzFile, err := os.Open(gzFilePath)
	if err != nil {
		return "", err
	}
	defer gzFile.Close()

	xmlFile, err := os.Create(xmlFilePath)
	if err != nil {
		return "", err
	}
	defer xmlFile.Close()

	gzReader, err := gzip.NewReader(gzFile)
	if err != nil {
		return "", err
	}
	defer gzReader.Close()

	_, err = io.Copy(xmlFile, gzReader)
	return xmlFilePath, err
}

func readXMLFile(filePath string) (string, error) {
	xmlData, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(xmlData), nil
}
