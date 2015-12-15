//function timer0() {
//	   var today = new Date();
//	   var h = today.getHours();
//       var m = today.getMinutes();
//       var s = today.getSeconds();
//       m = checkTime(m);
//       s = checkTime(s);
//       document.getElementById('timer').innerHTML = "考试倒计时："+ h + ":" + m + ":" + s;
//       var t = setTimeout(timer0, 500);
//	}
	
//function checkTime(i) {
//       if (i < 10) {i = "0" + i};  // add zero in front of numbers < 10
//       return i;
//}
 
//function timer1(){
//	var ts = (new Date(2018, 11, 11, 9, 0, 0)) - (new Date());//计算剩余的毫秒数
//	var dd = parseInt(ts / 1000 / 60 / 60 / 24, 10);//计算剩余的天数
//	var hh = parseInt(ts / 1000 / 60 / 60 % 24, 10);//计算剩余的小时数
//	var mm = parseInt(ts / 1000 / 60 % 60, 10);//计算剩余的分钟数
//	var ss = parseInt(ts / 1000 % 60, 10);//计算剩余的秒数
//	dd = checkTime(dd);
//	hh = checkTime(hh);
//	mm = checkTime(mm);
//	ss = checkTime(ss);
//	document.getElementById("timer").innerHTML = dd + "天" + hh + "时" + mm + "分" + ss + "秒";
//	setInterval("timer1()",1000);
//}

//function checkTime(i) {
//       if (i < 10) {i = "0" + i};  // add zero in front of numbers < 10
//       return i;
//} 

var ms = 5400000 // 剩余的毫秒数
function timer(){		
	var hh = parseInt(ms / 1000 / 60 / 60 % 24, 10);//计算剩余的小时数
	var mm = parseInt(ms / 1000 / 60 % 60, 10);//计算剩余的分钟数
	var ss = parseInt(ms / 1000 % 60, 10);//计算剩余的秒数	
	h = checkTime(hh);
	m = checkTime(mm);
	s = checkTime(ss);
	ms = ms - 1000;	
	document.getElementById("timer").innerHTML = "考试倒计时：" + h + ":" + m + ":" + s;	
	var t = setTimeout(timer, 1000)
}			

function checkTime(i) {
       if (i < 10) {i = "0" + i};  // add zero in front of numbers < 10
       return i;
}
