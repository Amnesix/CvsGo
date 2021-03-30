#!/usr/bin/bash
cc -o solvec -O2 solvec.c
go build solvego.go
echo "Exécution en c"
time ./solvec
echo "Exécution en go"
time ./solvego
