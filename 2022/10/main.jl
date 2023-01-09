using Printf

println("---  Day 9   ---")

content = read(open(ARGS[1], "r"), String)
lines = filter(l -> l != "", split(content, "\n"))

X = 1
cycle = 0
strengths = 0

for line in lines
    cmd = match(r"\S+ ?(?<num>-?\d+)?", line)

    if (cycle - 20) % 40 == 0
        @printf "Cycle:%d, X:%d, S: %d\n" cycle X (X * cycle)
        global strengths += X * cycle;
    end

    if cmd["num"] != nothing
        global cycle += 1

        if (cycle - 20) % 40 == 0
            @printf "Cycle:%d, X:%d, S: %d\n" cycle X (X * cycle)
            global strengths += X * cycle;
        end

        global X += parse(Int, cmd["num"])
    end

    global cycle += 1

end

println(strengths)

# println("--- Part Two ---")
