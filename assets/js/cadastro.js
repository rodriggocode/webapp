let createUser = document.querySelector(".btn-create");

let password = document.querySelector("#password");

let confirm_password = document.querySelector("#confirm_password");

const CreateUser = async () => {
  try {
    const headers = {
      "Content-Type": "application/json",
    };

    const dados = {
      method: "POST",
      headers: headers,
      body: JSON.stringify({
        user_name: "user_name",
        nick: "nick",
        email: "email",
        password: "password",
      }),
    };
    const response = await fetch(
      "https://devbook-zqaw.onrender.com/usuarios/create",
      dados,
    );
    console.log(response);
  } catch (erro) {
    console.log("Erro");
  }
};

createUser.addEventListener("click", (event) => {
  event.preventDefault();

  console.log("enviando formulario");

  if (password.value != confirm_password.value) {
    console.log("senha errada");
  } else {
    CreateUser();
  }
});
