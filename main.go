package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/valyala/fasthttp"
)

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	now := time.Now()
	defer log.Println(time.Since(now))
	var v ArenaUpdate
	req := ctx.PostBody()
	json.Unmarshal(req, &v)
	resp := play(&v)
	fmt.Fprint(ctx, resp)
	res, _ := json.Marshal(v)
	log.Println(string(res))
	log.Println(resp)
}

func main() {
	fasthttp.ListenAndServe(":8080", fastHTTPHandler)
}

func fire(me, en *PlayerState) (resp string) {
	if me.X == en.X {
		if me.Y > en.Y {
			switch me.Direction {
			case N:
				return "T"
			case S, E:
				return "L"
			case W:
				return "R"
			}
		} else {
			switch me.Direction {
			case S:
				return "T"
			case N, E:
				return "R"
			case W:
				return "L"
			}
		}
	} else {
		if me.X > en.X {
			switch me.Direction {
			case W:
				return "T"
			case S, E:
				return "R"
			case N:
				return "L"
			}
		} else {
			switch me.Direction {
			case E:
				return "T"
			case N, W:
				return "R"
			case S:
				return "L"
			}
		}
	}
	return "T"
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func move(me, en *PlayerState) (resp string) {
	if abs(me.X-en.X) < abs(me.Y-en.Y) {
		if me.X > en.X {
			switch me.Direction {
			case W:
				return "F"
			case S, E:
				return "R"
			case N:
				return "L"
			}
		} else {
			switch me.Direction {
			case E:
				return "F"
			case N, W:
				return "R"
			case S:
				return "L"
			}
		}
	} else {
		if me.Y > en.Y {
			switch me.Direction {
			case N:
				return "F"
			case S, E:
				return "L"
			case W:
				return "R"
			}
		} else {
			switch me.Direction {
			case S:
				return "F"
			case N, E:
				return "R"
			case W:
				return "L"
			}
		}
	}
	return "F"
}
func distance(me, en *PlayerState) (dist int) {
	sum := abs(me.X-en.X) + abs(me.Y-en.Y)
	if me.X == en.X {
		sum--
	}
	if me.Y == en.Y {
		sum--
	}
	if me.X > en.X {
		switch me.Direction {
		case S, N:
			sum++
		case E:
			sum += 2
		}
	} else if me.X < en.X {
		switch me.Direction {
		case S, N:
			sum++
		case W:
			sum += 2

		}
	}
	if me.Y > en.Y {
		switch me.Direction {
		case E, W:
			sum++
		case S:
			sum += 2
		}
	} else {
		switch me.Direction {
		case E, W:
			sum++
		case N:
			sum += 2
		}
	}
	return sum
}
func canFire(me, en *PlayerState) bool {
	if me.X != en.X && me.Y != en.Y {
		return false
	}
	if me.X == en.X {
		return abs(me.Y-en.Y) <= 3
	}
	return abs(me.X-en.X) <= 3
}

var lastName string = ""

func play(input *ArenaUpdate) string {
	meName := input.Links.Self.Href

	arena := &input.Arena
	me := arena.State[meName]
	last := arena.State[lastName]

	if lastName != "" && canFire(&me, &last) {
		return fire(&me, &last)
	}

	var nemesis PlayerState
	dist := 99999
	for i, k := range arena.State {
		if i == meName {
			continue
		}
		d := distance(&me, &k)

		if d < dist {
			dist = d
			nemesis = k
			lastName = i
		}
	}
	if canFire(&me, &nemesis) {

		resp := fire(&me, &nemesis)

		return resp
	}
	resp := move(&me, &nemesis)

	return resp
}
