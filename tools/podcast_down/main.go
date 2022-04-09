package main

import (
	"bufio"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	ConfigDir = "rss.conf"
)

func main() {
	fmt.Println("work")

	urls, err := getUrls()
	if err != nil {
		logErr(err)
		return
	}

	for _, url := range urls {
		if err := download(url); err != nil {
			logErr(err)
		}
	}
}
func download(url string) error {
	fmt.Println("download: ", url)

	//
	rssStr, err := get(url)
	if err != nil {
		return err
	}
	//
	rss := Rss{}
	if err := xml.Unmarshal(rssStr, &rss); err != nil {
		return err
	}
	//
	if err := mkdir(rss.Channel.Title); err != nil {
		return err
	}
	//
	for _, v := range rss.Channel.Item {
		format := getFormat(v.Enclosure.URL)
		path := fmt.Sprintf("%s/%s.%s", rss.Channel.Title, v.Title, format)
		fmt.Println(fmt.Sprintf("download: %s", path))
		if err := save(v.Enclosure.URL, path); err != nil {
			logErr(err)
			continue
		}
	}
	return nil
}

func logErr(err error) {
	fmt.Printf("error: %s", err)
	fmt.Println()
}

func getFormat(url string) string {
	strs := strings.Split(url, ".")
	if len(strs) < 2 {
		return ""
	}
	return strs[len(strs)-1]
}

func get(url string) ([]byte, error) {
	if url == "" {
		return nil, errors.New("url nil")
	}
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return robots, nil
}

func save(url string, path string) error {
	if exist, _ := pathExists(path); exist {
		return errors.New("path exist")
	}
	if url == "" {
		return errors.New("url nil")
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}
	return nil
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func mkdir(name string) error {
	if name == "" {
		return errors.New("name nil")
	}
	if exist, _ := pathExists(name); exist {
		return nil
	}
	return os.Mkdir(name, os.ModePerm)
}

func getUrls() ([]string, error) {
	urls := []string{}
	fi, err := os.Open(ConfigDir)
	if err != nil {
		return urls, err
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		urls = append(urls, string(a))
	}
	return urls, nil
}

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Atom    string   `xml:"atom,attr"`
	Itunes  string   `xml:"itunes,attr"`
	Version string   `xml:"version,attr"`
	Channel struct {
		Text  string `xml:",chardata"`
		Title string `xml:"title"`
		Link  struct {
			Text string `xml:",chardata"`
			Href string `xml:"href,attr"`
			Rel  string `xml:"rel,attr"`
			Type string `xml:"type,attr"`
		} `xml:"link"`
		Description string `xml:"description"`
		Generator   string `xml:"generator"`
		WebMaster   string `xml:"webMaster"`
		Author      string `xml:"author"`
		Explicit    string `xml:"explicit"`
		Language    string `xml:"language"`
		Image       struct {
			Text  string `xml:",chardata"`
			URL   string `xml:"url"`
			Title string `xml:"title"`
			Link  string `xml:"link"`
		} `xml:"image"`
		LastBuildDate string `xml:"lastBuildDate"`
		Ttl           string `xml:"ttl"`
		Item          []struct {
			Text        string `xml:",chardata"`
			Title       string `xml:"title"`
			Description string `xml:"description"`
			PubDate     string `xml:"pubDate"`
			Guid        struct {
				Text        string `xml:",chardata"`
				IsPermaLink string `xml:"isPermaLink,attr"`
			} `xml:"guid"`
			Link  string `xml:"link"`
			Image struct {
				Text string `xml:",chardata"`
				Href string `xml:"href,attr"`
			} `xml:"image"`
			Enclosure struct {
				Text string `xml:",chardata"`
				URL  string `xml:"url,attr"`
				Type string `xml:"type,attr"`
			} `xml:"enclosure"`
			Duration string `xml:"duration"`
		} `xml:"item"`
	} `xml:"channel"`
}
