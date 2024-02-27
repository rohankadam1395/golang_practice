console.log("hellow world")


let getToodos=async (event)=>{
    console.log("button clicked")
    const todos = await fetch("/todos").then((res) => res.json())
    console.log("test")
    console.log(todos)
    let list = document.getElementById("todos")
    list.replaceChildren()
    for (element of todos){
        console.log(element)
        let li = document.createElement("li")
        li.textContent = element.title
        list.appendChild(li)
    }

}   