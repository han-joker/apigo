package commands

import (
	"bytes"
	"flag"
	"fmt"
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
		Mod string
	}
}

func (cmd *CmdInit) initFlagSet() {
	cmd.FlagSet = flag.NewFlagSet(cmd.Name, flag.ExitOnError)
	cmd.FlagSet.StringVar(&cmd.Flags.Mod, "mod", "", "the mod name of project")
}

func NewCmdInit(name, intro string) *CmdInit {
	cmd := &CmdInit{}
	cmd.Name = name
	cmd.Intro = intro
	cmd.initFlagSet()

	return cmd
}

func (cmd *CmdInit) Run() {
	dst := cmd.FlagSet.Arg(0)
	if dst == "" {
		dst = "."
		cmd.FlagSet.Parse(os.Args[2:])
	} else {
		cmd.FlagSet.Parse(os.Args[3:])
	}
	if dst == "." {
		log.Println(". is not allowed in testing")
		return
	}

	if cmd.Flags.Mod == "" {
		cmd.Flags.Mod = dst
	}
	const PhMN = "PH_ModuleName_PH"
	replaces := map[string]string{
		PhMN: cmd.Flags.Mod,
	}

	templatePath := "../apigo_layout"
	if err := CopyLayout(templatePath, dst, replaces); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("init success")
}

func CopyLayout(src, dst string, replaces map[string]string) error {
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

			// replace placeholder
			srcBytes, err := os.ReadFile(src + string(os.PathSeparator) + path)
			if err != nil {
				return err
			}
			for f, r := range replaces {
				srcBytes = bytes.ReplaceAll(srcBytes, []byte(f), []byte(r))
			}

			// write new file
			if err := os.WriteFile(dst+string(os.PathSeparator)+path, srcBytes, 0750); err != nil {
				return err
			}
		}
		return nil
	})
	return err
}
