function controll(id){
  var SingleChoiceType = document.getElementById("SingleChoiceType")
  var MultipleChoiceType = document.getElementById("MultipleChoiceType")
  var TrueFalseType = document.getElementById("TrueFalseType")

  var SingleChoiceForm = document.getElementById("SingleChoiceForm")
  var MultipleChoiceForm = document.getElementById("MultipleChoiceForm")
  var TrueFalseForm = document.getElementById("TrueFalseForm")

  switch(id)
  {
  case "SingleChoiceType":
    MultipleChoiceForm.style.display = "none";
    TrueFalseForm.style.display = "none";

	SingleChoiceForm.style.display = "";
    break;
  case "MultipleChoiceType":
    SingleChoiceForm.style.display = "none";
    TrueFalseForm.style.display = "none";
	 
	MultipleChoiceForm.style.display = "";
    break;
  case "TrueFalseType":
    SingleChoiceForm.style.display = "none";
    MultipleChoiceForm.style.display = "none";

	TrueFalseForm.style.display = "";
    break;
  default:
    break;
  }
}