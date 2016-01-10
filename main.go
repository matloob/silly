package main

import (
	"fmt"
	"os"
	"strconv"
	"text/scanner"
)

//go:generate go tool yacc -o silly.go -p "silly" silly.y

type sillyLex struct {
	s scanner.Scanner
}

func (x *sillyLex) Lex(yylval *sillySymType) int {
	tok := x.s.Scan()
	switch {
	case tok > 0:
		return int(tok)
	case tok == scanner.EOF:
		return 0
	case tok == scanner.Ident:
		yylval.ident = x.s.TokenText()
		return IDENT
	case tok == scanner.Int:
		yylval.val, _ = strconv.Atoi(x.s.TokenText())
		return VAL
	}

	panic("unsupported token: " + x.s.TokenText())
}

type function struct{}

type env struct {
	bindings  map[string]int
	functions map[string]*function
	next      *env
}

func newEnv(next *env) *env {
	return &env{make(map[string]int), make(map[string]*function), next}
}

func (x *sillyLex) Error(message string) {
	panic(message)
}

type node interface {
	eval(e *env) int
}

type intNode struct {
	val int
}

func (n *intNode) eval(e *env) int {
	return n.val
}

type binopNode struct {
	op          string
	left, right node
}

func (n *binopNode) eval(e *env) int {
	ll := n.left.eval(e)
	rr := n.right.eval(e)
	switch n.op {
	case "+":
		return ll + rr
	case "-":
		return ll - rr
	case "/":
		return ll / rr
	case "*":
		return ll * rr
	}
	panic("unsupported op: " + n.op)
}

type assignNode struct {
	ident string
	right node
}

func (n *assignNode) eval(e *env) int {
	v := n.right.eval(e)
	e.bindings[n.ident] = v
	return v
}

type identNode struct {
	ident string
}

func (n *identNode) eval(e *env) int {
	for e != nil {
		if v, ok := e.bindings[n.ident]; ok {
			return v
		}
		e = e.next
	}
	return 0
}

type stmtList struct {
	c    node
	next *stmtList
}

func (n *stmtList) eval(e *env) int {
	v := 0
	for n != nil {
		v = n.c.eval(e)
		n = n.next
	}
	return v
}

var r node

func main() {
	var s scanner.Scanner
	s.Init(os.Stdin)
	sillyParse(&sillyLex{s})
	fmt.Println(r.eval(newEnv(nil)))
}
