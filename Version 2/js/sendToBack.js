askButton=document.getElementById("ask-Button")
direction=document.getElementById("Direction")
date=document.getElementById("Date")
clas=document.getElementById("Class")
number=document.getElementById("Number")
startdate=document.getElementById("start-date")
enddate=document.getElementById("end-date")
output=document.getElementById("output")


askButton.addEventListener("click", function () {
    let data = {
        Direction:direction.value,
        Date:date.value,
        Class:clas.value,
        Number:number.value,
        StartDate:startdate.value,
        EndDate:enddate.value
    };
    // Number:
    fetch("/get_time", {
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        method: "POST",
        body: JSON.stringify(data)
    }).then((response) => {
        response.text().then(function (data) {
            let result = JSON.parse(data);
            output.textContent="вывод"+result["Count"]+result["Date"]
        });
    }).catch(() => {
    });
})