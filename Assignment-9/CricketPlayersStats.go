package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Player struct {
	PlayerName    string
	Period        string
	MatchesPlayed int
	TotalScore    int
	AvgScore      string
}

func main() {
	file, err := os.Open("CricketPlayersStats.csv")
	if err != nil {
		fmt.Println("Error while opening the csv file:", err)
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error while reading csv file:", err)
		return
	}
	for _, record := range records {
		fmt.Println(record)
	}

	//add runs and matches played and update the data and write to the new file

	players := make([]Player, 0, len(records))
	for i, record := range records {
		if i == 0 {
			continue
		}
		matchesPlayed, _ := strconv.Atoi(record[2])
		totalScore, _ := strconv.Atoi(record[3])
		players = append(players, Player{
			PlayerName:    record[0],
			Period:        record[1],
			MatchesPlayed: matchesPlayed + 10,
			TotalScore:    totalScore + 100,
			AvgScore:      record[4],
		})
	}

	updatedFile, err := os.Create("UpdatesFile.csv")
	if err != nil {
		fmt.Println("Error while creating updated csv file:", err)
		return
	}
	defer updatedFile.Close()

	writer := csv.NewWriter(updatedFile)
	defer writer.Flush()

	header := []string{"PlayerName", "Period", "MatchedPlayed", "TotalScore", "AvgScore"}
	err = writer.Write(header)
	if err != nil {
		fmt.Println("Error while writing the header:", err)
		return
	}

	for _, player := range players {
		row := []string{
			player.PlayerName,
			player.Period,
			strconv.Itoa(player.MatchesPlayed),
			strconv.Itoa(player.TotalScore),
			player.AvgScore,
		}
		err = writer.Write(row)
		if err != nil {
			fmt.Println("error while writing players data to csv:", err)
			return
		}
	}

	fmt.Println("Data updated and saved successfully!")

}
