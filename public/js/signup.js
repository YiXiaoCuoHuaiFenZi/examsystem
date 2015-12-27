function controll(id){	
  var SignUpTypeOne = document.getElementById("SignUpTypeOne")
  var SignUpTypeTwo = document.getElementById("SignUpTypeTwo")

  var ManualAdd = document.getElementById("ManualAdd")
  var BatchAdd = document.getElementById("BatchAdd")

  switch(id)
  {
  case "SignUpTypeOne":  
    BatchAdd.style.display = "none";
	
	ManualAdd.style.display = "";    
    ManualAdd.style.visibility = "visible"; 	    
    break;
  case "SignUpTypeTwo":
    ManualAdd.style.display = "none"; 
	
	BatchAdd.style.display = "";    
    BatchAdd.style.visibility = "visible";
    break;
  default:
    break;
  }
}