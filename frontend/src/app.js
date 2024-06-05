import {RegisterForm} from "./js/register.js"

console.log("hello there");
let app = document.querySelector(".app");

if (app) {
    let register = new RegisterForm("register-container", app);
}