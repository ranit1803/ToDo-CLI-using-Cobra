package utils

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/mergestat/timediff"
	"github.com/ranit1803/ToDo-CLI-using-Cobra/internal/models"
)

func PrintTasks(tasks []models.Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tTitle\tCompleted\tCreated At\t")

	for _, task := range tasks {
		status := "❌"
		if task.Completed {
			status = "✅"
		}
		 createdAgo := timediff.TimeDiff(task.CreatedAt)
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", task.ID, task.Title, status, createdAgo)
	}
	w.Flush()
}
