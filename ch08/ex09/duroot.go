package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

type DirInfo struct {
	name   string
	nbytes int64
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan DirInfo)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		var _dirInfo DirInfo = DirInfo{name: root}
		go walkDir(root, &n, fileSizes, _dirInfo)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}
	var dirInfoSlice []DirInfo = nil
loop:
	for {
		select {
		case _dirInfo, ok := <-fileSizes:
			if !ok {
				break loop
			}
			dirInfoSlice = append(dirInfoSlice, _dirInfo)
		case <-tick:
			printDiskUsage(dirInfoSlice)
		}
	}

	printDiskUsage(dirInfoSlice)
}

type SumDir struct {
	nbytes int64
	nfiles int64
}

var _sum = map[string]SumDir{}

func printDiskUsage(dirInfoSlice []DirInfo) {
	for _, dirInfo := range dirInfoSlice {
		tmp := _sum[dirInfo.name]
		tmp.nbytes += dirInfo.nbytes
		tmp.nfiles += 1
		_sum[dirInfo.name] = tmp
	}

	for k, v := range _sum {
		fmt.Printf("dir name:%s  %d files  %.1f GB\n", k, v.nfiles, float64(v.nbytes)/1e9)
	}
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- DirInfo, dirInfo DirInfo) {
	dirInfo.nbytes = 0
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes, dirInfo)
		} else {
			dirInfo.nbytes = entry.Size()
			fileSizes <- dirInfo
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
