"use strict";
/*
 * ATTENTION: An "eval-source-map" devtool has been used.
 * This devtool is neither made for production nor for readable output files.
 * It uses "eval()" calls to create a separate source file with attached SourceMaps in the browser devtools.
 * If you are trying to read the output file, select a different devtool (https://webpack.js.org/configuration/devtool/)
 * or disable the default devtool with "devtool: false".
 * If you are looking for production-ready output files, see mode: "production" (https://webpack.js.org/configuration/mode/).
 */
exports.id = "vendor-chunks/canonicalize";
exports.ids = ["vendor-chunks/canonicalize"];
exports.modules = {

/***/ "(rsc)/./node_modules/canonicalize/lib/canonicalize.js":
/*!*******************************************************!*\
  !*** ./node_modules/canonicalize/lib/canonicalize.js ***!
  \*******************************************************/
/***/ ((module) => {

eval("/* jshint esversion: 6 */\n/* jslint node: true */\n\n\nmodule.exports = function serialize (object) {\n  if (object === null || typeof object !== 'object' || object.toJSON != null) {\n    return JSON.stringify(object);\n  }\n\n  if (Array.isArray(object)) {\n    return '[' + object.reduce((t, cv, ci) => {\n      const comma = ci === 0 ? '' : ',';\n      const value = cv === undefined || typeof cv === 'symbol' ? null : cv;\n      return t + comma + serialize(value);\n    }, '') + ']';\n  }\n\n  return '{' + Object.keys(object).sort().reduce((t, cv, ci) => {\n    if (object[cv] === undefined ||\n        typeof object[cv] === 'symbol') {\n      return t;\n    }\n    const comma = t.length === 0 ? '' : ',';\n    return t + comma + serialize(cv) + ':' + serialize(object[cv]);\n  }, '') + '}';\n};\n//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiKHJzYykvLi9ub2RlX21vZHVsZXMvY2Fub25pY2FsaXplL2xpYi9jYW5vbmljYWxpemUuanMiLCJtYXBwaW5ncyI6IkFBQUE7QUFDQTtBQUNhOztBQUViO0FBQ0E7QUFDQTtBQUNBOztBQUVBO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQSxLQUFLO0FBQ0w7O0FBRUEsV0FBVztBQUNYO0FBQ0E7QUFDQTtBQUNBO0FBQ0E7QUFDQTtBQUNBLEdBQUcsVUFBVTtBQUNiIiwic291cmNlcyI6WyIvVXNlcnMvYWhhcnBlci9pbm5nZXN0L2lubmdlc3Qtc2FuZGJveC90cy1uZXh0L25vZGVfbW9kdWxlcy9jYW5vbmljYWxpemUvbGliL2Nhbm9uaWNhbGl6ZS5qcyJdLCJzb3VyY2VzQ29udGVudCI6WyIvKiBqc2hpbnQgZXN2ZXJzaW9uOiA2ICovXG4vKiBqc2xpbnQgbm9kZTogdHJ1ZSAqL1xuJ3VzZSBzdHJpY3QnO1xuXG5tb2R1bGUuZXhwb3J0cyA9IGZ1bmN0aW9uIHNlcmlhbGl6ZSAob2JqZWN0KSB7XG4gIGlmIChvYmplY3QgPT09IG51bGwgfHwgdHlwZW9mIG9iamVjdCAhPT0gJ29iamVjdCcgfHwgb2JqZWN0LnRvSlNPTiAhPSBudWxsKSB7XG4gICAgcmV0dXJuIEpTT04uc3RyaW5naWZ5KG9iamVjdCk7XG4gIH1cblxuICBpZiAoQXJyYXkuaXNBcnJheShvYmplY3QpKSB7XG4gICAgcmV0dXJuICdbJyArIG9iamVjdC5yZWR1Y2UoKHQsIGN2LCBjaSkgPT4ge1xuICAgICAgY29uc3QgY29tbWEgPSBjaSA9PT0gMCA/ICcnIDogJywnO1xuICAgICAgY29uc3QgdmFsdWUgPSBjdiA9PT0gdW5kZWZpbmVkIHx8IHR5cGVvZiBjdiA9PT0gJ3N5bWJvbCcgPyBudWxsIDogY3Y7XG4gICAgICByZXR1cm4gdCArIGNvbW1hICsgc2VyaWFsaXplKHZhbHVlKTtcbiAgICB9LCAnJykgKyAnXSc7XG4gIH1cblxuICByZXR1cm4gJ3snICsgT2JqZWN0LmtleXMob2JqZWN0KS5zb3J0KCkucmVkdWNlKCh0LCBjdiwgY2kpID0+IHtcbiAgICBpZiAob2JqZWN0W2N2XSA9PT0gdW5kZWZpbmVkIHx8XG4gICAgICAgIHR5cGVvZiBvYmplY3RbY3ZdID09PSAnc3ltYm9sJykge1xuICAgICAgcmV0dXJuIHQ7XG4gICAgfVxuICAgIGNvbnN0IGNvbW1hID0gdC5sZW5ndGggPT09IDAgPyAnJyA6ICcsJztcbiAgICByZXR1cm4gdCArIGNvbW1hICsgc2VyaWFsaXplKGN2KSArICc6JyArIHNlcmlhbGl6ZShvYmplY3RbY3ZdKTtcbiAgfSwgJycpICsgJ30nO1xufTtcbiJdLCJuYW1lcyI6W10sImlnbm9yZUxpc3QiOlswXSwic291cmNlUm9vdCI6IiJ9\n//# sourceURL=webpack-internal:///(rsc)/./node_modules/canonicalize/lib/canonicalize.js\n");

/***/ })

};
;