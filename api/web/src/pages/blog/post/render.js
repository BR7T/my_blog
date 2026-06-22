function getQueryParam(param) {
    const urlParams = new URLSearchParams(window.location.search);
    return urlParams.get(param);
}

async function FindPost(idParam){
  try{
    let uq = `/post/get?id=${idParam}`
    const resp = await fetch(uq)

    if (!resp.ok) {
        console.error('Status:', resp.status, resp.statusText)
        return
    }

    let j = await resp.json()

    return j
  }catch(err){
    console.error(err)
  }buratinivictor
}


let titleDiv = document.querySelector(".titlePost")
let textDiv = document.querySelector(".textArea")

function CreatePage(obj){
  let titleE = document.createElement("h2")
  document.title = obj.Title
  let dateTimeE = document.createElement("p")
  titleE.innerHTML = obj.Title
  dateTimeE.innerHTML = FormatDate(obj.CreatedAt)
  dateTimeE.classList.add("dateField")
  titleDiv.appendChild(titleE)
  titleDiv.appendChild(dateTimeE)
  textDiv.innerHTML = Render(obj.Content)
}

function FormatDate(date){
  try{
    let d = String(new Date(date).getDate()).padStart(2 , '0')
    let m = String(new Date(date).getMonth()).padStart(2 , '0')
    let y = new Date(date).getFullYear()
    const months = ["Janeiro","Fevereiro","Março","Abril","Maio","Junho","Julho","Agosto","Setembro","Outubro","Novembro","Dezembro"]

    return `${d} de ${months[m-1]}, ${y} `
  }catch(err){
    return "NAN"
  }
}

function Render(text){
  let html = ""
  let sep = String(text).split("\n")

  for(let i = 0 ; i < sep.length ; i++){
    let line = sep[i]

    if(line.trim()==""){
      continue
    }
    const match = line.match(/^(#{1,6})\s/)
    if(match){
      const level = match[1].length
      const content = line.replace(/^#{1,6}\s/, "")
      line =`<h${level}>${content}</h${level}>`
      html += line
      continue
    }
    if(line == "---"){
      html += "<hr class='hr'/>"
      continue
    }
    const image = line.match(/^!\[([^\]]*)\]\(([^)]+)\)$/)
    if(image){
      html += `<div class='img-wrapper'><img src='/images/${image[2]}' alt='${image[1]}'></div>`
      continue
    }
    line = line
    .replace(/\[([^\]]+)\]\(([^)]+)\)/g,"<a href='$2'>$1</a>")
    .replace(/\*\*([^*]+)\*\*/g,"<strong>$1</strong>")
    .replace(/\*([^*]+)\*/g,"<em>$1</em>")
    .replace(/`([^`]+)`/g,"<code>$1</code>")
    .replace(/\=\=([^`]+)\=\=/g,"<span>$1</span>")

    
    html += `<p>${line}</p>`
  }
  return html
}

document.addEventListener("DOMContentLoaded" , async ()=>{
  const idParam = getQueryParam('id')
  const data = await FindPost(idParam)
  CreatePage(data["message"])
})


document.getElementById("home").addEventListener("click" , ()=>{
  window.location.href = "/home"
})