console.log("JS Launched");

const form = document.querySelector('form');
const inputNoValid = document.querySelector('#test1');
const send = document.querySelector('#send')

form.addEventListener('submit', event => {
    event.preventDefault();

    const inputUser = document.getElementById("send")
    console.log(inputUser.value)

    fetch('http://localhost:8080/post', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
                inputUser : inputUser.value,
            })
    })
        .then(response => response.json())
        .then(data => {
            console.log(data.Display)
            console.log(data);
            document.getElementById("test").innerHTML = data.Display;
            if (data.Life <= 0) {
                window.location.href = "/lose"
            }
            if (data.Life > 0 && data.Display == data.Word) {
                window.location.href = "/win"
            }
            send.value = "";
        });
});