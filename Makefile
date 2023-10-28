grammar:
ifeq ($(OS),Windows_NT)
	@echo off
	echo "Windows Grammar Builder"
	.\scripts\grammar.bat
else
	echo 'Linux/Darwin/Unix-like Grammar Builder'
	./scripts/grammar.sh
endif

stat:
ifeq ($(OS),Windows_NT)
	@echo off
	echo "Not supported"
else
	./scripts/stat.sh
endif
