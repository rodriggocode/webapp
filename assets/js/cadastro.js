const createUser = document.getElementById("btn-create");

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

    const date = {
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
      date,
    );
    console.log(response); // depois tirar
    if (response.status == 201) {
      Toastify({
        text: "Cadastro realizado com sucesso!",
        duration: 3000,
        close: false,
        gravity: "top",
        position: "right",
        stopOnFocus: true,
        style: {
          background: "#FF6B2B",
          borderRadius: "5px",
          border: "1px solid #5A5A5A",
          fontFamily: "DM Sans, sans-serif",
          fontSize: "0.7rem",
          color: "#5A5A5A",
          fontWeight: "bold",
        },
      }).showToast();
      setTimeout(() => window.location.assign("/"), 3000);
    } else if (response.status == 409) {
      Toastify({
        text: "Email ou Nick ja cadastrados, tente novamente!",
        duration: 3000,
        position: "right",
        gravity: "top",
        style: {
          background: "#FF0000",
          borderRadius: "5px",
          border: "1px solid #5A5A5A",
          fontFamily: "DM Sans, sans-serif",
          fontSize: "0.7rem",
          color: "#5A5A5A",
          fontWeight: "bold",
        },
      }).showToast();
    } else {
      Toastify({
        text: "Algo deu errado, tente novamente!",
        duration: 3000,
        gravity: "top",
        position: "right",
        style: {
          background: "#FF0000",
          borderRadius: "5px",
          border: "1px solid #5A5A5A",
          fontFamily: "DM Sans, sans-serif",
          fontSize: "0.7rem",
          color: "#5A5A5A",
          fontWeight: "bold",
        },
      }).showToast();
    }
  } catch (erro) {
    Toastify({
      text: "Algo deu errado, tente novamente!",
      duration: 3000,
      position: "right",
      style: {
        background: "#FF0000",
        borderRadius: "5px",
        border: "1px solid #5A5A5A",
        fontFamily: "DM Sans, sans-serif",
        fontSize: "0.7rem",
        color: "#5A5A5A",
        fontWeight: "bold",
      },
    }).showToast();
  }
};

const validateName = () => {
  const validateName = document.querySelector("#user_name");
  const validate_Name = document.querySelector(".name_valid");
  if (validateName.value == "") {
    validate_Name.classList.remove("hidden");
    return false;
  } else {
    validate_Name.classList.add("hidden");
    return true;
  }
};

const validateNick = () => {
  const validateNick = document.querySelector("#nick");
  const validate_Nick = document.querySelector(".valid_nick");
  const validLengNick = document.querySelector(".valid_leng");
  let isValid = true;
  if (validateNick.value == "") {
    validate_Nick.classList.remove("hidden");
    isValid = false;
  } else {
    validate_Nick.classList.add("hidden");
  }

  // revisar
  if (validateNick.value.length < 5) {
    validLengNick.classList.remove("hidden");
    isValid = false;
  } else {
    validLengNick.classList.add("hidden");
  }

  return isValid;
};

const validateEmail = () => {
  const emailValidate = document.querySelector("#email");
  const validate_Email = document.querySelector(".validate_email");

  if (emailValidate.value == "") {
    validate_Email.classList.remove("hidden");
    return false;
  } else {
    validate_Email.classList.add("hidden");
    return true;
  }
};

const validarPassword = () => {
  const Inputpassword = document.querySelector("#password");

  const confirm_password = document.querySelector("#confirm_password");
  const passWordCaracter = document.querySelector(".password_caracter");
  let isValidPass = true;
  if (Inputpassword.value.length < 5 && confirm_password.value.length < 5) {
    passWordCaracter.classList.remove("hidden");
    isValidPass = false;
  } else {
    passWordCaracter.classList.add("hidden");
  }
  const error_password = document.querySelector(".erro_password");
  if (Inputpassword.value != confirm_password.value) {
    error_password.classList.remove("hidden");
    isValidPass = false;
  } else {
    error_password.classList.add("hidden");
  }

  return isValidPass;
};

createUser.addEventListener("click", (event) => {

  event.preventDefault();

  const isEmailvalid = validateEmail();
  const isPasswordValid = validarPassword();
  const isNameValid = validateName();
  const isNickValid = validateNick();


  if (isEmailvalid && isPasswordValid && isNameValid && isNickValid) {
    CreateUser();
  } else {
    console.log("Formulario invalido");
  }

});
