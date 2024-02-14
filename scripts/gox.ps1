cd src
go mod tidy
$flag = "-X git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/buildconst.CommitHash=$hash -X git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/buildconst.BuildTime=$date"
$flag = "$flag -s -w"
$flag = "-ldflags=`"$flag`""

go build -v $flag -o out/interpreter.exe main/cli/main.go
cd ..