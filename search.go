package main

import (
	"fmt"
	"frederik/search/arglist"
	"frederik/search/linkedlist"
	"io/fs"
	"io/ioutil"
	"path"
	"strings"
)

func nameMatches(s string, searchTerm string) bool {
	return strings.Contains(strings.ToLower(strings.Trim(s, " ")), searchTerm)
}

func findDFS(wd string, args arglist.Args) []fs.FileInfo {
	var fileInfos []fs.FileInfo
	files, err := ioutil.ReadDir(wd)
	if err != nil {
		return fileInfos
	}

	for _, file := range files {
		if nameMatches(file.Name(), args.SearchTerm) {
			fileInfos = append(fileInfos, file)
		}
		if args.Recursive && file.IsDir() {
			fileInfos = append(fileInfos, findDFS(path.Join(wd, file.Name()), args)...)
		}
	}

	return fileInfos
}

func findBFS(wd string, args arglist.Args) []fs.FileInfo {
	queue := linkedlist.New()
	queue.Queue(wd)
	var fileInfos []fs.FileInfo

	for e := queue.Head(); e != nil; e = e.Next() {
		var filePath string = string(e.Value)
		files, err := ioutil.ReadDir(filePath)
		if err == nil {
			for _, file := range files {
				if nameMatches(file.Name(), args.SearchTerm) {
					fileInfos = append(fileInfos, file)
				}
				if args.Recursive && file.IsDir() {
					queue.Queue(path.Join(filePath, file.Name()))
				}
			}
		}
	}
	return fileInfos
}

func findFile(wd string, args arglist.Args) []fs.FileInfo {
	if args.Strategy == arglist.DFS {
		return findDFS(wd, args)
	} else if args.Strategy == arglist.BFS {
		return findBFS(wd, args)
	}
	return []fs.FileInfo{}
}

func main() {

	args := arglist.New()

	files := findFile(args.Cwd, *args)

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir(), file.ModTime())
	}
}
