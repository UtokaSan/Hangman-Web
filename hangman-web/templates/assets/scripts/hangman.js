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
            if (data.InputUse === false) {
                let para = document.createElement("p");
                let node = document.createTextNode(inputUser.value);
                para.appendChild(node);
                let element = document.getElementById("test1"); // Utiliser une flex box pour garder le texte aligné
                element.appendChild(para);
            }
            if (data.Life <= 0) {
                window.location.href = "/lose"
            }
            //Corriger cette partie
            if (data.Life > 0 && data.Display !== "_") {
                window.location.href = "/win"
            }
            send.value = "";
        });
});