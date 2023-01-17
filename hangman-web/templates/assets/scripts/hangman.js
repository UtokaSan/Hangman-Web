console.log("JS Launched");

const form = document.querySelector('form');

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
            document.getElementById("test").innerHTML = data.Display;
        });
});
if (data.Display === "win") {
    document.location.href = 'http://localhost:8080/post';
}