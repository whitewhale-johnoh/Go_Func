package main

import (
	"flag"
	"fmt"
)

func main() {
	title := flag.String("title", "", "Movie Title")
	runtime := flag.Int("runtime", 0, "Movie runtime")
	rating := flag.Float64("rating", 0.0, "Rating")
	release := flag.Bool("release", false, "release")

	flag.Parse()

	if flag.NFlag() == 0 {
		flag.Usage()
		return
	}

	fmt.Printf(
		"Movie Title: %s\n Runtime: %dmin\nRating:%f\n",
		*title, *runtime, *rating,
	)

	if *release == true {
		fmt.Println("Release : True")
	} else {
		fmt.Println("Release : False")
	}
}
