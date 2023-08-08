package dpsd

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type DPSD struct {
	Model     string
	InputDir  string
	OutputDir string
	SrcFiles  []SrcFile
}

type SrcFile struct {
	DCFDirName  string
	Filename    string
	FullPath    string
	NewFilename string
	Model       string
	Suffix      string
}

type NewDPSDInput struct {
	Model     string
	InputDir  string
	OutputDir string
}

func NewDPSD(in NewDPSDInput) *DPSD {
	t := new(DPSD)
	t.Model = in.Model
	t.InputDir = in.InputDir
	t.OutputDir = in.OutputDir
	return t
}

func (t *DPSD) Run() bool {
	err := t.findFiles()
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		log.Fatalf("alert: %v", err)
	}

	fmt.Printf("インポート元: %s\n", t.InputDir)
	fmt.Printf("インポート先: %s\n", t.OutputDir)

	err = t.copyFiles()
	if err != nil {
		log.Fatalf("alert: %v", err)
	}

	fmt.Printf("%s を全て削除しますか？", t.InputDir)
	if ask4Confirm() {
		runCommand("rm", "-rf", t.InputDir)
	}

	fmt.Printf("%s をアンマウントしますか？", t.InputDir)
	if ask4Confirm() {
		runCommand("diskutil", "unmount", t.InputDir)
	}

	return true
}

func (t *DPSD) findFiles() error {
	t.SrcFiles = []SrcFile{}

	inDir := filepath.Join(t.InputDir, "DCIM")
	dcimDirs, err := os.ReadDir(inDir)
	if err != nil {
		return err
	}
	for _, dcimDirE := range dcimDirs {
		if !dcimDirE.IsDir() {
			continue
		}
		dcfDir := filepath.Join(inDir, dcimDirE.Name())
		dcfEnts, err := os.ReadDir(dcfDir)
		if err != nil {
			return err
		}
		for _, dcfDirE := range dcfEnts {
			if dcfDirE.IsDir() {
				continue
			}
			dcimNum, _ := strconv.Atoi(dcimDirE.Name()[:3])
			dcfNum, _ := strconv.Atoi(dcfDirE.Name()[4:8])
			model := dcfDirE.Name()[:4]
			suffix := dcfDirE.Name()[9:12]

			// 105SIGMA/DP2M1234.JPG -> DP2M_0051234.JPEG
			newFilename := fmt.Sprintf("%s_%03d%04d.%s",
				model, dcimNum-100, dcfNum, suffix)

			t.SrcFiles = append(t.SrcFiles, SrcFile{
				DCFDirName:  dcimDirE.Name(),
				Filename:    dcfDirE.Name(),
				FullPath:    filepath.Join(dcfDir, dcfDirE.Name()),
				Model:       model,
				Suffix:      suffix,
				NewFilename: newFilename,
			})
		}
	}
	return nil
}

func (t *DPSD) copyFiles() error {
	nds := nowDateString()

	for _, sf := range t.SrcFiles {
		typeName := sf.Suffix
		if typeName == "JPG" {
			typeName = "JPEG"
		}
		outputDir := filepath.Join(t.OutputDir, nds, typeName)
		outputFn := filepath.Join(outputDir, sf.NewFilename)

		if _, err := os.Stat(outputDir); err != nil {
			fmt.Printf("[MKDIR] %s\n", outputDir)
			if err := os.MkdirAll(outputDir, 0755); err != nil {
				return err
			}
		}

		fmt.Printf("[COPY] %s/%s -> %s/%s\n", sf.DCFDirName, sf.Filename, typeName, sf.NewFilename)

		err := fileCopy(sf.FullPath, outputFn)
		if err != nil {
			return err
		}
	}

	return nil
}

func nowDateString() string {
	now := time.Now()
	return fmt.Sprintf("%02d%02d%02d", now.Year()-2000, now.Month(), now.Day())
}

func fileCopy(srcFn string, dstFn string) error {
	srcFh, err := os.Open(srcFn)
	if err != nil {
		return err
	}
	defer srcFh.Close()
	dstFh, err := os.Create(dstFn)
	if err != nil {
		return err
	}
	defer dstFh.Close()
	_, err = io.Copy(dstFh, srcFh)
	if err != nil {
		return err
	}
	return nil
}

func ask4Confirm() bool {
	var s string

	fmt.Printf("(y/N): ")
	_, err := fmt.Scan(&s)
	if err != nil {
		panic(err)
	}

	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if s == "y" || s == "yes" {
		return true
	}
	return false
}

func runCommand(c string, p ...string) error {
	cmd := exec.Command(c, p...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Start(); err != nil {
		return err
	}
	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}
