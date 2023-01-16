const bdy  = document.querySelector("html");
const form = document.getElementById("send-form-post");
const difficultEasy = document.getElementById("send-difficult-easy");
const difficultMoyen = document.getElementById("send-difficult-moyen");
const difficultHard = document.getElementById("send-difficult-hard");

bdy.addEventListener("mousemove", function () {
    document.getElementById("mySound").play();
});

form.addEventListener('submit', event => {
    console.log(difficultEasy);
    if(event.submitter.name === 'difficulty-easy') {
        data = 'Easy'
    } else if(event.submitter.name === 'difficulty-moyen') {
        data = 'Moyen'
    } else if (event.submitter.name === 'difficulty-hard') {
        data = 'Hard'
    }
    fetch('http://localhost:8080/post-difficulty', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            inputUser : data,
        })
    })
        .then(response => response.json())
        .then(data => {
            console.log(data)
        });
});