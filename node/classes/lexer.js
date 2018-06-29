export default class Lexer {
	constructor(src, i, collector) {
		this.i = i;
		this.src = src;
		this.collector = collector;
	}
}
