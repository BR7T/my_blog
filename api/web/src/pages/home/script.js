
async function getP(){
    let q = window.location.search
    const urlParams = new URLSearchParams(q)

    let param = urlParams.get("size")
    if(param == null){
        param = 1
    }
    let u = "/post/get?page=1"
    try{
        let r = await fetch(u)
        let j = await r.json()
        if (j["error"]){
            console.error("Erro ao buscar dados no servidor")
        }
        return j["message"]
    }catch(err){
        return 
    }
    
}

const main = document.getElementById("MainPageSelector")


document.addEventListener("DOMContentLoaded" , async () => {
    /** @param {Array} data */
    let data = await getP()
    Add(data , document.getElementById("Pages"))
})

/** @param {Array} data  */
/** @param {HTMLElement} element  */

function Add(data , element){
    for(let i=0; i<data.length ; i++){
        let date = FormatDate(data[i].CreatedAt)
        let ID = data[i].ID
        let mainDiv = document.createElement("div")

        mainDiv.innerHTML = `
            <p>${date}</p>
            <u>${data[i].Title}</u>
        `
        mainDiv.classList.add("Post")

        mainDiv.addEventListener("mousedown",()=>{enterpage(ID)})
        element.appendChild(mainDiv)
        console.info(data[i])
    }
}

/** @param {number} id  */
function enterpage(id){
    window.location.href = `/post?id=${id}`
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