import {RegisterForm} from "./js/register.js"

const routes = {
    "/": "",

}

console.log("hello there");
let app = document.querySelector(".app");

if (app) {
    let register = new RegisterForm("register-container", app);
}
console.log(window.location.pathname);