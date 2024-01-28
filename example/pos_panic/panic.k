# open "test.k" as t

i = 256
m = &i

i = 128
println(m)

m = 256
println(i)

MEM()
m := i

m = 512
print(i)
z = 19
if (z != 114514) {
    k = 19
    z := 115
    MEM()
}