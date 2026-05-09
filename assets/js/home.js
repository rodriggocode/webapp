// Logout
const logout = document.querySelector(".logout");
logout.addEventListener("click", () => {
  window.location.href = "/";
});

// Like
document.querySelectorAll(".like").forEach((like) => {
  like.addEventListener("click", async () => {
    const postId = like.dataset.postId;

    const countElement = like.querySelector(".count");
    const icon = like.querySelector("path");

    try {
      const response = await fetch(`/publicacoes/${postId}/curtir`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        }
      });

      if (response.ok) {
        const current = parseInt(countElement.textContent);
        countElement.textContent = `${current + 1} curtidas`;
        icon.style.fill = "red";
      } else {
        console.log("Erro ao curtir:", response.status);
      }
    } catch (erro) {
      console.log("Falha na requisição:", erro);
    }
  });
});

let showForm = document.getElementById("posts");
showForm.addEventListener("click", () => {
  let showPost = document.querySelector("#new-post-form");
  showPost.classList.remove("hidden");
});
