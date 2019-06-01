package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {

	questionFile := flag.String("questionFile","problems.csv","The csv file containing all the problems")
	timeLimit:= flag.Int("timeLimit",20,"Time Limit, default is 20")
	flag.Parse()

	file,err := os.Open(*questionFile)

	if err != nil {
		exit(fmt.Sprintf("Cannot open %s \n",*questionFile))
	}

	r := csv.NewReader(file)
	questionsList,err := r.ReadAll()

 	if err != nil {
		exit(fmt.Sprintf("Cannot parse %s \n",*questionFile))
	}

	correct := 0
	questionsStorted:= csvParser(questionsList)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	for i , p := range questionsStorted{
		fmt.Printf("Question %d.) %s = ",i+1,p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select{
		case <-timer.C:
			fmt.Println()
			fmt.Printf("You got %d correct out of %d \n",correct,len(questionsStorted))
			return

		case answer := <-answerCh:
			if answer == p.a {
				correct ++
			}

		}


	}

}

type questions struct{
	q string
	a string
}

func csvParser(lines [][]string) []questions {
	ret := make([]questions, len(lines))
	for i,line := range lines{
		ret[i] = questions{
			q : line[0],
			a : line[1],
		}
	}
	return ret
}


func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)


}