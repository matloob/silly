//line silly.y:2
package main

import __yyfmt__ "fmt"

//line silly.y:3
//line silly.y:7
type sillySymType struct {
	yys   int
	val   int
	ident string
	node  node
	list  *stmtList
}

const VAL = 57346
const IDENT = 57347

var sillyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"'{'",
	"'}'",
	"';'",
	"'='",
	"'+'",
	"'/'",
	"'-'",
	"'*'",
	"'('",
	"')'",
	"'#'",
	"VAL",
	"IDENT",
}
var sillyStatenames = [...]string{}

const sillyEofCode = 1
const sillyErrCode = 2
const sillyInitialStackSize = 16

//line silly.y:46

//line yacctab:1
var sillyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const sillyNprod = 13
const sillyPrivate = 57344

var sillyTokenNames []string
var sillyStates []string

const sillyLast = 33

var sillyAct = [...]int{

	5, 6, 8, 7, 9, 10, 13, 2, 19, 1,
	12, 3, 4, 14, 15, 16, 17, 18, 21, 11,
	6, 8, 7, 9, 0, 0, 0, 0, 0, 0,
	0, 0, 20,
}
var sillyPact = [...]int{

	-4, -1000, -7, -1000, -2, -4, -4, -4, -4, -4,
	-4, 3, -4, 12, -7, -7, -7, -7, -7, -1000,
	-1000, -1000,
}
var sillyPgo = [...]int{

	0, 6, 10, 19, 9,
}
var sillyR1 = [...]int{

	0, 4, 1, 1, 1, 1, 1, 1, 1, 1,
	3, 3, 2,
}
var sillyR2 = [...]int{

	0, 1, 1, 3, 3, 3, 3, 3, 1, 3,
	0, 2, 2,
}
var sillyChk = [...]int{

	-1000, -4, -1, 15, 16, 4, 8, 10, 9, 11,
	7, -3, -2, -1, -1, -1, -1, -1, -1, 5,
	-3, 6,
}
var sillyDef = [...]int{

	0, -2, 1, 2, 8, 10, 0, 0, 0, 0,
	0, 0, 10, 0, 3, 4, 5, 6, 7, 9,
	11, 12,
}
var sillyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 14, 3, 3, 3, 3,
	12, 13, 11, 8, 3, 10, 3, 9, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 6,
	3, 7, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 4, 3, 5,
}
var sillyTok2 = [...]int{

	2, 3, 15, 16,
}
var sillyTok3 = [...]int{
	0,
}

var sillyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	sillyDebug        = 0
	sillyErrorVerbose = false
)

type sillyLexer interface {
	Lex(lval *sillySymType) int
	Error(s string)
}

type sillyParser interface {
	Parse(sillyLexer) int
	Lookahead() int
}

type sillyParserImpl struct {
	lval  sillySymType
	stack [sillyInitialStackSize]sillySymType
	char  int
}

func (p *sillyParserImpl) Lookahead() int {
	return p.char
}

func sillyNewParser() sillyParser {
	return &sillyParserImpl{}
}

const sillyFlag = -1000

func sillyTokname(c int) string {
	if c >= 1 && c-1 < len(sillyToknames) {
		if sillyToknames[c-1] != "" {
			return sillyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func sillyStatname(s int) string {
	if s >= 0 && s < len(sillyStatenames) {
		if sillyStatenames[s] != "" {
			return sillyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func sillyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !sillyErrorVerbose {
		return "syntax error"
	}

	for _, e := range sillyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + sillyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := sillyPact[state]
	for tok := TOKSTART; tok-1 < len(sillyToknames); tok++ {
		if n := base + tok; n >= 0 && n < sillyLast && sillyChk[sillyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if sillyDef[state] == -2 {
		i := 0
		for sillyExca[i] != -1 || sillyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; sillyExca[i] >= 0; i += 2 {
			tok := sillyExca[i]
			if tok < TOKSTART || sillyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if sillyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += sillyTokname(tok)
	}
	return res
}

func sillylex1(lex sillyLexer, lval *sillySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = sillyTok1[0]
		goto out
	}
	if char < len(sillyTok1) {
		token = sillyTok1[char]
		goto out
	}
	if char >= sillyPrivate {
		if char < sillyPrivate+len(sillyTok2) {
			token = sillyTok2[char-sillyPrivate]
			goto out
		}
	}
	for i := 0; i < len(sillyTok3); i += 2 {
		token = sillyTok3[i+0]
		if token == char {
			token = sillyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = sillyTok2[1] /* unknown char */
	}
	if sillyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", sillyTokname(token), uint(char))
	}
	return char, token
}

func sillyParse(sillylex sillyLexer) int {
	return sillyNewParser().Parse(sillylex)
}

func (sillyrcvr *sillyParserImpl) Parse(sillylex sillyLexer) int {
	var sillyn int
	var sillyVAL sillySymType
	var sillyDollar []sillySymType
	_ = sillyDollar // silence set and not used
	sillyS := sillyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	sillystate := 0
	sillyrcvr.char = -1
	sillytoken := -1 // sillyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		sillystate = -1
		sillyrcvr.char = -1
		sillytoken = -1
	}()
	sillyp := -1
	goto sillystack

ret0:
	return 0

ret1:
	return 1

sillystack:
	/* put a state and value onto the stack */
	if sillyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", sillyTokname(sillytoken), sillyStatname(sillystate))
	}

	sillyp++
	if sillyp >= len(sillyS) {
		nyys := make([]sillySymType, len(sillyS)*2)
		copy(nyys, sillyS)
		sillyS = nyys
	}
	sillyS[sillyp] = sillyVAL
	sillyS[sillyp].yys = sillystate

sillynewstate:
	sillyn = sillyPact[sillystate]
	if sillyn <= sillyFlag {
		goto sillydefault /* simple state */
	}
	if sillyrcvr.char < 0 {
		sillyrcvr.char, sillytoken = sillylex1(sillylex, &sillyrcvr.lval)
	}
	sillyn += sillytoken
	if sillyn < 0 || sillyn >= sillyLast {
		goto sillydefault
	}
	sillyn = sillyAct[sillyn]
	if sillyChk[sillyn] == sillytoken { /* valid shift */
		sillyrcvr.char = -1
		sillytoken = -1
		sillyVAL = sillyrcvr.lval
		sillystate = sillyn
		if Errflag > 0 {
			Errflag--
		}
		goto sillystack
	}

sillydefault:
	/* default state action */
	sillyn = sillyDef[sillystate]
	if sillyn == -2 {
		if sillyrcvr.char < 0 {
			sillyrcvr.char, sillytoken = sillylex1(sillylex, &sillyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if sillyExca[xi+0] == -1 && sillyExca[xi+1] == sillystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			sillyn = sillyExca[xi+0]
			if sillyn < 0 || sillyn == sillytoken {
				break
			}
		}
		sillyn = sillyExca[xi+1]
		if sillyn < 0 {
			goto ret0
		}
	}
	if sillyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			sillylex.Error(sillyErrorMessage(sillystate, sillytoken))
			Nerrs++
			if sillyDebug >= 1 {
				__yyfmt__.Printf("%s", sillyStatname(sillystate))
				__yyfmt__.Printf(" saw %s\n", sillyTokname(sillytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for sillyp >= 0 {
				sillyn = sillyPact[sillyS[sillyp].yys] + sillyErrCode
				if sillyn >= 0 && sillyn < sillyLast {
					sillystate = sillyAct[sillyn] /* simulate a shift of "error" */
					if sillyChk[sillystate] == sillyErrCode {
						goto sillystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if sillyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", sillyS[sillyp].yys)
				}
				sillyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if sillyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", sillyTokname(sillytoken))
			}
			if sillytoken == sillyEofCode {
				goto ret1
			}
			sillyrcvr.char = -1
			sillytoken = -1
			goto sillynewstate /* try again in the same state */
		}
	}

	/* reduction by production sillyn */
	if sillyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", sillyn, sillyStatname(sillystate))
	}

	sillynt := sillyn
	sillypt := sillyp
	_ = sillypt // guard against "declared and not used"

	sillyp -= sillyR2[sillyn]
	// sillyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if sillyp+1 >= len(sillyS) {
		nyys := make([]sillySymType, len(sillyS)*2)
		copy(nyys, sillyS)
		sillyS = nyys
	}
	sillyVAL = sillyS[sillyp+1]

	/* consult goto table to find next state */
	sillyn = sillyR1[sillyn]
	sillyg := sillyPgo[sillyn]
	sillyj := sillyg + sillyS[sillyp].yys + 1

	if sillyj >= sillyLast {
		sillystate = sillyAct[sillyg]
	} else {
		sillystate = sillyAct[sillyj]
		if sillyChk[sillystate] != -sillyn {
			sillystate = sillyAct[sillyg]
		}
	}
	// dummy call; replaced with literal code
	switch sillynt {

	case 1:
		sillyDollar = sillyS[sillypt-1 : sillypt+1]
		//line silly.y:25
		{
			r = sillyDollar[1].node
		}
	case 2:
		sillyDollar = sillyS[sillypt-1 : sillypt+1]
		//line silly.y:28
		{
			sillyVAL.node = &intNode{sillyDollar[1].val}
		}
	case 3:
		sillyDollar = sillyS[sillypt-3 : sillypt+1]
		//line silly.y:29
		{
			sillyVAL.node = &binopNode{"+", sillyDollar[1].node, sillyDollar[3].node}
		}
	case 4:
		sillyDollar = sillyS[sillypt-3 : sillypt+1]
		//line silly.y:30
		{
			sillyVAL.node = &binopNode{"-", sillyDollar[1].node, sillyDollar[3].node}
		}
	case 5:
		sillyDollar = sillyS[sillypt-3 : sillypt+1]
		//line silly.y:31
		{
			sillyVAL.node = &binopNode{"/", sillyDollar[1].node, sillyDollar[3].node}
		}
	case 6:
		sillyDollar = sillyS[sillypt-3 : sillypt+1]
		//line silly.y:32
		{
			sillyVAL.node = &binopNode{"*", sillyDollar[1].node, sillyDollar[3].node}
		}
	case 7:
		sillyDollar = sillyS[sillypt-3 : sillypt+1]
		//line silly.y:33
		{
			sillyVAL.node = &assignNode{sillyDollar[1].ident, sillyDollar[3].node}
		}
	case 8:
		sillyDollar = sillyS[sillypt-1 : sillypt+1]
		//line silly.y:34
		{
			sillyVAL.node = &identNode{sillyDollar[1].ident}
		}
	case 9:
		sillyDollar = sillyS[sillypt-3 : sillypt+1]
		//line silly.y:36
		{
			sillyVAL.node = sillyDollar[2].list
		}
	case 10:
		sillyDollar = sillyS[sillypt-0 : sillypt+1]
		//line silly.y:39
		{
			sillyVAL.list = nil
		}
	case 11:
		sillyDollar = sillyS[sillypt-2 : sillypt+1]
		//line silly.y:40
		{
			sillyVAL.list = &stmtList{sillyDollar[1].node, sillyDollar[2].list}
		}
	case 12:
		sillyDollar = sillyS[sillypt-2 : sillypt+1]
		//line silly.y:43
		{
			sillyVAL.node = sillyDollar[1].node
		}
	}
	goto sillystack /* stack new state and value */
}
