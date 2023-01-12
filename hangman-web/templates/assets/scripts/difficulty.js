const bdy  = document.querySelector("html");
const form = document.getElementById("send-form-post");
const difficult = document.getElementById("send-difficult");

bdy.addEventListener("mousemove", function () {
    document.getElementById("mySound").play();
});

difficult.addEventListener('click', event => {
    form.submit();
});

