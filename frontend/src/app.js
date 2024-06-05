import {RegisterForm, registerView} from "./js/register.js"

const routes = [
    {path: "/auth", view: registerView}
]

console.log("hello there");
let app = document.querySelector(".app");

if (app) {
    let register = new RegisterForm("register-container", app);
}