console.log("JS Launched");
const form = document.querySelector('form');
const inputNoValid = document.querySelector('#test');
const send = document.querySelector('#send')
var dataRecupère;
var previousLength = history.length;
console.log(dataRecupère)

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
            dataRecupère = data.Word
            console.log(data.Display)
            document.getElementById("life").innerHTML = data.Life;
            document.getElementById("test").innerHTML = data.Display;
            if (data.Life <= 0) {
                window.location.reload()
                window.location.href = "/lose"
            }
            if (data.Life > 0 && data.Display === data.Word) {
                window.location.reload()
                window.location.href = "/win"
            }
            send.value = "";
            console.log("rst : ",dataRecupère)
        });
});