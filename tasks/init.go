package tasks

import (
	"fmt"

	"github.com/astaxie/beego/toolbox"
)

func init() {
	tk1 := toolbox.NewTask("tk1", "0 12 * * * *", func() error { fmt.Println("tk1"); return nil })

	err := tk.Run()
	if err != nil {
		t.Fatal(err)
	}

	toolbox.AddTask("tk1", tk1)
	toolbox.StartTask()
	defer toolbox.StopTask()

}
