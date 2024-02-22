cd src
go mod tidy

$hash=git log --pretty=format:'%h' -n 1
$date = (Get-Date).ToString("dd/MM/yyyy-HH:mm:ss")
$branch = git branch --show-current

$flag = "-X git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/buildconst.CommitHash=$hash -X git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/buildconst.BuildTime=$date"
$flag = "$flag -X git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/buildconst.BuildBranch=$branch"

$flag = "$flag -s -w"
$flag = "-ldflags=`"$flag`""

go build -v $flag -o out/minimal.exe main/minimal/main.go
cd ..