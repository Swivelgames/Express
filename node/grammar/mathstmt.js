import Collector from '../classes/collector';
import DigitToken from './digit';
import MathOpToken from './mathOp';

export default class MathStmt extends Collector {
	static type = 'MATHSTMT';
	static expected = [
		[DigitToken, MathOpToken, MathStmt],
		[DigitToken, MathOpToken, DigitToken]
	];
}
