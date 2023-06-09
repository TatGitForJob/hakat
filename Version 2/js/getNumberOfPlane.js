 direction = document.getElementById("Direction");
 flightSelection = document.getElementById("Number");

direction.addEventListener("change", () => {
    flightSelection.innerHTML = "";
    let data = {
        Direction: direction.value,
    };
    // Number:
    fetch("/get_season_class", {
        headers: {
            Accept: "application/json",
            "Content-Type": "application/json",
        },
        method: "POST",
        body: JSON.stringify(data),
    })
        .then((response) => {
            return response.json();
        })
        .then((data) => {
            data.forEach((optionValue) => {
                const optionElement = document.createElement("option");
                optionElement.value = optionValue;
                optionElement.text = optionValue;
                flightSelection.appendChild(optionElement);
            });
        })
        .catch((error) => {
            console.log(error);
        });
});
