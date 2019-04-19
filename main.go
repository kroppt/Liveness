package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/kroppt/StringSet"
)

type block struct {
	Name  string
	Def   string
	Use   []string
	Succs []int
	Preds []int
	LVin  Set
	LVout Set
}

func main() {
	var blocks = []*block{}
	blockstr := map[string]int{}
	// parse input
	input := bufio.NewReader(os.Stdin)
	str, err := input.ReadString('\n')
	for err != nil {
		fmt.Fprintf(os.Stderr, "could not read line: %v\n", err.Error())
		os.Exit(1)
	}
	for err == nil {
		str = strings.ReplaceAll(str, "\r", "")
		str = strings.Trim(str, "\n")
		if str == "" {
			break
		}
		var strs = strings.Split(str, " ")
		if len(strs) != 3 {
			fmt.Fprintf(os.Stderr, "could not parse line \"%s\", requires 3 items\n", str)
			os.Exit(1)
		}
		name := strs[0]
		def := strs[1]
		use := strings.Split(strs[2], ",")
		ind := len(blocks)
		blockstr[name] = ind
		blocks = append(blocks, &block{
			Name:  name,
			Def:   def,
			Use:   use,
			LVin:  NewSet(),
			LVout: NewSet(),
		})
		str, err = input.ReadString('\n')
	}
	str, err = input.ReadString('\n')
	for err == nil {
		str = strings.ReplaceAll(str, "\r", "")
		str = strings.Trim(str, "\n")
		if str == "" {
			break
		}
		strs := strings.Split(str, " ")
		if len(strs) != 2 {
			fmt.Fprintf(os.Stderr, "could not parse line \"%s\", requires 2 items\n", str)
			os.Exit(1)
		}
		fi, ok := blockstr[strs[0]]
		if !ok {
			fmt.Fprintf(os.Stderr, "source node not recognized \"%s\"\n", strs[0])
			os.Exit(1)
		}
		fb := blocks[fi]
		tos := strings.Split(strs[1], ",")
		for _, to := range tos {
			ti, ok := blockstr[to]
			if !ok {
				fmt.Fprintf(os.Stderr, "destination node not recognized \"%s\"\n", to)
				os.Exit(1)
			}
			fb.Succs = append(fb.Succs, ti)
			tb := blocks[ti]
			tb.Preds = append(tb.Preds, fi)
		}
		str, err = input.ReadString('\n')
	}

	worklist := []*block{}
	for i := range blocks {
		worklist = append(worklist, blocks[len(blocks)-i-1])
	}
	// begin iterating
	for len(worklist) != 0 {
		out := NewSet()
		bl := worklist[0]
		worklist = worklist[1:]
		for _, s := range bl.Succs {
			out.Union(blocks[s].LVin)
		}
		in := out.Copy()
		in.Remove(bl.Def)
		for _, s := range bl.Use {
			in.Add(s)
		}
		in.Remove("")
		if !in.Equals(bl.LVin) {
			for _, p := range bl.Preds {
				contains := false
				for _, w := range worklist {
					if blocks[p] == w {
						contains = true
					}
				}
				if contains {
					worklist = append(worklist, blocks[p])
				}
			}
		}
		bl.LVin = in
		bl.LVout = out
	}
	// print results
	for _, b := range blocks {
		fmt.Printf("%s\n", b.Name)
		fmt.Printf("  LVin: %s\n", b.LVin.Print())
		fmt.Printf("  LVout: %s\n", b.LVout.Print())
	}
}
