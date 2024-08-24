package main

import (
	"log"

	"strconv"

	"github.com/gofiber/fiber/v3"
)

type Exercise struct {
	ID   string
	Name string
	Url  string
}

type ExerciseList struct {
	ID       string
	ListName string
	Img      string
	Details  string
	Video    string
}

func main() {
	app := fiber.New()
	app.Get("/ex", HandleGetExercises)
	app.Get("/elist", HandleGetExerciseList)

	log.Fatal(app.Listen(":3000"))

}

func HandleGetExercises(c fiber.Ctx) error {
	exercise := generateExerciseDummyData()
	return c.JSON(exercise)
}

func HandleGetExerciseList(c fiber.Ctx) error {
	exL := generateExerciseListDummyData()
	return c.JSON(exL)
}

func generateRandomExercise(id int) Exercise {
	idStr := strconv.Itoa(id)
  
	ex :=  Exercise{
		ID:   idStr,               // Serialized ID from 1 to 100
		Name: "Exercise " + idStr, // Name in the format "exercise X"
		Url:  "http://google.com", // Static URL
	}
    return ex
}

func generateRandomExerciseList(id int) ExerciseList {
	idStr := strconv.Itoa(id)
	exl := ExerciseList{
		ID:       idStr,                               // Serialized ID from 1 to 100
		ListName: "Exercise List" + idStr,             // ListName in the format "ExerciseList X"
		Img:      "http://google.com/img/" + idStr,    // Static base URL with ID appended
		Details:  "Details for Exercise List " + idStr, // Custom details for each list
		Video:    "http://google.com/video/" + idStr,  // Static base URL with ID appended
	}
    return exl
}

func generateExerciseDummyData() []Exercise {
	exercises := make([]Exercise, 100)
	for i := 1; i < 101; i++ {
		exercises[i] = generateRandomExercise(i)

	}

	return exercises
}

func generateExerciseListDummyData() []ExerciseList {
	exerciseLists := make([]ExerciseList, 100)
	for i := 0; i < 101; i++ {
		exerciseLists[i] = generateRandomExerciseList(i)
	}

	return exerciseLists
}
