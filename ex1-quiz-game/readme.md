TODO Part1:
1. Create a program that will read in a Quiz provided via a CSV file and then
   present it to the user to play
2. The quiz needs to keep track of how many questions the user got right/wrong
3. The quiz should proceed to the next question whether the answer was right or 
    wrong.
4. The CSV file should default to problems.csv (provided) but should be customizable 
    by passing a filename via a flag
5. The CSV file will be in a two column format. Column 1 will be the question, and
    then column two in the same row will be the answer
6. Can assume questions will have single word/number answers
7. At the end of the quiz the program should print out the total number of correct
    questions and how many questions there were in total


TODO Part2:
7. Add a timer that defaults to 30 seconds and is configurable by a cli flag
8. The quiz should hard stop once the timer ends
9. Users should be asked to press enter to begin the quiz and start the timer
10. BONUS: Add a cli flag to randomize the contents of the CSV file