# CvsGo
J'ai eu la surprise en réécrivant une application de Sudoku en Go de remarque
que l'algorithme de recherche de la solution semblait vraiment beaucoup plus
rapide qu'en C. Je me suis demandé si c'était une impression ou non. Du coup,
j'ai extrait juste la partie de recherche de la solution dans deux sources,
l'un en C, l'autre en Go.

Pour le C :

> $ cc -o solvec solvec.c -O3<br>
> $ time solvec

Pour le Go :

> $ go build solvego.go<br>
> $ time solvego

Le résultat sur mon PC : **~150ms en C conte ~15ms en Go !!!** Soit 10× plus
rapide en go !

Y-a-t'il un autre niveau d'optimisation pour le C que j'aurais raté ?
