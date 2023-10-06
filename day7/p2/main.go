package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	radix "github.com/armon/go-radix"
)

func main() {
	fs := radix.New()
	cwd := ""
	fileRegex, _ := regexp.Compile(`\d+ [a-z]+`)
	cdIntoDir, _ := regexp.Compile(`\$ cd [a-z]+`)

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			if strings.HasPrefix(line, "$ cd /") {
				fs.Insert("/", 0)
				cwd = "/"
			} else if strings.HasPrefix(line, "$ cd ..") {
				pathComponents := strings.Split(cwd, "/")
				cwd = fmt.Sprintf("%s/", strings.Join(pathComponents[:len(pathComponents)-2], "/"))
			} else if cdIntoDir.MatchString(line) {
				newCwd := strings.Split(line, " ")[2]
				cwd = fmt.Sprintf("%s%s/", cwd, newCwd)
			} else if strings.HasPrefix(line, "dir ") {
				dirName := strings.Split(line, " ")[1]
				fs.Insert(fmt.Sprintf("%s%s/", cwd, dirName), 0)
			} else if fileRegex.MatchString(line) {
				file := strings.Split(line, " ")
				fSize, _ := strconv.Atoi(file[0])
				fs.Insert(fmt.Sprintf("%s%s", cwd, file[1]), fSize)
				fs.WalkPath(cwd, func(s string, v interface{}) bool {
					newSize := v.(int) + fSize
					fs.Insert(s, newSize)
					return false
				})
			}
		}
	}
	f.Close()
	rootInode, _ := fs.Get("/")
	freeDiskSpaceRequired := 30000000 - (70000000 - rootInode.(int))
	removalCandidates := make(map[string]int)
	fs.WalkPrefix("/", func(s string, v interface{}) bool {
		if s[len(s)-1] == '/' && v.(int) >= freeDiskSpaceRequired {
			removalCandidates[s] = v.(int)
		}
		return false
	})
	fsTreeKeys := make([]string, 0, len(removalCandidates))
	for key := range removalCandidates {
		fsTreeKeys = append(fsTreeKeys, key)
	}
	sort.Slice(fsTreeKeys, func(i, j int) bool { return removalCandidates[fsTreeKeys[i]] < removalCandidates[fsTreeKeys[j]] })
	for _, key := range fsTreeKeys {
		fmt.Printf("%d\n", removalCandidates[key])
		break
	}
}
