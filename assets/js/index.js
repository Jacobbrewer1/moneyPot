// Get the deposit modal
var depositModal = document.getElementById("depositModal");

// Get the button that opens the modal
var depositBtn = document.getElementById("depositButton");

// Get the withdraw modal
var withdrawModal = document.getElementById("withdrawModal");

// Get the button that opens the modal
var withdrawBtn = document.getElementById("withdrawButton");

// Get the <span> element that closes the modal
var span = document.getElementsByClassName("close");

// When the user clicks the button, open the modal
depositBtn.onclick = function () {
    depositModal.style.display = "block";
}

withdrawBtn.onclick = function () {
    withdrawModal.style.display = "block";
}

// When the user clicks on <span> (x), close the modal
// iterate through each span element for each modal
for (var i = 0; i < span.length; i++) {
    span[i].onclick = function () {
        depositModal.style.display = "none";
        withdrawModal.style.display = "none";
    }
}

// When the user clicks anywhere outside of the modal, close it
window.onclick = function (event) {
    if (event.target == depositModal || event.target == withdrawModal) {
        depositModal.style.display = "none";
        withdrawModal.style.display = "none";
    }
}

$(document).ready(function () {
    var interval = 500;   //number of mili seconds between each call
    var refresh = function() {
        $.ajax({
            url: "/live/updates/amount",
            method: 'get',
            processData: false,
            contentType: false,
            cache: false,
            success: function(d) {
                //document.getElementById("liveAmount").innerHTML = "New text!";
                document.getElementById("liveAmount").innerHTML = "Amount in pot: " + d;
                setTimeout(function() {
                    refresh();
                }, interval);
            }
        });
    };
    refresh();
});

function withdrawMoneyToDb(evt) {
    evt.preventDefault();
    let form = evt.target;
    let formData = new FormData(form);
    let amount = formData.get("withdrawMoneyInput");
    console.log("amount received: " + amount)
    if (amount != null && amount != "") {
        $.ajax({
            url: '/withdrawMoney',
            method: 'post',
            data: formData,
            processData: false,
            contentType: false,
            success: () => {
                console.log("Amount withdrawn");
                form.reset();
                withdrawModal.style.display = "none"
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
