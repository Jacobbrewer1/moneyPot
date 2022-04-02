// Get the modal
var depositModal = document.getElementById("depositModal");

// Get the button that opens the modal
var depositBtn = document.getElementById("depositButton");

// Get the <span> element that closes the modal
var span = document.getElementsByClassName("close")[0];

// When the user clicks the button, open the modal
depositBtn.onclick = function () {
    depositModal.style.display = "block";
}

// When the user clicks on <span> (x), close the modal
span.onclick = function () {
    depositModal.style.display = "none";
}

// When the user clicks anywhere outside of the modal, close it
window.onclick = function (event) {
    if (event.target == depositModal) {
        depositModal.style.display = "none";
    }
}

function depositMoneyToDb(evt) {
    evt.preventDefault();
    let form = evt.target;
    let formData = new FormData(form);
    let amount = formData.get("depositMoneyInput");
    console.log("amount received: " + amount)
    if (amount != null && amount != "") {
        $.ajax({
            url: '/depositMoney',
            method: 'post',
            data: formData,
            processData: false,
            contentType: false,
            success: () => {
                console.log("Amount deposited");
                form.reset();
                depositModal.style.display = "none"
            },
            error: (d) => {
                console.log("An error occurred. Please try again");
                console.log(d)
                form.reset();
            }
        });
    }
    return false
}