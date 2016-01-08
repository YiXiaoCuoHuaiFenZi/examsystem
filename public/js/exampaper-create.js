function createMethodControll(id) {
    var ExamPaperCreateMethodRandom = document.getElementById("ExamPaperCreateMethodRandom")
    var ExamPaperCreateMethodUpload = document.getElementById("ExamPaperCreateMethodUpload")

    var ExamPaperRandomForm = document.getElementById("ExamPaperRandomForm")
    var ExamPaperUploadForm = document.getElementById("ExamPaperUploadForm")

    switch (id) {
        case "ExamPaperCreateMethodRandom":
            ExamPaperUploadForm.style.display = "none";

            ExamPaperRandomForm.style.display = "";
            break;
        case "ExamPaperCreateMethodUpload":
            ExamPaperRandomForm.style.display = "none";

            ExamPaperUploadForm.style.display = "";
            break;
        default:
            break;
    }
}

// 检查输入是否是数字，可以是小数
function numCheck(id) {
    var input = document.getElementById(id)
    if (isNaN(input.value)) {
        alert("请输入数字");
        input.value = "";
    }

}

function checkScore() {
    var totalScore = document.getElementById("totalScore")
    var examPaper_Score = document.getElementById("examPaper_Score")
    var examPaper_SCCount = document.getElementById("examPaper_SCCount")
    var examPaper_SCScore = document.getElementById("examPaper_SCScore")
    var examPaper_MCCount = document.getElementById("examPaper_MCCount")
    var examPaper_MCScore = document.getElementById("examPaper_MCScore")
    var examPaper_TFCount = document.getElementById("examPaper_TFCount")
    var examPaper_TFScore = document.getElementById("examPaper_TFScore")

    totalScore = Number(examPaper_SCCount.value) * Number(examPaper_SCScore.value) +
        Number(examPaper_MCCount.value) * Number(examPaper_MCScore.value) +
        Number(examPaper_TFCount.value) * Number(examPaper_TFScore.value);

    if (totalScore != Number(examPaper_Score.value)) {
        alert("预设试卷总分数：" + examPaper_Score.value + "\n实际题目总分数：" +
            totalScore + "\n预设试卷总分数与实际题目分数总和不一致，请重新设置");
    } else {
        alert("预设试卷总分数：" + examPaper_Score.value + "\n实际题目总分数：" +
            totalScore + "\n预设试卷总分数与实际题目分数总和一致")
    }
}

// 检查输入是否是数字，只允许整数
//function onlyNum(){  
//  if(!(event.keyCode==46)&&!(event.keyCode==8)&&!(event.keyCode==37)&&!(event.keyCode==39)) 
//  if(!((event.keyCode>=48&&event.keyCode<=57)||(event.keyCode>=96&&event.keyCode<=105))) 
//  event.returnValue=false; 
//}
