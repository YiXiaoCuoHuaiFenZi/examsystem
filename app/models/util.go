package models

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func ParseSingleChoiceFile(file *os.File, qType string) ([]SingleChoice, error) {
	r := bufio.NewReader(file)

	scs := []SingleChoice{}
	sc := SingleChoice{}
	for {
		line := make([]byte, 1024, 1024)
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}

		l := strings.TrimSpace(string(line))

		if strings.HasPrefix(l, "题目：") {
			sc.Discription = strings.TrimPrefix(l, "题目：")
		}
		if strings.HasPrefix(l, "A.") {
			sc.A = strings.TrimPrefix(l, "A.")
		}
		if strings.HasPrefix(l, "B.") {
			sc.B = strings.TrimPrefix(l, "B.")
		}
		if strings.HasPrefix(l, "C.") {
			sc.C = strings.TrimPrefix(l, "C.")
		}
		if strings.HasPrefix(l, "D.") {
			sc.D = strings.TrimPrefix(l, "D.")
		}
		if strings.HasPrefix(l, "答案：") {
			sc.Answer = strings.TrimPrefix(l, "答案：")
			switch strings.TrimSpace(sc.Answer) {
			case "A":
				sc.Answer = sc.A
			case "B":
				sc.Answer = sc.B
			case "C":
				sc.Answer = sc.C
			case "D":
				sc.Answer = sc.D
			default:
				break
			}
			sc.Type = qType
			scs = append(scs, sc)
		}
	}
	return scs, nil
}

func ParseMultipleChoiceFile(file *os.File, qType string) ([]MultipleChoice, error) {
	r := bufio.NewReader(file)

	mcs := []MultipleChoice{}
	mc := MultipleChoice{}
	for {
		line := make([]byte, 1024, 1024)
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}

		l := strings.TrimSpace(string(line))

		if strings.HasPrefix(l, "题目：") {
			mc.Discription = strings.TrimPrefix(l, "题目：")
		}
		if strings.HasPrefix(l, "A.") {
			mc.A = strings.TrimPrefix(l, "A.")
		}
		if strings.HasPrefix(l, "B.") {
			mc.B = strings.TrimPrefix(l, "B.")
		}
		if strings.HasPrefix(l, "C.") {
			mc.C = strings.TrimPrefix(l, "C.")
		}
		if strings.HasPrefix(l, "D.") {
			mc.D = strings.TrimPrefix(l, "D.")
		}
		if strings.HasPrefix(l, "E.") {
			mc.D = strings.TrimPrefix(l, "E.")
		}
		if strings.HasPrefix(l, "F.") {
			mc.D = strings.TrimPrefix(l, "F.")
		}
		if strings.HasPrefix(l, "答案：") {
			answers := strings.Split(strings.TrimPrefix(l, "答案："), ",")
			var as []string
			for _, a := range answers {
				switch strings.TrimSpace(a) {
				case "A":
					as = append(as, mc.A)
				case "B":
					as = append(as, mc.B)
				case "C":
					as = append(as, mc.C)
				case "D":
					as = append(as, mc.D)
				case "E":
					as = append(as, mc.E)
				case "F":
					as = append(as, mc.F)
				default:
					break
				}
			}
			mc.Answer = as
			mc.Type = qType
			mcs = append(mcs, mc)
		}
	}
	return mcs, nil
}

func ParseTrueFalseFile(file *os.File, qType string) ([]TrueFalse, error) {
	r := bufio.NewReader(file)

	tfs := []TrueFalse{}
	tf := TrueFalse{}
	for {
		line := make([]byte, 1024, 1024)
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}

		l := strings.TrimSpace(string(line))
		if strings.HasPrefix(l, "题目：") {
			tf.Discription = strings.TrimPrefix(l, "题目：")
		}
		if strings.HasPrefix(l, "答案：") {
			tf.Answer = strings.TrimPrefix(l, "答案：")
			tf.Type = qType
			tfs = append(tfs, tf)
		}
	}
	return tfs, nil
}
