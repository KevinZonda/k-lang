cd src
go mod tidy

$hash=git log --pretty=format:'%h' -n 1
$date = (Get-Date).ToString("dd/MM/yyyy")
$time = (Get-Date).ToString("HH:mm:ss")
$branch = git branch --show-current

$flag = "-X git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/buildconst.CommitHash=$hash"
$flag = "$flag -X git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/buildconst.BuildDate=$date"
$flag = "$flag -X git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/buildconst.BuildTime=$time"
$flag = "$flag -X git.cs.bham.ac.uk/projects-2023-24/xxs166/src/main/buildconst.BuildBranch=$branch"
$flag = "-ldflags=`"$flag`""

echo $flag
$Env:GOOS = "js"; $Env:GOARCH = "wasm"
go build -v $flag -o ../playground/core.wasm main/wasm/main_wasm.go
cd ..