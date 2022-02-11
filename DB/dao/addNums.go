package dao

import (
	"DB/model"
	"DB/predict"
	"fmt"
)

func AddNums(a model.Staus) {
	fmt.Print(a)
	go predict.Predict(a)
	if _, ok := model.Node[a.Id]; ok {
		model.Node[a.Id] = append(model.Node[a.Id], a)
	} else {

		model.Node[a.Id] = append(model.Node[a.Id], a)
	}

}
