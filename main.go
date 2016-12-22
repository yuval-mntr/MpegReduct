package main

import (
	//"encoding/json"
	// "bufio"

	"path/filepath"
	"strings"
	// "crypto/md5"
	// "encoding/hex"
	// "errors"
	//"fmt"
	// "io/ioutil"
	"log"
	// "net/http"
	"os"
	"os/exec"
	// "regexp"
	// "strconv"
	// "strings"
	// "time"
	// "github.com/gorilla/mux"
)

// ////////////////////////////////////////
func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

//////////////////////////////////////////
func run(args ...string) string {
	log.Printf("len is %d", len(args))
	cmd := exec.Command("ffmpeg", args...)

	//var out bytes.Buffer
	//cmd.Stdout = &out
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmerr := cmd.Start()
	if cmerr != nil {
		log.Fatal(cmerr)
		return "ffmpegError"
	}
	log.Printf("Waiting for command to finish...")
	err := cmd.Wait()

	//log.Printf("out = %s", out.String())
	log.Printf("+++ DONE +++")
	if err != nil {
		log.Fatal(err.Error())
		return "ffmpegErrorAfterExec"
	}
	//j.mpgFile = outFileStr
	//j.ffmpegTook = time.Since(start).String()
	return ""
}

////////////////////////////////////////
func main() {
	if len(os.Args) < 2 {
		log.Printf("ERR - Input file.mp4 was not provided")
		return
	}

	input := os.Args[1]
	outdir := "out_" + strings.Replace(input, ".mp4", "", 1)
	os.MkdirAll(outdir, 0777)
	RemoveContents(outdir)

	log.Printf("+++Convert mp4 for R-BOX")
	log.Printf("+++Input file = " + input)
	log.Printf("+++Outdir = " + outdir)

	// DESKTOP- ffmpeg -i %1 -vcodec libx264 -crf 30 -vf scale=320:-1 desktop-crf30-scl-320.mp4
	outFile := outdir + "/desktop-crf30-scl-320.mp4"
	run("-i", input, "-vcodec", "libx264", "-crf", "30", "-vf", "scale=320:-1", outFile)

	// MOBILE - ffmpeg -i %1 -f mpeg1video -crf 30 -vf scale=320:-1 mobile-crf30-scl-320.mpg
	outFile = outdir + "/mobile-crf30-scl-320.mpg"
	run("-i", input, "-f", "mpeg1video", "-crf", "30", "-vf", "scale=320:-1", outFile)

	//run(cmd)

	// MOBILE
	//exec.Command("ffmpeg -i", substring, "-f", "mpeg1video", outFileStr, "-y")
}

// func ffmpegFile(j *Job) error {
//     start := time.Now()
//     if len(j.ffmpegUrl) == 0 {
//         fs, err := os.Open(j.mp4File)
//         if err != nil {
//             j.err = "error reading file"
//             return err
//         }
//         defer fs.Close()
//     }

//     outFileStr := "mpeg/" + j.mp4ChunckMd5 + MPEG_EXT

//     var inFile string
//     var substring string
//     if len(j.ffmpegUrl) == 0 {
//         inFile = j.mp4File
//         substring = inFile[2:len(inFile)]
//     } else {
//         substring = j.ffmpegUrl
//     }

//     cmd := exec.Command("ffmpeg", "-i", substring, "-f", "mpeg1video", outFileStr, "-y", "")
//     //cmd.Dir = "/Users/cm/src/go-mpeg"
//     //cmd.Stdout = os.Stdout
//     //cmd.Stderr = os.Stderr
//
// }
