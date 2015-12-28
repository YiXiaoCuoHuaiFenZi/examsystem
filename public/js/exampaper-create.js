function controll(id){
  var ExamPaperCreateMethodRandom = document.getElementById("ExamPaperCreateMethodRandom")
  var ExamPaperCreateMethodSelect = document.getElementById("ExamPaperCreateMethodSelect")

  var ExamPaperCreateRandomInputs = document.getElementById("ExamPaperCreateRandomInputs")
  var ExamPaperCreateSelectTable = document.getElementById("ExamPaperCreateSelectTable")

  switch(id)
  {
  case "ExamPaperCreateMethodRandom":
    ExamPaperCreateSelectTable.style.display = "none";
	
	ExamPaperCreateRandomInputs.style.display = "";
    break;
  case "ExamPaperCreateMethodSelect":
    ExamPaperCreateRandomInputs.style.display = "none"; 
	
	ExamPaperCreateSelectTable.style.display = "";
    break;
  default:
    break;
  }
}