function controll(id){
  var SignUpTypeOne = document.getElementById("SignUpTypeOne")
  var SignUpTypeTwo = document.getElementById("SignUpTypeTwo")

  var ManualAdd = document.getElementById("ManualAdd")
  var BatchAdd = document.getElementById("BatchSignUp")

  switch(id)
  {
  case "SignUpTypeOne":  
    BatchAdd.style.display = "none";
	
	ManualAdd.style.display = "";
    break;
  case "SignUpTypeTwo":
    ManualAdd.style.display = "none"; 
	
	BatchAdd.style.display = "";
    break;
  default:
    break;
  }
}