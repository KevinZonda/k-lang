.DEFAULT_GOAL := go

grammar:
ifeq ($(OS),Windows_NT)
	@echo '=== Windows Grammar Builder ==='
	.\scripts\grammar.bat
else
	echo '=== Linux/Darwin/Unix-like Grammar Builder ==='
	./scripts/grammar.sh
endif

stat:
ifeq ($(OS),Windows_NT)
	@echo 'Not supported!'
else
	./scripts/stat.sh
endif

go:
ifeq ($(OS),Windows_NT)
	@echo '=== Source Code Builder ==='
	powershell .\scripts\src.ps1
else
	echo '=== Source Code Builder ==='
	./scripts/src.sh
endif

gui:
ifeq ($(OS),Windows_NT)
	@echo '=== Source Code Builder ==='
	powershell .\scripts\gui.ps1
else
	echo '=== Source Code Builder ==='
	./scripts/gui.sh
endif

gox:
ifeq ($(OS),Windows_NT)
	@echo '=== Source Code Builder X ==='
	powershell .\scripts\gox.ps1
else
	echo '=== Source Code Builder X ==='
	./scripts/gox.sh
endif

wasm:
ifeq ($(OS),Windows_NT)
	@echo '=== Source Code Builder X ==='
	powershell .\scripts\wasm.ps1
else
	echo '=== Source Code Builder X ==='
	./scripts/wasm.sh
endif

tidy:
ifeq ($(OS),Windows_NT)
	@echo '=== Tidy ==='
	.\scripts\tidy.bat
else
	echo '=== Tidy ==='
	./scripts/tidy.sh
endif

vsix:
ifeq ($(OS),Windows_NT)
	@echo '=== VSIX Builder ==='
	.\scripts\vsix.bat
else
	echo '=== VSIX Builder ==='
	./scripts/vsix.sh
endif

all:
	make grammar
	make go

out:
ifeq ($(OS),Windows_NT)
	explorer .\src\out
else
	open .\src\out
endif

test:
ifeq ($(OS),Windows_NT)
	@echo '=== Test ==='
	.\scripts\test.bat
else
	echo '=== Test ==='
	./scripts/test.sh
endif