function pageClick(o) {
    pageNumber = o.innerText;
    pageList = document.getElementsByName("pageList");

    flag = false;
    if (pageNumber == "上一页" || pageNumber == "下一页") {
        $('.pagination>li').each(function() {
            var status = $(this).prop("class")
            if (status == "active") {
                if ($(this).text() == (pageList.length - 2).toString() && pageNumber == "下一页") {
                    flag = true;
                }

                if (($(this).text() == "1") && pageNumber == "上一页") {
                    flag = true;
                }

                var IdNumber = parseInt($(this).text())
                if (pageNumber == "上一页") {
                    pageNumber = (IdNumber - 1).toString();
                } else {
                    pageNumber = (IdNumber + 1).toString();
                }
            }
        });
    }

    if (flag)
        return;

    for (var k = 1; k <= pageList.length; k++) {
        pn = k.toString();
        els = document.getElementsByName("page_" + pn)
        if (pn == pageNumber) {
            for (var i = 0; i < els.length; i++) {
                els[i].style.display = "block";
            }
        } else {
            for (var i = 0; i < els.length; i++) {
                els[i].style.display = "none";
            }
        }
    }

    for (var i = 0; i < pageList.length; i++) {
        id = "#" + pageList[i].id;
        if (pageList[i].innerText == pageNumber) {
            $(id).addClass('active');
        } else {
            id = "#" + pageList[i].id
            $(id).removeClass('active');
        }
    }

    $("#controll_page_Previous").removeClass('disabled');
    $("#controll_page_Next").removeClass('disabled');

    if (pageNumber == "1") {
        $("#controll_page_Previous").addClass('disabled');
    }

    if (pageNumber == (pageList.length - 2).toString()) {
        $("#controll_page_Next").addClass('disabled');
    }
}
