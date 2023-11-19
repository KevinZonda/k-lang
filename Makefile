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
	.\scripts\src.bat
else
	echo '=== Source Code Builder ==='
	./scripts/src.sh
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