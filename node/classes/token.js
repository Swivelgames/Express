// @TODO: Add flag to consume whitespace

export default class Token {
	constructor(cfg) {
		const { type, expected } = this.constructor;
		if (typeof type !== 'string') {
			throw new TypeError(`${this.constructor.name} Token: 'type' must be specified as a string. Invalid value: ${type}`);
		}

		if (!Array.isArray(expected)) {
			throw new TypeError(`${this.constructor.name} Token: 'expected' must be an array. Invalid value: ${expected}`);
		}

		if (expected.length < 1) {
			throw new RangeError(`${this.constructor.name} Token: 'expected' must be an array of strings or regular expresses. It cannot be empty.`);
		}

		this.type = this.constructor.type.toUpperCase();
		this.expected = this.constructor.expected.slice();
		this.collected = [];
	}

	accepts(char) {
		const comp = `${this.collected.join('')}${char}`;
		const { expected } = this;
		const series = [].concat(expected);
		return series.reduce((b, e) => {
			console.error(e, char, comp, b);
			if (b !== false) return true;
			if (e instanceof RegExp) return e.test(comp);
			return e.indexOf(comp) === 0 ? true : b;
		}, false);
	}

	collect(char) {
		console.error(this.constructor.name, char);
		const accepted = this.accepts(char);
		console.error(this.constructor.name, 'accepted', accepted);
		if (!accepted) {
			if (this.collected.length === 0) {
				console.error(this.collected);
				throw new RangeError(`Unexpected character '${char}' for ${this.type} token. Expected: ${
					[].concat(this.expected).map(e=>`'${e}'`).join(', ')
				}.`);
			}
			return false;
		}

		this.collected.push(char);

		return true;
	}
}
