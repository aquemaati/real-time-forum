// register.js
export class RegisterForm {
    constructor(containerId, parent) {
        this.containerId = containerId;
        this.renderForm(parent, this.containerId);
    }

    renderForm(parent, containerId) {
        const registerForm = document.createElement("div");
        registerForm.id = containerId;
        registerForm.innerHTML = `
            <h2>Register</h2>
            <form id="registerForm" action="/register" method="post">
                <label for="nickname">Nickname:</label>
                <input type="text" id="nickname" name="nickname" required><br>
                <label for="age">Age:</label>
                <input type="number" id="age" name="age" required><br>
                <label for="gender">Gender:</label>
                <input type="text" id="gender" name="gender" required><br>
                <label for="firstName">First Name:</label>
                <input type="text" id="firstName" name="firstName" required><br>
                <label for="lastName">Last Name:</label>
                <input type="text" id="lastName" name="lastName" required><br>
                <label for="email">E-mail:</label>
                <input type="email" id="email" name="email" required><br>
                <label for="password">Password:</label>
                <input type="password" id="password" name="password" required><br>
                <button type="submit">Register</button>
            </form>
        `;
        parent.appendChild(registerForm);
    }

    setupFormSubmission() {
        const registerForm = document.getElementById("registerForm");
        registerForm.addEventListener("submit", async (event) => {
            event.preventDefault();
            const formData = new FormData(registerForm);
            const formDataObject = Object.fromEntries(formData.entries());

            try {
                const response = await fetch(registerForm.action, {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify(formDataObject),
                });

                if (response.ok) {
                    console.log("Registration successful");
                } else {
                    console.error("Registration failed");
                }
            } catch (error) {
                console.error("Error:", error);
            }
        });
    }

    activateSubmission() {
        const activateButton = document.getElementById("activateSubmit");
        activateButton.addEventListener("click", () => {
            this.setupFormSubmission();
            activateButton.disabled = true; // Disable the button to prevent multiple activations
        });
    }
}
