function KeyDown() { //屏蔽鼠标右键、Ctrl+n、shift+F10、F5刷新、退格键  
    if ((window.event.altKey) &&
        ((window.event.keyCode == 37) || //屏蔽   Alt+   方向键   ←  
            (window.event.keyCode == 39))) { //屏蔽   Alt+   方向键   →  
        alert("不准你使用ALT+方向键前进或后退网页！");
        event.returnValue = false;
    }
    if ((event.keyCode == 8) || //屏蔽退格删除键  
        (event.keyCode == 116) || //屏蔽   F5   刷新键  
        (event.keyCode == 112) || //屏蔽   F1   刷新键  
        (event.ctrlKey && event.keyCode == 82)) { //Ctrl   +   R  
        event.keyCode = 0;
        event.returnValue = false;
    }
    if ((event.ctrlKey) && (event.keyCode == 78)) //屏蔽   Ctrl+n  
        event.returnValue = false;
    if ((event.shiftKey) && (event.keyCode == 121)) //屏蔽   shift+F10  
        event.returnValue = false;
    if (window.event.srcElement.tagName == "A" && window.event.shiftKey)
        window.event.returnValue = false; //屏蔽   shift   加鼠标左键新开一网页  
    if ((window.event.altKey) && (window.event.keyCode == 115)) { //屏蔽Alt+F4  
        window.showModelessDialog("about:blank", "", "dialogWidth:1px;dialogheight:1px");
        return false;
    }
    if ((window.event.altkey) && (window.event.keyCode == 27)) {
        alert("认真答题！");
    }
}

function Showhelp() {
    alert("认真答题！");
    return false;
}


var iPreX;
var iPreY;

function MouseMove() {
    var isFlow = true;
    var iX = parseInt(event.clientX);
    var iY = parseInt(event.clientY);
    var iDivX = parseInt(document.getElementById("divT").style.left);
    var iDivY = parseInt(document.getElementById("divT").style.top);
    var iDivW = parseInt(document.getElementById("divT").style.width);
    var iDivH = parseInt(document.getElementById("divT").style.height);
    if (iX < iDivX || iX > (iDivX + iDivW)) {
        alert("超出范围"); //这里试试能不能把鼠标定在最大尺度.
    }
    if (iY < iDivY && iY > (iDivY + iDivH)) {
        alert("超出范围");
    }
}

function OpenPopUpWindow(url) {
    window.open(url, 'popUpWindow', 'fullscreen=yes, resizable=no,scrollbars=yes, toolbar = no, menubar = no, location = no, directories = no, status = no ');
}
