# String Design

String may have multiple way to process. This is a design for string process.

## Parser Mode

1. Standard: `"Hello World {x}\n"` -> `Hello World {x}` + NewLine
2. Constant: `@"Hello World {x}\n"` -> `Hello World {x}\n`
3. Variable: `$"Hello World {x}\n"` -> `Hello World ` + var x + NewLine

## Escape Character

| Escape Character | Meaning                    |
|------------------|----------------------------|
| `\a`             | Alert                      |
| `\n`             | NewLine                    |
| `\t`             | Tab                        |
| `\r`             | Return                     |
| `\\`             | Slash                      |
| `\"`             | Quote                      |
| `{{`             | LeftBrace (Variable Mode)  |
| `}}`             | RightBrace (Variable Mode) |