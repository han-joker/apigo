package commands

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func init() {
	Register(NewCmdInit("init", initIntro))
}

type CmdInit struct {
	Cmd
	Flags struct {
		Path string
	}
}

func (cmd *CmdInit) initFlagSet() {
	cmd.FlagSet = flag.NewFlagSet(cmd.Name, flag.ContinueOnError)
}

func NewCmdInit(name, intro string) *CmdInit {
	cmd := &CmdInit{}
	cmd.Name = name
	cmd.Intro = intro
	cmd.initFlagSet()

	return cmd
}

func (cmd *CmdInit) Run() {

	if err := cmd.FlagSet.Parse(os.Args[2:]); err != nil {
		return
	}

	dst := cmd.FlagSet.Arg(0)
	if dst == "" {
		dst = "."
	}

	templatePath := "./templates/gin"
	if err := CopyDir(templatePath, dst); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("init success")
}

func CopyDir(src, dst string) error {
	src, dst = filepath.Clean(src), filepath.Clean(dst)

	if dst != "." {
		if err := os.Mkdir(dst, 0750); err != nil {
			return err
		}
	}

	filesystem := os.DirFS(src)
	root := "."
	err := fs.WalkDir(filesystem, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path == root {
			return nil
		}
		if d.IsDir() {
			if err := os.Mkdir(dst+string(os.PathSeparator)+path, 0750); err != nil {
				return err
			}
		} else {
			srcFile, err := os.Open(src + string(os.PathSeparator) + path)
			if err != nil {
				return err
			}
			defer srcFile.Close()
			dstFile, err := os.OpenFile(dst+string(os.PathSeparator)+path, os.O_WRONLY|os.O_CREATE, 0750)
			if err != nil {
				return err
			}
			defer dstFile.Close()
			if _, err := io.Copy(dstFile, srcFile); err != nil {
				return err
			}
		}
		return nil
	})
	log.Println(err)
	return err
}
