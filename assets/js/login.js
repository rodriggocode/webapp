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
    const responseData = await response.json();
    if (response.ok) {
      console.log("Login bem-sucedido!");

      //window.location.assign("/home");  aqui eu tenho que mudar para a url que vem do fly.io https://webapp-snowy-flower-2545.fly.dev/home
      setTimeout(function () {
        window.location.href = responseData.redirect;
      }, 5000);
    } else {
      const errorData = await response.json();
      console.log("nao autorizado");
    }
  } catch (erro) {
    console.log("Erro na requisicao");
  }
};

loginUser.addEventListener("submit", (event) => {
  event.preventDefault();
  LoginForm();
});
