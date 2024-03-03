// Adapted from https://cdnjs.cloudflare.com/ajax/libs/codemirror/6.65.7/mode/go/go.min.js
// Original Work is done by CodeMirror Authors
// Modified by Xiang Shi (xxs166@student.bham.ac.uk)

! function(e) {
    "object" == typeof exports && "object" == typeof module ? e(require("../../lib/codemirror")) : "function" == typeof define && define.amd ? define(["../../lib/codemirror"], e) : e(CodeMirror)
}(function(p) {
    "use strict";
    p.defineMode("k", function(e) {
        var i, o = e.indentUnit,
            r = {
                break: !0,
                case: !0,
                continue: !0,
                default: !0,
                else: !0,
                for: !0,
                fn: !0,
                if: !0,
                open: !0,
                as: !0,
                map: !0,
                return: !0,
                struct: !0,
                match: !0,
                bool: !0,
                num: !0,
                int: !0,
                string: !0,
                try: !0,
                catch : !0
            },
            a = {
                true: !0,
                false: !0,
                nil: !0,
                print: !0,
                println: !0,
                MEM: !0,
                len: !0,
                panic: !0,
                range: !0,
            },
            c = /[+\-*&^%:=<>!|\/]/;

        function u(e, t) {
            var o, n = e.next();
            if ('"' == n || "'" == n) return t.tokenize = (o = n, function(e, t) {
                for (var n, r = !1, i = !1; null != (n = e.next());) {
                    if (n == o && !r) {
                        i = !0;
                        break
                    }
                    r = !r && false && "\\" == n
                }
                return (i || !r && false) && (t.tokenize = u), "string"
            }), t.tokenize(e, t);
            if (/[\d\.]/.test(n)) return "." == n ? e.match(/^[0-9]+([eE][\-+]?[0-9]+)?/) : "0" == n ? e.match(/^[xX][0-9a-fA-F]+/) || e.match(/^0[0-7]+/) : e.match(/^[0-9]*\.?[0-9]*([eE][\-+]?[0-9]+)?/), "number";
            if (/[\[\]{}\(\),;\:\.]/.test(n)) return i = n, null;
            if ("#" == n) return e.skipToEnd(), "comment"
            if (c.test(n)) return e.eatWhile(c), "operator";
            e.eatWhile(/[\w\$_\xa1-\uffff]/);
            t = e.current();
            return r.propertyIsEnumerable(t) ? ("case" != t && "default" != t || (i = "case"), "keyword") : a.propertyIsEnumerable(t) ? "atom" : "variable"
        }

        function f(e, t, n, r, i) {
            this.indented = e, this.column = t, this.type = n, this.align = r, this.prev = i
        }

        function s(e, t, n) {
            e.context = new f(e.indented, t, n, null, e.context)
        }

        function d(e) {
            var t;
            e.context.prev && (")" != (t = e.context.type) && "]" != t && "}" != t || (e.indented = e.context.indented), e.context = e.context.prev)
        }
        return {
            startState: function(e) {
                return {
                    tokenize: null,
                    context: new f((e || 0) - o, 0, "top", !1),
                    indented: 0,
                    startOfLine: !0
                }
            },
            token: function(e, t) {
                var n = t.context;
                if (e.sol() && (null == n.align && (n.align = !1), t.indented = e.indentation(), t.startOfLine = !0, "case" == n.type && (n.type = "}")), e.eatSpace()) return null;
                i = null;
                var r = (t.tokenize || u)(e, t);
                return "comment" == r || (null == n.align && (n.align = !0), "{" == i ? s(t, e.column(), "}") : "[" == i ? s(t, e.column(), "]") : "(" == i ? s(t, e.column(), ")") : "case" == i ? n.type = "case" : ("}" == i && "}" == n.type || i == n.type) && d(t), t.startOfLine = !1), r
            },
            indent: function(e, t) {
                if (e.tokenize != u && null != e.tokenize) return p.Pass;
                var n = e.context,
                    r = t && t.charAt(0);
                if ("case" == n.type && /^(?:case|default)\b/.test(t)) return e.context.type = "}", n.indented;
                t = r == n.type;
                return n.align ? n.column + (t ? 0 : 1) : n.indented + (t ? 0 : o)
            },
            electricChars: "{}):",
            closeBrackets: "()[]{}''\"\"",
            fold: "brace",
            lineComment: "#"
        }
    }), p.defineMIME("text/x-k", "k")
});
