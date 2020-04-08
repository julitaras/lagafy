import CONSTANTS from "./constants";

const Service = token => {
  const config = {
    baseURL: CONSTANTS.URL_API,
    headers: new Headers({
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json"
    })
  };

  return {
    getCurrentTravels: () => {
      return fetch(`${config.baseURL}/travels/current`, {
        method: "get",
        headers: config.headers
      });
    },
    getTravelDetails: id => {
      return fetch(`${config.baseURL}/travel/${id}`, {
        method: "get",
        headers: config.headers
      });
    },
    putCheckin: id => {
      return fetch(`${config.baseURL}reservation/status/${id}`, {
        method: "put",
        headers: config.headers
      });
    },
    putReservation: travelId => {
      return fetch(`${config.baseURL}/reservation/${travelId}`, {
        method: "post",
        headers: config.headers
      });
    },
    getMyReservations: () => {
      return fetch(`${config.baseURL}/reservation/myreservations`, {
        method: "get",
        headers: config.headers
      });
    },
    deleteTravel: id => {
      return fetch(`${config.baseURL}/travel/${id}`, {
        method: "delete",
        headers: config.headers
      });
    },
    updateTravel: travel => {
      return fetch(`${config.baseURL}/travel/${travel.id}`, {
        method: "put",
        headers: config.headers,
        body: JSON.stringify(travel)
      });
    },
    newTravel: travel => {
      return fetch(`${config.baseURL}/travel`, {
        method: "post",
        headers: config.headers,
        body: JSON.stringify(travel)
      });
    }
  };
};

export default Service;
