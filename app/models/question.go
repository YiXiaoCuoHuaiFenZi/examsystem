package models

import (
	"errors"

	"gopkg.in/mgo.v2/bson"
)

func (this *DBManager) GetSingleChoiceByDiscription(discription string) ([]SingleChoice, error) {
	t := this.session.DB(DBName).C(SingleChoiceCollection)

	ss := []SingleChoice{}
	err := t.Find(bson.M{"discription": discription}).All(&ss)

	return ss, err
}

func (this *DBManager) AddSingleChoice(s *SingleChoice) error {
	t := this.session.DB(DBName).C(SingleChoiceCollection)

	scs, err := this.GetSingleChoiceByDiscription(s.Discription)
	if err != nil {
		return err
	}

	for _, v := range scs {
		if v.Type == s.Type &&
			v.Discription == s.Discription &&
			v.A == s.A &&
			v.B == s.B &&
			v.C == s.C &&
			v.D == s.D {
			return errors.New("新增失败，该题目已经存在")
		}
	}

	err = t.Insert(s)
	if err != nil {
		return err
	}

	return nil
}

func (this *DBManager) GetMultipleChoiceByDiscription(discription string) ([]MultipleChoice, error) {
	t := this.session.DB(DBName).C(MultipleChoiceCollection)

	ms := []MultipleChoice{}
	err := t.Find(bson.M{"discription": discription}).All(&ms)

	return ms, err
}

func (this *DBManager) AddMultipleChoice(m *MultipleChoice) error {
	t := this.session.DB(DBName).C(MultipleChoiceCollection)

	mcs, err := this.GetMultipleChoiceByDiscription(m.Discription)
	if err != nil {
		return err
	}

	for _, v := range mcs {
		if v.Type == m.Type &&
			v.Discription == m.Discription &&
			v.A == m.A &&
			v.B == m.B &&
			v.C == m.C &&
			v.D == m.D &&
			v.E == m.E &&
			v.F == m.F {
			return errors.New("新增失败，该题目已经存在")
		}
	}

	err = t.Insert(m)
	if err != nil {
		return err
	}

	return nil
}

func (this *DBManager) GetTrueFalseByDiscription(discription string) ([]TrueFalse, error) {
	t := this.session.DB(DBName).C(TrueFalseCollection)

	ts := []TrueFalse{}
	err := t.Find(bson.M{"discription": discription}).All(&ts)

	return ts, err
}

func (this *DBManager) AddTrueFalse(f *TrueFalse) error {
	t := this.session.DB(DBName).C(TrueFalseCollection)

	tfs, err := this.GetTrueFalseByDiscription(f.Discription)
	if err != nil {
		return err
	}

	for _, v := range tfs {
		if v.Type == f.Type && v.Discription == f.Discription {
			return errors.New("新增失败，该题目已经存在")
		}
	}

	err = t.Insert(f)
	if err != nil {
		return err
	}

	return nil
}
