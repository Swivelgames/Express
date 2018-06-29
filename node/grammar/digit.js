import Token from '../classes/token';

export default class DigitToken extends Token {
	static type = 'DIGIT';
	static expected = [/^\d(\.\d+)?$/];
}
