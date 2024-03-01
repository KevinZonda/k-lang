/*
Language: Klang
Modifier: Xiang Shi <xxs166@student.bham.ac.uk>
Author: Stephan Kountso aka StepLg <steplg@gmail.com>
Contributors: Evgeny Stepanischev <imbolk@gmail.com>
Description: K Lang Highligh. Derived from go.js. Google go language (golang). For info about language
Website: http://golang.org/
Category: common, system
*/

export function klang(hljs: any) {
    const K_KEYWORDS = {
        keyword:
            'break default match case map struct else open if ' +
            'continue for break return ' +
            'bool num int string string',
        literal:
            'true false nil',
        built_in:
            'panic print println typeOf MEM'
    };
    return {
        name: 'K Language',
        aliases: ['klang'],
        keywords: K_KEYWORDS,
        illegal: '</',
        contains: [
            hljs.HASH_COMMENT_MODE,
            {
                className: 'string',
                variants: [
                    {
                        begin: '\'',
                        end: '\'',
                        contains: [
                            hljs.BACKSLASH_ESCAPE,
                        ]
                    },
                    {
                        begin: '"',
                        end: '"',
                        contains: [
                            hljs.BACKSLASH_ESCAPE,
                        ]
                    }
                ]
            },
            {
                className: 'number',
                variants: [
                    {
                        begin: hljs.C_NUMBER_RE + '[i]',
                        relevance: 1
                    },
                    hljs.C_NUMBER_MODE
                ]
            },
            {
                begin: /:=/ // relevance booster
            },
            {
                className: 'function',
                beginKeywords: 'fn',
                end: '\\s*(\\{|$)',
                excludeEnd: true,
                contains: [
                    hljs.TITLE_MODE,
                    {
                        className: 'params',
                        begin: /\(/,
                        end: /\)/,
                        keywords: K_KEYWORDS,
                        illegal: /["']/
                    }
                ]
            }
        ]
    };
}

export default klang;
