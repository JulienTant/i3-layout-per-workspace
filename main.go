package main

import (
	"fmt"
	"log"
	"os"

	"go.i3wm.org/i3/v4"
)

func main() {
	args := os.Args[1:]
	mapping := map[string]i3.Layout{}
	log.Println("building pairs")
	for i := 0; i < len(args); i += 2 {
		// handle the possiblity that an argument is missing
		var layout i3.Layout = "NOT FOUND"
		if i+1 < len(args) {
			layout = i3.Layout(args[i+1])
		}

		mapping[args[i]] = layout
		log.Printf("%s => %s", args[i], layout)
	}

	if len(args)%2 != 0 {
		log.Fatalln("error: number of argument must be even")
	}

	recv := i3.Subscribe(i3.WorkspaceEventType)
	for recv.Next() {
		ev := recv.Event().(*i3.WorkspaceEvent)

		l, exist := mapping[ev.Current.Name]
		if ev.Change == "focus" && ev.Current.Layout != i3.Tabbed && exist {
			_, err := i3.RunCommand(fmt.Sprintf(`[workspace="%s"] layout %s`, ev.Current.Name, string(l)))
			if err != nil {
				if i3.IsUnsuccessful(err) {
					log.Printf("[ERR] command was unsuccessful %s", err)
					continue
				}
				log.Printf("[ERR] sending command %s", err)
				continue
			}
		}
	}

	log.Fatal(recv.Close())
}
