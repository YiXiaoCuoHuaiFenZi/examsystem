function createMethodControll(id) {
    var ManualCreateRadio = document.getElementById("ManualCreateRadio")
    var BatchCreateRadio = document.getElementById("BatchCreateRadio")

    var ManualCreate = document.getElementById("ManualCreate")
    var BatchCreate = document.getElementById("BatchCreate")

    switch (id) {
        case "ManualCreateRadio":
            BatchCreate.style.display = "none";

            ManualCreate.style.display = "";
            break;
        case "BatchCreateRadio":
            ManualCreate.style.display = "none";

            BatchCreate.style.display = "";
            break;
        default:
            break;
    }
}

function batchCreateFormsContorll(id) {
    var BatchSingleChoiceType = document.getElementById("BatchSingleChoiceType")
    var BatchMultipleChoiceType = document.getElementById("BatchMultipleChoiceType")
    var BatchTrueFalseType = document.getElementById("BatchTrueFalseType")

    var BatchSingleChoiceForm = document.getElementById("BatchSingleChoiceForm")
    var BatchMultipleChoiceForm = document.getElementById("BatchMultipleChoiceForm")
    var BatchTrueFalseForm = document.getElementById("BatchTrueFalseForm")

    switch (id) {
        case "BatchSingleChoiceType":
            BatchMultipleChoiceForm.style.display = "none";
            BatchTrueFalseForm.style.display = "none";

            BatchSingleChoiceForm.style.display = "";
            break;
        case "BatchMultipleChoiceType":
            BatchSingleChoiceForm.style.display = "none";
            BatchTrueFalseForm.style.display = "none";

            BatchMultipleChoiceForm.style.display = "";
            break;
        case "BatchTrueFalseType":
            BatchSingleChoiceForm.style.display = "none";
            BatchMultipleChoiceForm.style.display = "none";

            BatchTrueFalseForm.style.display = "";
            break;
        default:
            break;
    }
}

function manualCreateFormsContorll(id) {
    var SingleChoiceType = document.getElementById("SingleChoiceType")
    var MultipleChoiceType = document.getElementById("MultipleChoiceType")
    var TrueFalseType = document.getElementById("TrueFalseType")

    var SingleChoiceForm = document.getElementById("SingleChoiceForm")
    var MultipleChoiceForm = document.getElementById("MultipleChoiceForm")
    var TrueFalseForm = document.getElementById("TrueFalseForm")

    switch (id) {
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
