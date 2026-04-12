const loginUser = document.querySelector(".login");
const loginEmail = document.querySelector("#email");

const form = document.querySelector(".forms");
const LoginForm = async () => {
    const email = form.elements["email"].value;
    const password = form.elements["password"].value;
    try {
        const headers = {
            "Content-Type": "application/json",
        };
        const date = {
            method: "POST",
            headers: headers,
            credentials: "include",
            body: JSON.stringify({
                email: email,
                password: password,
            }),
        };
        const apiUrl = "/login";
        const response = await fetch(apiUrl, date);
        if (response.ok) {
            console.log("Login bem-sucedido!");
            const data = await response.json();
            localStorage.setItem("token", data.access_token)
            window.location.assign("/home");
        } else {
            const errorData = await response.json();
        }
    } catch (erro) {
        console.log("Erro na requisicao");
    }
};

loginUser.addEventListener("click", (event) => {
    event.preventDefault();
    LoginForm();
});
