package parse

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pkg/errors"
	"github.com/scottshotgg/Express/token"
)

var (
	declaredType  string
	declaredName  string
	declaredValue token.Value
	declarations  = map[string]token.Value{}
)

// CheckType checks the usage of the type
func (m *Meta) CheckType() {
	fmt.Println("CheckType")

	switch m.NextToken.Type {
	case token.Ident:

		m.CollectCurrentToken()
		m.Shift()
		m.CollectCurrentToken()

		//TODO: change all of these to be assign
		switch m.NextToken.Type {
		case token.Init:
			fallthrough
		case token.Set:
			fallthrough
		case token.Assign:
			fmt.Println("found an ASSIGN type")
			m.Shift()
			m.CollectCurrentToken()

			fmt.Println(m.NextToken)
			if m.NextToken.Type != token.Literal && m.NextToken.Type != token.Ident {
				fmt.Println("did not find a literal")
				os.Exit(0)
			}

			m.Shift()
			m.CollectCurrentToken()
		}
	}
}

// CheckIdent check the usage of the ident
func (m *Meta) CheckIdent() {
	fmt.Println(token.Ident)

	m.CollectCurrentToken()

	switch m.NextToken.Type {
	case token.Init:
		fallthrough
	case token.Set:
		fallthrough
	case token.Assign:
		fmt.Println("found an ASSIGN type")
		m.Shift()
		m.CollectCurrentToken()

		// if the form is [ident] [= | : | :=] then expect an expression
		m.GetExpression()

	default:
		fmt.Println("vert da ferk")
		os.Exit(9)
	}
}

// GetFactor returns the next factor in the sequence
func (m *Meta) GetFactor() {
	m.Shift()
	switch m.CurrentToken.Type {
	case token.Ident:
		fmt.Println("found an ident")
		tValue, ok := declarations[m.CurrentToken.Value.String]
		if !ok {
			fmt.Println("Undefined variable reference")
			os.Exit(9)
		}
		if tValue.Type != declaredType {
			fmt.Println("Variable type mismatch")
			fmt.Println("Expected", declaredType, "got", tValue.Type)
			os.Exit(9)
		}
		// TODO: we may actually want to say that this is in actuallality the variable and not the value so that it can do optimizations
		declaredValue = tValue
		fmt.Println("declareds", declaredName, declaredType, declaredValue)
		m.CollectToken(token.Token{
			ID: 1,
			// Type:, // TODO: not sure what to put here
			Value: tValue,
		})

	case token.Literal:
		fmt.Println("found a literal")
		if m.CurrentToken.Value.Type != declaredType {
			fmt.Println("Variable type mismatch")
			fmt.Println("Expected", declaredType, "got", m.CurrentToken.Value.Type)
			os.Exit(9)
		}
		declaredValue = m.CurrentToken.Value
		m.CollectCurrentToken()

	case token.LParen:
		fmt.Println("found an expr")
		m.GetExpression()

	default:
		fmt.Println("ERROR getting factor")
		fmt.Println("Expected factor, got", m.CurrentToken)
		os.Exit(9)
	}
}

// GetTerm gets the next term in the sequence
func (m *Meta) GetTerm() {
	m.GetFactor()
}

// GetSecOp gets a secondary operation; + and -
func (m *Meta) GetSecOp() {
	fmt.Println("current", m.CurrentToken)
}

// AddOperands returns the addition of two operands based on their type
func (m *Meta) AddOperands(left, right token.Value) (token.Value, error) {
	var valueToken token.Value
	leftType := left.Type
	rightType := right.Type
	fmt.Println("firsttime", left, right, leftType, rightType)

	if leftType == rightType {
		valueToken.Type = leftType

		switch leftType {
		case token.IntType:
			valueToken.True = left.True.(int) + right.True.(int)
			valueToken.String = strconv.Itoa(valueToken.True.(int))

		case token.StringType:
			valueToken.True = left.True.(string) + right.True.(string)
			valueToken.String = valueToken.True.(string)

		case token.FloatType:
			valueToken.True = left.True.(float64) + right.True.(float64)
			// TODO: need to count the decimal place if we start using this
			valueToken.String = strconv.FormatFloat(valueToken.True.(float64), 'f', 5, 64)

		case token.BoolType:
			valueToken.True = left.True.(bool) || right.True.(bool)
			valueToken.String = strconv.FormatBool(valueToken.True.(bool))

		case token.CharType:
			// TODO: we will need to take into account the character encoding here and overflowing
			valueToken.True = string(rune(left.True.(string)[0]) + rune(right.True.(string)[0]))
			valueToken.String = valueToken.True.(string)

		// TODO: this will need some more thinking
		// case token.Byte:

		default:
			fmt.Println("Type not declared for AddOperands", left, right, leftType, rightType)
			os.Exit(9)
		}

		return valueToken, nil
	}
	// switch declaredType {
	// case token.IntType:
	// 	fmt.Println(left, right)
	// }

	err := errors.New("Could not perform AddOperand on operands")
	fmt.Println(err, left, right, leftType, rightType)
	return token.Value{}, err
}

// GetOperationValue returns the value of the operation being performed in the statement
func (m *Meta) GetOperationValue(left token.Token, right token.Token, op token.Token) token.Value {
	// token.Value {
	// 	// Type: m.DetermineTypeFromOperation(), // TODO: save this until later when we want to support multi type operations
	// 	Type:
	// },
	// // valueToken.True
	// // valueToken.String
	// m.CollectToken(valueToken)

	// fmt.Println(left.Value.True + right.Value.True)

	// if leftType == rightType {
	// switch on the op
	switch op.Value.String {
	case "+":
		value, err := m.AddOperands(left.Value, right.Value)
		if err != nil {
			fmt.Println("could not add operands idk wtf happened", left.Value, right.Value)
			os.Exit(9)
		}
		return value

	default:
		fmt.Println("Invalid operand", op)
	}

	// } else {
	// 	// were gonna have to do something else
	// 	// usupported for now
	// 	fmt.Println("Unsupported operation detected")
	// 	fmt.Println("typeof", leftType, op, rightType)
	// 	os.Exit(9)
	// }

	return token.Value{}
}

// GetExpression gets the next expression
func (m *Meta) GetExpression() {
	m.GetTerm()
	value1 := m.LastCollectedToken
	m.RemoveLastCollectedToken()
	fmt.Println("last", m.LastCollectedToken)

	// FIXME: need to make something to evaluate the statement
	if m.NextToken.Type == token.SecOp {
		m.Shift()
		op := m.CurrentToken

		m.GetExpression()
		value2 := m.LastCollectedToken
		m.RemoveLastCollectedToken()

		// FIXME: TODO: really should do this with some sort of eval or using reflection
		switch op.Value.String {
		case "+":
			fmt.Println("value1, value2", value1, value2)

			valueToken := token.Token{
				ID:   1,
				Type: token.Literal,
				// Expected: "",
				Value: m.GetOperationValue(value1, value2, op),
			}
			fmt.Println(valueToken)
			m.CollectToken(valueToken)

			declaredValue = valueToken.Value

		default:
			fmt.Println("Operator not defined")
			fmt.Println("Found operator:", op)
			os.Exit(9)
		}
	}
}

// GetAssignmentStatement gets the next assignment statement in the sequence
func (m *Meta) GetAssignmentStatement() error {
	m.Shift()
	switch m.CurrentToken.Type {
	// Get the TYPE
	case token.Type:
		fmt.Println("found a type")
		// switch m.CurrentToken.Value.String {
		// 	case "int"
		// }
		declaredType = m.CurrentToken.Value.String
		m.CollectCurrentToken()

		// Get the IDENT
		m.Shift()
		if m.CurrentToken.Type != token.Ident {
			fmt.Println("Syntax error getting assignment_stmt")
			fmt.Println("Expected IDENT, got", m.CurrentToken)
			os.Exit(9)
		}
		_, ok := declarations[m.CurrentToken.Value.String]
		if ok {
			fmt.Println("Variable already declared")
			os.Exit(9)
		}
		declaredName = m.CurrentToken.Value.String
		m.CollectCurrentToken()

		// Get the assignemnt operator
		m.Shift()
		if m.CurrentToken.Type != token.Assign && m.CurrentToken.Type != token.Init && m.CurrentToken.Type != token.Set {
			fmt.Println("Syntax error getting assignment_stmt")
			fmt.Println("Expected assign_op, got", m.CurrentToken)
			os.Exit(9)
		}
		m.CollectCurrentToken()

		// FIXME: this should return an error that we can check
		m.GetExpression()

		declarations[declaredName] = declaredValue
		declaredType = ""
		declaredName = ""
		declaredValue = token.Value{}
		fmt.Println(declarations)

	case token.Ident:
		fmt.Println("i spy an ident")
		currentIdent := m.CurrentToken

		if current, ok := declarations[currentIdent.Value.String]; ok {
			declaredName = currentIdent.Value.String
			declaredType = current.Type
		} else {
			fmt.Println("Variable reference not found", currentIdent)
			os.Exit(9)
		}

		m.Shift()
		if m.CurrentToken.Type == token.Assign {

			fmt.Println(m.CurrentToken)

			// FIXME: this should return an error that we can check
			m.GetExpression()

			declarations[declaredName] = declaredValue
			declaredType = ""
			declaredName = ""
			declaredValue = token.Value{}
			fmt.Println(declarations)
		}

	default:
		fmt.Println("ERROR getting assignement statement")
		fmt.Println("expected assignement statement beginning, got", m.CurrentToken)
		return errors.New("blah")
	}

	// m.GetExpression()
	return nil
}

// GetStatement gets the next statement in the sequence
func (m *Meta) GetStatement() {
	err := m.GetAssignmentStatement()
	// FIXME: ideally we should do a switch on the error but w/e for now
	// TODO: we need some way to backtrack to before this operation ...
	if err != nil {
		fmt.Println("error getting statement", err)
		// TODO: woah we could do partial compilation
		os.Exit(9)
	}
}

// CheckBlock check the usage of the block
func (m *Meta) CheckBlock() {
	for {
		m.GetStatement()

		// current := m.CurrentToken
		// switch current.Type {
		// case token.Type:
		// 	// at this point we need to think about the different options
		// 	// int name			: simplest
		// 	// int name = 6 : next simplest
		// 	// name int = 6 : would be taken care of by a random ident
		// 	// we would probably also have other ones for function params and stuff
		// 	m.CheckType()

		// case token.Ident:
		// 	m.CheckIdent()
		// }

		if m.NextToken == (token.Token{}) {
			fmt.Println("returning")
			return
		}
	}

}

// Semantic runs a semantic parse on the tokens
func Semantic(tokens []token.Token) ([]token.Token, error) {
	// Auto inject the brackets to ensure that they are there
	meta := Meta{
		IgnoreWS:         true,
		Tokens:           tokens[0].Value.True.([]token.Token),
		Length:           len(tokens[0].Value.True.([]token.Token)),
		CheckOptmization: true,
	}
	meta.Shift()

	fmt.Println(tokens)

	meta.CheckBlock()
	fmt.Println("tokens", meta.EndTokens)

	fmt.Println(declarations)

	// Here we are continuously applying semantic pressure to squash the tokens and furthur
	// simplify the tokens generated
	// for meta.CheckOptmization {
	// 	fmt.Println("Optimizing", meta.OptimizationAttempts)
	// 	// meta.CollectTokens(meta.ParseBlock().Value.True.([]token.Token))
	// 	fmt.Println("endTokens", meta.EndTokens)

	// 	fmt.Println(meta.EndTokens)
	// 	metaTokens := meta.EndTokens[0].Value.True.([]token.Token)
	// 	metaTokens = append(append([]token.Token{token.TokenMap["{"]}, metaTokens...), token.TokenMap["}"])
	// 	fmt.Println("metaTokens", len(metaTokens), len(meta.EndTokens))

	// 	// endTokens = meta.EndTokens

	// 	// TODO: FIXME: w/e this works for now
	// 	// Fix this from pulling off only the top one
	// 	// Only apply SemanticPressure once for now until we figure out the recursion more
	// 	if meta.OptimizationAttempts > 0 {
	// 		break
	// 	}

	// 	// fmt.Println("meta.CheckOptimization", meta.CheckOptmization)

	// 	// if !meta.CheckOptmization {
	// 	// 	break
	// 	// }

	// 	// if len(meta.EndTokens) < len(meta.Tokens) {
	// 	// 	break
	// 	// }

	// 	meta = Meta{
	// 		// FIXME: do we need to fix this hack?
	// 		// Tokens: ,
	// 		Tokens:               metaTokens,
	// 		Length:               len(metaTokens),
	// 		CheckOptmization:     meta.CheckOptmization,
	// 		OptimizationAttempts: meta.OptimizationAttempts + 1,
	// 	}
	// }

	return meta.EndTokens, nil
}
