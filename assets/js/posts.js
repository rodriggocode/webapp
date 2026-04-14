const post = document.getElementById("btn-post");

const postForm = document.querySelector("#send-form");

const CreatePostForm = async () => {
  const title = postForm.elements["title"].value;
  const content = postForm.elements["content"].value;

  try {
    const headers = {
      "Content-Type": "application/json",
    };

    const date = {
      method: "POST",
      headers: headers,
      body: JSON.stringify({
        title: title,
        content: content,
      }),
    };
    const response = await fetch(
      "/create/post", date);
    if (response.ok) {
      Toastify({
        text: "Publicacao criada",
        duration: 3000,
      }).showToast();
      window.location.href = "/home"
    } else {
      console.log(`erro: ${response.status}`)
    }
  } catch (erro) {
    console.log("Deu ruim alguma coisa aqui");
  }
}

post.addEventListener("click", (event) => {
  event.preventDefault()
  CreatePostForm();
});
