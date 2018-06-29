const lex = require('./lib/lexer.js');
module.exports.default = lex.default;
module.exports.lexer = lex.lexer;
module.exports.Collector = require('./lib/classes/collector.js').default;
module.exports.Token = require('./lib/classes/token.js').default;
