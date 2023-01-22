const takeValuePassword = document.querySelector("#password");
const takeValueLogin = document.querySelector("#login");
const button = document.querySelector("#submit");
const bdy  = document.querySelector("html")
const form = document.querySelector("#form")

bdy.addEventListener("mousemove", function () {
    document.getElementById("myAudio").play();
});

form.addEventListener('submit', event => {
    fetch('http://localhost:8080/post-login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            Mail : takeValueLogin.value,
            Password : takeValuePassword.value,
        })
    })
        .then(response => response.json())
        .then(data => {
            if(data == "IsAccountok") {
                window.location.href = "/difficulty"
            } else {
                alert("Incorrect Mail or Password")
            }
        });
});
