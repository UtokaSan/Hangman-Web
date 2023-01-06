console.log("JS Launched");

const form = document.querySelector('form');

form.addEventListener('submit', event => {
    event.preventDefault();

    const formData = new FormData(form);

    fetch('http://localhost:8080/get', {
        method: 'POST',
        body: formData
    })
        .then(response => response.text())
        .then(data => {
            document.getElementById("test").innerHTML = data;
            console.log("")
        });
});