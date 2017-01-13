package main;

import (
	"fmt"
	"os"
	"net/http"
	"strconv"
	"io/ioutil"
	"strings"
)

const DIR = "Join-AutoStart";
const WEBINFO = "Hello =)\n" +
				"I am being used by Jointercept right now.\n" +
				"You can change my port by supplying them in my arguments.\n" +
				"Thank you!";
const DIRINFO = "All programs in here are automatically ran by Jointercept\n" +
				"when a Join message is received. It is ran with the\n" +
				"message itself as a command line argument. Have fun!";

func handler(w http.ResponseWriter, r *http.Request){
	msg := r.FormValue("message");
	
	if(len(msg) > 0){
		fmt.Println("Intercepted \"" + msg + "\"");

		mkdir();
		programs, _ := ioutil.ReadDir(DIR);
		for _, program := range programs{
			name := "./" + program.Name();
			if(strings.HasSuffix(name, ".txt")){
				continue;
			}

			cmd := makecmd(name, msg);
			cmd.Dir = "Join-AutoStart";
			//out, err := cmd.Output();
			err := cmd.Start();

			if(err != nil){
				fmt.Fprintln(os.Stderr, "Warning: Couldn't start " + name + ": \"" + err.Error() + "\"");
			}
			//fmt.Println(string(out));
		}
	}

	fmt.Fprintln(w, WEBINFO);
}

func mkdir(){
	_ = os.Mkdir(DIR, 0777);
	_ = ioutil.WriteFile(DIR + "/README.txt", []byte(DIRINFO), 0644);
}

func main(){
	port := "8080";

	args := os.Args[1:];
	if(len(args) >= 1){
		_, err := strconv.Atoi(args[0]);
		if(err == nil){
			port = args[0];
		}
	}
	fmt.Println("Service starting at port " + port + "!");
	fmt.Println("Visible at http://localhost:" + port + "/. Change port in arguments.");
	mkdir();

	http.HandleFunc("/", handler);
	err := http.ListenAndServe(":" + port, nil);
	if(err != nil){
		fmt.Fprintln(os.Stderr, "Error! Couldn't serve website!");
	}
}
