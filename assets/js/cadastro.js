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
    console.log(response);
    if (response.status == 201) {
      window.location.assign("/");
    } else {
      window.alert("erro");
    }
  } catch (erro) {
    console.log("Erro");
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
