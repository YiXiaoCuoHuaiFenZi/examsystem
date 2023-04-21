package models

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (manager *DBManager) GetRandomSingleChoice(qtype string, count int) ([]SingleChoice, error) {
	t := manager.GetSingleChoiceCollection()

	var ss []SingleChoice
	cursor, err := t.Find(context.TODO(), bson.M{"type": qtype})
	for cursor.Next(context.TODO()) {
		var sc = SingleChoice{}
		err = cursor.Decode(&sc)
		if err != nil {
			log.Println(err)
		}
		ss = append(ss, sc)
	}

	c := len(ss)
	if c < count {
		log.Println("随机数大于单选题库题目数")
		return nil, errors.New("随机数大于单选题库题目数")
	}

	results := []SingleChoice{}
	for i := 0; i < count; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		rn := r.Intn(len(ss))

		results = append(results, ss[rn])
		ss = append(ss[:rn], ss[rn+1:]...)
	}

	return results, err
}

func (manager *DBManager) GetRandomMultipleChoice(qtype string, count int) ([]MultipleChoice, error) {
	t := manager.GetMultipleChoiceCollection()

	var mc []MultipleChoice
	cursor, err := t.Find(context.TODO(), bson.M{"type": qtype})
	for cursor.Next(context.TODO()) {
		var m = MultipleChoice{}
		err = cursor.Decode(&m)
		if err != nil {
			log.Println(err)
		}
		mc = append(mc, m)
	}

	c := len(mc)
	if c < count {
		log.Println("随机数大于单选题库题目数")
		return nil, errors.New("随机数大于单选题库题目数")
	}

	var results []MultipleChoice
	for i := 0; i < count; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		rn := r.Intn(len(mc))

		results = append(results, mc[rn])
		mc = append(mc[:rn], mc[rn+1:]...)
	}

	return results, err
}

func (manager *DBManager) GetRandomTrueFalse(qtype string, count int) ([]TrueFalse, error) {
	t := manager.GetTrueFalseCollection()

	tf := []TrueFalse{}
	cursor, err := t.Find(context.TODO(), bson.M{"type": qtype})
	for cursor.Next(context.TODO()) {
		var f = TrueFalse{}
		err = cursor.Decode(&f)
		if err != nil {
			log.Println(err)
		}
		tf = append(tf, f)
	}

	c := len(tf)
	if c < count {
		log.Println("随机数大于单选题库题目数")
		return nil, errors.New("随机数大于单选题库题目数")
	}

	results := []TrueFalse{}
	for i := 0; i < count; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		rn := r.Intn(len(tf))

		results = append(results, tf[rn])
		tf = append(tf[:rn], tf[rn+1:]...)
	}

	return results, err
}

func (manager *DBManager) GetSingleChoiceByDescription(description string) ([]SingleChoice, error) {
	t := manager.GetSingleChoiceCollection()

	var ss []SingleChoice
	cursor, err := t.Find(context.TODO(), bson.M{"description": description})
	for cursor.Next(context.TODO()) {
		var sc = SingleChoice{}
		err = cursor.Decode(&sc)
		if err != nil {
			log.Println(err)
		}
		ss = append(ss, sc)
	}

	return ss, err
}

func (manager *DBManager) AddSingleChoice(s *SingleChoice) error {
	t := manager.GetSingleChoiceCollection()

	scs, err := manager.GetSingleChoiceByDescription(s.Description)
	if err != nil {
		return err
	}

	for _, v := range scs {
		if v.Type == s.Type &&
			v.Description == s.Description &&
			v.A == s.A &&
			v.B == s.B &&
			v.C == s.C &&
			v.D == s.D {
			return errors.New("新增失败，该题目已经存在")
		}
	}

	_, err = t.InsertOne(context.TODO(), s)
	if err != nil {
		return err
	}

	return nil
}

func (manager *DBManager) GetMultipleChoiceByDiscription(description string) ([]MultipleChoice, error) {
	t := manager.GetMultipleChoiceCollection()

	var ms []MultipleChoice
	cursor, err := t.Find(context.TODO(), bson.M{"description": description})
	for cursor.Next(context.TODO()) {
		var m = MultipleChoice{}
		err = cursor.Decode(&m)
		if err != nil {
			log.Println(err)
		}
		ms = append(ms, m)
	}

	return ms, err
}

func (manager *DBManager) AddMultipleChoice(m *MultipleChoice) error {
	t := manager.GetMultipleChoiceCollection()

	mcs, err := manager.GetMultipleChoiceByDiscription(m.Description)
	if err != nil {
		return err
	}

	for _, v := range mcs {
		if v.Type == m.Type &&
			v.Description == m.Description &&
			v.A == m.A &&
			v.B == m.B &&
			v.C == m.C &&
			v.D == m.D &&
			v.E == m.E &&
			v.F == m.F {
			return errors.New("新增失败，该题目已经存在")
		}
	}

	_, err = t.InsertOne(context.TODO(), m)
	if err != nil {
		return err
	}

	return nil
}

func (manager *DBManager) GetTrueFalseByDiscription(description string) ([]TrueFalse, error) {
	t := manager.GetTrueFalseCollection()

	var ts []TrueFalse
	cursor, err := t.Find(context.TODO(), bson.M{"description": description})
	for cursor.Next(context.TODO()) {
		var f = TrueFalse{}
		err = cursor.Decode(&f)
		if err != nil {
			log.Println(err)
		}
		ts = append(ts, f)
	}

	return ts, err
}

func (manager *DBManager) AddTrueFalse(f *TrueFalse) error {
	t := manager.GetTrueFalseCollection()

	tfs, err := manager.GetTrueFalseByDiscription(f.Description)
	if err != nil {
		return err
	}

	for _, v := range tfs {
		if v.Type == f.Type && v.Description == f.Description {
			return errors.New("新增失败，该题目已经存在")
		}
	}

	_, err = t.InsertOne(context.TODO(), f)
	if err != nil {
		return err
	}

	return nil
}

func (manager *DBManager) GetSingleChoiceSummary() (map[string]int, error) {
	t := manager.GetSingleChoiceCollection()

	var scs []SingleChoice
	cursor, err := t.Find(context.TODO(), bson.M{})
	for cursor.Next(context.TODO()) {
		var s = SingleChoice{}
		err = cursor.Decode(&s)
		if err != nil {
			log.Println(err)
		}
		scs = append(scs, s)
	}

	if err != nil {
		return nil, err
	}

	results := make(map[string]int)
	for _, sc := range scs {
		if v, ok := results[sc.Type]; ok {
			results[sc.Type] = v + 1
		} else {
			results[sc.Type] = 1
		}
	}

	return results, err
}

func (manager *DBManager) GetMultipleChoiceSummary() (map[string]int, error) {
	t := manager.GetMultipleChoiceCollection()

	var mcs []MultipleChoice
	cursor, err := t.Find(context.TODO(), bson.M{})
	for cursor.Next(context.TODO()) {
		var m = MultipleChoice{}
		err = cursor.Decode(&m)
		if err != nil {
			log.Println(err)
		}
		mcs = append(mcs, m)
	}

	results := make(map[string]int)
	for _, mc := range mcs {
		if v, ok := results[mc.Type]; ok {
			results[mc.Type] = v + 1
		} else {
			results[mc.Type] = 1
		}
	}

	return results, err
}

func (manager *DBManager) GetTrueFalseSummary() (map[string]int, error) {
	t := manager.GetTrueFalseCollection()

	var tfs []TrueFalse
	cursor, err := t.Find(context.TODO(), bson.M{})
	for cursor.Next(context.TODO()) {
		var f = TrueFalse{}
		err = cursor.Decode(&f)
		if err != nil {
			log.Println(err)
		}
		tfs = append(tfs, f)
	}

	results := make(map[string]int)
	for _, tf := range tfs {
		if v, ok := results[tf.Type]; ok {
			results[tf.Type] = v + 1
		} else {
			results[tf.Type] = 1
		}
	}

	return results, err
}
