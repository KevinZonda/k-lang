open "std/string" as strs

pos = 0
mem = [0]

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
    for (x : instructions) {
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

            }
            case "]": {

            }
        }

    }

}

eval("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++.")
println()
println("Register -> ", mem)
printMem()