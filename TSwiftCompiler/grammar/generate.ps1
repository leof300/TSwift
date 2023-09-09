<#
antlr4-parse TSParser_rules.g4 start -gui input.txt
antlr4 -o ../visitor -visitor -no-listener -package visitor -Dlanguage=Go -XdbgST -Xlog *.g4
 #>

cd C:/Users/erikf/OneDrive/Documentos/GitHub/TSwift/TSwiftCompiler/grammar/
antlr4-parse TSParser_rules.g4 start -gui ./ArchivosPrueba/Basicos/Intermedias.swift

