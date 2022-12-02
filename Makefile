TODAY := $(shell date +%-d)

default: day${TODAY}/input | day${TODAY}

day%/input: | day%
	curl -fsSL --cookie session=$$SESSION https://adventofcode.com/2022/day/$*/input > $@

day%:
	mkdir -p $@
