import React, { useState, useEffect } from "react";
import {
  timeParser,
  dateParser,
  isReservationTime
} from "../../../../helpers/dateHelper";
import Service from "../../../../configs/services";
import { userContext } from "../../../../context";
import css from "../../../../../style.scss";
import cssReservas from "./index.scss";
import moment from "moment";

export const TravelDetails = ({ setSelectedTravel, travel: { id } }) => {
  const { token, user } = userContext();
  const [errorMsg, setErrorMsg] = useState(null);
  const [travelDetails, setTravelDetails] = useState(null);
  const [loading, setLoading] = useState(true);
  const [reserved, setReserved] = useState(false);

  useEffect(() => {
    getTravelDetails();
  }, []);

  const getTravelDetails = () => {
    setErrorMsg(null);
    Service(token)
      .getTravelDetails(id)
      .then(res => {
        if (res.status === 200) {
          return res.json();
        } else {
          setErrorMsg(`Error:`);
        }
      })
      .then(res => {
        if (errorMsg === null) {
          setTravelDetails(res);
        } else {
          setErrorMsg(`${errorMsg} ${res.message}`);
        }
        setLoading(false);
      });
  };

  const handleReservation = () => {
    setErrorMsg(null);
    setLoading(true);
    if (id && isReservationTime(travelDetails.departure)) {
      Service(token)
        .putReservation(id)
        .then(res => res.json())
        .then(res => {
          if (res.status === "confirmed") {
            setReserved(true);
            setLoading(false);
          } else {
            // setErrorMsg(CONSTANTS.SERVER_ERROR_MESSAGE);
            setErrorMsg(`Error: ${res.message}`);
            setLoading(false);
          }
        });
    } else {
      setErrorMsg("No puedes reservar en este horario");
      setLoading(false);
    }
  };

  const isReserved = () =>
    travelDetails.reservations.find(t => t.passenger.email === user.userName);

  return (
    <>
      {loading ? (
        <div className={css.containerSpinner}>
          <div className="spinner-border text-primary" role="status">
            <span className="sr-only">Loading...</span>
          </div>
        </div>
      ) : (
        <div className="container">
          <div className={cssReservas.containerReservas}>
            <h1 className="bold">
              <span onClick={() => setSelectedTravel(null)}>
                <i className="icon-Back-00 bold"></i>
              </span>
              Reservas
            </h1>
            <div className={cssReservas.datesTravels}>
              <p>{`Recorrido ${travelDetails.origin} a ${travelDetails.destination}`}</p>
              <p>{`Partida: ${timeParser(travelDetails.departure)} - ${
                travelDetails.departureAddress
              }`}</p>
              <p>{`Llegada: ${timeParser(travelDetails.arrival)} - ${
                travelDetails.arrivalAddress
              }`}</p>
              <p>{`Fecha: ${dateParser(travelDetails.departure)}`}</p>
            </div>

            {errorMsg !== null && <p className={css.errorMsg}>{errorMsg}</p>}
          </div>

          <div className="text-center">
            <button
              className={`btn ${css.btnSet} `}
              disabled={
                reserved ||
                !isReservationTime(travelDetails.departure) ||
                isReserved()
              }
              onClick={handleReservation}
            >
              <i className="icon-Checkmark-00"></i>
              {reserved || isReserved() ? "Reservado" : "Reservar mi lugar"}
            </button>
            {isReserved() || reserved ? (
              <>
                <button
                  className={`btn ${cssReservas.btnEmpty}`}
                  onClick={() => setSelectedTravel(null)}
                >
                  Volver al incio
                </button>
              </>
            ) : (
              <>
                <div className={cssReservas.reservaAlert}>
                  <div className="alert alert-primary" role="alert">
                    <i className="icon-Alert-00"></i>
                    {`Recorda que tenes hasta las ${timeParser(
                      moment(travelDetails.departure).subtract(30, "m")
                    )} para reservar`}
                  </div>
                </div>
              </>
            )}
          </div>
        </div>
      )}
    </>
  );
};
export default TravelDetails;
