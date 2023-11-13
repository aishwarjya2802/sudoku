# Sudoku

The program reads a Sudoku puzzle from command-line arguments, validates the input, solves the puzzle and prints the solution or an error message.

## Usage

To test the program, run the below-mentioned commands in your CLI
```bash
cd ~/
git clone https://github.com/aishwarjya2802/sudoku.git
cd sudoku
go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e
```
or the `tester.sh` file can be used to test the program.
```bash
bash tester.sh
```
Expected output:
```code
3 9 6 2 4 5 7 8 1$
1 7 8 3 6 9 5 2 4$
5 2 4 8 1 7 3 9 6$
2 8 7 9 5 1 6 4 3$
9 3 1 4 8 6 2 7 5$
4 6 5 7 2 3 9 1 8$
7 1 2 6 3 8 4 5 9$
6 5 9 1 7 4 8 3 2$
8 4 3 5 9 2 1 6 7$
```
For checking other inputs that produce errors the `errortester.sh` file can be used
```bash
bash errortester.sh
```
Expected output:
```code
Error$
Error$
Error$
```
