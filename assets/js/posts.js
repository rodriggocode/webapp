const post = document.getElementById("btn-post");

const postForm = document.querySelector("#send-form");

const CreatePostForm = async () => {
 const title = postForm.elements["title"].value;
 const content = postForm.elements["content"].value;

 try{
  const headers = {
   "Content-Type": "application/json",
   "Authorization": "Bearer " + localStorage.getItem("token"),
  };

  const date = {
   method: "POST",
   headers: headers,
   credentials: "include",
   body: JSON.stringify({
     title: title,
     content: content,
   }),
  };
  const response = await fetch(
   "https://devbook-zqaw.onrender.com/publicacao/criar", date);
  console.log(response);
 }catch(erro){
  console.log("Deu ruim alguma coisa aqui");
 }
}

post.addEventListener("click", (event) =>{
 event.preventDefault()
 CreatePostForm();
});
