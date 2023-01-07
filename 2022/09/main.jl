using Printf

content = read(open(ARGS[1], "r"), String)
lines = filter(l -> l != "", split(content, "\n"))

println("---  Day 9   ---")

dirs = Dict("U" => [0, 1], "D" => [0, -1],
            "R" => [1, 0], "L" => [-1, 0])

struct Knot
    value::Vector{Int}
    next::Union{Knot, Nothing}
end

function f(n)
    s = [0, 0]
    posh = s
    post = s
    visited = Set{Vector{Int}}();

    for line ∈ lines
        dir = split(line)[1]
        steps = parse(Int, split(line)[2])

        for i ∈ 1:steps
            next = posh + dirs[dir]

            vec = next - post
            if abs(vec[1]) > 1 || abs(vec[2]) > 1
                post = posh
            end

            posh = next

            push!(visited, post)
        end
    end

    return visited
end

println(length(f(1)))
