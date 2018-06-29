import { Token } from 'express-parse';

export default class DigitToken extends Token {
	static type = 'DIGIT';
	static expected = [/^\d(\.\d+)?$/];
}
