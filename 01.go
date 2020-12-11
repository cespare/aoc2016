package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func init() {
	addSolutions(1, problem1)
}

func problem1(ctx *problemContext) {
	b, err := ioutil.ReadAll(ctx.f)
	if err != nil {
		log.Fatal(err)
	}
	var dirs []cityDirections
	for _, field := range bytes.Fields(b) {
		s := strings.TrimRight(string(field), ",")
		var dir cityDirections
		switch s[0] {
		case 'L':
		case 'R':
			dir.right = true
		default:
			log.Fatal("bad dir")
		}
		d, err := strconv.ParseInt(s[1:], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		dir.d = d
		dirs = append(dirs, dir)
	}
	ctx.reportLoad()

	x, y := evalCityDirections(dirs)
	ctx.reportPart1(abs(x) + abs(y))

	x, y = evalCityDirectionsTwice(dirs)
	ctx.reportPart2(abs(x) + abs(y))
}

type cityDirections struct {
	right bool
	d     int64
}

var dirVecs = [][2]int64{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func evalCityDirections(dirs []cityDirections) (x, y int64) {
	var idir int
	for _, dir := range dirs {
		if dir.right {
			idir = (idir + 1) % 4
		} else {
			idir = (idir + 3) % 4
		}
		vec := dirVecs[idir]
		x += vec[0] * dir.d
		y += vec[1] * dir.d
	}
	return x, y
}

func evalCityDirectionsTwice(dirs []cityDirections) (x, y int64) {
	m := map[[2]int64]struct{}{
		[2]int64{0, 0}: {},
	}
	var idir int
	for _, dir := range dirs {
		if dir.right {
			idir = (idir + 1) % 4
		} else {
			idir = (idir + 3) % 4
		}
		vec := dirVecs[idir]
		for i := int64(0); i < dir.d; i++ {
			x += vec[0]
			y += vec[1]
			v := [2]int64{x, y}
			if _, ok := m[v]; ok {
				return x, y
			}
			m[v] = struct{}{}
		}
	}
	panic("unreached")
}
