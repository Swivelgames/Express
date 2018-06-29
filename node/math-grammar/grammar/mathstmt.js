'use strict';

Object.defineProperty(exports, "__esModule", {
	value: true
});

var _expressParse = require('express-parse');

var _digit = require('./digit');

var _digit2 = _interopRequireDefault(_digit);

var _mathOp = require('./mathOp');

var _mathOp2 = _interopRequireDefault(_mathOp);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _possibleConstructorReturn(self, call) { if (!self) { throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); } return call && (typeof call === "object" || typeof call === "function") ? call : self; }

function _inherits(subClass, superClass) { if (typeof superClass !== "function" && superClass !== null) { throw new TypeError("Super expression must either be null or a function, not " + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass; }

var MathStmt = function (_Collector) {
	_inherits(MathStmt, _Collector);

	function MathStmt() {
		_classCallCheck(this, MathStmt);

		return _possibleConstructorReturn(this, (MathStmt.__proto__ || Object.getPrototypeOf(MathStmt)).apply(this, arguments));
	}

	return MathStmt;
}(_expressParse.Collector);

MathStmt.type = 'MATHSTMT';
MathStmt.expected = [[_digit2.default, _mathOp2.default, MathStmt], [_digit2.default, _mathOp2.default, _digit2.default]];
exports.default = MathStmt;