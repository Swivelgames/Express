import { Token } from 'express-parse';

export default class MathOpToken extends Token {
	static type = 'MATHOP';
	static expected = ['+','-','*','/'];
}
