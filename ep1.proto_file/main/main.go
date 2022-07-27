package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"main/user"
)

func main() {
	ReadFile()
}

func ReadFile() {
	in, err := ioutil.ReadFile("0727.group")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	g := &user.Group{}
	if err = proto.Unmarshal(in, g); err != nil {
		log.Fatalln("Failed to parse group:", err)
	}
	fmt.Printf("%v", len(g.User))
	fmt.Printf("%+v", g)
}

func WiterFile() {
	u := user.User{
		Name:  "Eli Wu",
		Age:   24,
		Phone: "13870919778",
	}
	g := user.Group{
		User: []*user.User{
			{
				Name:  "Jamse Zhang",
				Age:   18,
				Phone: "13888888888",
			},
			&u,
		},
	}
	out, err := proto.Marshal(&g)
	if err != nil {
		log.Fatalln("Failed to encode group:", err)
	}
	err = ioutil.WriteFile("0727.group", out, 0644)
	if err != nil {
		log.Fatalln("Failed to write group file")
	}
}
