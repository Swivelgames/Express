import fs from 'fs';
import path from 'path';
import util from 'util';
import MathStmt from './grammar/mathstmt';
import lexer from './lexer';

const INPUT_FILE = 'example.expr';

const program = [];
const pointer = [];
const starting = MathStmt;

const src = fs.readFileSync(
	path.join(__dirname, INPUT_FILE)
).toString('utf8');

console.error('src', src);

// process.exit(127);

const tree = lexer(src, 0, MathStmt);

console.error(
	util.inspect(tree, { depth: null, colors: true })
);
