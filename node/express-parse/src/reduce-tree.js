import Collector from './classes/collector';

const reduceToken = (token) => ({
	type: token.type,
	value: token.collected.join('')
});

const reduceCollector = (col) => ({
	type: col.type,
	tokens: col.collected.map(r => {
		if (r instanceof Collector) {
			return reduceCollector(r);
		} else {
			return reduceToken(r);
		}
	})
});

export default (tree) => reduceCollector(tree);
