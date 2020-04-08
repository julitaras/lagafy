import React, { useState, useEffect } from "react";
import Service from "../../../configs/services";
import { userContext } from "../../../context";
import ReservationCard from "./reservationCard";
import cssTravel from "../index.scss";

const reservationsMock = [
  {
    id: 1,
    status: "confirmed",
    origin: "LAGASH",
    destination: "MELI DOT",
    departure: "2020-03-04T09:00:00-03:00",
    arrival: "2020-03-04T10:00-03:00"
  },
  {
    id: 2,
    status: "confirmed",
    origin: "MELI DOT",
    destination: "LAGASH",
    departure: "2020-03-04T17:00:00-03:00",
    arrival: "2020-03-04T18:00:00-03:00"
  }
];

export const ReservationContainer = ({ loading, setLoading, isLoading }) => {
  const { token } = userContext();
  const [reservations, setReservations] = useState([]);

  useEffect(() => {
    getReservations();
  }, []);

  const getReservations = () => {
    Service(token)
      .getMyReservations()
      .then(res => {
        if (res.status === 200) {
          return res.json();
        } else {
          console.log("failed to call travels/current");
          throw "error";
        }
      })
      .then(res => {
        setReservations(res);
        setLoading(false);
      })
      .catch(err => console.log("Server Error"));
    //setReservations(reservationsMock);
  };

  return (
    <div>
      {!loading && reservations.length === 0 ? (
        <div className="text-center">
          <i className="icon-Sad-Calendar-00"></i>
          <p> Todavía no tenés reservas!</p>
        </div>
      ) : (
        <div>
          {!isLoading() && (
            <div className="row justify-content-md-center">
              {reservations.map(t => (
                <ReservationCard
                  key={t.id}
                  idReservation={t.id}
                  idTravel={t.TravelID}
                  origin={t.Travel.origin.toUpperCase()}
                  destination={t.Travel.destination.toUpperCase()}
                  departure={t.Travel.departure}
                  arrival={t.Travel.arrival}
                  status={t.status}
                />
              ))}
            </div>
          )}
        </div>
      )}
    </div>
  );
};
export default ReservationContainer;
