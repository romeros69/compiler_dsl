run:
	go build && ./compiler_dsl

lex:
	golex -o lexer.go lexer.l

yacc:
	goyacc -o parser.go pasrser.y

clear:
	rm compiler_dsl