package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/9d4/mytodo/todo"
	"github.com/asdine/storm"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

const dbName = "mytodo.db"

type data struct {
	db    *storm.DB
	store todo.ToDoStore
}

type (
	cobraFunc  func(cmd *cobra.Command, args []string)
	actionFunc func(cmd *cobra.Command, args []string, d *data)
)

func run(fn actionFunc) cobraFunc {
	return func(cmd *cobra.Command, args []string) {
		d := &data{}

		dbpath := getDBPath()
		db, err := storm.Open(dbpath)
		if err != nil {
			log.Fatalln("error opening database: ", err)
		}

		d.db = db
		d.store = todo.NewStore(db)
		fn(cmd, args, d)
	}
}

// returns db path
func getDBPath() string {
	// check in current dir first
	wd, _ := os.Getwd()
	inWd := filepath.Join(wd, dbName)
	if _, err := os.Stat(inWd); err == nil {
		fmt.Println("INFO:", "using database in current working directory.")
		return inWd
	}

	// if in the current directory not found,
	// then return homedir concatenated with dbpath and dbName
	// where dbpath is .config/.
	home, err := homedir.Dir()
	if err != nil {
		// if there is error here then we should use current dir instead
		return inWd
	}

	return filepath.Join(home, ".config", dbName)
}
