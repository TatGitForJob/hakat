seasonsButton =document.getElementById("seasonsButton")
homepageButton=document.getElementById("homepageButton")
profil1Button =document.getElementById("proButton")
profil2Button=document.getElementById("proTwoButton")








seasonsButton.addEventListener("click", function () {
    window.location.href = "/seasons"
})
homepageButton.addEventListener("click", function () {
    window.location.href = "/homepage"
})
profil1Button.addEventListener("click", function () {
    window.location.href = "/profile"
})
profil2Button.addEventListener("click", function () {
    window.location.href = "/predict"
});