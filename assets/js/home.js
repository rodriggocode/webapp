logout = document.querySelector(".logout");
logout.addEventListener("click", () => {
  window.location.href = "/";
});


let like = document.querySelector(".like");

const Like = () => {
  let cont = document.querySelector("#count");
  cont.innerHTML = parseInt(cont.innerHTML) + 1;

  let icon = document.getElementById("like-icon");
 icon.style.fill = "red";
}



like.addEventListener("click", () => {
  Like();
});


let showForm = document.getElementById("posts");
showForm.addEventListener("click", () => {
  console.log("chegou aqui");
  let showPost = document.querySelector("#new-post-form");
  showPost.classList.remove("hidden");

});
