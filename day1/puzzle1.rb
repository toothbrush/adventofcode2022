#!/usr/bin/env ruby

def max(a,b)
  a>b ? a : b
end

calories_this_elf = 0
calories_max = 0

$stdin.read.lines.each do |l|
  this_line = l.strip
  if this_line == ""
    calories_max = max(calories_max, calories_this_elf)
    calories_this_elf = 0
  else
    calories_this_elf += Integer(this_line)
  end
end

puts calories_max
