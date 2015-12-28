function controll(id){
  var QuestionTypeOne = document.getElementById("QuestionTypeOne")
  var QuestionTypeTwo = document.getElementById("QuestionTypeTwo")
  var QuestionTypeThree = document.getElementById("QuestionTypeThree")

  var SingleAnswer = document.getElementById("SingleAnswer")
  var MultipleAnswer = document.getElementById("MultipleAnswer")
  var JudgeAnswer = document.getElementById("JudgeAnswer")

  var OptionA = document.getElementById("OptionA")
  var OptionB = document.getElementById("OptionB")
  var OptionC = document.getElementById("OptionC")
  var OptionD = document.getElementById("OptionD")
  var OptionE = document.getElementById("OptionE")
  var OptionF = document.getElementById("OptionF")

  switch(id)
  {
  case "QuestionTypeOne":
    MultipleAnswer.style.display = "none";
    JudgeAnswer.style.display = "none";
    OptionE.style.display = "none";
    OptionF.style.display = "none";
	
	OptionA.style.display = "";
	OptionB.style.display = "";
	OptionC.style.display = "";
	OptionD.style.display = "";
	SingleAnswer.style.display = "";
    break;
  case "QuestionTypeTwo":
    SingleAnswer.style.display = "none";
    JudgeAnswer.style.display = "none";
	 
	OptionA.style.display = "";
	OptionB.style.display = "";
	OptionC.style.display = "";
	OptionD.style.display = "";
	OptionE.style.display = "";
	OptionF.style.display = "";
	MultipleAnswer.style.display = "";
    break;
  case "QuestionTypeThree":
    SingleAnswer.style.display = "none";
    MultipleAnswer.style.display = "none";
    OptionA.style.display = "none";
    OptionB.style.display = "none";
    OptionC.style.display = "none";
    OptionD.style.display = "none";
    OptionE.style.display = "none";
    OptionF.style.display = "none";
	
	JudgeAnswer.style.display = "";
    break;
  default:
    break;
  }
}