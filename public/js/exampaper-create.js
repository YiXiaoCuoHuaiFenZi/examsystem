function controll(id){	
  var questionRadio1 = document.getElementById("QuestionTypeOne")
  var questionRadio2 = document.getElementById("QuestionTypeTwo")
  var questionRadio3 = document.getElementById("QuestionTypeThree")

  var singleAnswer = document.getElementById("SingleAnswer")
  var multipleAnswer = document.getElementById("MultipleAnswer")
  var judgeAnswer = document.getElementById("JudgeAnswer")

  var optionA = document.getElementById("OptionA")
  var optionB = document.getElementById("OptionB")
  var optionC = document.getElementById("OptionC")
  var optionD = document.getElementById("OptionD")
  var optionE = document.getElementById("OptionE")
  var optionF = document.getElementById("OptionF")

  switch(id)
  {
  case "QuestionTypeOne":  
    multipleAnswer.style.display = "none";    
    judgeAnswer.style.display = "none"; 
	optionE.style.display = "none";    
    optionF.style.display = "none"; 
	
	optionA.style.display = "";    
    optionA.style.visibility = "visible"; 
	optionB.style.display = "";    
    optionB.style.visibility = "visible"; 	
	optionC.style.display = "";    
    optionC.style.visibility = "visible"; 
	OptionD.style.display = "";    
    OptionD.style.visibility = "visible"; 
	singleAnswer.style.display = "";
    singleAnswer.style.visibility = "visible";    
    break;
  case "QuestionTypeTwo":
    singleAnswer.style.display = "none";    
    judgeAnswer.style.display = "none";
	 
	optionA.style.display = "";    
    optionA.style.visibility = "visible"; 
	optionB.style.display = "";    
    optionB.style.visibility = "visible"; 	
	optionC.style.display = "";    
    optionC.style.visibility = "visible"; 
	OptionD.style.display = "";    
    OptionD.style.visibility = "visible"; 	
	optionE.style.display = "";    
    optionE.style.visibility = "visible"; 
	optionF.style.display = "";    
    optionF.style.visibility = "visible"; 	
	multipleAnswer.style.display = "";
    multipleAnswer.style.visibility = "visible";	  
    break;
  case "QuestionTypeThree": 
    singleAnswer.style.display = "none";    
    multipleAnswer.style.display = "none";
	optionA.style.display = "none";    
    optionB.style.display = "none";
	optionC.style.display = "none";    
    OptionD.style.display = "none";	
	optionE.style.display = "none";    
    optionF.style.display = "none";  
	
	judgeAnswer.style.display = "";
    judgeAnswer.style.visibility = "visible";
    break;
  default:
    break;
  }
}