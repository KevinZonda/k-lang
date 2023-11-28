# Patches

Some library need to be patched before it can be compiled.


| Patch Name | Description | Path | Source | Affect |
| --- | --- | --- | --- | --- |
| `0001-fix-missing-RestoreWindowsConsole-EnableWindowsANSIC.patch` | Original library does not compile on Windows. | Windows: `%PATH%\pkg\mod\github.com\tliron\kutil@v0.3.11\terminal`, macOS/Linux/Unix: No action needed` | [GitHub Commit](https://github.com/KevinZonda/kutil/commit/2fb99362a1d3c201f532600fb3fe760083a69034) [GitHub PR](https://github.com/tliron/kutil/pull/1/files) | go lib `github.com\tliron\kutil` <= v0.3.12 |