using Printf

println("---  Day 9   ---")

content = read(open(ARGS[1], "r"), String)
lines = filter(l -> l != "", split(content, "\n"))

dirs = Dict("U" => [0, 1], "D" => [0, -1],
            "R" => [1, 0], "L" => [-1, 0])

mutable struct Knot
    pos::Vector{Int}
    prev::Vector{Int}
    next::Union{Knot, Nothing}
end

function f(n)
    s = [0, 0]
    rope = Knot(s, s, nothing)
    head = rope
    for i ∈ 1 : n
        head.next = Knot(s, s, nothing)
        head = head.next
    end

    visited = Set{Vector{Int}}();

    for line ∈ lines
        dir = split(line)[1]
        steps = parse(Int, split(line)[2])

        tail = nothing

        for i ∈ 1 : steps
            rope.prev = rope.pos
            rope.pos += dirs[dir]
            head = rope

            tail = head

            while head.next != nothing
                tail = head.next

                delta = head.pos - head.prev
                diff  = head.pos - tail.pos

                if abs(diff[1]) > 1 || abs(diff[2]) > 1
                    tail.prev = tail.pos

                    if head.prev[1] == tail.pos[1] ||
                       head.prev[2] == tail.pos[2]
                        tail.pos += delta
                    else
                        tail.pos += [sign(diff[1]) * min(abs(diff[1]), 1),
                                     sign(diff[2]) * min(abs(diff[2]), 1)]
                    end
                end

                head = tail
            end

            push!(visited, tail.pos)
        end
    end

    return visited
end

println(length(f(1)))

println("--- Part Two ---")
println(length(f(9)))
