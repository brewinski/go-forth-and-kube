package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/rivo/tview"
)

const (
	indexPage = "*index*" // The name of the Finder page.
	wrkDir    = "/Users/chris.brewin/Development"
	githubDir = "/Users/chris.brewin/Documents/github"
)

var (
	app         *tview.Application // The tview application.
	pages       *tview.Pages       // The application pages.
	finderFocus tview.Primitive    // The primitive in the Finder that last had focus.
)

func main() {
	app = tview.NewApplication()

	setupNavigation(app)

	if err := app.Run(); err != nil {
		fmt.Printf("Error running application: %s\n", err)
	}
}

func setupNavigation(app *tview.Application) {
	subCommands := tview.NewList().
		ShowSecondaryText(false)
	subCommands.SetBorder(true).
		SetTitle("Sub Commands")

	brewinskiCLICommands := tview.NewList().
		ShowSecondaryText(false).
		AddItem("open with vscode", "Search for repos on device", 0, func() {
			app.SetFocus(subCommands)
		}).
		AddItem("brewinski-cli", "Brewinski CLI", 0, nil)
	brewinskiCLICommands.SetBorder(true).
		SetTitle("Brewinski CLI Commands")

	subCommands.SetFocusFunc(func() {
		subCommands.Clear()
		repos, err := listFoldersInDirectory(githubDir)
		if err != nil {
			log.Fatalf("unable to list repos in directory: %w", err)
		}

		subCommands.AddItem("Back", "", 0, func() {
			app.SetFocus(brewinskiCLICommands)
			subCommands.Clear()
		})

		addArrayItemsToList(subCommands, repos)
	})

	// create the layout
	flex := tview.NewFlex().
		AddItem(brewinskiCLICommands, 0, 1, true).
		AddItem(subCommands, 0, 1, true)

	// Set up the pages and show the Finder.
	pages = tview.NewPages().
		AddPage(indexPage, flex, true, true)

	app.SetRoot(pages, true)
}

func addArrayItemsToList(list *tview.List, items []FolderListItem) {
	for _, item := range items {
		command := createOpenWithVsCodeCommand(item.Path)
		list.AddItem(item.Name, "", 0, func() {
			if err := command.Run(); err != nil {
				log.Fatal(err)
			}
		})
	}
}

type FolderListItem struct {
	Name string
	Path string
}

func listFoldersInDirectory(dir string) ([]FolderListItem, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %w", err)
	}

	directories := []FolderListItem{}
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		directories = append(directories, FolderListItem{
			Name: entry.Name(),
			Path: fmt.Sprintf("%s/%s", dir, entry.Name()),
		})
	}

	return directories, nil
}

func createOpenWithVsCodeCommand(folder string) *exec.Cmd {
	cmd := exec.Command("code", folder)
	return cmd
}
