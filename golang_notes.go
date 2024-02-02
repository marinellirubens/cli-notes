package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"

	cli "github.com/urfave/cli/v2"
)

var HOME string

func zipSource(source, target string) error {
	currDir, _ := os.Getwd()
	os.Chdir(get_user_home())
	tempZipName := ".notes.zip"
	// 1. Create a ZIP file and zip.Writer
	f, err := os.Create(tempZipName)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := zip.NewWriter(f)
	defer writer.Close()

	files, err := os.ReadDir(HOME)
	if err != nil {
		log.Fatal("Error trying to read folder", err)
		return err
	}
	// filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
	for _, file := range files {
		path := filepath.Join(".notes", file.Name())
		fmt.Println(path)
		if err != nil {
			log.Fatal("Error trying to read folder", err)
			continue
		}
		info, err := file.Info()

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			log.Fatal("Error trying to create file header", err)
			continue
		}
		fmt.Println(header.Name)

		// set compression
		header.Method = zip.Deflate

		// 5. Create writer for the file header and save content of the file
		headerWriter, err := writer.CreateHeader(header)
		if err != nil {
			continue
		}

		if info.IsDir() {
			continue
		}
		fmt.Println(path)

		f, err := os.Open(path)
		if err != nil {
			log.Fatal("Error trying to open file", err)
			continue
		}
		defer f.Close()
		// fmt.Println(headerWriter)
		_, err = io.Copy(headerWriter, f)
		if err != nil {
			log.Fatal("Error trying to copy file", err)
			continue
		}
	}
	os.Chdir(currDir)
	os.Rename(filepath.Join(get_user_home(), tempZipName), target)
	return nil
}

func zipFolder(cCtx *cli.Context) error {
	fileExport := cCtx.String("output")
	if err := zipSource(HOME, fileExport); err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Notes exported to file " + cCtx.String("output"))
	return nil
}

// Reads a note from the directory ~/.notes
func read_note(filename string) error {
	if filename == "" {
		list_notes(nil)
		return nil
	}

	fullPath := HOME + "/" + filename
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		fmt.Println("Note not found")
		return err
	}

	// read note from file
	data, err := os.ReadFile(fullPath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println(string(data))
	return nil
}

// Returns the user home directory
func get_user_home() string {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return user.HomeDir
}

// Deletes a note from the directory ~/.notes
func detele_note(filename string) error {
	fullPath := HOME + "/" + filename
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		fmt.Println("Note not found")
		return err
	}

	err := os.Remove(fullPath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("Note " + filename + " deleted")
	return nil
}

// Initializes the directory ~/.notes
func init() {
	// check if directory exists
	path := get_user_home() + "/.notes"
	HOME = path

	if _, err := os.Stat(path); os.IsNotExist(err) {
		// create directory
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
}

// Reads the stdin content
func get_stdin() []byte {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		var stdin []byte
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			stdin = append(stdin, scanner.Bytes()...)
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("stdin = %s\n", stdin)
		return stdin
	}
	return nil
}

// Initializes the editor with the note name
func init_editor(note_name string) error {
	if note_name == "" {
		fmt.Println("No note name provided")
		return nil
	}
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}
	var cmd string
	cmd = editor + " " + get_user_home() + "/.notes/" + note_name
	cmd_exec := exec.Command("sh", "-c", cmd)
	cmd_exec.Stdout = os.Stdout
	cmd_exec.Stderr = os.Stderr
	cmd_exec.Stdin = os.Stdin

	err := cmd_exec.Run()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// Lists all notes on the directory ~/.notes
func list_notes(cCtx *cli.Context) error {
	// list all notes on directory ~/.notes
	files, err := os.ReadDir(get_user_home() + "/.notes")
	if err != nil {
		return err
	}
	if len(files) == 0 {
		fmt.Println("No notes found")
		return nil
	}

	fmt.Println("Notes:")
	for _, f := range files {
		fmt.Println("	" + f.Name())
	}
	return nil
}

func main() {
	app := &cli.App{
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:  "add",
				Usage: "Adds a note to the list",
				Action: func(cCtx *cli.Context) error {
					init_editor(cCtx.Args().First())
					fmt.Println("added task: ", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:  "delete",
				Usage: "Deletes a note from the list",
				Action: func(cCtx *cli.Context) error {
					for _, arg := range cCtx.Args().Slice() {
						detele_note(arg)
					}
					return nil
				},
			},
			{
				Name:   "list",
				Usage:  "Lists all notes",
				Action: list_notes,
			},
			{
				Name:  "get",
				Usage: "Gets a note from the list",
				Action: func(cCtx *cli.Context) error {
					read_note(cCtx.Args().First())
					return nil
				},
			},
			{
				Name:   "export",
				Usage:  "Exports all notes to a file",
				Action: zipFolder,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Value:   ".notes.zip",
						Usage:   "Output file",
					},
				},
			},
			{
				Name:  "import",
				Usage: "imports file with notes exported with export command",
				Action: func(cCtx *cli.Context) error {
					fileImport := cCtx.String("input")
					// check if file exists
					if _, err := os.Stat(fileImport); os.IsNotExist(err) {
						fmt.Println("File " + fileImport + " does not exists")
						return nil
					}

					// create zip file with all notes from ~/.notes
					// fmt.Println("unzip", "-o", fileImport, "-d", HOME)
					cmd := exec.Command("unzip", "-o", fileImport, "-d", get_user_home())
					// cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					cmd.Stdin = os.Stdin
					err := cmd.Run()
					if err != nil {
						log.Fatal(err)
					}
					// create zip file with all notes from ~/.notes
					fmt.Println(cCtx.String("input"))
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "input",
						Aliases: []string{"i"},
						Value:   ".notes.zip",
						Usage:   "Input file",
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
