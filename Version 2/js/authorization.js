form = document.getElementById("form")
pass = document.getElementById("pass")
login = document.getElementById("login")


form.addEventListener("submit", function () {
    let data = {
        Login: login.value,
        Password: pass.value
    };
    // Number:
    fetch("/", {
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        method: "POST",
        body: JSON.stringify(data)
    })
        .then(function (response) {
                window.location.href = "/homepage"
        })
        .catch(function (error) {
            console.error("Ошибка при отправке запроса:", error);
        });
});