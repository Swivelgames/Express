import Collector from './classes/collector';

function lexer(src, curI, CollectorClass, parent = false) {
	console.error('lexer', CollectorClass.name);
	console.error('src', src);
	let i = curI;
	let token;
	const collector = new CollectorClass();
	while(!collector.isSatisfied()) {
		let TokenClass = collector.nextExpected();
		console.error('TokenClass', TokenClass.name, i);
		if (Object.getPrototypeOf(TokenClass) === Collector) {
			const [newI, newToken] = lexer(src, i, TokenClass, true);
			if (!newToken.isSatisfied()) {
				if(!collector.tryAltPath()) break;
				else {
					i = newI;
					continue;
				}
			}
			i = newI;
			token = newToken;
		} else {
			token = new TokenClass();
			try {
				while(token.collect(src[i])) {
					i++;
				}
			} catch(e) {
				if (!collector.tryAltPath()) {
					i = curI;
					break;
				}
			}
		}
		if (token) collector.collect(token);
	}
	if (parent) return [i, collector];
	return collector;
}

export default lexer;
