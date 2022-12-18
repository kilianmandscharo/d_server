function handleClick() {
  const button = document.querySelector(".open-button")
  const present = document.querySelector(".present")
  const prelude = document.querySelector(".prelude")
  button.classList.add("disappear")
  present.classList.add("disappear")
  prelude.classList.add("appear")
}

function addButtonFunctionality() {
  const button = document.querySelector(".open-button")
  button.addEventListener("click", handleClick)
}

function createSnowflakes() {

}

addButtonFunctionality()
