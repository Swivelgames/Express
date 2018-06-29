const doesTokenMatchType = (token, expected) => typeof expected === 'string' ? expected === token.type : expected.type === token.type;

export default class Collector {
	constructor(cfg) {
		const { type, expected } = this.constructor;
		if (typeof type !== 'string') {
			throw new TypeError(`${this.constructor.name} Collector: 'type' must be specified as a string. Invalid value: ${type}`);
		}

		if (!Array.isArray(expected)) {
			throw new TypeError(`${this.constructor.name} Collector: 'expected' must be an array. Invalid value: ${expected}`);
		}

		if (expected.length < 1) {
			throw new RangeError(`${this.constructor.name} Collector: 'expected' must be an array of strings, token types, or collector types. It cannot be empty.`);
		}

		this.type = this.constructor.type.toUpperCase();
		this.expected = this.constructor.expected.slice();
		this.collected = [];
		this.path = -1;
	}

	accepts(token, path = this.path) {
		if (path === -1) {
			return this.expected.find((p, i) => this.accepts(token, i));
		}
		return !!this.expected[path].find(t => doesTokenMatchType(token, t));
	}

	collect(token) {
		const nextToken = this.nextExpected();
		const isExpected = doesTokenMatchType(token, nextToken);
		if (!isExpected) {
			throw new TypeError(`Unexpected ${token.type} token in ${this.type} collector. Expected ${nextToken}.`);
		}

		this.collected.push(token);

		return true;
	}

	nextExpected() {
		if (this.isSatisfied()) return [];
		const { path, collected } = this;
		if (path === -1) {
			this.path = 0;
			return this.expected[0][0];
		}
		return this.expected[path][collected.length];
	}

	tryAltPath() {
		const { path, expected, collected } = this;
		if (path === -1) return false;
		if (path === expected.length - 1) return false;
		for (var i=path+1;i<expected.length;i++) {
			const goodPath = collected.find((t,ci) => doesTokenMatchType(t, expected[i][ci]));
			if (goodPath) {
				this.path = i;
				return true;
			}
		}
		return false;
	}

	isSatisfied() {
		const { path } = this;
		if (path === -1) return false;

		const { expected, collected } = this;
		const curPath = expected[path];
		if (curPath.length !== collected.length) return false;

		return true;
	}
}
