import React, { useState } from "react";
import Link from "next/link";
import { dateParser, timeParser } from "../../../../helpers/dateHelper";
import { userContext } from "../../../../context";
import cssReservation from "./index.scss";
import css from "../../../../../style.scss";
import Service from "../../../../configs/services";
import ButtonSpinner from "../../../uiKit/spinnerButton";
import { isCheckinStarted } from "../../../../helpers/dateHelper";

export const ReservationCard = ({
  origin,
  destination,
  departure,
  idTravel,
  idReservation,
  status,
  arrival
}) => {
  const { token } = userContext();
  const [checkined, setCheckined] = useState(status === "onboard");
  const [errorMsg, setErrorMsg] = useState(null);
  const [loading, setLoading] = useState(false);

  const handleCheckin = () => {
    setLoading(true);
    setErrorMsg(null);

    if (idReservation) {
      Service(token)
        .putCheckin(idReservation)
        .then(res => res.json())
        .then(res => {
          setLoading(false);
          if (res.status === "onboard") {
            setCheckined(true);
          } else {
            setErrorMsg(`Error: ${res.message}`);
          }
        })
        .catch(err => {
          setLoading(false);
          setErrorMsg(
            <p className={css.errorMsg}>"Hubo un error intente mas tarde"</p>
          );
        });
    }
  };

  return (
    <div className="col-6 col-lg-2">
      <div className={`${cssReservation.card} card`}>
        <div
          className={`${cssReservation.cardReservation} card-body text-center`}
        ><Link
            href="/reservation/[...params]"
            as={`/reservation/${idTravel}/${idReservation}/${
              checkined ? "onboard" : status
            }`}
          >

          
            <div>
            <p
            className={`${cssReservation.titleReservation} semiBold`}
          >{`${origin} a ${destination}`}</p>
              <p
                className={`${cssReservation.datesReservation} light`}
              >{`Partida: ${timeParser(departure)}`}</p>
              <p
                className={cssReservation.datesReservation}
              >{`Llegada: ${timeParser(arrival)}`}</p>
              <p
                className={cssReservation.datesReservation}
              >{`Fecha: ${dateParser(departure)}`}</p>
            </div>
          </Link>
          <ButtonSpinner
            className={`btn ${css.btnSet}`}
            disabled={checkined || !isCheckinStarted(departure)}
            onClick={handleCheckin}
            loading={loading}
            checkined={checkined}
          >
            {isCheckinStarted(departure) ? "Ya me subí" : "Fuera de horario"}
          </ButtonSpinner>
        </div>
      </div>
    </div>
  );
};
export default ReservationCard;
