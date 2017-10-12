"use strict";

// Name input related elements
var input = document.getElementById("name");
var button = document.getElementById("submit");
var nameResponse = document.getElementById("response");

button.addEventListener("click", function() {
    var url = "http://localhost:4004/zips/" + input.value;
    fetch(url)
        .then((resp) => resp.json())
        .then(renderZips)
        .catch(handleError);
});

function renderZips(data) {
    console.log(data);
    
    nameResponse.innerHTML = "Zip: " + data[0].code + ", City: " + data[0].city + ", State: " + data[0].state;
}

function handleError(err) {
    console.error(err);
    alert(err.message);
}