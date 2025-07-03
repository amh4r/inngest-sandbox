"use strict";
/*
 * ATTENTION: An "eval-source-map" devtool has been used.
 * This devtool is neither made for production nor for readable output files.
 * It uses "eval()" calls to create a separate source file with attached SourceMaps in the browser devtools.
 * If you are trying to read the output file, select a different devtool (https://webpack.js.org/configuration/devtool/)
 * or disable the default devtool with "devtool: false".
 * If you are looking for production-ready output files, see mode: "production" (https://webpack.js.org/configuration/mode/).
 */
exports.id = "vendor-chunks/ansi-regex";
exports.ids = ["vendor-chunks/ansi-regex"];
exports.modules = {

/***/ "(rsc)/./node_modules/ansi-regex/index.js":
/*!******************************************!*\
  !*** ./node_modules/ansi-regex/index.js ***!
  \******************************************/
/***/ ((module) => {

eval("\nmodule.exports = (options)=>{\n    options = Object.assign({\n        onlyFirst: false\n    }, options);\n    const pattern = [\n        '[\\\\u001B\\\\u009B][[\\\\]()#;?]*(?:(?:(?:(?:;[-a-zA-Z\\\\d\\\\/#&.:=?%@~_]+)*|[a-zA-Z\\\\d]+(?:;[-a-zA-Z\\\\d\\\\/#&.:=?%@~_]*)*)?\\\\u0007)',\n        '(?:(?:\\\\d{1,4}(?:;\\\\d{0,4})*)?[\\\\dA-PR-TZcf-ntqry=><~]))'\n    ].join('|');\n    return new RegExp(pattern, options.onlyFirst ? undefined : 'g');\n};\n//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiKHJzYykvLi9ub2RlX21vZHVsZXMvYW5zaS1yZWdleC9pbmRleC5qcyIsIm1hcHBpbmdzIjoiQUFBYTtBQUViQSxPQUFPQyxPQUFPLEdBQUdDLENBQUFBO0lBQ2hCQSxVQUFVQyxPQUFPQyxNQUFNLENBQUM7UUFDdkJDLFdBQVc7SUFDWixHQUFHSDtJQUVILE1BQU1JLFVBQVU7UUFDZjtRQUNBO0tBQ0EsQ0FBQ0MsSUFBSSxDQUFDO0lBRVAsT0FBTyxJQUFJQyxPQUFPRixTQUFTSixRQUFRRyxTQUFTLEdBQUdJLFlBQVk7QUFDNUQiLCJzb3VyY2VzIjpbIi9Vc2Vycy9haGFycGVyL2lubmdlc3QvaW5uZ2VzdC1zYW5kYm94L3RzLW5leHQvbm9kZV9tb2R1bGVzL2Fuc2ktcmVnZXgvaW5kZXguanMiXSwic291cmNlc0NvbnRlbnQiOlsiJ3VzZSBzdHJpY3QnO1xuXG5tb2R1bGUuZXhwb3J0cyA9IG9wdGlvbnMgPT4ge1xuXHRvcHRpb25zID0gT2JqZWN0LmFzc2lnbih7XG5cdFx0b25seUZpcnN0OiBmYWxzZVxuXHR9LCBvcHRpb25zKTtcblxuXHRjb25zdCBwYXR0ZXJuID0gW1xuXHRcdCdbXFxcXHUwMDFCXFxcXHUwMDlCXVtbXFxcXF0oKSM7P10qKD86KD86KD86KD86O1stYS16QS1aXFxcXGRcXFxcLyMmLjo9PyVAfl9dKykqfFthLXpBLVpcXFxcZF0rKD86O1stYS16QS1aXFxcXGRcXFxcLyMmLjo9PyVAfl9dKikqKT9cXFxcdTAwMDcpJyxcblx0XHQnKD86KD86XFxcXGR7MSw0fSg/OjtcXFxcZHswLDR9KSopP1tcXFxcZEEtUFItVFpjZi1udHFyeT0+PH5dKSknXG5cdF0uam9pbignfCcpO1xuXG5cdHJldHVybiBuZXcgUmVnRXhwKHBhdHRlcm4sIG9wdGlvbnMub25seUZpcnN0ID8gdW5kZWZpbmVkIDogJ2cnKTtcbn07XG4iXSwibmFtZXMiOlsibW9kdWxlIiwiZXhwb3J0cyIsIm9wdGlvbnMiLCJPYmplY3QiLCJhc3NpZ24iLCJvbmx5Rmlyc3QiLCJwYXR0ZXJuIiwiam9pbiIsIlJlZ0V4cCIsInVuZGVmaW5lZCJdLCJpZ25vcmVMaXN0IjpbMF0sInNvdXJjZVJvb3QiOiIifQ==\n//# sourceURL=webpack-internal:///(rsc)/./node_modules/ansi-regex/index.js\n");

/***/ })

};
;