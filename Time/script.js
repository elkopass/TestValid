function generateText(input) {
    return true
}

function addElementOnPage() {
    //получение значения
    let inputValue = document.querySelector("#input_text").value


    // место, куда будем вставлять текст
    let parentNode = document.querySelector("#container")
    //клонируем шаблон текста
    let newElement = document.querySelector(".output_template").cloneNode()
    //вставляем в него контент (результат твоей функции)
    newElement.textContent = generateText(inputValue);
    // делаем его видимым
    newElement.style.display = "block";
    //добавляем к тому месту
    parentNode.appendChild(newElement);
}

document.querySelector("#btn_save").addEventListener('click', addElementOnPage)
