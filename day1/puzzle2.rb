#!/usr/bin/env ruby

calories_this_elf = 0
calories_max = [0]

$stdin.each_line do |l|
  this_line = l.strip

  if this_line == ""
    if calories_max.min < calories_this_elf
      calories_max.push(calories_this_elf)
      if calories_max.length > 3
        calories_max = calories_max.sort.reverse.take(3)
      end
    end

    calories_this_elf = 0
  else
    calories_this_elf += Integer(this_line)
  end
end

# And a fixup for if the last elf is a max-elf.
if calories_max.min < calories_this_elf
  calories_max.push(calories_this_elf)
  if calories_max.length > 3
    calories_max = calories_max.sort.reverse.take(3)
  end
end

puts "------------------"
pp calories_max
puts calories_max.sum
