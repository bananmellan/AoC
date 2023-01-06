content = read(open(ARGS[1], "r"), String)
lines = split(content, "\n")

println("Part 1:")

cd = [""]
dirs = Dict{Vector{String}, Int}()
doneDir = false

for line ∈ lines
    parseline = match(r"\$ (?<cmd>[\S\W]+)", line)

    if parseline != nothing
        cmd = parseline["cmd"]
        args = split(cmd)

        if args[1] == "cd"
            if args[2] == "/"
                global cd = [""]
            elseif args[2] == ".."
                global cd = cd[1 : length(cd) - 1]
            else
                global cd = [cd ; args[2]]
            end

            if haskey(dirs, cd)
                global doneDir = true
            else
                global doneDir = false
                dirs[cd] = 0
            end
        end
    elseif !doneDir
        file = match(r"(?<size>\d+) (\S+)", line)
        dir  = match(r"dir (?<dir>\S+)", line)

        if file != nothing
            # dsize = get!(dirs, cd, 0)
            dirs[cd] += parse(Int, file["size"])
        end
    end
end

ndirs = Dict{Vector{String}, Int}()

for (dir, dsize) ∈ dirs
    get!(ndirs, dir, 0)

    for i ∈ 1 : length(dir) - 1
        path = dir[1 : i]
        psize = get!(ndirs, path, 0)
        ndirs[path] = psize + dirs[dir]
    end

    ndirs[dir] += dirs[dir]
end

println(sum(f -> f.second, filter(f -> f.second <= 100000, ndirs)))

println("\nPart 2:")

spaceNeeded = ndirs[[""]] - 40000000
smallest = [""]

for (dir, dsize) ∈ ndirs
    if dsize >= spaceNeeded && dsize < ndirs[smallest]
        global smallest = dir
    end
end

println(ndirs[smallest])
