const takeValue = document.querySelector("#password");
const takeValueLogin = document.querySelector("#login");
const button = document.querySelector("#submit");
const bdy  = document.querySelector("html")
button.addEventListener("click", function () {
    if (takeValue.value == "toto" && takeValueLogin.value == "toto@gmail.com") {
        window.open("menu.html","_self");
    } else {
        alert("Password Incorrect or Login Incorrect");
    }
});

bdy.addEventListener("mousemove", function () {
    document.getElementById("myAudio").play();
});