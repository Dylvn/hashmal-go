/*
    Submit ajax change password form
*/
const changePass = () => {
    changePassForm.addEventListener("submit", (e) => {
        e.preventDefault();
        const url = changePassForm.attributes["action"].value;
        const method = changePassForm.attributes["method"].value;
        const data = new FormData(changePassForm);
    
        fetch(url, {method: method, body: data}).then(res => {
            return res.text();
        }).then(data => {
            changePassContainer.innerHTML = data;
        });
    });
}

const changePassForm = document.querySelector("#js-change-password");
const changePassContainer = document.querySelector("#js-block-change-password");

if (changePassForm != null && changePassContainer != null) {
    changePass();
}