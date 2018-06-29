'use strict';

Object.defineProperty(exports, "__esModule", {
	value: true
});

var _expressParse = require('express-parse');

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _possibleConstructorReturn(self, call) { if (!self) { throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); } return call && (typeof call === "object" || typeof call === "function") ? call : self; }

function _inherits(subClass, superClass) { if (typeof superClass !== "function" && superClass !== null) { throw new TypeError("Super expression must either be null or a function, not " + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass; }

var DigitToken = function (_Token) {
	_inherits(DigitToken, _Token);

	function DigitToken() {
		_classCallCheck(this, DigitToken);

		return _possibleConstructorReturn(this, (DigitToken.__proto__ || Object.getPrototypeOf(DigitToken)).apply(this, arguments));
	}

	return DigitToken;
}(_expressParse.Token);

DigitToken.type = 'DIGIT';
DigitToken.expected = [/^\d(\.\d+)?$/];
exports.default = DigitToken;