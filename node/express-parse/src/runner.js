import fs from 'fs';
import path from 'path';
import util from 'util';
import yargs from 'yargs';
import lexer from './lexer';

const argv = yargs
	.options({
		g: {
			alias: 'grammar',
			demandOption: true,
			describe: 'Directory containing grammar definitions',
			normalize: true,
			type: 'string'
		},
		f: {
			alias: ['in-file'],
			demandOption: true,
			describe: 'Source file',
			normalize: true,
			type: 'string'
		},
		s: {
			alias: 'start-collector',
			demandOption: false,
			describe: 'Starting token',
			default: 'program',
			type: 'string'
		}
	}).argv;

const startPath = path.join(path.resolve(argv.g), `${argv.s}.js`);
if (!fs.existsSync(startPath)) {
	console.error(`Cannot locate start point for lexer: ${startPath}`);
	process.exit(1);
}

const inputFile = path.resolve(argv.f);
if (!fs.existsSync(inputFile)) {
	console.error(`Cannot locate source file: ${inputFile}`);
	process.exit(1);
}

const START = require(startPath).default;
const INPUT_FILE = fs.readFileSync(inputFile).toString('utf8');
const tree = lexer(INPUT_FILE, 0, START);

console.error(
	util.inspect(tree, { depth: null, colors: true })
);
