<#
antlr4-parse TSParser_rules.g4 start -gui input.txt
antlr4 -o ../visitor -visitor -no-listener -package visitor -Dlanguage=Go -XdbgST -Xlog *.g4
C:/Users/erikf/OneDrive/Documentos/GitHub/TSwift/TSwiftCompiler/grammar/

cd C:/Users/erikf/OneDrive/Documentos/GitHub/TSwift/TSwiftCompiler/grammar/
antlr4-parse TSParser_rules.g4 start -gui ./ArchivosPrueba/Arreglos/Vectores.swift

 #>

param (
    [string]$pathToChangeTo
)

cd $pathToChangeTo
antlr4-parse ../grammar/TSParser_rules.g4 start -gui ./tempfiles/temp.swift

