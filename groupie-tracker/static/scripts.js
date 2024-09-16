// showModalScript shows the modal when it is loaded into the DOM.
function showModalScript() {
    const modal = document.querySelector("#modal-container #modal");
    modal.showModal();
}

// closeModalScript closes the modal when the confirm button is clicked.
function closeModalScript() {
    const btn = document.getElementById("modal-confirm-btn");
    btn.addEventListener("click", function(e) {
        const modal = document.getElementById("confirm-modal");
        modal.close();
    });
}