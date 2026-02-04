let createUser = document.querySelector(".btn-create");

let Inputpassword = document.querySelector("#password");

let confirm_password = document.querySelector("#confirm_password");

let form = document.querySelector(".forms");

const CreateUser = async () => {
  let user_name = form.elements["user_name"].value;
  let nick = form.elements["nick"].value;
  let email = form.elements["email"].value;
  let password = form.elements["password"].value;
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

createUser.addEventListener("click", (event) => {
  event.preventDefault();

  alert("enviando formulario");

  if (Inputpassword.value != confirm_password.value) {
    console.log("senha errada");
  } else {
    CreateUser();
  }

  console.log(email.value);
});
