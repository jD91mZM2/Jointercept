package main

import (
	"fmt"
	"github.com/legolord208/stdutil"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type logItem struct {
	Msg  string
	Time string
}

var log = make([]logItem, 0)

const TIME_FORMAT = "03:04:05PM"
const DIR = "Join-AutoStart"

var README = filepath.Join(DIR, "README.txt")

const DIRINFO = "All programs in here are automatically ran by Jointercept\n" +
	"when a Join message is received. It is ran with the\n" +
	"message itself as a command line argument. Have fun!"

var TEMPLATE = template.Must(template.New("main").Parse(TEMPLATE_CODE))
var TEMPLATE_TABLE = template.Must(template.New("table").Parse(TEMPLATE_CODE_TABLE))

func handler(w http.ResponseWriter, r *http.Request) {
	msg := r.FormValue("message")

	if msg != "" {
		now := time.Now()

		fmt.Println("Intercepted \"" + msg + "\"")
		log = append(log, logItem{
			Msg:  msg,
			Time: now.Format(TIME_FORMAT),
		})

		mkdir()
		programs, err := ioutil.ReadDir(DIR)
		if err != nil {
			stdutil.PrintErr("Couldn't read directory", err)
			return
		}

		for _, program := range programs {
			name := program.Name()
			if strings.HasSuffix(name, ".txt") {
				continue
			}

			cmd := makecmd(name, msg)
			cmd.Dir = "Join-AutoStart"
			err := cmd.Start()

			if err != nil {
				stdutil.PrintErr("Warning: Couldn't start "+name, err)
			}
		}
	} else {
		TEMPLATE.Execute(w, log)
	}
}
func handlerTable(w http.ResponseWriter, r *http.Request) {
	TEMPLATE_TABLE.Execute(w, log)
}

func mkdir() {
	err := os.Mkdir(DIR, 0755)
	if err != nil && !os.IsExist(err) {
		stdutil.PrintErr("Error creating directory", err)
	}

	err = ioutil.WriteFile(README, []byte(DIRINFO), 0666)
	if err != nil && !os.IsExist(err) {
		stdutil.PrintErr("Error creating README", err)
	}
}

func main() {
	port := "8080"

	args := os.Args[1:]
	if len(args) >= 1 {
		_, err := strconv.Atoi(args[0])
		if err == nil {
			port = args[0]
		} else {
			stdutil.PrintErr("Argument is not a number", nil)
		}
	}
	fmt.Println("Service starting at port " + port + "!")
	fmt.Println("Visible at http://localhost:" + port + "/. Change port in arguments.")
	mkdir()

	http.HandleFunc("/", handler)
	http.HandleFunc("/table", handlerTable)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		stdutil.PrintErr("Error! Couldn't serve website", err)
	}
}
