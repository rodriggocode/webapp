const createUser = document.querySelector(".btn-create");

const form = document.querySelector(".forms");

const CreateUser = async () => {
  const user_name = form.elements["user_name"].value;
  const nick = form.elements["nick"].value;
  const email = form.elements["email"].value;
  const password = form.elements["password"].value;
  try {
    const headers = {
      "Content-Type": "application/json",
    };

    const dados = {
      method: "POST",
      headers: headers,
      body: JSON.stringify({
        user_name: user_name,
        nick: nick,
        email: email,
        password: password,
      }),
    };
    const response = await fetch(
      "https://devbook-zqaw.onrender.com/create/user",
      dados,
    );
    console.log(response);
  } catch (erro) {
    console.log("Erro");
  }
};

const validarPassword = () => {
  const Inputpassword = document.querySelector("#password");

  const confirm_password = document.querySelector("#confirm_password");
  const passWordCaracter = document.querySelector(".password_caracter");
  if (Inputpassword.value.length < 5 && confirm_password.value.length < 5) {
    passWordCaracter.classList.remove("hidden");
    return;
  } else {
    passWordCaracter.classList.add("hidden");
  }
  const error_password = document.querySelector(".erro_password");
  if (Inputpassword.value != confirm_password.value) {
    error_password.classList.remove("hidden");
  } else {
    error_password.classList.add("hidden");
  }

  CreateUser();
};

createUser.addEventListener("click", (event) => {
  event.preventDefault();

  validarPassword();
});
