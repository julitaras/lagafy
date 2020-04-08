const fetch = require("node-fetch");
const jwt = require("jsonwebtoken");
const dotenv = require("dotenv");
global.Headers = fetch.Headers;

dotenv.config({ path: ".env" });
var urlTravel = "https://lagafyapi.azurewebsites.net/apigroup/travel";
var urlTemplates = "https://lagafyapi.azurewebsites.net/apigroup/travels/templates";
var travels = [];

function generateToken() {
    return jwt.sign({}, process.env.JwtSecret);
}
var token = generateToken();
token = "Bearer " + token;

function getData() {
    fetch(urlTemplates, {
            method: "GET",
            headers: new Headers({
                Authorization: token,
                "Content-Type": "application/json"
            })
        })
        .then(resp => resp.json())
        .then(r => createArray(r))
        .catch(error => console.log("Error:", error))
}

function createArray(data) {
    for (let i = 0; i < data.length; i++) {
        let departure = new Date(data[i].departure);
        let arrival = new Date(data[i].arrival);
        const now = new Date();
        if (
            now.getUTCHours() >= 18 &&
            departure.getUTCHours() >= 9 &&
            departure.getUTCHours() < 15
        ) {
            let nextTravelDay = 1;
            if (now.getDate() === 5) {
                nextTravelDay = 3;
            }

            departure.setDate(now.getDate() + nextTravelDay);
            departure.setMonth(now.getMonth());
            departure.setFullYear(now.getFullYear());

            arrival.setDate(now.getDate() + nextTravelDay);
            arrival.setMonth(now.getMonth());
            arrival.setFullYear(now.getFullYear());
        }
        if (
            now.getUTCHours() >= 13 &&
            now.getUTCHours() < 18 &&
            departure.getUTCHours() >= 15
        ) {
            departure.setDate(now.getDate());
            departure.setMonth(now.getMonth());
            departure.setFullYear(now.getFullYear());

            arrival.setDate(now.getDate());
            arrival.setMonth(now.getMonth());
            arrival.setFullYear(now.getFullYear());
        }

        if (new Date(departure).getUTCDate() !== new Date(data[i].departure).getUTCDate() || new Date(departure).getUTCMonth() !== new Date(data[i].departure).getUTCMonth() || new Date(departure).getUTCFullYear() !== new Date(data[i].departure).getUTCFullYear()) {
            obj = {
                hasWifi: data[i].hasWifi,
                capacity: data[i].capacity,
                driver: data[i].driver,
                departure: departure,
                arrival: arrival,
                origin: data[i].origin,
                destination: data[i].destination,
                arrivalAddress: data[i].arrivalAddress,
                departureAddress: data[i].departureAddress,
                status: data[i].status,
                template: false
            };
            travels.push(obj);
        }
    }
    sendData();
}

getData();

function sendData() {
    fetch(urlTravel, {
            method: "POST",
            body: JSON.stringify(travels),
            headers: new Headers({
                Authorization: token,
                "Content-Type": "application/json",
                "userMail": "johannac@lagash.com"
            })
        })
        .then(res => res.json())
        .catch(error => console.log("Error:", error))
        .then(response => console.log("Success:", response));
}