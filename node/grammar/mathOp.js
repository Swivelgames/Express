import Token from '../classes/token';

export default class MathOpToken extends Token {
	static type = 'MATHOP';
	static expected = ['+','-','*','/'];
}
