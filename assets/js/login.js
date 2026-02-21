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
    const apiUrl = "https://devbook-zqaw.onrender.com/publicacoes";
    const response = await fetch(apiUrl, date);
    const responseData = await response.json();
    if (response.ok) {
      console.log("Login bem-sucedido!");

      window.location.href = responseData.redirect;
    } else {
      const errorData = await response.json();
      const erroMessage =
        errorData.Erro || "O correu um erro desconheciado no servidor";
      alert(erroMessage);
    }
  } catch (erro) {
    console.log("Erro na requisicao");
  }
};

loginUser.addEventListener("submit", (event) => {
  event.preventDefault();
  LoginForm();
});
