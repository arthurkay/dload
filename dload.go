package dload

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/ermanimer/progress_bar"
)

func progress(currentProgress float64) {
	//create parameters
	output := os.Stdout
	schema := "({bar}) ({percent}) ({current} of {total} completed)"
	filledCharacter := "="
	blankCharacter := "-"
	var length float64 = 60
	var totalValue float64 = 100
	//create new progress bar
	pb := progress_bar.NewProgressBar(output, schema, filledCharacter, blankCharacter, length, totalValue)
	//start
	err := pb.Start()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//update
	for value := currentProgress; value <= 100; value++ {
		time.Sleep(20 * time.Millisecond)
		err := pb.Update(float64(value))
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}
}

func downloadPercent(done chan int64, path string, total int64) {
	var stop bool = false

	for {
		select {
		case <-done:
			stop = true
		default:
			file, err := os.Open(path)
			if err != nil {
				panic(err)
			}

			fi, er := file.Stat()
			if er != nil {
				panic(er)
			}

			size := fi.Size()

			// Make sure size is not equal to zero to prevent 0 division error
			if size == 0 {
				size = 1
			}

			var percent float64 = float64(size) / float64(total) * 100
			progress(percent)
			fmt.Println()
		}

		if stop {
			break
		}

		time.Sleep(time.Second)
	}
}

// Download takes in a string url and a string
// file download destination, then spawns multiple
// goroutines to download the file with a progress status
// bar showing the chunk downloads
func Download(url, dest string) {
	file := path.Base(url)

	var path bytes.Buffer
	path.WriteString(dest)
	path.WriteString("/")
	path.WriteString(file)

	start := time.Now()
	out, err := os.Create(path.String())
	if err != nil {
		panic(err)
	}

	defer out.Close()

	headResp, er := http.Head(url)
	if er != nil {
		panic(er)
	}
	defer headResp.Body.Close()

	size, e := strconv.Atoi(headResp.Header.Get("Content-Length"))

	if e != nil {
		panic(e)
	}

	done := make(chan int64)
	go downloadPercent(done, path.String(), int64(size))

	resp, erro := http.Get(url)

	if erro != nil {
		panic(erro)
	}

	n, eror := io.Copy(out, resp.Body)
	if eror != nil {
		panic(eror)
	}

	done <- n

	elapsed := time.Since(start)
	log.Printf("Download completed in %s", elapsed)
}
