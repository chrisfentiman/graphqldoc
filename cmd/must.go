package cmd

import "log"

func must(v interface{}, err error) interface{} {
	if err != nil {
		log.Fatalln(err)
	}
	return v
}
