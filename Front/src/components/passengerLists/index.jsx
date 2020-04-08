import React, { useState, useEffect, useContext, Fragment } from "react";
import StatusList from "./statusList";
import userContext from "../../context/userContext";
import css from "./index.scss";
import cssGral from "../../../style.scss";
import Service from "../../configs/services";
import moment from "moment";
import CONSTANTS from "../../configs/constants";
import {
  isCheckinStarted,
  dateParser,
  timeParser
} from "../../helpers/dateHelper";

const Index = ({ toggleTravel, selectedTravel, list }) => {
  const { user, token } = useContext(userContext);
  const [actualPassenger, setActualPassenger] = useState(null);
  const [checkin, setCheckin] = useState(false);
  const [reserve, setReserve] = useState(false);
  const [reservations, setReservations] = useState(list.reservations);
  const [loading, setLoading] = useState(false);
  const [errorMsg, setErrorMsg] = useState(null);

  useEffect(() => {
    checkStatus();
  }, [reservations]);

  const checkStatus = () => {
    var pg = list.reservations.find(p => p.passenger.email === user.userName);
    if (pg !== undefined) {
      setActualPassenger(pg);
      if (pg.status === "confirmed") {
        setReserve(true);
      } else if (pg.status === "onboard") {
        setCheckin(true);
      }
    }
    setReservations(reservations);
  };

  const handleCheckin = () => {
    setErrorMsg(null);
    setLoading(true);
    if (isCheckinStarted(list.departure)) {
      Service(token)
        .putCheckin(actualPassenger.id)
        .then(res => {
          if (res.status === 200) {
            var aux = reservations;
            var pg = aux.find(p => p.passenger.email === user.userName);
            if (pg !== undefined) {
              pg.status = "onboard";
            }
            setReservations(aux);
            setCheckin(true);
          } else {
            setErrorMsg(CONSTANTS.SERVER_ERROR_MESSAGE);
          }
        })
        .then(setLoading(false));
    } else {
      setErrorMsg(
        "The check in time is from the departure time to 15 minutes later"
      );
      setLoading(false);
    }
  };

  const handleReservation = () => {
    setErrorMsg(null);
    setLoading(true);
    if (selectedTravel !== null && selectedTravel !== undefined) {
      Service(token)
        .putReservation(selectedTravel)
        .then(res => {
          if (res.status === 200) {
            setReserve(true);
            return res.json();
          } else {
            setErrorMsg(CONSTANTS.SERVER_ERROR_MESSAGE);
          }
        })
        .then(res => {
          if (res !== undefined) {
            setReservations([...reservations, res]);
            setActualPassenger(res);
          }
          setLoading(false);
        });
    }
  };

  const handleToggle = () => {
    setErrorMsg(null);
    toggleTravel();
  };

  return (
    <Fragment>
      {loading ? console.count("loading") : null}
      {loading ? (
        <div className={css.containerSpinner}>
          <div className="spinner-border text-primary" role="status">
            <span className="sr-only">Loading...</span>
          </div>
        </div>
      ) : (
        <div className="container">
          {!reservations.length > 0 ? (
            <div className="container">
              <div className={cssGral.NoTravels}>
                <h3 className="text-center">
                  There are no reservations on this trip
                </h3>
              </div>
            </div>
          ) : (
            <Fragment>
              <div className={css.datesTravels}>
                <p>
                  <span className="font-weight-bold"> Origin : </span>
                  {list.origin}
                </p>
                <p>
                  <span className="font-weight-bold"> Destination : </span>
                  {list.destination}
                </p>
                <p>
                  <span className="font-weight-bold"> Capacity : </span>
                  {list.capacity}
                </p>
                <p>
                  <span className="font-weight-bold"> Departure : </span>
                  {dateParser(list.departure)} {timeParser(list.departure)}
                </p>
                <p>
                  <span className="font-weight-bold"> Arrival : </span>
                  {dateParser(list.arrival)} {timeParser(list.arrival)}
                </p>
              </div>

              <StatusList
                status="Reservas"
                passengers={reservations}
              ></StatusList>
              {errorMsg !== null ? (
                <p className={cssGral.errorMsg}>{errorMsg}</p>
              ) : null}
              {!checkin && reserve && (
                <div className={css.btnSucces}>
                  <button
                    className="btn btn-outline-success"
                    onClick={handleCheckin}
                  >
                    Checkin
                  </button>
                </div>
              )}
            </Fragment>
          )}

          <div className={css.actionButtons}>
            {!reserve && !checkin && (
              <button
                className="btn btn-outline-success"
                onClick={handleReservation}
              >
                Reservations
              </button>
            )}
            <button className="btn btn-outline-danger" onClick={handleToggle}>
              Back
            </button>
          </div>
        </div>
      )}
    </Fragment>
  );
};

export default Index;
