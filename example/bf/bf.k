open "std/string" as strs

pos = 0
mem = [0]
instructions = []
loopLvl = 0

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
        # print("Running -> ", x, " ")
        # printMem()
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
                if (mem[pos] != 0) {
                    continue
                }
                loopLvl = loopLvl + 1
                for (instructions[i] != "]" || loopLvl != 0) {
                    i = i + 1
                    if (instructions[i] == "[") {
                        loopLvl = loopLvl + 1
                    }
                    if (instructions[i] == "]") {
                        loopLvl = loopLvl - 1
                    }
                }
            }
            case "]": {
                if (mem[pos] == 0) {
                    continue
                }
                loopLvl = loopLvl + 1
                for (instructions[i] != "[" || loopLvl != 0) {
                    i = i - 1
                    if (instructions[i] == "]") {
                        loopLvl = loopLvl + 1
                    }
                    if (instructions[i] == "[") {
                        loopLvl = loopLvl - 1
                    }
                }
                
            }
        }

    }

}


eval(">++++++++[<+++++++++>-]<.>++++[<+++++++>-]<+.+++++++..+++.>>++++++[<+++++++>-]<+
+.------------.>++++++[<+++++++++>-]<+.<.+++.------.--------.>>>++++[<++++++++>-
]<+.")

println()
printMem()