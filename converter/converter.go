package converter

import (
	"crypto/sha256"
	"fmt"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"golang.org/x/image/webp"
)

type Converter struct {
	// Output Path
	Output string `yaml:"OutputPath"`
	// Should the Original WebP File be Removed, not yet implemented
	RMOriginal bool `yaml:"RemoveOriginal"`

	// Remember last File's Hash so there will be only 1 output
	LastFile string
}

func New() *Converter {
	homedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	converter := &Converter{
		Output:     filepath.FromSlash(homedir + "/Downloads/"),
		RMOriginal: false,
		LastFile:   "",
	}
	return converter
}

func (c *Converter) GetFilename(dec string) string {
	arr := strings.Split(dec, "\\")
	fmt.Println("Got Filename")
	return arr[len(arr)-1]
}

func (c *Converter) CreateOutput(filename string) *os.File {
	// Check if file with the name exists
	fsplit := strings.Split(filename, ".")
	outputfilename := fsplit[0]
	fulloutpath := c.Output + outputfilename + ".png"
	fmt.Println(fulloutpath)

	i := 1
	_, err := os.Stat(fulloutpath) // if file does not exist error gets defined
	for err == nil {
		outputfilename = fsplit[0] + "(" + strconv.Itoa(i) + ")"
		fulloutpath = c.Output + outputfilename + ".png"
		i++
		_, err = os.Stat(fulloutpath)
		fmt.Println(fulloutpath)
		fmt.Println(err)

	}

	out, err := os.Create(fulloutpath)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Opened File :", fulloutpath)
	return out

}

func (c *Converter) Hash(dec string) string {
	file, err := os.Open(dec)
	if err != nil {
		fmt.Println(err)
	}

	hash := sha256.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		fmt.Println(err)
	}
	file.Close()
	return string(hash.Sum(nil))

}

func (c *Converter) WaitWriteComplete(dec string) bool {
	for {
		filestat, err := os.Stat(dec)
		if err != nil {
			fmt.Println(err)
		}

		timepast := time.Now().Sub(filestat.ModTime())
		if timepast > time.Second*1 {
			return true
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}

func (c *Converter) Convert(s chan string) {
	for {
		select {
		case dec := <-s:

			// Wait for Write being Complete
			c.WaitWriteComplete(dec)

			// Check if last file is identical
			filehash := c.Hash(dec)
			if filehash == c.LastFile {
				continue
			} else {
				c.LastFile = filehash
			}

			time.Sleep(1 * time.Second)
			reader, err := os.Open(dec)
			if err != nil {
				fmt.Println(err)
				continue
			}

			pngout := c.CreateOutput(c.GetFilename(dec))

			image, err := webp.Decode(reader)
			reader.Close()

			err = png.Encode(pngout, image)
			pngout.Close()
			if err != nil {
				fmt.Println(err)
			}
		}

	}

}
