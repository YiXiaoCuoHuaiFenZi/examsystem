package models

import (
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
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

func ParseExamPaperFile(exampaperFilePath string) (ExamPaper, string, string, string, error) {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	exp := ExamPaper{}
	scFilePath := dir + "/scs.txt"
	mcFilePath := dir + "/mcs.txt"
	tfFilePath := dir + "/tfs.txt"

	f, err := os.Open(exampaperFilePath)
	scf, err := os.Create(scFilePath)
	if err != nil {
		log.Println(err)
		return exp, "", "", "", err
	}
	defer scf.Close()
	mcf, err := os.Create(mcFilePath)
	if err != nil {
		log.Println(err)
		return exp, "", "", "", err
	}
	defer mcf.Close()
	tff, err := os.Create(tfFilePath)
	if err != nil {
		log.Println(err)
		return exp, "", "", "", err
	}
	defer tff.Close()

	r := bufio.NewReader(f)
	head, sc, mc, tf := false, false, false, false
	for {
		line := make([]byte, 1024, 1024)
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		l := string(line)
		if l == "********************试卷信息********************" {
			head, sc, mc, tf = true, false, false, false
		} else if l == "********************单选题********************" {
			head, sc, mc, tf = false, true, false, false
		} else if l == "********************多选题********************" {
			head, sc, mc, tf = false, false, true, false
		} else if l == "********************判断题********************" {
			head, sc, mc, tf = false, false, false, true
		}

		if head {
			l := strings.TrimSpace(string(line))
			if strings.HasPrefix(l, "试卷标题：") {
				exp.Title = strings.TrimPrefix(l, "试卷标题：")
			}
			if strings.HasPrefix(l, "试卷描述：") {
				exp.Discription = strings.TrimPrefix(l, "试卷描述：")
			}
			if strings.HasPrefix(l, "总分值：") {
				i, err := strconv.Atoi(strings.Trim(strings.TrimPrefix(l, "总分值："), "分"))
				if err != nil {
					log.Println(err)
					return exp, "", "", "", err
				}
				exp.Score = i
			}
			if strings.HasPrefix(l, "考试时间：") {
				i, err := strconv.Atoi(strings.Trim(strings.TrimPrefix(l, "考试时间："), "分钟"))
				if err != nil {
					log.Println(err)
					return exp, "", "", "", err
				}
				exp.Time = i
			}
			if strings.HasPrefix(l, "单选题每题分值：") {
				i, err := strconv.Atoi(strings.Trim(strings.TrimPrefix(l, "单选题每题分值："), "分"))
				if err != nil {
					log.Println(err)
					return exp, "", "", "", err
				}
				exp.SCScore = float32(i)
			}
			if strings.HasPrefix(l, "多选题每题分值：") {
				i, err := strconv.Atoi(strings.Trim(strings.TrimPrefix(l, "多选题每题分值："), "分"))
				if err != nil {
					log.Println(err)
					return exp, "", "", "", err
				}
				exp.MCScore = float32(i)
			}
			if strings.HasPrefix(l, "判断题每题分值：") {
				i, err := strconv.Atoi(strings.Trim(strings.TrimPrefix(l, "判断题每题分值："), "分"))
				if err != nil {
					log.Println(err)
					return exp, "", "", "", err
				}
				exp.TFScore = float32(i)
			}
		}
		if sc {
			n, err := scf.WriteString(l + "\r\n")
			if err != nil {
				log.Println(n, err)
				return exp, "", "", "", err
			}
		}
		if mc {
			n, err := mcf.WriteString(l + "\r\n")
			if err != nil {
				log.Println(n, err)
				return exp, "", "", "", err
			}
		}
		if tf {
			n, err := tff.WriteString(l + "\r\n")
			if err != nil {
				log.Println(n, err)
				return exp, "", "", "", err
			}
		}
	}
	return exp, scFilePath, mcFilePath, tfFilePath, nil
}
