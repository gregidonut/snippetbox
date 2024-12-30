const curPathName = new URL(window.location.href).pathname
const navList = document.querySelectorAll("nav > ul > li")

navList.forEach(function (e) {
    if (curPathName === e.querySelector("a").getAttribute("href")) {
        const text = e.querySelector("a").innerText
        const strongEl = document.createElement("strong")
        strongEl.innerText = text

        e.replaceWith(strongEl)
    }
})

