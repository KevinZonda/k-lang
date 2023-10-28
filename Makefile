grammar:
ifeq ($(OS),Windows_NT)
	@echo off
	echo "Windows Grammar Builder"
	.\grammar.bat
else
	echo 'Linux/Darwin/Unix-like Grammar Builder'
	./grammar.sh
endif
