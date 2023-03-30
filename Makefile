run:
	go build && ./compiler_dsl

lex:
	golex -o lexer.go lexer.l

yacc:
	goyacc -o parser.go parser.y

clear:
	rm compiler_dsl