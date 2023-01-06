content = read(open("input", "r"))

function f(n)
    for i in 1 : length(content) - n - 1
        pos = i + n - 1
        four = collect(content[i : pos])
        if unique(four) == four; return pos; end
    end
end

println("Part 1:")
println(f(4))

println("\nPart 2:")
println(f(14))
