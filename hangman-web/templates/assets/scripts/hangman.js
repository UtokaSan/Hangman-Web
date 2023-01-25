console.log("JS Launched");
const form = document.querySelector('form');
const send = document.querySelector('#send')
//POSITION HIDDEN
document.getElementById("position-bar").style.visibility="hidden";
document.getElementById("position-head").style.visibility="hidden";
document.getElementById("position-body").style.visibility="hidden";
document.getElementById("position-hand-right").style.visibility="hidden";
document.getElementById("position-hand-left").style.visibility="hidden";
document.getElementById("position-hand-left").style.visibility="hidden";
document.getElementById("position-leg-left").style.visibility="hidden";
document.getElementById("position-leg-right").style.visibility="hidden";

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
            switch (data.Life) {
                case 6 :
                    document.getElementById("position-bar").style.visibility="visible";
                    break;
                case 5:
                    document.getElementById("position-head").style.visibility="visible";
                    break;
                case 4:
                    document.getElementById("position-body").style.visibility="visible";
                    break;
                case 3:
                    document.getElementById("position-hand-left").style.visibility="visible";
                    break;
                case 2:
                    document.getElementById("position-hand-right").style.visibility="visible";
                    break;
                case 1:
                    document.getElementById("position-leg-left").style.visibility="visible";
                    break;
                case 0:
                    document.getElementById("position-leg-right").style.visibility="visible";
                    break;
            }
            send.value = "";
        });
});
