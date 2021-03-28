#include <stdio.h>
#include <stdlib.h>

#define tst294 "....12.8.5..8...........73..85......6..9.3....2..6..4..6..2....7..3.9..2..2..19.."
#define tst295 "....14....3....2...7..........9...3.6.1.............8.2.....1.4....5.6.....7.8..."
#define CASE_VIDE '.'

static char sudoku[9][9];

int isOk(int r, int c, int val)
{
    int i, j;
    for (i=0; i<9; i++)
        if (sudoku[i][c] == val)
            return 0;
    for (j=0; j<9; j++)
        if (sudoku[r][j] == val)
            return 0;
    for (i=3*(r/3); i<3*(1+r/3); i++)
        for (j=3*(c/3); j<3*(1+c/3); j++)
                if (sudoku[i][j] == val)
                    return 0;
    return 1;
}

int chercheVide(int *x, int *y)
{
    int r, c;
    for (c=0; c<9; c++)
        for (r=0; r<9; r++)
            if (sudoku[r][c] == CASE_VIDE)
            {
                *x=r;
                *y=c;
                return 1;
            }
    return 0;
}

int Solve()
{
    char i;
    int r=0, c=0;
    if (!chercheVide(&r, &c))
        return 1;
    for (i='1'; i<='9'; i++)
    {
        if (isOk(r, c, i))
        {
            sudoku[r][c] = i;
            if (Solve())
                return 1;
            sudoku[r][c] = CASE_VIDE;
        }
    }
    return 0;
}

void affiche()
{
    for (int i=0; i<9; i++) {
        for (int j=0; j<9; j++) {
            printf("%c ", sudoku[i][j]);
        }
        printf("\n");
    }
}

int main()
{
    for (int i=0; i<9; i++) {
        for (int j=0; j<9; j++) {
            sudoku[i][j] = tst295[i*9+j];
        }
    }
    affiche();
    printf("\n");
    if (Solve())
        affiche();
    return 0;
}
