content = read(open(ARGS[1], "r"), String)
lines = split(content, "\n")

println("---  Day 8   ---")

rows = length(filter(l -> l != "", lines))
cols = length(first(lines))
grid = zeros(Int, rows, cols)

for (i, line) = zip(1 : rows, lines),
    (j, char) = zip(1 : cols,  line)

    grid[i, j] = parse(Int, char)
end

visibles = 0
highscore = 1

for i = 1 : rows, j = 1 : cols
    counted = false
    score = 1

    for (dx, dy) ∈ [(1, 0), (-1, 0), (0, 1), (0, -1)]
        visible = true
        x = i
        y = j

        while 1 < x < rows && 1 < y < cols
            x += dx; y += dy

            if grid[i, j] <= grid[x, y]
                visible = false
                break
            end
        end

        score *= abs(i - x) + abs(j - y)

        global highscore = score > highscore ? score : highscore

        if visible && !counted
            counted = true
            global visibles += 1
        end
    end
end

println(visibles)

println("--- Part Two ---")
println(highscore)
