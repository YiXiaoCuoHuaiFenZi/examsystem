var ms = 5400000 // 剩余的毫秒数
function timer() {
    var hh = parseInt(ms / 1000 / 60 / 60 % 24, 10); //计算剩余的小时数
    var mm = parseInt(ms / 1000 / 60 % 60, 10); //计算剩余的分钟数
    var ss = parseInt(ms / 1000 % 60, 10); //计算剩余的秒数    
    h = checkTime(hh);
    m = checkTime(mm);
    s = checkTime(ss);
    ms = ms - 1000;
    document.getElementById("timer").innerHTML = "倒计时间：" + h + ":" + m + ":" + s;
    var t = setTimeout(timer, 1000)
}

function checkTime(i) {
    if (i < 10) {
        i = "0" + i
    }; // add zero in front of numbers < 10
    return i;
}

function SubmitExamPaper() {
    var result = "";
    $.ajax({
        cache: true,
        type: "POST",
        url: "/Examinee/PostExam",
        data: $('#ExamForm').serialize(), // 你的formid
        async: false,
        error: function(request) {
            alert("Connection error");
        },
        success: function(data) {}
    });
}

loseFocusTimes = 0
window.onblur = function() {
    loseFocusTimes += 1;
    if (loseFocusTimes > 3) {
        SubmitExamPaper();
        window.location.reload();
        alert("你切屏超过3次，视为作弊，已经自动交卷！");
    } else {
        alert("第" + loseFocusTimes + "次离开考场，超过3次将视为作弊，自动交卷！");
    }
}

// 鼠标移动事件，如果移出考试区域则弹出提示
// 响应速度太慢，不满足要求，暂时停用
// var iPreX;
// var iPreY;
// window.onmousemove = function(e) {
//     e = e || window.event;
//     var x = e.clientX;
//     var y = e.clientY;
//     var w = document.body.clientWidth;
//     var h = document.body.clientHeight;

//     // if (x > 0 && x <= w && y > 0 && y <= h) {
//     //     window.focus();
//     //     return false;
//     // }

//     flag = false
//     // 往左移出
//     if (x < 0 && x < iPreX) {
//         flag = true;

//     }
//     // 往上移出
//     if (y < 0 && y < iPreY) {
//         flag = true;

//     }
//     // 往右移出
//     if (x > iPreX && x > w - 2) {
//         flag = true;

//     }
//     // 往下移出
//     if (y > iPreY && y > h - 2) {
//         flag = true;

//     }
//     if (flag) {
//         alert("请勿离开考场，否则将视为作弊，自动交卷！");
//     }

//     iPreX = x;
//     iPreY = y;
// }

window.onkeydown = function() { //屏蔽鼠标右键、Ctrl+n、shift+F10、F5刷新、退格键
    //e = e || window.event; // 兼容IE和FF
    e = event ? event : (window.event ? window.event : null); // ie firefox都可以使用的事件
    if ((e.altKey) &&
        ((e.keyCode == 37) || //屏蔽   Alt+   方向键   ←  
            (e.keyCode == 39))) { //屏蔽   Alt+   方向键   →  
        alert("已禁用ALT+方向键快捷键前进或后退网页，请认真答题！");
        e.returnValue = false;
    }
    if ((e.keyCode == 8) || //屏蔽退格删除键  
        (e.keyCode == 116) || //屏蔽   F5   刷新键  
        (e.keyCode == 112) || //屏蔽   F1   刷新键  
        (e.ctrlKey && e.keyCode == 82)) { //Ctrl   +   R  
        e.keyCode = 0;
        alert("已经禁用刷新功能，请认真答题！");
        e.returnValue = false;
    }
    if ((e.ctrlKey) && (e.keyCode == 78)) { //屏蔽   Ctrl+n  
        alert("请勿开启新的浏览器窗口，认真答题，否则将视为作弊，自动交卷！")
        e.returnValue = false;
    }

    if ((e.shiftKey) && (e.keyCode == 121)) { //屏蔽   shift+F10 
        alert("请认真答题，否则将视为作弊，自动交卷！");
        e.returnValue = false;
    }

    if (e.srcElement.tagName == "A" && e.shiftKey) {
        alert("请认真答题，否则将视为作弊，自动交卷！");
        e.returnValue = false; //屏蔽   shift   加鼠标左键新开一网页  
    }

    if ((e.altKey) && (e.keyCode == 115)) { //屏蔽Alt+F4  
        window.showModelessDialog("about:blank", "", "dialogWidth:1px;dialogheight:1px");
        return false;
    }
    if ((e.altkey) && (e.keyCode == 27)) {
        alert("请认真答题，否则将视为作弊，自动交卷！");
    }
}
