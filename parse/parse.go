package parse

import (
	"errors"
	"fmt"
	"os"

	"github.com/sgg7269/tokenizer/token"
)

var endTokens []token.Token

// FIXME: This will probably need to be a map to an actual scope which will then be
// a map[string]token.Value{}
var identMap = map[string]token.Value{}

func getLastParsedToken() (token.Token, int, error) {
	fmt.Println("getLastParsedToken")
	fmt.Println(endTokens[len(endTokens)-1])

	if parseIndex > 0 {
		return endTokens[len(endTokens)-1], parseIndex, nil
	}

	return token.Token{}, parseIndex, errors.New("parseIndex less than 1")
}

// TODO: take another look at the returns on this function later
func getFactor(i int, tokens []token.Token) (token.Token, error) {
	fmt.Println("getting factor")

	// FIXME: we should first add a check for parens
	// FIXME: this will only work for spaces between them; hence the +2
	// lookAhead := tokens[i]

	lookAheadNext, i := getNextNonWSToken(parseIndex)
	fmt.Println("next", lookAheadNext)

	// 3 + 9 * 4
	// 9 * 4 + 3

	// lookAheadNext, _ = getNextNonWSToken(parseIndex)
	// parseIndex = parseIndex + i
	// fmt.Println("next", lookAheadNext)

	// lookAheadNext, _ = getNextNonWSToken(parseIndex)
	// fmt.Println("next", lookAheadNext)

	// FIXME: this should be a switch case
	if lookAheadNext.Type == "LITERAL" {
		fmt.Println("Found a literal")
		endTokens = append(endTokens, lookAheadNext)
		parseIndex = parseIndex + i
	} else if lookAheadNext.Type == "IDENT" {
		// gonna need to look up the ident; this could be a var or a keyword, but we need to check if its actually a function and stuff
		if ident, ok := identMap[lookAheadNext.Value.String]; ok {
			fmt.Println("woah i found the var", ident)
		} else {
			fmt.Println("wtf mom the var isnt there")
			return token.Token{}, errors.New("undefined variable reference")
		}

		fmt.Println("Found an ident")
		endTokens = append(endTokens, lookAheadNext)
		parseIndex = parseIndex + i
		// lookAhead.Expected = ""
	} else if lookAheadNext.Type == "L_PAREN" {
		fmt.Println("Found a paren")
		endTokens = append(endTokens, lookAheadNext)
		parseIndex = parseIndex + i
		expr, err := getExpr(parseIndex, tokens)
		if err != nil {
			fmt.Println("woah bro an error", expr, err)
		}
		rParen, i := getNextNonWSToken(parseIndex)
		fmt.Println(rParen, i)
		if rParen.Type != "R_PAREN" {
			// throw up a fuckin err or something, idk 2 lazy rn brah
		}
		endTokens = append(endTokens, rParen)
		parseIndex = parseIndex + i

	} else if lookAheadNext.Type == "L_BRACE" {
		blockTok, err := getBlock()
		if err != nil {
			fmt.Println("something when wrong in the block", blockTok, err)
		}
	} else {
		return lookAheadNext, errors.New("Didn't find a factor")
	}

	return lookAheadNext, nil
}

// func getTerm(i int, tokens []token.Token) (token.Token, error) {
// 	fmt.Println("getting term")

// 	return getFactor(i, tokens)
// }

// FIXME: fix this guys return later
func getTerm(i int, tokens []token.Token) (token.Token, error) {
	fmt.Println("getting term")

	// FIXME: need to look at how this will effect negative mult/div statements
	factor, err := getFactor(i, tokens)
	if err != nil {
		return factor, err
	}
	fmt.Println("factor", factor)

	for {
		opTerm, err := getPriOp()
		if err != nil {
			// FIXME: fix this shoehorned EOF
			if opTerm.Type == "EOF" {
				fmt.Println("APPENDING TO EOF", opTerm)
				endTokens = append(endTokens, opTerm)
				return opTerm, nil
			}
			fmt.Println("ERROR:", err)
			return opTerm, nil
		}
		fmt.Println("opTerm", opTerm)

		factor, err := getFactor(i, tokens)
		if err != nil {
			fmt.Println("ERROR:", err)
			return factor, err
		}
		fmt.Println("factor", factor)
	}

	// TODO: this needs to check EOS
	// its fine now since we only have one statement
	return token.Token{
		Type: "IDK WTD TD",
	}, errors.New("how 2 even get here")
}

func getPriOp() (token.Token, error) {
	fmt.Println("getPriOp")

	op, i := getNextNonWSToken(parseIndex)
	if op.Type == "PRI_OP" {
		endTokens = append(endTokens, op)
		parseIndex = parseIndex + i
		return op, nil
	}

	fmt.Println("pri op, i", op, i)

	return op, errors.New("Did not find ze pri op")
}

func getBlock() (token.Token, error) {
	fmt.Println("getBlock")

	lBrace, i := getNextNonWSToken(parseIndex)
	fmt.Println("lBrace, i", lBrace, i)
	endTokens = append(endTokens, lBrace)
	// need to do a get statements thing
	// savedParseIndex = parseIndex+1
	// TODO: we should also make this return an error
	Parse(tokensGlobal[parseIndex+1:], "namerino")

	// rBrace, i := getNextNonWSToken(parseIndex)
	rBrace, i, err := getLastParsedToken() //parseIndex)
	fmt.Println("rBrace, i, err", rBrace, i, err)
	// endTokens = append(endTokens, rBrace)
	if rBrace.Type == "R_BRACE" {
		fmt.Println("found an r brace again")
	}

	return rBrace, nil
}

func getSecOp() (token.Token, error) {
	fmt.Println("getSecOp")

	op, i := getNextNonWSToken(parseIndex)
	if op.Type == "SEC_OP" {
		endTokens = append(endTokens, op)
		parseIndex = parseIndex + i
		return op, nil
	}

	fmt.Println("sec op, i", op, i)

	return op, errors.New("Did not find ze op")
}

// FIXME: fix this guys return later
func getExpr(i int, tokens []token.Token) (token.Token, error) {
	fmt.Println("getting expr")

	// Find a negative or positive
	// TODO: should check the error later
	_, err := getSecOp()
	if err == nil {
		fmt.Println("no error...? what to do")
	}

	// // TODO: idk need to solve this later, 3 lazy 5 u
	// if err != nil {

	// }

	for {
		termToken, err := getTerm(i, tokens)
		if err != nil {
			// FIXME: fix this shoehorned EOF
			if termToken.Type == "EOF" {
				fmt.Println("APPENDING TO EOF", termToken)
				endTokens = append(endTokens, termToken)
				return termToken, nil
			}
			fmt.Println("ERROR:", err)
			return termToken, err
		}
		fmt.Println("termToken", termToken)

		opTerm, err := getSecOp()
		if err != nil {
			fmt.Println("ERROR:", err)
			return opTerm, nil
		}
		fmt.Println("opTerm", opTerm)
	}

	// TODO: this needs to check EOS
	// its fine now since we only have one statement
	return token.Token{
		Type: "IDK WTD TD",
	}, errors.New("how 2 even get here")
}

// Meta ...
type Meta struct {
	Expected  string // TODO: this should change to an int later after we properly assign IDs
	LastToken token.Token
}

// func equals(t token.Token, meta Meta, i int, tokens []token.Token) bool {
// 	fmt.Println("New token:", t, meta, i)
// 	// TODO: this is whats not working
// 	tok := token.TokenMap[meta.LastToken.Value.String+t.Value.String]
// 	fmt.Println(tok)
// 	meta.Expected = "EXPR"
// 	meta.LastToken = tok
// 	endTokens[len(endTokens)-1] = tok

// 	ge := getExpr(i, tokens)
// 	fmt.Println()
// 	fmt.Println("ge", ge)
// 	fmt.Println()

// 	if ge {
// 		i = i + 1
// 		meta.Expected = "EOS"
// 		meta.LastToken = tokens[i]
// 		endTokens = append(endTokens, tokens[i])
// 		// FIXME: clean this shit up
// 		return true
// 	}
// 	// TODO: this would be an error
// 	fmt.Println()
// 	fmt.Println("Syntax ERROR")
// 	fmt.Println()
// 	// TODO: need to put actual error codes here
// 	// FIXME: we shouldn't os.exit here, instead return an error, handle it, probably should have some kind of map lookup for the specific error shit
// 	// FIXME: we also need to print out debuf information about the current parse information
// 	os.Exit(666)

// 	return false
// }

// // Parse ...
// func Parse(lexTokens []token.Token, name string) []token.Token {

// 	fmt.Println()
// 	fmt.Println("Outtputting")

// 	meta := Meta{}

// 	// FIXME: Need to make this not a range over
// 	for i := 0; i < len(lexTokens); i++ {
// 		t := lexTokens[i]
// 		// strip out WS tokens for now
// 		if t.Type == "WS" || t.Type == "EOF" {
// 			continue
// 		}

// 		// FIXME: this should be an int after the change
// 		if meta.Expected != "" {
// 			if t.Type == meta.Expected {
// 				fmt.Println("Wow we were actually expected")
// 				fmt.Println(t)
// 				// TODO: run some function, do stuff
// 				meta.Expected = t.Expected
// 				meta.LastToken = t
// 				endTokens = append(endTokens, t)
// 				continue
// 			} else {
// 				fmt.Println("wtf why mom")
// 				// TODO: need to handle this

// 				// TODO: why are we doing this based on the last token?
// 				// TODO: it might be more useful if we compare the current types of the token and the meta.LastToken
// 				// TODO: this is where we could have functions already plug and play defined that have the token check the 'nextToken' and then return the token that should be used
// 				switch meta.LastToken.Value.String {
// 				case ":":
// 					// TODO: this could be recursive
// 					switch t.Value.String {
// 					case "=":
// 						// FIXME: these are very hacky right now because the function actually os.Exits out
// 						// FIXME: WS assumption here
// 						if equals(t, meta, i+2, lexTokens) {
// 							i++
// 							continue
// 						}
// 					}
// 				case "=":
// 					if equals(t, meta, i, lexTokens) {
// 						i++
// 						continue
// 					}
// 				case ";":
// 					endTokens = append(endTokens, t)
// 				default:
// 					fmt.Println()
// 					fmt.Printf("Syntax ERROR: default case hit %+v %+v %d\n", t, meta, i)
// 					fmt.Println()
// 					// TODO: need to put actual error codes here
// 					// FIXME: we shouldn't os.exit here, instead return an error, handle it, probably should have some kind of map lookup for the specific error shit
// 					// FIXME: we also need to print out debuf information about the current parse information
// 					os.Exit(666)
// 				}
// 			}
// 		}
// 		fmt.Println(t)
// 		meta.Expected = t.Expected
// 		meta.LastToken = t
// 		fmt.Println(t)
// 		endTokens = append(endTokens, t)
// 	}

// 	tokenFilename := name + ".tokens"

// 	// For more granular writes, open a file for writing.
// 	tokenFile, err := os.Create(tokenFilename)
// 	defer func() {
// 		if err = tokenFile.Close(); err != nil {
// 			fmt.Println("ERROR: Could not close file:", tokenFilename)
// 		}
// 	}()
// 	if err != nil {
// 		fmt.Println("ERROR: Could not open token output file:", tokenFilename)
// 		os.Exit(9)
// 	}
// 	tokenWriter := bufio.NewWriter(tokenFile)

// 	// TODO: this needs to be outputted as program.expr.parse
// 	fmt.Println()
// 	fmt.Println("End Tokens:")
// 	for i := 0; i < len(endTokens); i++ {
// 		t := endTokens[i]
// 		fmt.Println(t)

// 		// TODO: should make a function specifically for writing the tokens
// 		tokenJSON, jerr := json.Marshal(t)
// 		if jerr != nil {
// 			fmt.Println("ERROR: Could not marshal token JSON: ", t)
// 		}

// 		_, err = tokenWriter.WriteString(string(tokenJSON) + "\n")
// 		if err != nil || jerr != nil {
// 			fmt.Println("ERROR: Could not write token data: ", tokenJSON)
// 		}
// 	}

// 	err = tokenWriter.Flush()
// 	if err != nil {
// 		fmt.Println("ERROR: Could not flush writer, data may be missing:", tokenFilename)
// 	}

// 	return endTokens
// }

// TODO: we will need this later
func getNextNonWSToken(i int) (token.Token, int) {
	tokens := tokensGlobal[i+1:]

	for i = 0; i < len(tokens); i++ {
		if i >= parseLen-1 {
			return token.Token{}, i
		}

		if tokens[i].Type != "WS" {
			return tokens[i], i + 1
		}
	}

	return token.Token{}, -1
}

// func getNextToken(i int) token.Token {
// 	if i < parseLen-1 {
// 		return tokensGlobal[i+1]
// 	}
// 	return token.Token{}
// }

// FIXME: we need to clean up all errors

// TODO: think of a different name
// For now just return the token like this
func getType(t token.Token) (token.Token, error) {
	nextToken, i := getNextNonWSToken(parseIndex)

	switch nextToken.Type {
	case "IDENT":
		// TODO: this needs to have some logic to say if it is already there then wtf to do
		// TODO: this needs to do keywords too
		if _, ok := identMap[nextToken.Value.String]; !ok {
			fmt.Println("DID NOT FOUND THE VARIABLE")
			identMap[nextToken.Value.String] = token.Value{
				Type: t.Value.String,
				True: "",
				// String: "",
			}
		}
		endTokens = append(endTokens, nextToken)
		parseIndex = parseIndex + i
		return nextToken, nil
	default:
		return nextToken, errors.New("Didn't find a valid token")
	}
}

// func getEOS() (token.Token, error) {
// 	nextToken, i := getNextNonWSToken(parseIndex)

// 	switch nextToken.Type {
// 	case "EOF":
// 		endTokens = append(endTokens, nextToken)
// 		parseIndex = parseIndex + i
// 		return nextToken, nil
// 	default:
// 		return nextToken, errors.New("Didn't find a valid token")
// 	}
// }

func getAssign() (token.Token, error) {
	nextToken, i := getNextNonWSToken(parseIndex)

	// FIXME: this should necessarily expect a LITERAL, should expect probably another EXPR or atleast something that can be evaluated
	switch nextToken.Type {
	case "LITERAL":
		endTokens = append(endTokens, nextToken)
		parseIndex = parseIndex + i + 1
		return nextToken, nil
	default:
		return nextToken, errors.New("Didn't find a valid token")
	}
}

var parseIndex = 0
var parseLen = 0
var tokensGlobal []token.Token

// Parse ...
func Parse(tokens []token.Token, name string) []token.Token {
	tokensGlobal = tokens
	parseLen = len(tokens) - 1

	// FIXME: we should start off with things like GetStatement(), GetExpr(), GetTerm(), etc
	for {
		t := tokens[parseIndex]

		if t.Type != "WS" {
			// TODO: would be more efficient to just make a 'stripWSTokens()' function
			switch t.Type {
			case "TYPE":
				endTokens = append(endTokens, t)
				token, err := getType(t)
				if err != nil {
					fmt.Printf("ERROR: %s\nFound: %#v\n", err, token)
					os.Exit(666)
				}

			case "IDENT":
				endTokens = append(endTokens, t)
				fmt.Println("found an ident")

			// TODO: in the case for ":" (SET), there n$eeds to be some checking for the assign/equals/set tokens
			case "ASSIGN":
				endTokens = append(endTokens, t)
				fmt.Println("found an equals")
				// token, err := getAssign()
				// if err != nil {
				// 	fmt.Printf("ERROR: %s\nFound: %#v\n", err, token)
				// 	os.Exit(666)
				// }

				fmt.Println("getExpression")
				exprToken, err := getExpr(parseIndex+1, tokens[parseIndex+1:])
				if err != nil {
					fmt.Println("got a fucking error dude")
					return endTokens
				}
				fmt.Println("exprToken", exprToken)

			case "NEWLINE":
				// endTokens = append(endTokens, t)
				// fmt.Println("found newline")

			case "L_BRACE":
				fmt.Println("left found a thingerooni", t)

			case "R_BRACE":
				// This should not be hit for a standard object definition
				fmt.Println("right found a thingerooni", t)
				endTokens = append(endTokens, t)
				return endTokens

			case "EOS":
				endTokens = append(endTokens, t)
				fmt.Println("found an EOS")
			// We might need this later for something else if we reuse the semicolon
			// token, err := getEOS()
			// if err != nil {
			// 	fmt.Printf("ERROR: %s\nFound: %#v\n", err, token)
			// 	os.Exit(666)
			// }

			case "EOF":
				endTokens = append(endTokens, t)
				fmt.Println("found EOF")
				// TODO: FIXME: might need to make something to check enclosing, variable mappings, shit, etc

			default:
				fmt.Println("I did not recognize this token")
				fmt.Println(t)
				return endTokens
			}

			fmt.Println(t)
			fmt.Println()
		}

		if parseIndex >= parseLen {
			break
		}
		parseIndex++
	}

	fmt.Println("identMap", identMap)

	return endTokens
}

// TODO: FIXME: we need to implement something that will track the statement and origanize the data in such a away that will make it easy to to build the variable map
