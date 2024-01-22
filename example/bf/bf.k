open "std/string" as strs

pos = 0
mem = [0]
loop = []
loopPos = -1

fn printMem() {
    cur = 0
    print("MEM -> ")
    for (b : mem) {
        if (cur == pos) {
            print("(", b, ")")
        } else {
            print(b)
        }
        print(" ")
        cur = cur + 1
    }
    println()
}

fn eval(code) {
    instructions = strs.split(code, "")
    println(instructions)
    for (i = 0; i < len(instructions); i = i + 1) {
        x = instructions[i]
        print("Running -> ", x, " ")
        printMem()
        match(x) {
            case ">": {
                pos = pos + 1
                if (pos >= len(mem)) {
                    mem = mem + 0
                }
            }
            case "<": {
                pos = pos - 1
            }
            case "+": {
                mem[pos] = mem[pos] + 1
            }
            case "-": {
                mem[pos] = mem[pos] - 1
            }
            case ".": {
                print(strs.fromAsci(mem[pos]))
            }
            case ",": {
                # TODO: INPUT
            }
            case "[": {
                if (mem[pos] > 0) {
                    end = findEndPair(i)
                    print(instructions[end])
                    loopPos = loopPos + 1
                    if (loopPos >= len(loop)) { loop = loop + 0}
                    loop[loopPos] = pos
                } else {
                    findEndPair(i)
                    # find it's pair ]
                }

            }
            case "]": {

            }
        }

    }

}

fn findEndPair(pos) {
    # FIXME: baseline: frame protection
    lvl = 1
    pos = pos + 1
    
    for (i = pos; i < len(instructions); i = i + 1) {
        cur = instructions[i]
        if (cur == "[") {
            lvl = lvl + 1
        } 
        
        if (cur == "]") {
            lvl = lvl - 1
            if (lvl == 0) {
                return i
            }
        }
    }
    return -1
}

eval("+++++++++++++++++++++++++++++[++++++++++++++++]++++++++++++++++++++.")
println()
printMem()