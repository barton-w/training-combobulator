package main

import (
	"encoding/json"
	"log"
	"training-combobulator/config"
	"training-combobulator/dal"
	"training-combobulator/handlers"
)

type answers struct {
	Q1 uint32 `json:"q1"`
	Q2 uint32 `json:"q2"`
	Q3 string `json:"q3"`
	Q4 uint32 `json:"q4"`
}

func main() {
	// initialize config
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %s", err.Error())
	}
	log.Println("configuration initialized")

	// connect to the database
	dbClient, err := dal.Connect(cfg.Database)
	if err != nil {
		log.Fatalf("failed to connected to database: %s", err.Error())
	}
	log.Println("connected to database")

	// The intent is that handlers could be converted to http route handlers,
	// parsing client query data to fetch various results.
	// While these handlers do answer the questions in the assignment,
	// they are user and exercise agnostic, and can also provide time based aggregations

	// q1 - In total, how many pounds have these athletes Bench Pressed?
	answer1, err := handlers.GetExerciseTotalWeight(dbClient, "Bench Press")
	if err != nil {
		// don't fatal on these, continue through all answers
		log.Printf("error: %s", err.Error())
	}

	// q2 - How many pounds did Barry Moore Back Squat in 2016?
	aggregates, err := handlers.GetUserYearlyExerciseTotals(dbClient, "Barry", "Moore", "Back Squat")
	if err != nil {
		log.Printf("error: %s", err.Error())
	}

	yearTotals, err := handlers.GetTotalYearlyWeight(aggregates)
	if err != nil {
		log.Printf("error: %s", err.Error())
	}
	answer2 := yearTotals[2016]

	// q3 - In what month of 2017 did Barry Moore Back Squat the most total weight?
	maxMonths, err := handlers.GetMaxWeightMonths(aggregates)
	answer3 := maxMonths[2017]

	// q4 - What is Abby Smith's Bench Press PR weight?
	answer4, err := handlers.GetUserPR(dbClient, "Abby", "Smith", "Bench Press")

	// respond by logging answers in JSON
	ans := answers{
		Q1: answer1,
		Q2: answer2,
		Q3: answer3,
		Q4: answer4,
	}

	b, err := json.Marshal(ans)
	if err != nil {
		log.Fatalf("error marshalling answers to json: %s", err)
	}

	log.Printf("RESULTS OF THE TRAINING COMBOBULATOR:\n%s\n", string(b))
}
