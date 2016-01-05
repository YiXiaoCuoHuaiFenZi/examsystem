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

// 检查输入是否是数字，可以是小数
function numCheck(id){
  var input =document.getElementById(id)
  if (isNaN(input.value)){
    alert("请输入数字");
	input.value="";
	}
}

// 检查输入是否是数字，只允许整数
//function onlyNum(){  
//  if(!(event.keyCode==46)&&!(event.keyCode==8)&&!(event.keyCode==37)&&!(event.keyCode==39)) 
//  if(!((event.keyCode>=48&&event.keyCode<=57)||(event.keyCode>=96&&event.keyCode<=105))) 
//  event.returnValue=false; 
//} 