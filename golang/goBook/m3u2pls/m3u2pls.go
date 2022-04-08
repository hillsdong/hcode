package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 1 ||
		(!strings.HasSuffix(os.Args[1], ".m3u") && !strings.HasSuffix(os.Args[1], ".pls")) {
		fmt.Printf("Usage: %s <file.m3u>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	if raw, err := ioutil.ReadFile(os.Args[1]); err == nil {
		if strings.HasSuffix(os.Args[1], ".m3u") {
			printPls(readM3u(string(raw)))
		} else {
			printM3u(readPls(string(raw)))
		}
	} else {
		log.Fatal(err)
	}
}

// Song struc
type Song struct {
	title   string
	url     string
	seconds int
}

func readM3u(content string) (songs []Song) {
	var song Song
	for _, line := range strings.Split(content, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#EXTM3U") {
			continue
		}
		if strings.HasPrefix(line, "#EXTINF") {
			if i := strings.IndexAny(line, "-0123456789"); i > -1 {
				line = line[i:]
				if j := strings.Index(line, ","); j > -1 {
					song.title = line[j+1:]
					if seconds, err := strconv.Atoi(line[:j]); err != nil {
						fmt.Printf("parse %s's senconds error :%v\n", song.title, err)
						song.seconds = -1
					} else {
						song.seconds = seconds
					}
				}
			}
		} else {
			song.url = line
		}
		if song.url != "" && song.title != "" && song.seconds != 0 {
			songs = append(songs, song)
			song = Song{}
		}
	}
	return songs
}
func printPls(songs []Song) {
	fmt.Printf("[playlist]\n")
	for i, song := range songs {
		fmt.Printf("File%d=%s\n", i+1, song.url)
		fmt.Printf("Title%d=%s\n", i+1, song.title)
		fmt.Printf("Length%d=%d\n", i+1, song.seconds)
	}
	fmt.Printf("NumberOfEntries=%d\n", len(songs))
	fmt.Printf("Version=2\n")
}

func readPls(content string) (songs []Song) {
	var song Song
	for _, line := range strings.Split(content, "\n") {
		line = strings.TrimSpace(line)

		if i := strings.Index(line, "="); i > -1 {
			if strings.HasPrefix(line, "File") {
				song.url = line[i+1:]
			} else if strings.HasPrefix(line, "Title") {
				song.title = line[i+1:]
			} else if strings.HasPrefix(line, "Length") {
				if seconds, err := strconv.Atoi(line[i+1:]); err != nil {
					fmt.Printf("parse %s's senconds error :%v\n", song.title, err)
					song.seconds = -1
				} else {
					song.seconds = seconds
				}
			}
		}

		if song.url != "" && song.title != "" && song.seconds != 0 {
			songs = append(songs, song)
			song = Song{}
		}
	}
	return songs
}
func printM3u(songs []Song) {
	fmt.Printf("#EXTM3U\n")
	for _, song := range songs {
		fmt.Printf("#EXTINF:%d,%s\n%s\n", song.seconds, song.title, song.url)
	}
}
