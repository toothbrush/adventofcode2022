TODAY := $(shell date +%-d)

default: day${TODAY}/input

day%/input: day%
	curl -fsSL --cookie session=$$SESSION https://adventofcode.com/2022/day/$*/input > $@

day%:
	mkdir -p $@
